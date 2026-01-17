package config

import (
	"context"

	"github.com/go-redis/redis/v8"
	logrus "github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

func NewRedis(cfg *Env, db int) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		logrus.Fatal("Failed to parse Redis URL: " + err.Error())
	}
	opts.DB = db // Ensure we are using the correct database

	client := redis.NewClient(opts)
	if err := client.Ping(context.Background()).Err(); err != nil {
		logrus.Fatal("Failed to connect to Redis: " + err.Error())
	}
	return client
	
}
