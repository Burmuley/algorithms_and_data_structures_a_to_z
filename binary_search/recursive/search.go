package recursive

import "cmp"

func BinarySearch[T cmp.Ordered](arr []T, value T) int {
	var inSearch func(low, high int) int

	inSearch = func(low, high int) int {
		if low >= high {
			return -1
		}

		mid := (low + high) / 2
		if arr[mid] == value {
			return mid
		}

		if arr[mid] < value {
			return inSearch(mid+1, high)
		} else {
			return inSearch(low, mid)
		}
	}

	return inSearch(0, len(arr))
}
