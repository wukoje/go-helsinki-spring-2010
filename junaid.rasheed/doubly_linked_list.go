package main

import "fmt"


type list struct {
	e     string
	left  *list
	right *list
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	node := new(list)
	node.e = e
	aux := l
	for j := 0; l != nil && j <= i; l, j = l.right, j+1 {
		aux = l
	}

	node.left, node.right = aux, aux.right
	aux.right = node
	if aux.right != nil {
		aux.right.left = node
	}
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	aux := l
	for j := 0; l != nil && j <= i; l, j = l.right, j+1 {
		aux = l
	}
	if l != nil {
		aux.right = l.right
		if aux.right != nil {
			aux.right.left = aux
		}
	}
}

func (l *list) print() {
	for l = l.right; l != nil; l = l.right {
		fmt.Printf("\"%v\" ", l.e)
	}
	fmt.Printf("\n")
}

func main() {

	list1 := new(list)

	list1.insert(0, "Helsinki")
	list1.insert(1, "Oulu")
	list1.insert(2, "Nokia")
	list1.insert(3, "Turku")
	list1.insert(4, "Tampere")
	list1.print()

	list1.delete(1)
	list1.print()

	list1.delete(2)
	list1.print()
}

