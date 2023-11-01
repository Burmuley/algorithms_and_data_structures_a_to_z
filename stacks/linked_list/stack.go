package linked_list

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type node[T any] struct {
	next  *node[T]
	value T
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}
}

type Stack[T any] struct {
	head  *node[T]
	count int
}

func NewStack[T any]() *Stack[T] {
	stack := &Stack[T]{}
	return stack
}

func (stack *Stack[T]) Count() int {
	return stack.count
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.count == 0 {
		var zero T
		return zero, ErrEmptyStack
	}

	tmp := stack.head
	stack.head = stack.head.next
	stack.count--

	return tmp.value, nil
}

func (stack *Stack[T]) Push(value T) {
	tmp := newNode(value)
	tmp.next = stack.head
	stack.head = tmp
	stack.count++
}

func (stack *Stack[T]) Peek() (T, error) {
	if stack.count == 0 {
		var zero T
		return zero, ErrEmptyStack
	}

	return stack.head.value, nil
}
