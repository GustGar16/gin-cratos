package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client = RedisConnect()
var ctx = context.Background()
var redisAddress = myEnv["REDIS_URL"] + ":" + myEnv["REDIS_PORT"]
var redisDB, _ = strconv.Atoi(myEnv["REDIS_DB"])

func RedisConnect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: myEnv["REDIS_PASS"],
		DB:       redisDB,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic("Error to connect with redis " + err.Error())
	}
	fmt.Println("Redis Connected")

	return client
}

func SetRedisData(key string, data string) bool {
	err := Redis.Set(ctx, key, data, 0).Err()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func GetRedisData(key string) string {
	val, err := Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("key " + key + " does not exist")
		return ""
	} else if err != nil {
		panic(err)
	}
	return val
}
