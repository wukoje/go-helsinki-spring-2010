// two-way sentinel linked list
package main

import (
	"fmt"
)

type node struct {
	e    string
	next *node
	prev *node
}

type list struct {
	len int
	s   node // sentinel node
}


// insert adds the element e at index i in the list l

func (l *list) insert(i int, e string) {
	if i < 0 || i > (*l).len {
		// this can't be good
		panic("ERROR: List insertion out of bounds")
	}

	newNode := &node{e, nil, nil}

	prev := &(*l).s
	for j := 0; j < i; j++ {
		prev = (*prev).next
	}

	var next *node
	if (*prev).next != nil {
		next = (*prev).next
		(*next).prev = newNode
	}
	(*prev).next = newNode

	(*newNode).prev = prev
	(*newNode).next = next

	(*l).len++
}


// delete removes the element at index i in the list l

func (l *list) delete(i int) {
	if i < 0 || i > (*l).len {
		// this can't be good
		panic("ERROR: List deletion out of bounds")
	}

	delNode := (*l).s.next
	for j := 0; j < i; j++ {
		delNode = (*delNode).next
	}
	prev := (*delNode).prev
	(*prev).next = (*delNode).next

	if (*delNode).next != nil {
		next := (*delNode).next
		(*next).prev = prev
	}
}

func (l *list) print() {
	n := &(*l).s
	fmt.Print("List: ")
	for i := 0; (*n).next != nil; i++ {
		n = (*n).next
		fmt.Printf("%v:'%v' ", i, (*n).e)
	}
	fmt.Println()
}

func main() {
	var l list
	fmt.Println("Do some inserts..")
	fmt.Println("Inserting '1st' to 0")
	l.insert(0, "1st")
	l.print()
	fmt.Println("Inserting '2nd' to 1")
	l.insert(1, "2nd")
	l.print()
	fmt.Println("Inserting '3rd' to 1")
	l.insert(1, "3rd")
	l.print()
	fmt.Println("Inserting '4th' to 0")
	l.insert(0, "4th")
	l.print()
	fmt.Println("Inserting '5th' to 4")
	l.insert(4, "5th")
	l.print()
	fmt.Println("Inserting '6th' to 2")
	l.insert(2, "6th")
	l.print()

	// will not work
	//l.insert(-1,"7th")
	//l.insert(7,"7th")
	//l.delete(-1)
	//l.delete(7)

	fmt.Println("Delete all..")
	fmt.Println("Deleting 2")
	l.delete(2)
	l.print()
	fmt.Println("Deleting 4")
	l.delete(4)
	l.print()
	fmt.Println("Deleting 0")
	l.delete(0)
	l.print()
	fmt.Println("Deleting 1")
	l.delete(1)
	l.print()
	fmt.Println("Deleting 0")
	l.delete(0)
	l.print()
	fmt.Println("Deleting 0")
	l.delete(0)
	l.print()
}
