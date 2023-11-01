package binary_search_tree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNode_Get(t *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		n    *Node[T]
		args args[T]
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "single node (found)",
			n:    NewNode[int](10),
			args: args[int]{value: 10},
			want: NewNode[int](10),
		},
		{
			name: "single node (not found)",
			n:    NewNode[int](10),
			args: args[int]{value: 20},
			want: nil,
		},
		{
			name: "one larger child (found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				return node
			}(),
			args: args[int]{value: 20},
			want: NewNode[int](20),
		},
		{
			name: "one larger child (not found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				return node
			}(),
			args: args[int]{value: 30},
			want: nil,
		},
		{
			name: "one smaller child (found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(5)
				return node
			}(),
			args: args[int]{value: 5},
			want: NewNode[int](5),
		},
		{
			name: "one smaller child (not found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(5)
				return node
			}(),
			args: args[int]{value: 3},
			want: nil,
		},
		{
			name: "two children (found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(5)
				node.Insert(20)
				return node
			}(),
			args: args[int]{value: 5},
			want: NewNode[int](5),
		},
		{
			name: "two children (not found)",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(5)
				node.Insert(20)
				return node
			}(),
			args: args[int]{value: 30},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Get(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		n    *Node[T]
		args args[T]
		want *Node[T]
	}
	tests := []testCase[int]{
		{
			name: "single node",
			n:    NewNode[int](10),
			args: args[int]{value: 10},
			want: NewNode[int](10),
		},
		{
			name: "one larger child",
			n:    NewNode[int](10),
			args: args[int]{value: 20},
			want: func() *Node[int] {
				node := NewNode[int](10)
				node.right = NewNode(20)
				return node
			}(),
		},
		{
			name: "one lower child",
			n:    NewNode[int](10),
			args: args[int]{value: 5},
			want: func() *Node[int] {
				node := NewNode[int](10)
				node.left = NewNode(5)
				return node
			}(),
		},
		{
			name: "two children",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				return node
			}(),
			args: args[int]{value: 5},
			want: func() *Node[int] {
				node := NewNode[int](10)
				node.left = NewNode(5)
				node.right = NewNode(20)
				return node
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.Insert(tt.args.value)
			if !reflect.DeepEqual(tt.n, tt.want) {
				t.Errorf("Insert() = %v, want %v", tt.n, tt.want)
			}

		})
	}
}

func TestNode_Max(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		n    *Node[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "single node",
			n:    NewNode[int](10),
			want: 10,
		},
		{
			name: "one larger child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.right = NewNode(20)
				return node
			}(),
			want: 20,
		},
		{
			name: "one lower child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.left = NewNode(5)
				return node
			}(),
			want: 10,
		},
		{
			name: "two children",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				node.Insert(5)
				return node
			}(),
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Max(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Min(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		n    *Node[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "single node",
			n:    NewNode[int](10),
			want: 10,
		},
		{
			name: "one larger child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.right = NewNode(20)
				return node
			}(),
			want: 10,
		},
		{
			name: "one lower child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.left = NewNode(5)
				return node
			}(),
			want: 5,
		},
		{
			name: "two children",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				node.Insert(5)
				return node
			}(),
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_TraverseInOrder(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		n    *Node[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "single node",
			n:    NewNode[int](10),
			want: []int{10},
		},
		{
			name: "one larger child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.right = NewNode(20)
				return node
			}(),
			want: []int{10, 20},
		},
		{
			name: "one lower child",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.left = NewNode(5)
				return node
			}(),
			want: []int{5, 10},
		},
		{
			name: "two children",
			n: func() *Node[int] {
				node := NewNode[int](10)
				node.Insert(20)
				node.Insert(5)
				return node
			}(),
			want: []int{5, 10, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.TraverseInOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TraverseInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
