package main

import (
	"fmt"
)

const (
	LAST_ELEMENT = 1<<16 // insert to the end of list
	FIRST_ELEMENT = 0 // insert to the beginning
)

type list struct {
	value string
	next *list // next node
	prev *list // previous node
	
	first *list // head of the list
	last *list // tail of the list
}

func (l *list) String() string {
	l = l.first
	s := fmt.Sprintf ("%s - ", l.value)

	l = l.next
	for i := 0; l.last != l; i++{
		s += fmt.Sprintf("%s - ", l.value)
		l = l.next
	}

	l = l.last
	s += fmt.Sprintf ("%s", l.value)

	return s
}


// insert adds the element e at index i in the list l
// first element is found from index=0
func (l *list) insert(i int, e string) {
	l = l.first // jump to the beginning
	for j := 0; j < i; j++ {
		l = l.next
		if l == l.last {
			l = l.prev
			break // reached the end of list
		}
	}

	next := l.next
	prev := l
	
	node := &list{e, next, prev, l.first, l.last}
	next.prev = node // append a new node
	prev.next = node

}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	l = l.first.next // jump to the fisrt element
	for j := 0; j < i; j++ {
		l = l.next
		if l == l.last {
			l = l.prev
			break // reached the end of list
		}
	}

	// drop selected node index from the list
	// and adjust pointers
	next := l.next
	prev := l.prev
	prev.next = next
	next.prev = prev

}

func createLinkedList() *list {
	// init sentinel nodes (first and last node)
	first := &list{"BEGIN", nil, nil, nil, nil} // start of the list
	last := &list{"END", nil, nil, nil, nil} // end of the list

	first.prev = first // init first node
	first.next = last
	first.first = first // pointer to the first node
	first.last = last // pointer to the last node

	last.prev = first // init end node
	last.next = last
	last.first = first
	last.last = last
	return first
}

func main() {
	fmt.Printf("Doubly linked list\n")
	l := createLinkedList()

	l.insert(0, "element1")
	fmt.Printf ("1 %v\n", l)
	l.delete(0)
	fmt.Printf ("1 %v\n", l)

	l.insert(1, "element2")
	fmt.Printf ("2 %v\n", l)
	l.insert(2, "element3")
	fmt.Printf ("3 %v\n", l)
	l.insert(3, "element4")
	fmt.Printf ("4 %v\n", l)
	l.insert(4, "element5")
	fmt.Printf ("5 %v\n", l)
	l.insert(2, "middle1")
	fmt.Printf ("6 %v\n", l)
	l.insert(4, "middle2")
	fmt.Printf ("7 %v\n", l)

	l.insert(LAST_ELEMENT, "last item 1")
	l.insert(LAST_ELEMENT, "last item 2")
	l.insert(FIRST_ELEMENT, "first item 1")
	l.insert(FIRST_ELEMENT, "first item 2")

	fmt.Printf ("%v\n", l)

	l.delete(FIRST_ELEMENT)
	fmt.Printf ("D1 %v\n", l)
	l.delete(2)
	fmt.Printf ("D2 %v\n", l)
	l.delete(LAST_ELEMENT)
	fmt.Printf ("D3 %v\n", l)

}
