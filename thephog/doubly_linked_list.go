package main

import "fmt"
import "strconv"

type node struct {
	prev, next *node
	data       string
}

type list struct {
	first, last *node
}

func (l *list) getNode(index int) *node {
	target := l.first
	for i := 0; i < index; i++ {

		if target == nil {
			break
		}

		target = target.next

	}

	return target
}

func (l *list) insert(i int, e string) {
	target := l.getNode(i)
	if target == nil {
		if l.first == nil && l.last == nil {
			l.first = &node{nil, nil, e}
			l.last = l.first
		} else {
			n := &node{l.last, nil, e}
			l.last.next = n
			l.last = n
		}
		return
	}

	n := &node{target.prev, target, e}
	if target.prev != nil {
		target.prev.next = n
	}
	target.prev = n

}

func (l *list) delete(i int) {
	target := l.getNode(i)
	if target == nil {
		return
	}

	if target.prev != nil {
		target.prev.next = target.next
	} else {
		l.first = target.next
	}

	if target.next != nil {
		target.next.prev = target.prev
	} else {
		l.last = target.prev
	}

}

func (l *list) String() string {
     	s := ""
	node := l.first
	for {
		if node == nil {
			break
		}

		s += node.data + "\n"
		node = node.next
	}
	return s
}

func main() {
	fmt.Printf("Creating list...\n")
	l := &list{nil, nil}

	fmt.Printf("Adding values...\n")
	for i := 0; i < 66; i++ {
		e := "SuperTest" + strconv.Itoa(i)
		l.insert(i, e)
	}

	fmt.Printf("List now contains:\n%v\n", l)

	fmt.Printf("Deleting values..\n")
	for i := 65; i > 33; i-- {
		l.delete(i)
	}

	fmt.Printf("List now contains:\n%v\n", l)
	
	fmt.Printf("Deleting value at position 5..\n")
	l.delete(5)
	
	fmt.Printf("List now contains:\n%v\n", l)

	fmt.Printf("Deleting values..\n")
	for i := 33; i >= 0; i-- {
		l.delete(i)
	}

	fmt.Printf("List now contains:\n%v\n", l)
}
