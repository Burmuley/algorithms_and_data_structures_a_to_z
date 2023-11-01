package binary_search

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			args: args[int, int]{
				key:   0,
				value: 100,
			},
			want: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(0, 100)
				return st
			}(),
		},
		{
			name: "adding to non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(22, 33)
				st.Add(44, 55)
				return st
			}(),
			args: args[int, int]{
				key:   0,
				value: 100,
			},
			want: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(22, 33)
				st.Add(44, 55)
				st.Add(0, 100)
				return st
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.st.Add(tt.args.key, tt.args.value); !(slices.Compare(tt.st.keys, tt.want.keys) == 0) || !(slices.Compare(tt.st.keys, tt.want.keys) == 0) {
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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			want: 0,
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want: 1,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			args: args[int]{
				key: 100,
			},
			want: false,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			args: args[int]{
				key: 100,
			},
			want:  0,
			want1: false,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			args: args[int]{
				key: 100,
			},
			wantRes: false,
			want:    NewSymbolTable[int, int](cmp.Compare[int]),
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			wantRes: true,
			want:    NewSymbolTable[int, int](cmp.Compare[int]),
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
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
				st := NewSymbolTable[int, int](cmp.Compare[int])
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
			if got := tt.st.Remove(tt.args.key); got != tt.wantRes && !(slices.Compare(tt.st.keys, tt.want.keys) == 0) || !(slices.Compare(tt.st.keys, tt.want.keys) == 0) {
				t.Errorf("Remove() = %v, %v want %v, %v", got, tt.st.keys, tt.wantRes, tt.want.keys)
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
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			want: []int{},
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want: []int{10},
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
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

func TestSymbolTable_Rank(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name string
		st   *SymbolTable[K, V]
		args args[K]
		want int
	}
	tests := []testCase[int, int]{
		{
			name: "empty table",
			st:   NewSymbolTable[int, int](cmp.Compare[int]),
			args: args[int]{
				key: 10,
			},
			want: 0,
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			want: 0,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args: args[int]{
				key: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Rank(tt.args.key); got != tt.want {
				t.Errorf("Rank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolTable_Min(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		want     K
		wantCond bool
	}
	tests := []testCase[int, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[int, int](cmp.Compare[int]),
			want:     0,
			wantCond: false,
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want:     10,
			wantCond: true,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want:     0,
			wantCond: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, cond := tt.st.Min()
			if cond != tt.wantCond {
				t.Errorf("Min() condition = %v, wantCond %v", cond, tt.wantCond)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSymbolTable_Max(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		want     K
		wantCond bool
	}
	tests := []testCase[int, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[int, int](cmp.Compare[int]),
			want:     0,
			wantCond: false,
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want:     10,
			wantCond: true,
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want:     19,
			wantCond: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.Max()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Max() condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}

func TestSymbolTable_RemoveMin(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name    string
		st      *SymbolTable[K, V]
		want    bool
		wantRes *SymbolTable[K, V]
	}
	tests := []testCase[int, int]{
		{
			name:    "empty table",
			st:      NewSymbolTable[int, int](cmp.Compare[int]),
			want:    false,
			wantRes: NewSymbolTable[int, int](cmp.Compare[int]),
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want: true,
			wantRes: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				return st
			}(),
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want: true,
			wantRes: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				st.Remove(0)
				return st
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.st.RemoveMin()
			if got != tt.want {
				t.Errorf("RemoveMin() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.st.keys, tt.wantRes.keys) || !reflect.DeepEqual(tt.st.values, tt.wantRes.values) {
				t.Errorf("Min() got keys = %v values = %v, want keys = %v values = %v", tt.st.keys, tt.st.values, tt.wantRes.keys, tt.wantRes.values)
			}
		})
	}
}

func TestSymbolTable_RemoveMax(t *testing.T) {
	type testCase[K cmp.Ordered, V any] struct {
		name    string
		st      *SymbolTable[K, V]
		want    bool
		wantRes *SymbolTable[K, V]
	}
	tests := []testCase[int, int]{
		{
			name:    "empty table",
			st:      NewSymbolTable[int, int](cmp.Compare[int]),
			want:    false,
			wantRes: NewSymbolTable[int, int](cmp.Compare[int]),
		},
		{
			name: "single element table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			want: true,
			wantRes: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				return st
			}(),
		},
		{
			name: "non-empty table",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			want: true,
			wantRes: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				st.Remove(19)
				return st
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.st.RemoveMax()
			if got != tt.want {
				t.Errorf("RemoveMin() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(tt.st.keys, tt.wantRes.keys) || !reflect.DeepEqual(tt.st.values, tt.wantRes.values) {
				t.Errorf("Max() got keys = %v values = %v, want keys = %v values = %v", tt.st.keys, tt.st.values, tt.wantRes.keys, tt.wantRes.values)
			}
		})
	}
}

func TestSymbolTable_Select(t *testing.T) {
	type args struct {
		index int
	}
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		args     args
		want     K
		wantCond bool
	}
	tests := []testCase[int, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[int, int](cmp.Compare[int]),
			args:     args{index: 3},
			wantCond: false,
			want:     0,
		},
		{
			name: "single element table (not found)",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			args:     args{index: 3},
			want:     0,
			wantCond: false,
		},
		{
			name: "single element table (found)",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				st.Add(10, 100)
				return st
			}(),
			args:     args{index: 0},
			want:     10,
			wantCond: true,
		},
		{
			name: "non-empty table (found)",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args:     args{index: 10},
			want:     10,
			wantCond: true,
		},
		{
			name: "non-empty table (not found)",
			st: func() *SymbolTable[int, int] {
				st := NewSymbolTable[int, int](cmp.Compare[int])
				for i := 0; i < 20; i++ {
					st.Add(i, i)
				}
				return st
			}(),
			args:     args{index: 44},
			want:     0,
			wantCond: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.Select(tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Select() got condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}

func TestSymbolTable_Floor(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		args     args[K]
		want     K
		wantCond bool
	}
	tests := []testCase[string, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[string, int](cmp.Compare[string]),
			args:     args[string]{key: "c"},
			want:     "",
			wantCond: false,
		},
		{
			name: "multiple elements #1",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "g"},
			want:     "f",
			wantCond: true,
		},
		{
			name: "multiple elements #1",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "e"},
			want:     "e",
			wantCond: true,
		},
		{
			name: "multiple elements #1",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "a"},
			want:     "",
			wantCond: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.Floor(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Floor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Floor() got condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}

func TestSymbolTable_Ceiling(t *testing.T) {
	type args[K cmp.Ordered] struct {
		key K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		args     args[K]
		want     K
		wantCond bool
	}
	tests := []testCase[string, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[string, int](cmp.Compare[string]),
			args:     args[string]{key: "c"},
			want:     "",
			wantCond: false,
		},
		{
			name: "multiple elements #1",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "b"},
			want:     "b",
			wantCond: true,
		},
		{
			name: "multiple elements #2",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "a"},
			want:     "b",
			wantCond: true,
		},
		{
			name: "multiple elements #3",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{key: "g"},
			want:     "",
			wantCond: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.Ceiling(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ceiling() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Ceiling() got condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}

func TestSymbolTable_Range(t *testing.T) {
	type args[K cmp.Ordered] struct {
		startKey K
		endKey   K
	}
	type testCase[K cmp.Ordered, V any] struct {
		name     string
		st       *SymbolTable[K, V]
		args     args[K]
		want     []K
		wantCond bool
	}
	tests := []testCase[string, int]{
		{
			name:     "empty table",
			st:       NewSymbolTable[string, int](cmp.Compare[string]),
			args:     args[string]{startKey: "c", endKey: "f"},
			want:     []string{},
			wantCond: false,
		},
		{
			name: "multiple elements #1",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{startKey: "a", endKey: "g"},
			want:     []string{"b", "c", "d", "e", "f"},
			wantCond: true,
		},
		{
			name: "multiple elements #2",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{startKey: "a", endKey: "d"},
			want:     []string{"b", "c", "d"},
			wantCond: true,
		},
		{
			name: "multiple elements #3",
			st: func() *SymbolTable[string, int] {
				st := NewSymbolTable[string, int](cmp.Compare[string])
				for i, v := range []string{"b", "c", "d", "e", "f"} {
					st.Add(v, i)
				}
				return st
			}(),
			args:     args[string]{startKey: "d", endKey: "z"},
			want:     []string{"d", "e", "f"},
			wantCond: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.st.Range(tt.args.startKey, tt.args.endKey)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantCond {
				t.Errorf("Range() got condition = %v, want %v", got1, tt.wantCond)
			}
		})
	}
}
