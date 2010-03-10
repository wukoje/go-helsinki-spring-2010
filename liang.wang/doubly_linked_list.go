package main

import "fmt"

func main() {
	head := new(list)
	head.insert(0, "hello")
	head.print()
	head.insert(1, "world")
	head.print()
	head.insert(0, "Say")
	head.print()
	head.insert(3, "Again")
	head.print()
	
	head.delete(0)
	head.print()
	head.delete(2)
	head.print()
}

type list struct {
	value string
	prev *list
	next *list
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	node := new(list)
	node.value = e
	p:= l
	for j:=0; l!=nil && j<=i; l,j=l.next,j+1 { p=l }

	node.prev, node.next = p, p.next
	p.next = node
	if p.next!=nil { p.next.prev = node }
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	p := l
	for j:=0; l!=nil && j<=i; l,j=l.next,j+1 { p=l }
	if l!=nil {
		p.next = l.next
		if p.next != nil {
			p.next.prev = p
		}
	}
}

// output the list
func (l *list) print() {
	fmt.Printf("[ ")
	for l=l.next;l!=nil;l=l.next {
		fmt.Printf("\"%v\" ", l.value)
	}
	fmt.Printf("]\n")
}
