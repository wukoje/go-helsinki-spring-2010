package main

import (
	"fmt"
)

type list struct {
	length int   // length of list
	head   *node // first node in list
	tail   *node // last node in list
}

type node struct {
	next *node  // next node
	prev *node  // previous node
	elem string // payload
}

func main() {
	l := &list{0, nil, nil}

	// Populate list...
	fmt.Println("Empty list:\t\t", l)
	l.insert(0, "bar")
	fmt.Println("Inserted one node:\t", l)
	l.insert(0, "foo")
	fmt.Println("Replaced first node:\t", l)
	l.insert(2, "baz")
	fmt.Println("Added node at end:\t", l)
	fmt.Println("Traverse in reverse:\t", l.reverseString())
	l.insert(2, "frob")
	fmt.Println("Add node in middle:\t", l)
	fmt.Println("Traverse in reverse:\t", l.reverseString())

	// ...and empty it
	l.delete(2)
	fmt.Println("Removed central node:\t", l)
	l.delete(2)
	fmt.Println("Removed node from end:\t", l)
	l.delete(0)
	fmt.Println("Removed first node:\t", l)
	l.delete(0)
	fmt.Println("Removed final node:\t", l)
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	if l.length >= i {
		l.length++
		if i == 0 {
			l.head = &node{l.head, nil, e}
			if l.head.next != nil {
				l.head.next.prev = l.head
			}
			if l.length == 1 {
				l.tail = l.head
			}
		} else {
			current := l.head
			for j := 0; j < i-1; j++ {
				current = current.next
			}
			following := current.next
			current.next = &node{following, current, e}
			if following != nil {
				following.prev = current.next
			} else {
				l.tail = current.next
			}
		}
	}
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	if i < l.length {
		if i == 0 {
			l.head = l.head.next
			if l.head != nil {
				l.head.prev = nil
			} else {
				l.tail = nil
			}
		} else {
			current := l.head
			for j := 0; j < i-1; j++ {
				current = current.next
			}
			removed := current.next
			current.next = removed.next
			if removed.next != nil {
				removed.next.prev = current
			} else {
				l.tail = current
			}
		}
		l.length--
	}
}

// return string representation of l
func (l *list) String() string {
	s := "[ "
	for i, n := 0, l.head; i < l.length; i, n = i+1, n.next {
		s += n.elem
		if i < l.length-1 {
			s += ", "
		}
	}
	return s + " ]"
}

// return string representation of reversed l
func (l *list) reverseString() string {
	s := "[ "
	for i, n := l.length-1, l.tail; i >= 0; i, n = i-1, n.prev {
		s += n.elem
		if i > 0 {
			s += ", "
		}
	}
	return s + " ]"
}
