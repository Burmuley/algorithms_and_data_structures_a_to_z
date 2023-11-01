package binary_search_tree

import (
	"cmp"
	"reflect"
	"testing"
)

func TestBinarySearchTree_Get(t1 *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name    string
		t       *BinarySearchTree[T]
		args    args[T]
		want    *Node[T]
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       NewBST[int](),
			args:    args[int]{value: 10},
			want:    nil,
			wantErr: true,
		},
		{
			name: "single element (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				return tree
			}(),
			args:    args[int]{value: 10},
			want:    NewNode(10),
			wantErr: false,
		},
		{
			name: "single element (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				return tree
			}(),
			args:    args[int]{value: 20},
			want:    nil,
			wantErr: true,
		},
		{
			name: "one larger element (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(20)
				return tree
			}(),
			args:    args[int]{value: 20},
			want:    NewNode(20),
			wantErr: false,
		},
		{
			name: "one larger element (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(20)
				return tree
			}(),
			args:    args[int]{value: 30},
			want:    nil,
			wantErr: true,
		},
		{
			name: "one smaller element (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			}(),
			args:    args[int]{value: 5},
			want:    NewNode(5),
			wantErr: false,
		},
		{
			name: "one smaller element (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			}(),
			args:    args[int]{value: 30},
			want:    nil,
			wantErr: true,
		},
		{
			name: "multiple elements (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(24)
				tree.Insert(28)
				tree.Insert(31)
				tree.Insert(29)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(20)
				return tree
			}(),
			args: args[int]{value: 15},
			want: func() *Node[int] {
				n := NewNode(15)
				n.left = NewNode(12)
				n.right = NewNode(20)
				return n
			}(),
			wantErr: false,
		},
		{
			name: "multiple elements (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(24)
				tree.Insert(15)
				tree.Insert(28)
				tree.Insert(31)
				tree.Insert(29)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(20)
				return tree
			}(),
			args:    args[int]{value: 30},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.Get(tt.args.value)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Insert(t1 *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		t    *BinarySearchTree[T]
		args args[T]
		want *BinarySearchTree[T]
	}
	tests := []testCase[int]{
		{
			name: "single node",
			t:    NewBST[int](),
			args: args[int]{value: 10},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](10)
				return tree
			}(),
		},
		{
			name: "one larger child",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				return tree
			}(),
			args: args[int]{value: 40},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				tree.root.right = NewNode[int](40)
				return tree
			}(),
		},
		{
			name: "one lower child",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				return tree
			}(),
			args: args[int]{value: 20},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				tree.root.left = NewNode[int](20)
				return tree
			}(),
		},
		{
			name: "two children",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(40)
				return tree
			}(),
			args: args[int]{value: 20},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				tree.root.left = NewNode[int](20)
				tree.root.right = NewNode[int](40)
				return tree
			}(),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Max(t1 *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name    string
		t       *BinarySearchTree[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       NewBST[int](),
			want:    0,
			wantErr: true,
		},
		{
			name: "single element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				return tree
			}(),
			want:    10,
			wantErr: false,
		},
		{
			name: "one larger element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(20)
				return tree
			}(),
			want:    20,
			wantErr: false,
		},
		{
			name: "one smaller element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			}(),
			want:    10,
			wantErr: false,
		},
		{
			name: "multiple elements",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(24)
				tree.Insert(28)
				tree.Insert(31)
				tree.Insert(29)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(20)
				return tree
			}(),
			want:    37,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.Max()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Max() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Min(t1 *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name    string
		t       *BinarySearchTree[T]
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       NewBST[int](),
			want:    0,
			wantErr: true,
		},
		{
			name: "single element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				return tree
			}(),
			want:    10,
			wantErr: false,
		},
		{
			name: "one larger element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(20)
				return tree
			}(),
			want:    10,
			wantErr: false,
		},
		{
			name: "one smaller element",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			}(),
			want:    5,
			wantErr: false,
		},
		{
			name: "multiple elements",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(24)
				tree.Insert(28)
				tree.Insert(31)
				tree.Insert(29)
				tree.Insert(15)
				tree.Insert(12)
				tree.Insert(20)
				return tree
			}(),
			want:    12,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.Min()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Min() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Min() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Remove(t1 *testing.T) {
	type args[T cmp.Ordered] struct {
		value T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		t    *BinarySearchTree[T]
		args args[T]
		want *BinarySearchTree[T]
	}
	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    NewBST[int](),
			args: args[int]{value: 10},
			want: NewBST[int](),
		},
		{
			name: "single node (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](10)
				return tree
			}(),
			args: args[int]{value: 10},
			want: NewBST[int](),
		},
		{
			name: "single node (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](10)
				return tree
			}(),
			args: args[int]{value: 20},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](10)
				return tree
			}(),
		},
		{
			name: "one larger child (found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(40)
				return tree
			}(),
			args: args[int]{value: 40},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				return tree
			}(),
		},
		{
			name: "one larger child (not found)",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(40)
				return tree
			}(),
			args: args[int]{value: 50},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				tree.root.right = NewNode(40)
				return tree
			}(),
		},
		{
			name: "one lower child",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(20)
				return tree
			}(),
			args: args[int]{value: 20},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				return tree
			}(),
		},
		{
			name: "two children",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(37)
				tree.Insert(40)
				tree.Insert(20)
				return tree
			}(),
			args: args[int]{value: 20},
			want: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.root = NewNode[int](37)
				tree.root.right = NewNode[int](40)
				return tree
			}(),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Remove(tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_TraverseInOrder(t1 *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		t    *BinarySearchTree[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "emtpy tree",
			t:    NewBST[int](),
			want: []int{},
		},
		{
			name: "single node",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				return tree
			}(),
			want: []int{10},
		},
		{
			name: "one larger child",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(20)
				return tree
			}(),
			want: []int{10, 20},
		},
		{
			name: "one lower child",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				return tree
			}(),
			want: []int{5, 10},
		},
		{
			name: "two children",
			t: func() *BinarySearchTree[int] {
				tree := NewBST[int]()
				tree.Insert(10)
				tree.Insert(5)
				tree.Insert(20)
				return tree
			}(),
			want: []int{5, 10, 20},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.TraverseInOrder(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("TraverseInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
