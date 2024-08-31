package main

import "slices"

type SliceSortedOrderList struct {
	slice []Order
}

func NewSliceSortedOrderList() *SliceSortedOrderList {
	return &SliceSortedOrderList{ slice: []Order{} }
}

func (l *SliceSortedOrderList) Slice() []Order {
	return l.slice
}

func (l *SliceSortedOrderList) Len() int {
	return len(l.slice)
}

func (l *SliceSortedOrderList) Get(index int) Order {
	return l.slice[index]
}

func (l *SliceSortedOrderList) RemoveFirst() {
	l.slice = l.slice[1:]
}

func (l *SliceSortedOrderList) RemoveLast() {
	l.slice = l.slice[:len(l.slice) - 1]
}

func (l *SliceSortedOrderList) Insert(o Order) {
	i := 0
	for i < len(l.slice) && o.Price > l.slice[i].Price {
		i++
	}
	l.slice = slices.Insert(l.slice, i, o)
}
