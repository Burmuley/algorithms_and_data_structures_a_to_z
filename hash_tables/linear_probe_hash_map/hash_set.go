package linear_probe_hash_map

import (
	"cmp"
	"fmt"
	"github.com/spaolacci/murmur3"
	"slices"
)

const defaultCapacity = 10

type hashFunc[K comparable] func(key K, capacity uint64) uint64
type Comparer[T cmp.Ordered] func(x, y T) int

func murmurHash[K comparable](v K, m uint64) uint64 {
	mur := murmur3.New64()
	_, _ = mur.Write([]byte(fmt.Sprint(v)))
	index := mur.Sum64() & 0x7fffffff % m
	return index
}

type keyVal[K cmp.Ordered, V any] struct {
	key    K
	value  V
	filled bool
}

type HashSet[K cmp.Ordered, V any] struct {
	entries  []keyVal[K, V]
	hash     hashFunc[K]
	comparer Comparer[K]
	count    int
	capacity int
	//prime    Prime
}

func NewHashSet[K cmp.Ordered, V any]() *HashSet[K, V] {
	hs := &HashSet[K, V]{}
	hs.hash = murmurHash[K]
	hs.comparer = cmp.Compare[K]
	//hs.prime = NewPrime()
	//hs.capacity = int(hs.prime.getPrime(uint(defaultCapacity)))
	hs.capacity = defaultCapacity
	hs.entries = make([]keyVal[K, V], hs.capacity)

	return hs
}

func NewHashSetWithCapacity[K cmp.Ordered, V any](capacity int) *HashSet[K, V] {
	hs := &HashSet[K, V]{}

	hs.hash = murmurHash[K]
	hs.comparer = cmp.Compare[K]
	//hs.prime = NewPrime()
	//hs.capacity = int(hs.prime.getPrime(uint(capacity)))
	hs.capacity = capacity
	hs.entries = make([]keyVal[K, V], hs.capacity)
	return hs
}

func (hs *HashSet[K, V]) Count() int {
	return hs.count
}

func (hs *HashSet[K, V]) Get(key K) (V, bool) {
	hash := hs.hash(key, uint64(hs.capacity))
	for i := hash; hs.entries[i].filled != false; i = (i + 1) % uint64(hs.capacity) {
		if hs.comparer(key, hs.entries[i].key) == 0 {
			return hs.entries[i].value, true
		}
	}
	var zero V
	return zero, false
}

func (hs *HashSet[K, V]) Add(key K, value V) {
	//if hs.count >= 10*hs.capacity {
	//	hs.Resize(int(hs.prime.ExpandPrime(uint(hs.capacity))))
	//}

	if hs.count >= (hs.capacity / 2) {
		hs.Resize(hs.capacity * 2)
	}

	hash := hs.hash(key, uint64(hs.capacity))
	var i uint64
	for i = hash; hs.entries[i].filled != false; i = (i + 1) % uint64(hs.capacity) {
		if hs.comparer(key, hs.entries[i].key) == 0 {
			hs.entries[i].value = value
			return
		}
	}

	hs.entries[i] = keyVal[K, V]{
		filled: true,
		key:    key,
		value:  value,
	}
	hs.count++
}

func (hs *HashSet[K, V]) Contains(key K) bool {
	hash := hs.hash(key, uint64(hs.capacity))
	for i := hash; hs.entries[i].filled != false; i = (i + 1) % uint64(hs.capacity) {
		if hs.comparer(key, hs.entries[i].key) == 0 {
			return true
		}
	}
	return false
}

func (hs *HashSet[K, V]) tryGet(key K) (uint, bool) {
	hash := hs.hash(key, uint64(hs.capacity))
	for i := hash; hs.entries[i].filled != false; i = (i + 1) % uint64(hs.capacity) {
		if hs.comparer(key, hs.entries[i].key) == 0 {
			return uint(i), true
		}
	}
	return 0, false
}

func (hs *HashSet[K, V]) Remove(key K) bool {
	idx, found := hs.tryGet(key)
	if !found {
		return false
	}

	var zeroKey K
	var zeroVal V

	hs.entries[idx] = keyVal[K, V]{
		key:    zeroKey,
		value:  zeroVal,
		filled: false,
	}
	hs.count--
	idx = (idx + 1) % uint(hs.capacity)

	for hs.entries[idx].filled != false {
		tmp := hs.entries[idx]
		hs.entries[idx] = keyVal[K, V]{
			key:    zeroKey,
			value:  zeroVal,
			filled: false,
		}
		hs.count--
		hs.Add(tmp.key, tmp.value)
		idx = (idx + 1) % uint(hs.capacity)
	}

	if hs.count > 0 && hs.count <= (hs.capacity/8) {
		hs.Resize(hs.capacity / 2)
	}

	//if hs.capacity < defaultCapacity && hs.count <= (hs.capacity*2) {
	//	hs.Resize(int(hs.prime.ReducePrime(uint(hs.capacity))))
	//}

	return true
}

func (hs *HashSet[K, V]) Resize(capacity int) {
	newHS := NewHashSetWithCapacity[K, V](capacity)

	for _, v := range hs.entries {
		if !v.filled {
			continue
		}
		newHS.Add(v.key, v.value)
	}

	hs.capacity = newHS.capacity
	hs.entries = newHS.entries
	hs.count = newHS.count
}

func (hs *HashSet[K, V]) Keys() []K {
	keys := make([]K, 0, hs.count)
	for _, k := range hs.entries {
		if k.filled {
			keys = append(keys, k.key)
		}
	}

	slices.SortFunc(keys, hs.comparer)
	return keys
}
