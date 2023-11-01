package shell

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {
	gap := 1

	for gap < len(input)/3 {
		gap = 3*gap + 1
	}

	for gap >= 1 {
		for i := gap; i < len(input); i++ {
			for j := i; j >= gap && input[j] < input[j-gap]; j -= gap {
				input[j], input[j-gap] = input[j-gap], input[j]
			}
		}
		gap /= 3
	}

	return input
}
