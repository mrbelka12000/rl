package inmem

import "sync"

type Store struct {
	Data map[string]any
	sync.RWMutex
}

func New() *Store {
	return &Store{
		Data: make(map[string]any),
	}
}
