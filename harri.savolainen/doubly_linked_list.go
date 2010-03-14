package main

import fmt "fmt"

func main() {

	// Insert few
	l := new(list)
	l.insert(0, "foo")
	l.insert(1, "foo1")
	l.print()

	// Insert few more
	l.insert(2, "foo2")
	l.insert(3, "foo3")

	// Insert in between and in the beginnig
	l.insert(0, "newBeginingNew0")
	l.insert(2, "newInTheMiddleNew2")
	l.print()

	// Delete from the end
	l.delete(5)
	l.print()

	// Delete from the begining
	l.delete(0)

	// Delete from the middle few times
	l.delete(2)
	l.delete(2)
	l.print()
}

type list struct {
	begin *node
	nodes int
}

type node struct {
	value       string
	left, right *node
}

// insert adds the element e at index i in the list l

func (l *list) init() {
	l.nodes = 0
	l.begin = &node{"", nil, nil}
}

func (l *list) print() {

	if l.nodes == 0 {
		fmt.Println("empty")
		return
	}

	n := l.begin

	for j := 0; j < l.nodes; j++ {
		if n.left == nil {
			fmt.Printf("nil ")
		}
		fmt.Printf("- (Node %d) %s -", j, n.value)
		if n.right == nil {
			fmt.Printf(" nil\n")
		}
		n = n.right
	}

}

func (l *list) insert(i int, e string) {

	// Initialize list on first time
	if l.begin == nil {
		l.init()
	}

	// List overflow
	if i > l.nodes+1 {
		panic("invalid index")
	}

	// Adding first node
	if l.nodes == 0 {
		l.nodes = 1
		n := l.begin
		n.value = e
		return
	}

	// Adding to end of the list
	if i == l.nodes {
		n := l.begin

		for j := 0; j < l.nodes-1; j++ {
			n = n.right
		}

		nn := &node{e, n, nil}
		n.right = nn
		l.nodes++
		return
	}

	// Adding node to begin of the list
	if i == 0 {
		n := l.begin
		l.begin = &node{e, nil, n}
		n.left = l.begin
		l.nodes++
		return
	}

	// Adding node in somewhere in between
	b := l.begin

	for j := 0; j < i; j++ {
		b = b.right
	}

	a := b.left
	c := &node{e, a, b}
	a.right, b.left = c, c
	l.nodes++

}


// delete removes the element at index i in the list l

func (l *list) delete(i int) {

	if l.begin == nil {
		panic("list empty")
	}

	// List over/underflow
	if i > l.nodes || i < 0 {
		panic("not exists")
	}

	// Removing the last node
	if l.nodes == 1 && i == 0 {
		l.begin = nil
		l.nodes = 0
		return
	}

	// Removing at the end of the list
	if i == l.nodes-1 {
		n := l.begin

		for j := 0; j < l.nodes-1; j++ {
			n = n.right
		}

		n.left.right = nil
		n = nil
		l.nodes--
		return
	}

	// Removing the first node
	if i == 0 {
		n := l.begin.right
		l.begin = n
		l.begin.left = nil
		l.nodes--
		return
	}


	// Removing in somewhere between
	c := l.begin

	for j := 0; j < i; j++ {
		c = c.right
	}

	c.left.right, c.right.left = c.right, c.left
	l.nodes--
}


