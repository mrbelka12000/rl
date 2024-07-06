package fixed_window

import (
	"github.com/mrbelka12000/rl/workers"
)

type FixedWindow struct {
	worker workers.Cache
	limit  uint
}

func New(w workers.Cache, limit uint) *FixedWindow {
	return &FixedWindow{
		worker: w,
		limit:  limit,
	}
}

func (f FixedWindow) Lock(key string) error {
	return nil
}
