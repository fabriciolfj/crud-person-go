package service

import (
	"github.com/person/datasource"
	"github.com/person/model"

	"fmt"
)

func Save(p *model.Person) error {
	err := checkConnection()
	if err != nil {
		return err
	}

	p.GenerateUUID()
	result := datasource.DB.Create(p)

	SendMessage(p)
	return result.Error
}

func Find(uuid string) (*model.Person, error) {
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
