package timer

import (
	"errors"
	"sync"
	"time"
)

type stackEntry struct {
	name string
	time time.Time
}

type timeStack struct {
	c    int
	data []*stackEntry
	lock sync.Mutex
}

func (s *timeStack) Depth() int {
	return s.c
}

func (s *timeStack) Push(name string, t time.Time) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, &stackEntry{name, t})
	s.c++
}

func (s *timeStack) Empty() bool {
	return s.c == 0
}

func (s *timeStack) Pop() (*stackEntry, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.Empty() {
		return nil, errors.New("empty stack")
	}

	res := s.data[s.c-1]
	s.data = s.data[:s.c-1]
	s.c--
	return res, nil
}
