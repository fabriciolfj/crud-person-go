package service

import (
	"encoding/json"
	"fmt"
	"github.com/person/datasource"
	"github.com/person/model"
	"github.com/person/redis"
	"time"
)

func Save(p *model.Person) error {
	err := checkConnection()
	if err != nil {
		return err
	}

	p.GenerateUUID()
	tx := datasource.DB.Begin()
	result := tx.Create(p)

	if result.Error == nil {
		tx.Commit()
		SendMessage(p)

		json, _ := json.Marshal(p)
		if err := redis.RDB.Set(redis.CTX, p.Uuid, json, 30*time.Minute).Err(); err != nil {
			fmt.Println("fail put cache", err)
		}

		return result.Error
	}

	tx.Rollback()
	return result.Error
}

func Find(uuid string) (*model.Person, error) {
	p, err := redis.RDB.Get(redis.CTX, uuid).Result()
	if err == nil {
		var person model.Person
		err := json.Unmarshal([]byte(p), &person)
		if err != nil {
			panic(err)
		}

		fmt.Println("data found redis", person)
		return &person, nil
	}

	var person model.Person

	result := datasource.DB.Find(&person, "Uuid = ?", uuid)
	if result.Error != nil {
		return nil, result.Error
	}

	return &person, nil
}

func checkConnection() error {
	if datasource.DB == nil {
		return fmt.Errorf("DB is not initialized")
	}

	return nil
}
