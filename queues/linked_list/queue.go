package linked_list

import (
	"errors"
)

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	list *LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: NewLinkedList[T]()}
}

func (queue *Queue[T]) Count() int {
	return queue.list.Count()
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.list.AddLastValue(value)
}

func (queue *Queue[T]) Dequeue() error {
	if queue.Count() == 0 {
		return ErrEmptyQueue
	}

	queue.list.RemoveFirst()
	return nil
}

func (queue *Queue[T]) Peek() (T, error) {
	var zero T
	if queue.Count() == 0 {
		return zero, ErrEmptyQueue
	}

	return queue.list.Head().Value(), nil
}
