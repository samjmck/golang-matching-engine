package main

type LinkedOrderListNode struct {
	order Order
	next *LinkedOrderListNode
}

type OrderedLinkedOrderList struct {
	len int
	head *LinkedOrderListNode
}

func NewOrderedLinkedOrderList() *OrderedLinkedOrderList {
	return &OrderedLinkedOrderList{
		len: 0,
		head: nil,
	}
}

func (l *OrderedLinkedOrderList) Slice() []Order {
	slice := make([]Order, l.Len())
	node := l.head
	for i := 0; node != nil; i++ {
		slice[i] = node.order
		node = node.next
	}
	return slice
}

func (l *OrderedLinkedOrderList) Len() int {
	return l.len
}

func (l *OrderedLinkedOrderList) GetNode(index int) *LinkedOrderListNode {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (l *OrderedLinkedOrderList) Get(index int) Order {
	return l.GetNode(index).order
}

func (l *OrderedLinkedOrderList) RemoveFirst() {
	l.head = l.head.next
	l.len = l.len - 1
}

func (l *OrderedLinkedOrderList) RemoveLast() {
	if l.Len() == 1 {
		l.head = nil
	} else {
		secondToLastNode := l.GetNode(l.Len() - 2)
		secondToLastNode.next = nil
	}
	l.len = l.len - 1
}

func (l *OrderedLinkedOrderList) Insert(o Order) {
	if l.Len() == 0 {
		l.head = &LinkedOrderListNode{
			order: o,
			next: nil,
		}
	} else {
		node := l.head
		for node.next != nil && node.next.order.Price < o.Price {
			node = node.next
		}
		insertedNode := &LinkedOrderListNode{
			order: o,
			next: node.next,
		}
		node.next = insertedNode
	}

	l.len = l.len + 1
}
