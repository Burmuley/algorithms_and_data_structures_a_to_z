package chain_hash_map

import "math"

const (
	hashPrime           = 101
	maxPrimeArrayLength = 0x7FEFFFFD
)

type Prime struct {
	predefined []uint
	minPrime   uint
}

func NewPrime() Prime {
	prime := Prime{
		predefined: []uint{53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613, 393241, 786433,
			1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741},
	}
	prime.minPrime = prime.predefined[0]
	return prime
}

func (p Prime) getPrime(minp uint) uint {
	for i := 0; i < len(p.predefined); i++ {
		if minp >= p.predefined[i] {
			return p.predefined[i]
		}
	}

	for i := minp | 1; i < math.MaxUint; i += 2 {
		if isPrime(i) && (i-1)%hashPrime != 0 {
			p.predefined = append(p.predefined, i)
			return i
		}
	}

	return minp
}

func (p Prime) ExpandPrime(oldSize uint) uint {
	newSize := oldSize * 2
	if newSize > maxPrimeArrayLength && maxPrimeArrayLength > oldSize {
		return maxPrimeArrayLength
	}
	return p.getPrime(newSize)
}

func (p Prime) ReducePrime(oldSize uint) uint {
	newSize := oldSize / 2
	if newSize > maxPrimeArrayLength && maxPrimeArrayLength > oldSize {
		return maxPrimeArrayLength
	}
	return p.getPrime(newSize)
}

func isPrime(n uint) bool {
	if n%2 != 0 {
		limit := uint(math.Sqrt(float64(n)))
		for i := uint(3); i <= limit; i += 2 {
			if n%i == 0 {
				return false
			}
		}

		return true
	}

	return n == 2
}
