package sliding_window

import (
	"github.com/mrbelka12000/rl/workers"
)

type SlidingWindow struct {
	worker workers.Cache
	limit  uint
}

func New(w workers.Cache, limit uint) *SlidingWindow {
	return &SlidingWindow{
		worker: w,
		limit:  limit,
	}
}

func (s SlidingWindow) Lock(key string) error {
	return nil
}
