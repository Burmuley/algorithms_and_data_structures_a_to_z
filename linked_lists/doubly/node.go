package doubly

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (node *Node[T]) Value() T {
	return node.value
}

func (node *Node[T]) SetValue(value T) *Node[T] {
	node.value = value
	return node
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (node *Node[T]) SetNext(next *Node[T]) *Node[T] {
	node.next = next
	return node
}

func (node *Node[T]) Prev() *Node[T] {
	return node.prev
}

func (node *Node[T]) SetPrev(prev *Node[T]) *Node[T] {
	node.prev = prev
	return node
}
