package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	client *redis.Client
	ttl    time.Duration
}

func New(cl *redis.Client, ttl time.Duration) *Store {
	return &Store{
		client: cl,
		ttl:    ttl,
	}
}

func (s *Store) Incr(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Get(ctx context.Context, key string) (int, bool) {
	//TODO implement me
	panic("implement me")
}
