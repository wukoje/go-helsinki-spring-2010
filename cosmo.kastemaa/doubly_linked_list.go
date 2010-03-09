package main

import "fmt"

func main() {
	l := new(list)
	l.insert(0, "a")
	fmt.Printf("%v\n", l)
	l.insert(0, "b")
	fmt.Printf("%v\n", l)
	l.insert(2, "c")
	fmt.Printf("%v\n", l)
	l.delete(1)
	fmt.Printf("%v\n", l)
	l.insert(1, "d")
	fmt.Printf("%v\n", l)
	l.insert(2, "e")
	fmt.Printf("%v\n", l)
	l.delete(0)
	fmt.Printf("%v\n", l)
	l.delete(2)
	fmt.Printf("%v\n", l)
}

type list struct {
	length int
	first *link
	last *link
}

type link struct {
	value string
	previous *link
	next *link
}

func (l *list) insert(i int, e string) {
	f := func(l *list, i int, p *link, n *link) {
		added := &link{e, p, n}
		if p != nil {
			p.next = added
		}
		if n != nil {
			n.previous = added
		}
		if i == 0 {
			l.first = added
		}
		if i == l.length {
			l.last = added
		}
		l.length++
	}
	l.operate(i, f)
}

func (l *list) delete(i int) {
	f := func(l *list, i int, p *link, n *link) {
		if n != nil {
			if p != nil {
				p.next = n.next
			}
			if n.next != nil {
				n.next.previous = p
			}
			if i == 0 {
				l.first = n.next
			}
			if i == l.length - 1 {
				l.last = p
			}
			l.length--
		}
	}
	l.operate(i, f)
}

func (l *list) operate(i int, f func(*list, int, *link, *link)) {
	if 0 <= i && i <= l.length {
		var p, n *link
		if i <= l.length / 2 { // go forwards
			n = l.first
			for c := i; c > 0; c-- {
				p, n = n, n.next
			}
		} else { // go backwards
			p = l.last
			for c := l.length - i; c > 0; c-- {
				p, n = p.previous, p
			}
		}
		f(l, i, p, n)
	}
}

func (l *list) String() string {
	result := ""
	for ln := l.first; ln != nil; ln = ln.next {
		if len(result) > 0 {
			result += ", "
		}
		result += ln.value
	}
	return "[" + result + "]"
}
