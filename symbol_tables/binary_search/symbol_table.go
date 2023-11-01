package binary_search

import (
	"cmp"
)

const defaultCapacity = 10

type Comparer[T cmp.Ordered] func(x, y T) int

type SymbolTable[K cmp.Ordered, V any] struct {
	comparer Comparer[K]
	count    int
	keys     []K
	values   []V
}

func NewSymbolTable[K, V cmp.Ordered](comparer Comparer[K]) *SymbolTable[K, V] {
	return &SymbolTable[K, V]{
		comparer: comparer,
		keys:     make([]K, defaultCapacity),
		values:   make([]V, defaultCapacity),
		count:    0,
	}
}

func (st *SymbolTable[K, V]) Count() int {
	return st.count
}

func (st *SymbolTable[K, V]) Empty() bool {
	return st.count == 0
}

func (st *SymbolTable[K, V]) Add(key K, value V) {
	rank := st.Rank(key)
	if rank < st.Count() && st.comparer(st.keys[rank], key) == 0 {
		st.values[rank] = value
		return
	}

	if st.Count() == cap(st.keys) {
		largerKeysArr := make([]K, st.Count()*2, st.Count()*2)
		largerValArr := make([]V, st.Count()*2, st.Count()*2)
		copy(largerKeysArr, st.keys)
		copy(largerValArr, st.values)
		st.keys = largerKeysArr
		st.values = largerValArr
	}

	for i := st.Count() - 1; i > rank; i-- {
		st.keys[i] = st.keys[i-1]
		st.values[i] = st.values[i-1]
	}

	st.keys[rank] = key
	st.values[rank] = value
	st.count++
}

func (st *SymbolTable[K, V]) TryGet(key K) (V, bool) {
	var zero V
	if st.Count() == 0 {
		return zero, false
	}
	rank := st.Rank(key)
	if rank < st.Count() && st.comparer(st.keys[rank], key) == 0 {
		return st.values[rank], true
	}

	return zero, false
}

func (st *SymbolTable[K, V]) Contains(key K) bool {
	rank := st.Rank(key)

	if rank < st.Count() && st.comparer(st.keys[rank], key) == 0 {
		return true
	}

	return false
}

func (st *SymbolTable[K, V]) Remove(key K) bool {
	if st.Count() == 0 {
		return false
	}

	rank := st.Rank(key)
	if rank == st.Count() || st.comparer(st.keys[rank], key) != 0 {
		return false
	}

	for i := rank; i < st.Count()-1; i++ {
		st.keys[i] = st.keys[i+1]
		st.values[i] = st.values[i+1]
	}

	var zeroKey K
	var zeroVal V
	st.count--
	st.keys[st.count] = zeroKey
	st.values[st.count] = zeroVal
	return true
}

func (st *SymbolTable[K, V]) Keys() []K {
	keys := make([]K, 0, st.Count())
	for i := 0; i < st.Count(); i++ {
		keys = append(keys, st.keys[i])
	}
	return keys
}

func (st *SymbolTable[K, V]) Rank(key K) int {
	low, high := 0, st.Count()-1

	for low <= high {
		mid := low + (high-low)/2
		cmpKey := st.comparer(key, st.keys[mid])

		if cmpKey < 0 {
			high = mid - 1
		} else if cmpKey > 0 {
			low = mid + 1
		} else {
			return mid
		}
	}

	return low
}

func (st *SymbolTable[K, V]) Min() (K, bool) {
	var zero K
	if st.Count() == 0 {
		return zero, false
	}

	return st.keys[0], true
}

func (st *SymbolTable[K, V]) Max() (K, bool) {
	var zero K
	if st.Count() == 0 {
		return zero, false
	}

	return st.keys[st.Count()-1], true
}

func (st *SymbolTable[K, V]) RemoveMin() bool {
	if st.Count() == 0 {
		return false
	}

	_min, _ := st.Min()
	return st.Remove(_min)
}

func (st *SymbolTable[K, V]) RemoveMax() bool {
	if st.Count() == 0 {
		return false
	}

	_max, _ := st.Max()
	return st.Remove(_max)
}

func (st *SymbolTable[K, V]) Select(index int) (K, bool) {
	var zero K
	if index < 0 || index > st.Count()-1 || st.Count() == 0 {
		return zero, false
	}

	return st.keys[index], true
}

func (st *SymbolTable[K, V]) Ceiling(key K) (K, bool) {
	var zero K
	if st.Count() == 0 {
		return zero, false
	}

	rank := st.Rank(key)
	if rank == st.Count() {
		return zero, false
	}

	return st.keys[rank], true
}

func (st *SymbolTable[K, V]) Floor(key K) (K, bool) {
	var zero K
	if st.Count() == 0 {
		return zero, false
	}

	rank := st.Rank(key)
	cmpRes := st.comparer(key, st.keys[rank])
	if rank < st.Count() && cmpRes == 0 {
		return st.keys[rank], true
	}
	if rank == 0 {
		return zero, false
	}

	return st.keys[rank-1], true
}

func (st *SymbolTable[K, V]) Range(low, high K) ([]K, bool) {
	if st.Count() == 0 {
		return []K{}, false
	}

	lowRank := st.Rank(low)
	highRank := st.Rank(high)
	keys := make([]K, 0, highRank-lowRank+1)
	for i := lowRank; i < highRank; i++ {
		keys = append(keys, st.keys[i])
	}

	if st.Contains(high) {
		keys = append(keys, st.keys[highRank])
	}

	return keys, true
}
