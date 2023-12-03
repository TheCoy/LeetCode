package golang

import "fmt"

type MinStack struct {
	min   []any
	stack []any
}

func NewMinStack() *MinStack {
	return &MinStack{
		min:   make([]any, 0),
		stack: make([]any, 0),
	}
}

func (ms *MinStack) Pop() any {
	if len(ms.stack) <= 0 {
		return fmt.Errorf("empty")
	}
	result := ms.stack[len(ms.stack)-1]
	ms.min = ms.min[0 : len(ms.min)-1]
	ms.stack = ms.stack[0 : len(ms.stack)-1]

	return result
}

func (ms *MinStack) Top() any {
	if len(ms.stack) <= 0 {
		return fmt.Errorf("empty")
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Push(item any) {
	ms.stack = append(ms.stack, item)
	if len(ms.min) <= 0 {
		ms.min = append(ms.min, item)
	} else {
		latest := min(item.(int), ms.min[len(ms.min)-1].(int))
		ms.min = append(ms.min, latest)
	}
}
