package sieve_of_eratosthenes

func Sieve(max int) []int {
	composite := make([]bool, max+1)
	primes := make([]int, 0)

	for p := 2; p <= max; p++ {
		if composite[p] {
			continue
		}

		primes = append(primes, p)

		for i := p * p; i <= max; i += p {
			composite[i] = true
		}
	}

	return primes
}
