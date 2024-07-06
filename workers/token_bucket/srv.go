package token_bucket

import (
	"context"

	"github.com/mrbelka12000/rl/workers"
)

type TokenBucket struct {
	worker workers.Cache
	limit  uint
}

func New(w workers.Cache, limit uint) *TokenBucket {
	return &TokenBucket{
		worker: w,
		limit:  limit,
	}
}

func (t TokenBucket) Lock(key string) error {
	return t.worker.Incr(context.Background(), key)
}
