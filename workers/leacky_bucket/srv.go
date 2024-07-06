package leacky_bucket

import (
	"github.com/mrbelka12000/rl/workers"
)

type LeackyBucket struct {
	worker workers.Cache
	limit  uint
}

func New(w workers.Cache, limit uint) *LeackyBucket {
	return &LeackyBucket{
		worker: w,
		limit:  limit,
	}
}

func (l LeackyBucket) Lock(key string) error {
	return nil
}
