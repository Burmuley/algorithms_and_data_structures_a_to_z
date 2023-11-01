package insertion

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {
	for p := 1; p < len(input); p++ {
		curUnsorted := input[p]
		i := 0
		for i = p; i > 0 && input[i-1] > curUnsorted; i-- {
			input[i] = input[i-1]
		}
		input[i] = curUnsorted
	}
	return input
}
