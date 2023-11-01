package dummy

import "errors"

const defaultCapacity = 10

var ErrEmptyQueue = errors.New("queue is empty")

type Queue[T any] struct {
	head, tail int
	arr        []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{arr: make([]T, defaultCapacity)}
}

func NewQueueWithCapacity[T any](capacity int) *Queue[T] {
	return &Queue[T]{arr: make([]T, capacity)}
}

func (queue *Queue[T]) Count() int {
	return queue.tail - queue.head
}

func (queue *Queue[T]) Enqueue(value T) {
	if len(queue.arr) == queue.Count() {
		largerArr := make([]T, queue.Count()*2)
		copy(largerArr, queue.arr)
		queue.arr = largerArr
	}

	queue.arr[queue.tail] = value
	queue.tail++
}

func (queue *Queue[T]) Dequeue() error {
	if queue.Count() == 0 {
		return ErrEmptyQueue
	}

	var zero T
	queue.arr[queue.head] = zero
	queue.head++

	if queue.Count() == 0 {
		queue.head = 0
		queue.tail = 0
	}

	return nil
}

func (queue *Queue[T]) Peek() (T, error) {
	var zero T
	if queue.Count() == 0 {
		return zero, ErrEmptyQueue
	}

	return queue.arr[queue.head], nil
}
