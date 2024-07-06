package inmem

import (
	"time"
)

const (
	defaultFlushInterval = 5 * time.Second
)

// Flush starts GC for cache
func (s *Store) Flush() {
	go s.flusher()
}

func (s *Store) flusher() {
	for {
		<-time.After(defaultFlushInterval)

		s.RLock()

		now := time.Now().Unix()

		for k, expiresAt := range s.stash {
			if expiresAt < now {
				delete(s.stash, k)
				delete(s.cache, k)
			}
		}

		s.RUnlock()
	}
}
