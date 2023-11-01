package iterative

import "cmp"

func BinarySearch[T cmp.Ordered](arr []T, value T) int {
	low, high := 0, len(arr)

	for low < high {
		mid := (low + high) / 2

		if arr[mid] == value {
			return mid
		}

		if arr[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return -1
}
