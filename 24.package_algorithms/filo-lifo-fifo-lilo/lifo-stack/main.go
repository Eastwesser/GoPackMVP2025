package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop()) // is 2
	fmt.Println(stack.Pop()) // is 1, and if we Pop() more, it leads to panic
}
