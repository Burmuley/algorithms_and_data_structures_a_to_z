package test_data

import "cmp"

type args[T cmp.Ordered] struct {
	Input       []T
	SearchValue T
}
