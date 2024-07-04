package inmem

import (
	"context"
	"sync"
	"time"
)

type Store struct {
	Data map[string]any
	sync.RWMutex
}

func New() *Store {
	return &Store{
		Data: make(map[string]any),
	}
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
