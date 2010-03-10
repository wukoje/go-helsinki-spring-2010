package main

import (
	"fmt"
)

type list struct {
	len  int
	head *listNode
}

type listNode struct {
	data string
	prev *listNode
	next *listNode
}

func (n *listNode) String() string { return n.data }

func (l *list) String() string {
	str := "("
	tmp := l.head
	for tmp != nil {
		str = str + " " + fmt.Sprint(tmp)
		tmp = tmp.next
	}
	return str + ")"
}

// Inserts an new node to list in given index.
// If list is not long enough, inserts empty nodes for filling.
func (l *list) insert(i int, e string) {
	newNode := &listNode{e, nil, nil}

	if l.len == 0 {
		l.head = newNode
	} else {
		tmp := l.head
		for i > 0 {
			if tmp.next == nil { // empty node is inserted
				newEmpty := &listNode{"", tmp, nil}
				tmp.next = newEmpty
			}
			tmp = tmp.next
			i--
		}
		tmp.prev.next = newNode
		newNode.prev = tmp.prev
		tmp.prev = newNode
		newNode.next = tmp
	}
	l.len++
}

// Removes a node from a given index.
func (l *list) delete(i int) {
	if l.len == 0 {
		return
	}

	tmp := l.head
	for i > 0 {
		tmp = tmp.next
		if tmp == nil { // given index not in the list
			return
		}
		i--
	}
	tmp.prev.next, tmp.next.prev = tmp.next, tmp.prev
}

func main() {
	list := &list{0, nil}
	list.insert(0, "0")
	list.insert(1, "1")
	list.insert(2, "2")
	list.insert(1, "1/2")
	fmt.Println(list)
	list.delete(1)
	fmt.Println(list)
}
