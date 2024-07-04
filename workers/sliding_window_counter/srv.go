package sliding_window_counter

import (
	"time"

	"github.com/mrbelka12000/rl/workers"
)

type SlidingWindowCounter struct {
	worker workers.Cache
	limit  int
	ttl    time.Duration
}

func New(w workers.Cache, limit uint, ttl time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{}
}

func (s SlidingWindowCounter) Lock(key string) error {
	return nil
}
