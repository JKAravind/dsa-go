package lists

import (
	"fmt"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Display() {
	for index, val := range s.items {
		fmt.Println(index, val)
	}
}

func (s *Stack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]

	//this just modifies the range of the slice and the leftout lastelement is cleared by go's garbase collector
	//we can also remove using append() refer google
}
