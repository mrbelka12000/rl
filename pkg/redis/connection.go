package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetConnection(address string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   0,
	})

	err := client.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return client, nil
}
