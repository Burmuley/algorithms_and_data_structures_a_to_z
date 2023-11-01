package doubly

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
	} else {
		list.Head().Next().SetPrev(list.Head())
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

	node.SetPrev(list.tail)
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

	if list.head != nil {
		list.head.SetPrev(nil)
	}

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

	list.tail.Prev().SetNext(nil)
	list.tail = list.tail.Prev()
	list.count--

	return list
}

func (list *LinkedList[T]) Remove(node *Node[T]) *LinkedList[T] {
	if node == nil || list.count == 0 {
		return list
	}

	p, n := node.Prev(), node.Next()
	if p != nil {
		p.SetNext(n)
	}

	if n != nil {
		n.SetPrev(p)
	}

	if list.head == node {
		list.head = n
	}

	if list.tail == node {
		list.tail = p
	}

	node.SetNext(nil).SetPrev(nil)
	list.count--
	return list
}
