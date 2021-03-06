package main

import "github.com/cheekybits/genny/generic"

// NOTE: this is how easy it is to define a generic type
type Something generic.Type

// SomethingQueue is a queue of Somethings.
type SomethingQueue struct {
	items []Something
}

func NewSomethingQueue() *SomethingQueue {
	return &SomethingQueue{items: make([]Something, 0)}
}
func (q *SomethingQueue) Push(item Something) {
	q.items = append(q.items, item)
}
func (q *SomethingQueue) Pop() Something {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
