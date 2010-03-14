package main

import (
	"fmt"
)

type list struct {
	e		 string
	next *list
	prev *list
}

func (l *list) String() string {
	s := l.e
	if l.next == nil {
		return s
	}
	return s + " <-> " + l.next.String()
}

func (l *list) insert(i int, e string) {
	if l.next == nil && l.prev == nil && l.e == "" {
		l.e = e
	} else if i < 1 {
		n := new(list)
		n.e = l.e
		n.next = l.next
		n.prev = l
		if n.next != nil {
			n.next.prev = n
		}
		l.e = e
		l.next = n
	} else if l.next == nil {
		n := new(list)
		n.e = e
		n.next = nil
		n.prev = l
		l.next = n
	} else {
		l.next.insert(i-1, e)
	}
}

func (l *list) delete(i int) {
	if i < 1 || l.next == nil {
		if l.prev != nil {
			l.prev.next = l.next
		} else {
			l.e = l.next.e
			l.next = l.next.next
			l.prev = nil
		}
		if l.next != nil {
			l.next.prev = l.prev
		}
	} else {
		l.next.delete(i-1)
	}
}

func main() {
	s := []string{"a", "b", "c", "d", "e"}
	l1 := new(list)
	for _, k := range s {
		l1.insert(0, k)
		fmt.Printf("insert to first, list now: %v\n", l1)
	}
	l2 := new(list)
	for i, k := range s {
		l2.insert(i+318, k)
		fmt.Printf("insert to last, list now: %v\n", l2)
	}
	l3 := new(list)
	l3.insert(4, s[0])
	l3.insert(1, s[1])
	l3.insert(0, s[2])
	l3.insert(3, s[3])
	l3.insert(2, s[4])
	fmt.Printf("random inserts, result: %v\n", l3)
	l1.delete(2)
	fmt.Printf("list 1 after delete(2): %v\n", l1)
	l1.delete(0)
	fmt.Printf("list 1 after delete(0): %v\n", l1)
	l2.delete(3)
	fmt.Printf("list 2 after delete(3): %v\n", l2)
	l2.delete(88)
	fmt.Printf("list 2 after delete(88): %v\n", l2)
	l3.delete(2)
	l3.delete(3)
	fmt.Printf("list 3 after delete(2), delete(3): %v\n", l3)
}

