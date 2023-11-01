package bubble

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {
	for p := len(input) - 1; p > 0; p-- {
		for i := 0; i < p; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
			}
		}
	}

	return input
}
