package merge

import "cmp"

func Sort[T cmp.Ordered](input []T) []T {
	internalSort(0, len(input)-1, input)
	return input
}

func internalSort[T cmp.Ordered](low, high int, arr []T) {
	if high <= low {
		return
	}
	mid := (high + low) / 2
	internalSort[T](low, mid, arr)
	internalSort[T](mid+1, high, arr)
	merge(low, mid, high, arr)
	return
}

func merge[T cmp.Ordered](low, mid, high int, arr []T) {
	if arr[mid] <= arr[mid+1] {
		return
	}

	aux := make([]T, len(arr))
	copy(aux, arr)
	i, j := low, mid+1

	for k := low; k <= high; k++ {
		if i > mid {
			arr[k] = aux[j]
			j++
		} else if j > high {
			arr[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			arr[k] = aux[j]
			j++
		} else {
			arr[k] = aux[i]
			i++
		}
	}
}
