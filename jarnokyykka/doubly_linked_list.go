package main

import (
	"fmt"
)

type list struct {
	length int
	head *node
}

type node struct {
	next *node
	prev *node	
	value string
}

func main() {
	l := new(list)
	l.length = 0
	l.head = nil
	
	l.insert(0, "eka")
	l.insert(1, "toka")
	l.insert(2, "kolmas")
	l.insert(3, "neljas")
	l.insert(0, "EKA")
	l.insert(3, "KOLMAS")
	

	n := l.head

	fmt.Printf("\nPrinting list values: ")
	for {
		fmt.Printf("\n%s", n.value )
		if n.next == nil {
			break
		}
		n = n.next
	}

	fmt.Printf("\n\nRemoving item index 2\n\nPrinting list values: ")
	l.delete(2)
	n = l.head
	
	for {
		fmt.Printf("\n%s", n.value)
		if n.next == nil {
			break
		}
		n = n.next
	}
	

	fmt.Printf("\n\nPrinting in reverse order: ")
	n = l.head
	for n.next != nil {
		n = n.next
	}
		
	for {
		fmt.Printf("\n%s", n.value)
		if n.prev == nil {
			break
		}
		n = n.prev
	}
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	newNode := new(node)
	newNode.value = e
	current := l.head
	var last *node = nil

	if l.head == nil {
		l.head = newNode
		fmt.Printf("Added index: %d, value: %s\n", i, e)		
		l.length++
		return
	}

	if l.length <= i {
		last = current
		for current.next != nil {
			last = current
			current = current.next
		}
		current.next = newNode
		newNode.prev = current
		l.length++
		fmt.Printf("Added index: %d, value: %s\n", i, e)
		return
	}

	for j := 0; j != i; j++ {
		last = current
		current = current.next
	} 
	
	//link the new node
	if i != 0 {
		last.next = newNode
		newNode.prev = last
	} else {
		l.head = newNode
	}
	newNode.next = current
	current.prev = newNode	

	l.length++	
	fmt.Printf("Added index: %d, value: %s\n", i, e)	
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	//in list?
	if !(i < l.length && i >= 0) {
		return
	}

	current := l.head

	for j := 0; j != i; j++ {
		current = current.next
	}
	if current.prev != nil {
		current.prev.next = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	}
}
