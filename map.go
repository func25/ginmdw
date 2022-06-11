package ginmdw

import "sync"

// simple concurrent map many writes, 1 read
type syncMap[T comparable, K any] struct {
	m map[T]K
	l sync.Mutex
}

func (s *syncMap[T, K]) Get(key T) (K, bool) {
	s.l.Lock()
	defer s.l.Unlock()

	v, ok := s.m[key]
	return v, ok
}

func (s *syncMap[T, K]) Set(key T, value K) {
	s.l.Lock()
	defer s.l.Unlock()

	s.m[key] = value
}

func (s *syncMap[T, K]) GetSet(key T, f func(K) K) {
	s.l.Lock()
	defer s.l.Unlock()

	v, _ := s.m[key]
	v = f(v)
	s.m[key] = v
}
