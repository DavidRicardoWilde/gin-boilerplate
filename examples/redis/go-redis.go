package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	// create a client with go-redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "root",
		DB:       0,
	})
	ctx := context.Background()
	client.Conn(ctx)

	// set a key
	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println(err)
	}

	val := client.Get(ctx, "key")
	fmt.Println(val.Val())

	err = client.Close()
	if err != nil {
		log.Println(err)
	}
}
