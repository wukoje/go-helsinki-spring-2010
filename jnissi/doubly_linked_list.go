package main

import "fmt"

type node struct {
	pred, succ *node
	value      string
}

func (n *node) insertBefore(n1 *node) {
	if n.pred != nil {
		n.pred.succ = n1
	}
	n1.pred, n.pred, n1.succ = n.pred, n1, n
}

func (n *node) insertAfter(n1 *node) {
	if n.succ != nil {
		n.succ.pred = n1
	}
	n1.succ, n.succ, n1.pred = n.succ, n1, n
}

type list struct {
	head, tail *node
}

func (l *list) insert(i int, v string) {
	new := false
	if l.head == nil {
		l.head = &node{nil, nil, ""}
		l.tail = l.head
		new = true
	}

	curr := l.head
	for j := i; j > 0; j-- {
		if curr.succ == nil {
			new = true
			curr.insertAfter(&node{nil, nil, ""})
			l.tail = curr.succ
		}
		curr = curr.succ
	}
	if new {
		curr.value = v
	} else {
		curr.insertBefore(&node{nil, nil, v})
		if i == 0 {
			l.head = curr.pred
		}
	}
}

func (l *list) remove(i int) {
	curr := l.head
	if curr == nil {
		return
	}
	for ; i > 0; i-- {
		if curr.succ == nil {
			return
		}
		curr = curr.succ
	}
	if curr == l.head {
		l.head = curr.succ
	}
	if curr == l.tail {
		l.tail = curr.pred
	}
	if curr.succ != nil {
		curr.succ.pred = curr.pred
	}
	if curr.pred != nil {
		curr.pred.succ = curr.succ
	}
}

func (l *list) String() string {
	var ret string = "list: "
	if l.head == nil {
		return ret
	}
	curr := l.head
	for {
		ret += "\"" + curr.value + "\""
		if curr.succ == nil {
			break
		}
		ret += ", "
		curr = curr.succ
	}
	ret += " reverse: "
	curr = l.tail
	for {
		ret += "\"" + curr.value + "\""
		if curr.pred == nil {
			break
		}
		ret += ", "
		curr = curr.pred
	}
	return ret
}

func main() {
	list := list{nil, nil}
	list.insert(0, "one")
	list.insert(0, "two")
	list.insert(5, "three")
	list.remove(0)
	fmt.Printf("%s\n", list.String())
}
