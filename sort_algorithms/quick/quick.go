package quick

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {
	internalSort(0, len(input)-1, input)
	return input
}

func internalSort[T cmp.Ordered](low, high int, arr []T) {
	if high <= low {
		return
	}

	j := partition(low, high, arr)
	internalSort(low, j-1, arr)
	internalSort(j+1, high, arr)
}

func partition[T cmp.Ordered](low, high int, arr []T) int {
	i, j := low, high+1
	pivot := arr[low]

	for {
		for i++; arr[i] < pivot; i++ {
			if i == high {
				break
			}
		}

		for j--; pivot < arr[j]; j-- {
			if j == low {
				break
			}
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[low], arr[j] = arr[j], arr[low]

	return j
}
