package max_heap

import (
	"cmp"
	"errors"
)

const defaultCapacity = 4

var (
	ErrEmptyHeap    = errors.New("empty heap")
	ErrInvalidIndex = errors.New("index does not exists in the heap")
)

type Heap[T cmp.Ordered] struct {
	count int
	list  []T
}

func NewHeap[T cmp.Ordered]() *Heap[T] {
	return NewHeapWCapacity[T](defaultCapacity)
}

func NewHeapWCapacity[T cmp.Ordered](capacity int) *Heap[T] {
	h := &Heap[T]{
		count: 0,
		list:  make([]T, capacity),
	}

	return h
}

func (h *Heap[T]) Insert(value T) {
	if h.count == len(h.list) {
		newList := make([]T, h.count*2)
		copy(newList, h.list)
		h.list = newList
	}

	h.list[h.count] = value
	h.float(h.count)
	h.count++
	return
}

func (h *Heap[T]) Values() []T {
	vals := make([]T, 0, h.count)
	for i := 0; i < h.count; i++ {
		vals = append(vals, h.list[i])
	}
	return vals
}

func (h *Heap[T]) ValuesSorted() []T {
	vals := make([]T, h.count)
	lastIndex := h.count - 1
	for i := 0; i <= lastIndex; i++ {
		vals[lastIndex-i] = h.list[0]
		h.list[0], h.list[lastIndex-i] = h.list[lastIndex-i], h.list[0]
		h.sink(0, lastIndex-i-1)

	}
	return vals
}

func (h *Heap[T]) Remove() (T, error) {
	return h.RemoveIndex(0)
}

func (h *Heap[T]) RemoveIndex(index int) (T, error) {
	var zero T
	if h.count < 1 {
		return zero, ErrEmptyHeap
	}
	if index > h.count {
		return zero, ErrInvalidIndex
	}

	removedValue := h.list[index]
	h.list[index] = h.list[h.count-1]
	if index == 0 || cmp.Less(h.list[index], h.list[parentIndex(index)]) {
		h.sink(index, h.count-1)
	} else {
		h.float(index)
	}

	h.count--
	return removedValue, nil
}

func (h *Heap[T]) Peek() (T, error) {
	if h.count < 1 {
		var zero T
		return zero, ErrEmptyHeap
	}
	return h.list[0], nil
}

func (h *Heap[T]) float(index int) {
	newVal := h.list[index]

	for index > 0 && cmp.Less(h.list[parentIndex(index)], h.list[index]) {
		h.list[index] = h.list[parentIndex(index)]
		index = parentIndex(index)
		h.list[index] = newVal
	}

	h.list[index] = newVal
}

func (h *Heap[T]) sink(index, lastIndex int) {
	for index <= lastIndex {
		leftIndex := leftChildIndex(index)
		rightIndex := rightChildIndex(index)

		if leftIndex > lastIndex {
			break
		}

		// determine child index to swap
		childIndex := 0
		if rightIndex > lastIndex {
			childIndex = leftIndex
		} else {
			if cmp.Less(h.list[leftIndex], h.list[rightIndex]) {
				childIndex = rightIndex
			} else {
				childIndex = leftIndex
			}
		}

		if cmp.Less(h.list[index], h.list[childIndex]) {
			h.list[index], h.list[childIndex] = h.list[childIndex], h.list[index]
		} else {
			break
		}

		index = childIndex
	}
}

func parentIndex(index int) int {
	return (index - 1) / 2
}

func leftChildIndex(index int) int {
	return (2 * index) + 1
}

func rightChildIndex(index int) int {
	return (2 * index) + 2
}
