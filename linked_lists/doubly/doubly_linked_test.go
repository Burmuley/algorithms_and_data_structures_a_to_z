package doubly

import (
	"reflect"
	"testing"
)

func TestNewLinkedList_InitStateEmptyList(t *testing.T) {
	list := NewLinkedList[int]()
	if list.Count() != 0 {
		t.Errorf("Initial state: wrong count: %v, want %v", list.Count(), 0)
	}
	if list.Head() != nil || list.Tail() != nil {
		t.Errorf("Initial state: head or tail values: %v %v, want %v %v", list.Head(), list.Tail(), nil, nil)
	}
}

func TestNewLinkedList_InitStateSingleElementList(t *testing.T) {
	list := NewLinkedList[int]().AddFirst(NewNode(1))
	if list.Count() != 1 {
		t.Errorf("Initial state: wrong count: %v, want %v", list.Count(), 0)
	}
	if list.Head() == nil || list.Head() == nil {
		t.Errorf("Initial state: head or tail values: %v %v, want %v %v", list.Head(), list.Tail(), nil, nil)
	}
	if list.Head() != list.Tail() {
		t.Errorf("Initial state: head should be equal to tail: %v %v, want %v %v", list.Head(), list.Tail(), nil, nil)
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	type args[T any] struct {
		node *Node[T]
	}
	type testCase[T any] struct {
		name string
		list *LinkedList[T]
		args args[T]
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{
			name: "Empty list",
			list: NewLinkedList[int](),
			args: struct{ node *Node[int] }{node: NewNode(10)},
			want: func() *LinkedList[int] {
				node := NewNode(10)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
		},
		{
			name: "Single element list",
			list: func() *LinkedList[int] {
				node := NewNode(5)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
			args: struct{ node *Node[int] }{node: NewNode(10)},
			want: func() *LinkedList[int] {
				headnode := NewNode(10)
				tailnode := NewNode(5)
				headnode.next = tailnode
				tailnode.prev = headnode
				list := NewLinkedList[int]()
				list.head = headnode
				list.tail = tailnode
				list.count = 2
				return list
			}(),
		},
		{
			name: "More than one element",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			args: struct{ node *Node[int] }{node: NewNode(13)},
			want: func() *LinkedList[int] {
				vals := []int{13, 1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.AddFirst(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFirst() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	type args[T any] struct {
		node *Node[T]
	}
	type testCase[T any] struct {
		name string
		list *LinkedList[T]
		args args[T]
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{
			name: "Empty list",
			list: NewLinkedList[int](),
			args: struct{ node *Node[int] }{node: NewNode(10)},
			want: func() *LinkedList[int] {
				node := NewNode(10)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
		},
		{
			name: "Single element list",
			list: func() *LinkedList[int] {
				node := NewNode(10)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
			args: struct{ node *Node[int] }{node: NewNode(5)},
			want: func() *LinkedList[int] {
				headnode := NewNode(10)
				tailnode := NewNode(5)
				headnode.next = tailnode
				tailnode.prev = headnode
				list := NewLinkedList[int]()
				list.head = headnode
				list.tail = tailnode
				list.count = 2
				return list
			}(),
		},
		{
			name: "More than one element",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			args: struct{ node *Node[int] }{node: NewNode(13)},
			want: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8, 13}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.AddLast(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddLast() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	type testCase[T any] struct {
		name string
		list *LinkedList[T]
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{
			name: "Empty list",
			list: NewLinkedList[int](),
			want: NewLinkedList[int](),
		},
		{
			name: "Single element list",
			list: func() *LinkedList[int] {
				node := NewNode(10)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
			want: NewLinkedList[int](),
		},
		{
			name: "More than one element",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			want: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.RemoveLast(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLast() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	type testCase[T any] struct {
		name string
		list *LinkedList[T]
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{
			name: "Empty list",
			list: NewLinkedList[int](),
			want: NewLinkedList[int](),
		},
		{
			name: "Single element list",
			list: func() *LinkedList[int] {
				node := NewNode(10)
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
			want: NewLinkedList[int](),
		},
		{
			name: "More than one element",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			want: func() *LinkedList[int] {
				vals := []int{1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.RemoveFirst(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", *got, *tt.want)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type testCase[T any] struct {
		name string
		list *LinkedList[T]
		node *Node[T]
		want *LinkedList[T]
	}

	nodeToDelete_SingleElem := NewNode(1)
	nodeToDelete_MultiElemTail := NewNode(33)
	nodeToDelete_MultiElemHead := NewNode(44)
	tests := []testCase[int]{
		{
			name: "Empty list",
			list: NewLinkedList[int](),
			want: NewLinkedList[int](),
			node: nil,
		},
		{
			name: "Single element list",
			list: func() *LinkedList[int] {
				node := nodeToDelete_SingleElem
				list := NewLinkedList[int]()
				list.head = node
				list.tail = node
				list.count = 1
				return list
			}(),
			want: NewLinkedList[int](),
			node: nodeToDelete_SingleElem,
		},
		{
			name: "More than one element (delete tail)",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}
				list.AddLast(nodeToDelete_MultiElemTail)

				return list
			}(),
			want: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			node: nodeToDelete_MultiElemTail,
		},
		{
			name: "More than one element (delete head)",
			list: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}
				list.AddLast(nodeToDelete_MultiElemHead)

				return list
			}(),
			want: func() *LinkedList[int] {
				vals := []int{1, 1, 2, 3, 5, 8}
				list := NewLinkedList[int]()
				for _, v := range vals {
					list.AddLast(NewNode(v))
				}

				return list
			}(),
			node: nodeToDelete_MultiElemHead,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.list.Remove(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", *got, *tt.want)
			}
		})
	}
}
