package ring

import (
	"errors"
	"slices"
)

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
	if queue.head <= queue.tail {
		return queue.tail - queue.head
	}

	return queue.tail - queue.head + len(queue.arr)
}

func (queue *Queue[T]) Enqueue(value T) {
	if len(queue.arr)-1 == queue.Count() {
		prevCount := queue.Count()
		largerArr := make([]T, queue.Count()*2)
		copy(largerArr, queue.arr[queue.head:])
		slices.Insert(queue.arr, queue.head, queue.arr[0:queue.tail]...)
		queue.arr = largerArr
		queue.head = 0
		queue.tail = prevCount
	}

	queue.arr[queue.tail] = value
	if queue.tail < len(queue.arr)-1 {
		queue.tail++
	} else {
		queue.tail = 0
	}
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
	} else if queue.head == len(queue.arr) {
		queue.head = 0
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
