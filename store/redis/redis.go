package redis

import "github.com/redis/go-redis/v9"

type Store struct {
	client *redis.Client
}

func New(cl *redis.Client) *Store {
	return &Store{client: cl}
}
