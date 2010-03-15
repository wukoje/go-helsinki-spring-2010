package main

import "fmt"

type list struct {
	sentinel *node
}

type node struct {
	data string
	prev *node
	next *node
}

func newList() *list {
	sentinel := new(node)
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &list{sentinel}
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	p := l.sentinel

	for i > 0 {
		p = p.next
		i--
	}

	newNode := &node{e, p, p.next}

	p.next = newNode
	newNode.next.prev = newNode
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	p := l.sentinel

	for i >= 0 {
		p = p.next
		i--
	}

	p.prev.next = p.next
	p.next.prev = p.prev
}

func (l *list) printList() {
	p := l.sentinel.next
	for p != l.sentinel {
		fmt.Print(p.data + " ")
		p = p.next
	}

	fmt.Println()
}

func main() {
	l := newList()
	l.insert(0, "a")
	l.insert(1, "c")
	l.insert(2, "d")
	l.insert(3, "e")
	l.insert(1, "b")
	l.insert(0, "0")
	l.insert(3, "-")

	l.printList() // 0ab-cde

	l.delete(3)
	l.printList() // 0abcde

	l.delete(0)
	l.printList() // abcde

	l.delete(2)
	l.printList() // abde

	l.delete(0)
	l.delete(0)
	l.delete(0)
	l.delete(0)
	l.printList() // <empty>
}
