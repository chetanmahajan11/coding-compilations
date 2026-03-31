package main

import (
	"fmt"
	"sync"
)

type stack struct {
	st []int
	mu sync.Mutex
}

// push inserts new elements in the stack
func (s *stack) push(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.st = append(s.st, val)
}

// pop deletes the top of the stack and returns the previous top
func (s *stack) pop() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.st) == 0 {
		return 0, fmt.Errorf("stack already empty!!! Can't perform pop operation...")
	}
	top := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return top, nil
}

// top returns current top
func (s *stack) top() (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.st) == 0 {
		return 0, fmt.Errorf("stack is empty!!")
	}
	return s.st[len(s.st)-1], nil
}

// isEmpty check if the current stack is empty or not
func (s *stack) isEmpty() bool {
	return len(s.st) == 0
}
