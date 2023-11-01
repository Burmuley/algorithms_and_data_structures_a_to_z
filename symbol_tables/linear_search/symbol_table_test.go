package linear_search

import (
	"cmp"
	"reflect"
	"testing"
)

func compareLinkedLists[K, V any](a, b *Node[K, V]) bool {
	for curA, curB := a, b; curA != nil || curB != nil; curA, curB = curA.next, curB.next {
		if !reflect.DeepEqual(curA, curB) {
			return false
		}
	}

	return true
}

func TestSymbolTable_Add(t *testing.T) {
	type args[K cmp.Ordered, V any] struct {
		key   K
		value V
	}
	type testCase[K cmp.Ordered, V any] struct {
		name string
		st   *SymbolTable[K, V]
		args args[K, V]
		want *SymbolTable[K, V]
	}
	tests := []testCase[int, int]{
		{
			name: "adding to empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			args: args[int, int]{
				key:   0,
				value: 100,
			},
			want: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(0, 100)
				return st
			}(),
		},
		{
			name: "adding to non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(22, 33)
				st.Add(44, 55)
				return st
			}(),
			args: args[int, int]{
				key:   0,
				value: 100,
			},
			want: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(22, 33)
				st.Add(44, 55)
				st.Add(0, 100)
				return st
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.st.Add(tt.args.key, tt.args.value); !compareLinkedLists[int, int](tt.st.head, tt.want.head) {
				t.Errorf("Add() = %v, want %v", *tt.st, *tt.want)
			}
		})
	}
}

func TestSymbolTable_Count(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name string
		st   *SymbolTable[K, V]
		want int
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			want: 0,
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(10, 100)
				return st
			}(),
			want: 1,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Count(); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolTable_Contains(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name string
		st   *SymbolTable[K, V]
		args args[K]
		want bool
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			args: args[int]{
				key: 100,
			},
			want: false,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Contains(tt.args.key); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolTable_TryGet(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name  string
		st    *SymbolTable[K, V]
		args  args[K]
		want  V
		want1 bool
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			args: args[int]{
				key: 100,
			},
			want:  0,
			want1: false,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			want:  10,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.TryGet(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TryGet() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TryGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSymbolTable_Remove(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name    string
		st      *SymbolTable[K, V]
		args    args[K]
		wantRes bool
		want    *SymbolTable[K, V]
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			args: args[int]{
				key: 100,
			},
			wantRes: false,
			want:    NewSymbolTable[int, int](CompareOrdered[int]),
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(10, 100)
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			wantRes: true,
			want:    NewSymbolTable[int, int](CompareOrdered[int]),
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			wantRes: true,
			want: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				st.Remove(10)
				return st
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Remove(tt.args.key); got != tt.wantRes && !compareLinkedLists[int, int](tt.st.head, tt.want.head) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolTable_Keys(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name string
		st   *SymbolTable[K, V]
		want []K
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](CompareOrdered[int]),
			want: []int{},
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				st.Add(10, 100)
				return st
			}(),
			want: []int{10},
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](CompareOrdered[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want: func() []int {
				keys := make([]int, 0, 20)
				for i := 0; i < 20; i++ {
					keys = append(keys, i)
				}
				return keys
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}
