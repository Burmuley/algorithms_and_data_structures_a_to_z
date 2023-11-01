package chain_hash_map

import (
	"cmp"
	"fmt"
	"github.com/Burmuley/algorithms_and_data_structures_a_to_z/symbol_tables/linear_search"
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

type HashSet[K cmp.Ordered, V any] struct {
	chains   []*linear_search.SymbolTable[K, V]
	hash     hashFunc[K]
	comparer Comparer[K]
	count    int
	capacity int
	prime    Prime
}

func NewHashSet[K cmp.Ordered, V any]() *HashSet[K, V] {
	hs := &HashSet[K, V]{}
	hs.hash = murmurHash[K]
	hs.comparer = cmp.Compare[K]
	hs.prime = NewPrime()
	hs.capacity = int(hs.prime.getPrime(uint(defaultCapacity)))
	hs.chains = make([]*linear_search.SymbolTable[K, V], hs.capacity)
	for i := range hs.chains {
		hs.chains[i] = linear_search.NewSymbolTable[K, V](cmp.Compare[K])
	}
	return hs
}

func NewHashSetWithCapacity[K cmp.Ordered, V any](capacity int) *HashSet[K, V] {
	hs := &HashSet[K, V]{}

	hs.hash = murmurHash[K]
	hs.comparer = cmp.Compare[K]
	hs.prime = NewPrime()
	hs.capacity = int(hs.prime.getPrime(uint(capacity)))
	hs.chains = make([]*linear_search.SymbolTable[K, V], hs.capacity)
	for i := range hs.chains {
		hs.chains[i] = linear_search.NewSymbolTable[K, V](cmp.Compare[K])
	}
	return hs
}

func (hs *HashSet[K, V]) Count() int {
	return hs.count
}

func (hs *HashSet[K, V]) Get(key K) (V, bool) {
	hash := hs.hash(key, uint64(hs.capacity))
	return hs.chains[hash].TryGet(key)
}

func (hs *HashSet[K, V]) Add(key K, value V) {
	if hs.count >= 10*hs.capacity {
		hs.Resize(int(hs.prime.ExpandPrime(uint(hs.capacity))))
	}

	hash := hs.hash(key, uint64(hs.capacity))
	if !hs.chains[hash].Contains(key) {
		hs.count++
	}
	hs.chains[hash].Add(key, value)
}

func (hs *HashSet[K, V]) Contains(key K) bool {
	hash := hs.hash(key, uint64(hs.capacity))
	return hs.chains[hash].Contains(key)
}

func (hs *HashSet[K, V]) Remove(key K) bool {
	hash := hs.hash(key, uint64(hs.capacity))
	res := hs.chains[hash].Remove(key)
	if res {
		hs.count--
		if hs.capacity < defaultCapacity && hs.count <= (hs.capacity*2) {
			hs.Resize(int(hs.prime.ReducePrime(uint(hs.capacity))))
		}
	}

	return res
}

func (hs *HashSet[K, V]) Resize(capacity int) {
	newHS := NewHashSetWithCapacity[K, V](capacity)

	for i := range hs.chains {
		for _, k := range hs.chains[i].Keys() {
			v, _ := hs.chains[i].TryGet(k)
			newHS.Add(k, v)
		}
	}

	hs.capacity = newHS.capacity
	hs.chains = newHS.chains
	hs.count = newHS.count
}

func (hs *HashSet[K, V]) Keys() []K {
	keys := make([]K, 0, hs.count)
	for _, hv := range hs.chains {
		for _, k := range hv.Keys() {
			keys = append(keys, k)
		}
	}

	slices.SortFunc(keys, hs.comparer)
	return keys
}
