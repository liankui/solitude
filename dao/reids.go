package dao

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
)

var Redis *redis.Client

func openRedis(addr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis connnect error: %v", err)
	}
	return client
}

func NewRedis() *redis.Client {
	return openRedis(
		viper.GetString("Redis.Host"),
		viper.GetString("Redis.Password"),
		viper.GetInt("Redis.DB"),
		)
}