package inmem

import (
	"context"
	"sync"
	"time"
)

type (
	// Store in memory cache
	Store struct {
		cache map[string]item
		stash map[string]int64
		sync.RWMutex
		ttl time.Duration
	}

	item struct {
		value    int
		expireAt int64
	}
)

func New(ttl time.Duration) *Store {
	s := &Store{
		cache: make(map[string]item),
		stash: make(map[string]int64),
		ttl:   ttl,
	}

	s.Flush()

	return s
}

func (s *Store) Incr(ctx context.Context, key string) error {
	s.Lock()
	defer s.Unlock()

	i, ok := s.cache[key]
	if !ok {
		expires := time.Now().Add(s.ttl).Unix()
		i = item{
			value:    1,
			expireAt: expires,
		}
		s.cache[key] = i
		s.stash[key] = expires
		return nil
	}

	i.value++
	s.cache[key] = i

	return nil
}

func (s *Store) Get(ctx context.Context, key string) (int, bool) {
	s.RLock()
	defer s.RUnlock()

	i, ok := s.cache[key]
	if !ok {
		return 0, false
	}

	return i.value, true
}
