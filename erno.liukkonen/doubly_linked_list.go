package main

import fmt "fmt"

type list struct {
	count int
	head *node
}

type node struct {
	values string
	previous *node
	next	*node
}

func (l *list) insert (i int, e string) {
	nNode := &node{e, nil, nil}

	if l.count != 0 {

		temp := l.head
		for i > 0 {
			if temp.next == nil {
				temp.next = &node{"", temp, nil}
			}
			temp = temp.next
			i--
		}

		temp.previous.next =nNode
		nNode.previous = temp.previous
		temp.previous = nNode
		nNode.next = temp
	}else {
		l.head = nNode
	}
	l.count++
}

func (l *list) delete(i int) {
	if l.count != 0 {

		temp := l.head
		for i > 0 {
			temp = temp.next
			if temp == nil {
				return
			}
			i--
		}

	temp.previous.next, temp.next.previous = temp.next, temp.previous
	}else {
		return
	}
}


func main() {
	list := &list{0, nil}
	list.insert(0, "5")
	list.insert(1, "3")
	list.insert(2, "4")
	list.insert(3, "2")
	list.insert(4, "8")

	for node := list.head; node != nil; node = node.next {
		fmt.Printf("%s ", node.values)
	}
	fmt.Println()

	list.delete(4)

	for node := list.head; node != nil; node = node.next {
		fmt.Printf("%s ", node.values)
	}
	fmt.Println()

	list.delete(1)

	for node := list.head; node != nil; node = node.next {
		fmt.Printf("%s ", node.values)
	}
	fmt.Println()
}
