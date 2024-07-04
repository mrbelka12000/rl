package token_bucket

import (
	"time"

	"github.com/mrbelka12000/rl/workers"
)

type TokenBucket struct {
	worker workers.Cache
	limit  int
	ttl    time.Duration
}

func New(w workers.Cache, limit uint, ttl time.Duration) *TokenBucket {
	return &TokenBucket{}
}

func (t TokenBucket) Lock(key string) error {
	return nil
}
