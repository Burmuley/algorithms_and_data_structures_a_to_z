package linear_search

import (
	"cmp"
	"slices"
)

type Comparer[T cmp.Ordered] func(x, y T) int

func CompareOrdered[T cmp.Ordered](x, y T) bool {
	if x == y {
		return true
	}

	return false
}

type Node[K, V any] struct {
	next  *Node[K, V]
	key   K
	value V
}

type SymbolTable[K cmp.Ordered, V any] struct {
	comparer Comparer[K]
	head     *Node[K, V]
	count    int
}

func NewSymbolTable[K cmp.Ordered, V any](comparer Comparer[K]) *SymbolTable[K, V] {
	return &SymbolTable[K, V]{
		comparer: comparer,
	}
}

func (st *SymbolTable[K, V]) Count() int {
	return st.count
}

func (st *SymbolTable[K, V]) Empty() bool {
	return st.count == 0
}

func (st *SymbolTable[K, V]) Add(key K, value V) {
	for node := st.head; node != nil; node = node.next {
		if st.comparer(key, node.key) == 0 {
			node.value = value
			return
		}
	}

	st.head = &Node[K, V]{
		key:   key,
		value: value,
		next:  st.head,
	}
	st.count++

	return
}

func (st *SymbolTable[K, V]) TryGet(key K) (V, bool) {
	var zero V
	if st.count < 1 {
		return zero, false
	}

	for node := st.head; node != nil; node = node.next {
		if st.comparer(key, node.key) == 0 {
			return node.value, true
		}
	}

	return zero, false
}

func (st *SymbolTable[K, V]) Contains(key K) bool {
	if st.count < 1 {
		return false
	}

	for node := st.head; node != nil; node = node.next {
		if st.comparer(key, node.key) == 0 {
			return true
		}
	}

	return false
}

func (st *SymbolTable[K, V]) Remove(key K) bool {
	if st.count < 1 {
		return false
	}

	for cur, prev := st.head, st.head; cur != nil; prev, cur = cur, cur.next {
		if st.comparer(key, cur.key) == 0 {
			prev.next = cur.next
			st.count--
			return true
		}
	}

	return false
}

func (st *SymbolTable[K, V]) Keys() []K {
	keys := make([]K, 0, st.count)
	for node := st.head; node != nil; node = node.next {
		keys = append(keys, node.key)
	}
	slices.Reverse(keys)
	return keys
}
