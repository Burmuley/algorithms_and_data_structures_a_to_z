package linked_list

type Node[T any] struct {
	next  *Node[T]
	value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (node *Node[T]) Value() T {
	return node.value
}

func (node *Node[T]) SetValue(value T) *Node[T] {
	node.value = value
	return node
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (node *Node[T]) SetNext(next *Node[T]) *Node[T] {
	node.next = next
	return node
}

type LinkedList[T any] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		head:  nil,
		tail:  nil,
		count: 0,
	}
}

func (list *LinkedList[T]) Head() *Node[T] {
	return list.head
}

func (list *LinkedList[T]) Tail() *Node[T] {
	return list.tail
}

func (list *LinkedList[T]) Count() int {
	return list.count
}

func (list *LinkedList[T]) AddFirst(node *Node[T]) *LinkedList[T] {
	cn := list.head

	list.head = node
	node.SetNext(cn)
	list.count++

	if list.count == 1 {
		list.tail = list.head
	}

	return list
}

func (list *LinkedList[T]) AddLast(node *Node[T]) *LinkedList[T] {
	defer func() {
		list.count++
	}()

	if list.count == 0 {
		list.tail = node
		list.head = node
		return list
	}

	list.tail.SetNext(node)
	list.tail = node

	return list
}

func (list *LinkedList[T]) AddFirstValue(value T) *LinkedList[T] {
	list.AddFirst(NewNode(value))
	return list
}

func (list *LinkedList[T]) AddLastValue(value T) *LinkedList[T] {
	list.AddLast(NewNode(value))
	return list
}

func (list *LinkedList[T]) RemoveFirst() *LinkedList[T] {
	if list.count == 0 {
		return list
	}

	list.head = list.head.Next()
	if list.count == 1 {
		list.tail = nil
	}
	list.count--
	return list
}

func (list *LinkedList[T]) RemoveLast() *LinkedList[T] {
	if list.count == 0 {
		return list
	}

	if list.count == 1 {
		list.head = nil
		list.tail = nil
		list.count--
		return list
	}

	current := list.head
	for current.Next() != list.tail {
		current = current.Next()
	}

	current.SetNext(nil)
	list.tail = current
	list.count--

	return list
}
