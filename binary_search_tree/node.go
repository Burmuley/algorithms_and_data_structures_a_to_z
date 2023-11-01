package binary_search_tree

import "cmp"

type Node[T cmp.Ordered] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

func NewNode[T cmp.Ordered](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) Insert(value T) {
	c := cmp.Compare(value, n.value)

	if c == 0 {
		return
	}

	if c < 0 {
		if n.left == nil {
			n.left = NewNode[T](value)
			return
		}

		n.left.Insert(value)
		return
	}

	if c > 0 {
		if n.right == nil {
			n.right = NewNode[T](value)
			return
		}

		n.right.Insert(value)
		return
	}
}

func (n *Node[T]) Get(value T) *Node[T] {
	c := cmp.Compare(value, n.value)

	if c == 0 {
		return n
	}

	if c < 0 && n.left != nil {
		return n.left.Get(value)
	}

	if c > 0 && n.right != nil {
		return n.right.Get(value)
	}

	return nil
}

func (n *Node[T]) TraverseInOrder() []T {
	list := make([]T, 0)
	n.innerTraverseInOrder(&list)
	return list
}

func (n *Node[T]) innerTraverseInOrder(list *[]T) {
	if n.left != nil {
		n.left.innerTraverseInOrder(list)
	}

	*list = append(*list, n.value)

	if n.right != nil {
		n.right.innerTraverseInOrder(list)
	}
}

func (n *Node[T]) Min() T {
	if n.left != nil {
		return n.left.Min()
	}

	return n.value
}

func (n *Node[T]) Max() T {
	if n.right != nil {
		return n.right.Max()
	}

	return n.value
}

func (n *Node[T]) Value() T {
	return n.value
}
