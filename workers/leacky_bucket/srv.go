package leacky_bucket

import (
	"time"

	"github.com/mrbelka12000/rl/workers"
)

type LeackyBucket struct {
	worker workers.Cache
	limit  uint
	ttl    time.Duration
}

func New(w workers.Cache, limit uint, ttl time.Duration) *LeackyBucket {
	return &LeackyBucket{}
}

func (l LeackyBucket) Lock(key string) error {
	return nil
}
