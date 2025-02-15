package database

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var REDIS *redis.Client

func ConnectToRedis() {

	// Create a context
	ctx := context.Background()

	// Redis connect vars
	dsn := fmt.Sprintf("redis://:%s@%s:%s/0",
		os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))

	// Parsing dsn into redis url
	url, err := redis.ParseURL(dsn)
	if err != nil {
		log.Panic().Err(err)
	}

	// Connect to redis using url
	REDIS := redis.NewClient(url)

	// Test connection
	_, err = REDIS.Ping(ctx).Result()
	if err != nil {
		log.Panic().Err(err)
	} else {
		log.Info().Msg("Successfully connected to redis!")
	}
}
