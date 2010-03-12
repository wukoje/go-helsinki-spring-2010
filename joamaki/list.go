package main

import "fmt"

type entry struct {
	prev, next *entry
	data string
}

type list struct {
	head *entry
}

func (l *list) insert(i int, e string) {
	n := l.head  // next from i
	var p *entry // prev from i

	for i > 0 && n != nil {
		p = n
		n = n.next
		i--
	}

	if i > 0 {
		panic("list index out of range")
	}

	le := &entry{p, n, e}
	if p != nil { p.next = le }
	if n != nil { n.prev = le }
	if n == l.head { l.head = le }
}

func (l *list) delete(i int) {
	le := l.head

	for i > 0 && le != nil {
		le = le.next
		i--
	}

	if i > 0 {
		panic("list index out of range")
	}

	if le.prev != nil {
		le.prev.next = le.next
	} else {
		l.head = le.next
	}
	if le.next != nil {
		le.next.prev = le.prev
	}
}

func (l *list) String() string {
	le := l.head
	var res string
	for le != nil {
		res = res + le.data + " "
		le = le.next
	}
	return res
}

func main() {
	l  := new(list)
	l.insert(0, "a")
	l.insert(1, "c")
	l.insert(2, "f")
	l.insert(2, "e")
	l.insert(2, "d")
	l.insert(1, "b")
	l.insert(6, "g")

	fmt.Println(l)
	l.delete(5)
	l.delete(4)
	l.delete(2)
	l.delete(2)
	l.delete(1)
	l.delete(0)
	fmt.Println(l)
}
