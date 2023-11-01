package selection

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {

	for p := len(input) - 1; p > 0; p-- {
		largest := 0
		for i := 1; i <= p; i++ {
			if input[i] > input[largest] {
				largest = i
			}
		}

		input[largest], input[p] = input[p], input[largest]
	}

	return input
}
