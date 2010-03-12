package main

import (
	"fmt"
)

type list struct {
	first *node
	elements int
}

type node struct {
	key string
	prev, next *node
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	if i > l.elements || i < 0 {
		return
	}

	l.elements++
	n := new(node)
	n.key = e

	if l.first == nil {
		l.first = n
		return
	}

	pos := l.first
	for j := 1; j < i; j++ {
		pos = pos.next
	}

	n.prev = pos
	n.next = pos.next
	pos.next = n
	if n.next != nil {
		n.next.prev = n
	}

}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	if i >= l.elements || i < 0 {
		return
	}
	l.elements--

	pos := l.first
	for j := 1; j < i; j++ {
		pos = pos.next
	}

	if pos == l.first {
		l.first = pos.next
		return
	}

	pos.prev.next = pos.next
	pos.next.prev = pos.prev
}

func (l *list) print() {
	for i, pos := 0, l.first; pos != nil; i, pos = i+1, pos.next {
		fmt.Println(i, "-", pos.key)
	}
}

func main() {
	l := new(list)
	l.insert(0, "Alpha")
	l.insert(1, "Beta")
	l.insert(2, "Gaga")
	l.insert(1, "Foo")
	l.insert(2, "Bar")
	fmt.Println("Contents of the list:")
	l.print()

	fmt.Println("After removing 0:")
	l.delete(0)
	l.print()
	fmt.Println("After removing 3:")
	l.delete(3)
	l.print()
}
