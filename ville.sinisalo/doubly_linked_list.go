package main

import (
	"fmt"
	"rand"
	"time"
)

type node struct {
	v string
	next *node
	prev *node
}

type list struct {
	root *node
	tail *node
	size int
}

// get returns the value of the element at index i in list l
func (l *list) get(i int) string {
	if i < 0 || i >= l.size {
		panic("list index out of bounds")
	}
	n := l.root
	for ; i!=0; i-- {
		n = n.next
	}
	return n.v
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	if i < 0 || i > l.size {
		panic("list index out of bounds")
	}
	n := new(node)
	n.v = e
	if l.size == 0 {
		l.root = n
		l.tail = n
		l.size = 1
		return
	}

	mid := l.size/2
	if mid >= i {
		t := l.root
		for ; i!=0; i-- {
			t = t.next
		}
		if t.prev != nil {
			t.prev.next = n
		} else {
			l.root = n
		}
		n.prev = t.prev
		t.prev = n
		n.next = t
	} else {
		t := l.tail
		for i=l.size-i; i!=0; i-- {
			t = t.prev
		}
		if t.next != nil {
			t.next.prev = n
		} else {
			l.tail = n
		}
		n.next = t.next
		t.next = n
		n.prev = t
	}
	l.size++
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	if i < 0 || i >= l.size {
		panic("list index out of bounds")
	}

	var n *node
	mid := l.size/2
	if mid >= i {
		n = l.root
		for ; i!=0; i-- {
			n = n.next
		}
	} else {
		n = l.tail
		for i=l.size-i-1; i!=0; i-- {
			n = n.prev
		}
	}
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.root = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
	l.size--	
}

func main() {
	rand.Seed(time.Nanoseconds())
	
	l := new(list)
	
	for i:=0; i<5; i++ {
		l.insert(i, fmt.Sprintf("%d", i))
	}
	for i:=0; i<5; i++ {
		fmt.Printf("%d: %s\n", i, l.get(i))
	}
	fmt.Printf("..insert 0: 0..\n")
	l.insert(0, "0")
	for i:=0; i<6; i++ {
		fmt.Printf("%d: %s\n", i, l.get(i))
	}
	fmt.Printf("..insert 0: proo..\n")
	l.insert(0, "proo")
	for i:=0; i<7; i++ {
		fmt.Printf("%d: %s\n", i, l.get(i))
	}
	fmt.Printf("..insert 7: last..\n")
	l.insert(7, "last")
	for i:=0; i<8; i++ {
		fmt.Printf("%d: %s\n", i, l.get(i))
	}
	fmt.Printf("..insert 4: middle..\n")
	l.insert(4, "middle")
	for i:=0; i<9; i++ {
		fmt.Printf("%d: %s\n", i, l.get(i))
	}

	fail := rand.Intn(4)
	if fail == 0 {
		fmt.Printf("..insert 10: too much..\n")
		l.insert(10, "too much")
	} else if fail == 1 {
		fmt.Printf("..get 9..\n")
		fmt.Printf("%d: %s\n", 9, l.get(9))
	} else if fail == 2 {
		fmt.Printf("..delete 0, 4 & 8..\n")
		l.delete(8)
		l.delete(4)
		l.delete(0)
		for i:=0; i<6; i++ {
			fmt.Printf("%d: %s\n", i, l.get(i))
		}
	} else {
		fmt.Printf("..delete 9..\n")
		l.delete(9)
	}
}
