package main

import "fmt"

type list struct {
	head *node
}

type node struct {
	str  string
	next *node
	prev *node
}

func (l *list) insert(idx int, e string) {
	np := l.head
	nn := new(node)
	nn.str = e
	if np == nil {
		l.head = nn
		return
	}
	var i int
	for i = 0; np.next != nil && i < idx; i++ {
		np = np.next
	}
	if i < idx {
		nn.prev = np
		np.next = nn
	} else {
		nn.prev = np.prev
		nn.next = np
		np.prev = nn
		if nn.prev != nil {
			nn.prev.next = nn
		}
	}
	if l.head == np {
		l.head = nn
	}
}

func (l *list) delete(idx int) {
	np := l.head
	if np == nil {
		return
	}
	for i := 0; np.next != nil && i < idx; i++ {
		np = np.next
	}
	if np == nil {
		return
	}
	if np.next != nil {
		np.next.prev = np.prev
	}
	if np.prev != nil {
		np.prev.next = np.next
	}
	if l.head == np {
		l.head = np.next
	}
	// free np
}

func main() {
	l := new(list)
	l.insert(0, "4")
	l.insert(0, "3")
	l.insert(0, "2")
	l.insert(0, "1")
	l.insert(0, "0")
	l.insert(3, "3x")
	l.insert(6, "6x")
	l.insert(99, "99x")
	for np := l.head; np != nil; np = np.next {
		fmt.Printf("%s ", np.str)
	}
	fmt.Print("\n")
}
