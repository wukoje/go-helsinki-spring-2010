package main

import "fmt"

type list struct {
	prev *list
	next *list
	data string
}


// insert adds the element e at index i in the list 
func (l *list) insert(i int, e string) {
	node := new(list)
	node.data = e

	var prev *list = l
	for idx := 0; idx <= i && l != nil; l,idx = l.next, idx+1 {
		prev = l
	}
	l = prev

	node.prev = l
	node.next = l.next
	if l.next != nil {
		l.next.prev = node
	}
	l.next = node
}


// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	for idx := 0; idx <= i && l != nil; idx = idx+1 {
		l = l.next
	}
	if l != nil {
		l.prev.next = l.next
		if l.next != nil {
			l.next.prev = l.prev
		}
	}
}

func main() {
	var l *list = new(list) // use a dummy node 
	l.insert(0, "d")
	l.insert(0, "c")
	l.insert(0, "a")
	l.insert(1, "b")
	l.insert(2, "ZZZ")
	l.insert(555, "e") // insert at end
	l.insert(0, "YYY")
	
	// restore alphabet order
	l.delete(0) // YYY
	l.delete(2) // ZZZ
	l.delete(56) // do nothing (especially don't reference nil)

	fmt.Println("Forward:")
	prev := l
	for node := l; node != nil; node = node.next {
		prev = node
		fmt.Printf("%s ", node.data);
	}
	fmt.Printf("\nBackward: ")
	for node := prev; node.prev != nil; node = node.prev {
		fmt.Printf("%s ", node.data);
	}
}
 
