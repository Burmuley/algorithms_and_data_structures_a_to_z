package max_heap

import (
	"cmp"
	"reflect"
	"testing"
)

func TestHeap_Insert(t *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		h    *Heap[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "empty heap",
			h:    NewHeap[int](),
			args: args[int]{10},
			want: 10,
		},
		{
			name: "single element (larger)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			args: args[int]{10},
			want: 30,
		},
		{
			name: "single element (lower)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(5)
				return h
			}(),
			args: args[int]{10},
			want: 10,
		},
		{
			name: "multiple elements (lower)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(24)
				h.Insert(37)
				h.Insert(28)
				h.Insert(31)
				return h
			}(),
			args: args[int]{17},
			want: 37,
		},
		{
			name: "multiple elements (lower)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(24)
				h.Insert(17)
				h.Insert(28)
				h.Insert(31)
				return h
			}(),
			args: args[int]{37},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Insert(tt.args.value)
			p, _ := tt.h.Peek()
			if !(tt.want == p) {
				t.Errorf("Insert() head element = %v, wantErr %v", p, tt.want)
			}
		})
	}
}

func TestHeap_Peek(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name    string
		h       *Heap[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty heap",
			h:       NewHeap[int](),
			want:    0,
			wantErr: true,
		},
		{
			name: "single element",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			want: 30,
		},
		{
			name: "multiple elements #1",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 20; i > 0; i-- {
					h.Insert(i)
				}
				return h
			}(),
			want: 20,
		},
		{
			name: "multiple elements #2",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 0; i < 20; i++ {
					h.Insert(i)
				}
				return h
			}(),
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Remove(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name    string
		h       *Heap[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty heap",
			h:       NewHeap[int](),
			want:    0,
			wantErr: true,
		},
		{
			name: "single element",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			want: 30,
		},
		{
			name: "multiple elements #1",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 20; i > 0; i-- {
					h.Insert(i)
				}
				return h
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Remove()
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_RemoveIndex(t *testing.T) {
	type args struct {
		index int
	}
	type testCase[T cmp.Ordered] struct {
		name    string
		h       *Heap[T]
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty heap",
			h:       NewHeap[int](),
			args:    args{10},
			want:    0,
			wantErr: true,
		},
		{
			name: "single element (found)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			args: args{0},
			want: 30,
		},
		{
			name: "single element (not found)",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			args:    args{10},
			want:    0,
			wantErr: true,
		},
		{
			name: "multiple elements #1",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 20; i > 0; i-- {
					h.Insert(i)
				}
				return h
			}(),
			args: args{10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.RemoveIndex(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_Values(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		h    *Heap[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty heap",
			h:    NewHeap[int](),
			want: []int{},
		},
		{
			name: "single element",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			want: []int{30},
		},
		{
			name: "multiple elements #1",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 20; i > 0; i-- {
					h.Insert(i)
				}
				return h
			}(),
			want: func() []int {
				l := make([]int, 0, 19)
				for i := 20; i > 0; i-- {
					l = append(l, i)
				}
				return l
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_ValuesSorted(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		h    *Heap[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty heap",
			h:    NewHeap[int](),
			want: []int{},
		},
		{
			name: "single element",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				h.Insert(30)
				return h
			}(),
			want: []int{30},
		},
		{
			name: "multiple elements #1",
			h: func() *Heap[int] {
				h := NewHeap[int]()
				for i := 20; i > 0; i-- {
					h.Insert(i)
				}
				return h
			}(),
			want: func() []int {
				l := make([]int, 0, 19)
				for i := 1; i <= 20; i++ {
					l = append(l, i)
				}
				return l
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.ValuesSorted(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
