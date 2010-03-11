/*
 * doubly_linked_list.go: A simple doubly linked linked list
 * implementation supporting insertion and deletion. To retain the
 * method signature specified in the course assignment the methods
 * take strings, but the actual list can store arbitrary data.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"fmt"
)

type list struct {
	head, tail *element
	length     int
}

type element struct {
	data       interface{}
	next, prev *element
}

// Delete removes the list element at index i. Negative indices
// count back from the end of the list (-1 is the last element).
// Invalid indices are ignored.
func (l *list) delete(i int) {
	var node *element

	if i < 0 {
		i += l.length // Negative index, count back from the end
	}
	if i < 0 || i >= l.length {
		return // Ignore invalid index
	}

	if i > l.length/2 {
		// Closer to the end
		node = l.tail
		for atIndex := l.length - 1; atIndex > i; atIndex-- {
			node = node.prev
		}
	} else {
		// Closer to the beginning
		node = l.head
		for atIndex := 0; atIndex < i; atIndex++ {
			node = node.next
		}
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}

	l.length--
}

// Insert adds the string e into the list so that it will have
// index i. Negative indices count back from the end of the list
// (-1 is the last element). Indices outside the list cause the
// string to be added at the closest end.
func (l *list) insert(i int, e string) {
	node := new(element)
	node.data = e

	if i < 0 {
		i += l.length + 1 // Negative index, count back from the end
	}

	switch {
	case i <= 0:
		node.next = l.head
	case i >= l.length:
		node.prev = l.tail
	case i > l.length/2:
		// Closer to the end
		atElement := l.tail
		for atIndex := l.length - 1; atIndex > i; atIndex-- {
			atElement = atElement.prev
		}
		node.next, node.prev = atElement, atElement.prev
	default:
		// Closer to the beginning
		atElement := l.head
		for atIndex := 0; atIndex < i; atIndex++ {
			atElement = atElement.next
		}
		node.next, node.prev = atElement, atElement.prev
	}

	if node.next != nil {
		node.next.prev = node
	} else {
		l.tail = node
	}
	if node.prev != nil {
		node.prev.next = node
	} else {
		l.head = node
	}

	l.length++
}

// Reverse inverts the order of the list.
func (l *list) reverse() {
	for node := l.tail; node != nil; node = node.next {
		node.next, node.prev = node.prev, node.next
	}
	l.tail, l.head = l.head, l.tail
}

func (l *list) String() string {
	s := ""
	for node := l.head; node != nil; node = node.next {
		s += fmt.Sprintf("%v -> ", node.data)
	}
	return s + "nil"
}

// Test the list implementation by inserting and deleting strings
// in different ways (selected to cover the different cases inside
// the relevant methods).
func main() {
	l := new(list)
	fmt.Println(l)
	l.insert(0, "qux") // Inserting to an empty list
	fmt.Println(l)
	l.insert(0, "foo") // Inserting to the head
	fmt.Println(l)
	l.insert(1, "bar") // Inserting to the middle (first half)
	fmt.Println(l)
	l.insert(-2, "baz") // Inserting to the middle (second half)
	fmt.Println(l)
	l.insert(-1, "quux") // Inserting to the tail
	fmt.Println(l)
	l.reverse()
	fmt.Println(l)
	l.delete(-1) // Deleting the tail
	fmt.Println(l)
	l.delete(-2) // Deleting from the middle (second half)
	fmt.Println(l)
	l.delete(1) // Deleting from the middle (first half)
	fmt.Println(l)
	l.delete(0) // Deleting the head
	fmt.Println(l)
	l.delete(100) // Does nothing
	l.insert(100, "foo") // Inserting to the tail
	l.reverse()
	fmt.Println(l)

	// Clear the list
	for l.length > 0 {
		l.delete(0)
	}
	fmt.Println(l)
}
