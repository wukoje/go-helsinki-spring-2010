
package main

import "fmt"

type node struct {
	prev, next *node
	value string
}

type list struct {
	front, back *node
}

func (l *list) insert(index int, e string) {
	var n0 *node // previous node 
	n1 := l.front // next node

	for i := 0; i < index; i++ {
		n0 = n1
		n1 = n1.next
	}

	newNode := &node { n0, n1, e }
	if n0 != nil {
		n0.next = newNode
	} else {
		l.front = newNode
	}

	if n1 != nil {
		n1.prev = newNode
	} else {
		l.back = newNode
	}
}

func (l *list) delete(index int) {
	n := l.front // node at index

	for i := 0; i < index; i++ {
		n = n.next
	}

	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.front = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.back = n.prev
	}

}

func (l *list) String() (s string) {
	s += "[ "
	for n := l.front; n != nil; n = n.next {
		s += n.value + " "
	}
	s += "] "

	s += "[ "
	for n := l.back; n != nil; n = n.prev {
		s += n.value + " "
	}
	s += "]"
	return s
}

func main() {
	l1 := new(list)
	l1.insert(0, "0")
	l1.insert(1, "1")
	fmt.Println(l1)

	l2 := new(list)
	l2.insert(0, "1")
	l2.insert(0, "0")
	fmt.Println(l2)

	l3 := new(list)
	l3.insert(0, "0")
	l3.insert(1, "1")
	l3.insert(2, "2")
	fmt.Println(l3)

	l3.insert(0, "-1")
	l3.insert(3, "1.5")
	fmt.Println(l3)

	l3.delete(0)
	l3.delete(2)
	fmt.Println(l3)
}

