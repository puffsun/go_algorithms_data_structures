package stack

import (
	"sync"
)

type Stack struct {
	stack []interface{}
	len   int
	lock  sync.Mutex
}

func New() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0

	return s
}

func (s *Stack) Push(element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	prepend := make([]interface{}, 1)
	prepend[0] = element
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Pop() (element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len == 0 {
		return nil
	}

	element, s.stack = s.stack[0], s.stack[1:]
	s.len--
	return
}

func (s *Stack) Peek() (element interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.len == 0 {
		return nil
	}

	element = s.stack[0]
	return
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}

func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}
