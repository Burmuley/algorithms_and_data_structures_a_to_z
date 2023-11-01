package array

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack[T any] struct {
	arr   []T
	count int
}

func NewStack[T any]() *Stack[T] {
	stack := &Stack[T]{}
	stack.arr = make([]T, 10)
	return stack
}

func (stack *Stack[T]) Count() int {
	return stack.count
}

func (stack *Stack[T]) Pop() (T, error) {
	var zero T
	if stack.count < 1 {
		return zero, ErrEmptyStack
	}

	elem := stack.arr[stack.count-1]
	stack.arr[stack.count-1] = zero
	stack.count--
	return elem, nil
}

func (stack *Stack[T]) Push(value T) {
	if stack.count == len(stack.arr) {
		longerArr := make([]T, stack.count*2)
		copy(longerArr, stack.arr)
		stack.arr = longerArr
	}

	stack.arr[stack.count] = value
	stack.count++
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.count < 1 {
		var zero T
		return zero, ErrEmptyStack
	}
	elem := stack.arr[stack.count-1]
	return elem, nil
}
