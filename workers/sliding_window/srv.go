package sliding_window

import (
	"time"

	"github.com/mrbelka12000/rl/workers"
)

type SlidingWindow struct {
	worker workers.Cache
	limit  int
	ttl    time.Duration
}

func New(w workers.Cache, limit uint, ttl time.Duration) *SlidingWindow {
	return &SlidingWindow{}
}

func (s SlidingWindow) Lock(key string) error {
	return nil
}
