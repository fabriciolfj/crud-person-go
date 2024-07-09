package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var CTX = context.Background()
var RDB *redis.Client

func init() {

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := RDB.Set(CTX, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	_, err = RDB.Get(CTX, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("conexao redis ok")
}
