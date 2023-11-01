package binary_search_tree

import (
	"cmp"
	"errors"
)

var (
	ErrEmptyTree     = errors.New("empty tree")
	ErrValueNotFound = errors.New("value not found")
)

type BinarySearchTree[T cmp.Ordered] struct {
	root *Node[T]
}

func NewBST[T cmp.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (t *BinarySearchTree[T]) Get(value T) (*Node[T], error) {
	if t.root != nil {
		val := t.root.Get(value)
		if val != nil {
			return val, nil
		}

		return val, ErrValueNotFound
	}

	return nil, ErrEmptyTree
}

func (t *BinarySearchTree[T]) Min() (T, error) {
	var zero T
	if t.root == nil {
		return zero, ErrEmptyTree
	}

	return t.root.Min(), nil
}

func (t *BinarySearchTree[T]) Max() (T, error) {
	var zero T
	if t.root == nil {
		return zero, ErrEmptyTree
	}

	return t.root.Max(), nil
}

func (t *BinarySearchTree[T]) Insert(value T) {
	if t.root == nil {
		t.root = NewNode[T](value)
		return
	}

	t.root.Insert(value)
}

func (t *BinarySearchTree[T]) TraverseInOrder() []T {
	if t.root != nil {
		return t.root.TraverseInOrder()
	}

	return []T{}
}

func (t *BinarySearchTree[T]) Remove(value T) {
	t.root = t.prRemove(t.root, value)
}

func (t *BinarySearchTree[T]) prRemove(stRoot *Node[T], value T) *Node[T] {
	if stRoot == nil {
		return nil
	}

	c := cmp.Compare(value, stRoot.value)
	if c < 0 {
		stRoot.left = t.prRemove(stRoot.left, value)
	} else if c > 0 {
		stRoot.right = t.prRemove(stRoot.right, value)
	} else {
		if stRoot.left == nil {
			return stRoot.right
		}

		if stRoot.right == nil {
			return stRoot.left
		}

		stRoot.value = stRoot.right.Min()
		stRoot.right = t.prRemove(stRoot.right, stRoot.value)
	}

	return stRoot
}
