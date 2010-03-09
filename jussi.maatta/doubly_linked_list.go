package main

import "fmt"

type list struct {
	head *listNode
	size int
}

type listNode struct {
	prev, next *listNode
	val        string
}

func (l *list) String() string {
	s := "list[ "
	p := l.head
	for p != nil {
		s += p.val + " "
		p = p.next
	}
	s += "]"
	return s
}

func (l *list) insert(i int, e string) {
	switch { // making friendly assumptions about unfriendly arguments
	case i < 0:
		i = 0
	case i > l.size:
		i = l.size
	}

	newNode := &listNode{nil, nil, e}

	if l.head == nil { // empty list
		l.head = newNode
	} else if i == 0 { // it's a new head
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else { // nonempty list and i>0
		p := l.head
		for i > 1 {
			p = p.next
			i--
		}
		newNode.prev = p
		newNode.next = p.next
		p.next = newNode
		if newNode.next != nil {
			newNode.next.prev = newNode
		}
	}

	l.size++
}

func (l *list) delete(i int) {
	if i < 0 || (i+1) > l.size {
		return // giving the silent treatment
	}

	p := l.head
	for i > 0 {
		p = p.next
		i--
	}

	if p.prev != nil {
		p.prev.next = p.next
	}
	if p.next != nil {
		p.next.prev = p.prev
	}
	if l.head == p {
		l.head = p.next
	}

	l.size--
}

func main() {
	fmt.Println("Creating the initial list")
	a := &list{}
	a.insert(0, "3rd")
	a.insert(0, "1st")
	a.insert(1, "2nd")
	a.insert(3, "5th")
	a.insert(3, "4th")
	fmt.Printf("%v\n\n", a)

	for _, v := range []int{1, 2, 1, 0, 0} {
		fmt.Printf("delete(%v)\n", v)
		a.delete(v)
		fmt.Printf("%v\n", a)
	}
}
