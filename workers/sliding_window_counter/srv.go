package sliding_window_counter

import (
	"github.com/mrbelka12000/rl/workers"
)

type SlidingWindowCounter struct {
	worker workers.Cache
	limit  uint
}

func New(w workers.Cache, limit uint) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		worker: w,
		limit:  limit,
	}
}

func (s SlidingWindowCounter) Lock(key string) error {
	return nil
}
