package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	client *redis.Client
}

func New(cl *redis.Client) *Store {
	return &Store{client: cl}
}

func (s *Store) Incr(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Expire(ctx context.Context, key string, ttl time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Get(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}
