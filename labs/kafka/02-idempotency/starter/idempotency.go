package idempotency

import "sync"

type Store struct {
	mu   sync.Mutex
	seen map[string]struct{}
}

func New() *Store { return &Store{seen: map[string]struct{}{}} }
func (s *Store) Apply(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.seen[key]; ok {
		return false
	}
	s.seen[key] = struct{}{}
	return true
}
