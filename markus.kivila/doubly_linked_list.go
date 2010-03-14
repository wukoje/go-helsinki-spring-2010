/*
1 doubly_linked_list.go

Implement a doubly linked list that supports insert and delete.

type list struct {
// your code here
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string)

// delete removes the element at index i in the list l
func (l *list) delete(i int)
*/
package main

import (
	"fmt"
)

type list struct {
	next *list
	prev *list
	e string
}

func newList(next, prev *list, e string) *list {
	return &list{next, prev, e}
}

// insert adds the element e at index i in the list l
// indexing starts from 0
func (l *list) insert(i int, e string) {
	if i == 0 {
		n := newList(l, nil, e)
		l.prev = n
		return
	} else if i < 0 {
		fmt.Printf("ERR: Negative index\n")
		return
	}

	var prev *list = nil
	if prev = l.stepTo(i-1); prev == nil {
		return
	}

	var next *list = nil
	var n *list = nil
	if next = prev.next; next == nil {
		n = newList(nil, prev, e)
	} else {
		n = newList(next, prev, e)
		next.prev = n
	}
	prev.next = n
}

// walks the list to the given index and returns that node
// returns nil if list is shorter than i
func (l *list) stepTo(i int) *list {
	for j := 0; j < i; j++ {
		if l = l.next; l == nil {
			fmt.Printf("ERR: Index %d out of bounds (%d)\n", i, j)
			return nil
		}
	}
	return l
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	if i == 0 {
		if n := l.next; n != nil {
			n.prev = nil
		}
		l.next = nil
		return
	} else if i < 0 {
		fmt.Printf("ERR: Negative index\n")
		return
	}

	var node *list = nil
	if node = l.stepTo(i); node == nil {
		return
	}

	p := node.prev
	n := node.next

	if p != nil && n != nil {
		p.next = n
		n.prev = p
	} else if p != nil {
		p.next = nil
	} else if n != nil {
		n.prev = nil
	}
}

func (l *list) print() {
	for l != nil {
		fmt.Printf("%s", l.e)
		l = l.next
	}
	fmt.Printf("\n")
}

func main() {
	s := []string{ "E", "P", "I", "C", " ", "W", "I", "N" }
	l := newList(nil, nil, s[0])
	for i := 1; i < len(s); i++ {
		l.insert(i, s[i])
	}
	fmt.Printf("Inserted slice %v to list\n", s)
	fmt.Printf("List: ");
	l.print()

	last := len(s) - 1
	fmt.Printf("Removing node %d (last)\n", last)
	l.delete(last)
	l.print()

	fmt.Printf("Removing node 5\n")
	l.delete(5)
	l.print()

	// ok, removing first node is a bit annoying
	// should've wrapped the head in separate struct
	fmt.Printf("Removing node 0\n")
	newHead := l.next
	l.delete(0)
	newHead.print()
	if newHead.prev != nil {
		fmt.Printf("Failed to remove the first node\n")
	}

	fmt.Printf("Inserting %s to %d\n", s[1], 4);
	newHead.insert(4, s[1])
	newHead.print()

	fmt.Printf("Inserting %s to %d\n", s[3], 6);
	newHead.insert(6, s[3])
	newHead.print()

	fmt.Printf("Inserting %s to %d (Expecting FAIL)\n", s[3], 666);
	newHead.insert(666, s[3])
	newHead.print()

	fmt.Printf("Inserting %s to %d\n", s[0], 0);
	newHead.insert(0, s[0])
	newHead = newHead.prev
	newHead.print()
}
