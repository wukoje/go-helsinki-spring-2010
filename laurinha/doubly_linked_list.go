package main

import "fmt"

type listNode struct {
	next  *listNode
	prev  *listNode // this isn't actually needed
	value string
}

type list struct {
	head   *listNode
	length int
}

// returns the i'th node in l or nil if i is not a valid index
func (l *list) getNode(i int) *listNode {
	if i < 0 || i >= l.length {
		return nil
	}

	current := l.head
	for i > 0 {
		current = current.next
		i--
	}

	return current
}

// inserts e into l as the i'th member
func (l *list) insert(i int, e string) {
	if i < 0 || i > l.length {
		return
	}

	if l.head == nil {
		l.head = &listNode{nil,nil,e}
	} else if i == 0 {
		node := &listNode{l.head, nil, e}
		l.head.prev = node
		l.head = node
	} else {
		prev := l.getNode(i-1)
		next := prev.next
		node := &listNode{next, prev, e}
		prev.next = node
		if next != nil {
			next.prev = node
		}
	}

	l.length++
}

// removes the i'th element in l and returns it
// ok = false if given index is out of bounds, otherwise true
func (l *list) remove(i int) (e string, ok bool) {
	// the excercise description has no return value here, but that would
	// make it impossible to get anything out of the list so a valid
	// implementation for both insert and delete would be { return }
	node := l.getNode(i)
	if node == nil {
		return "", false
	}

	e = node.value
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	if node == l.head {
		l.head = next
	}
	l.length--

	return e, true
}

func main() {
	l := new(list)
	for i := 0; i < 5; i++ {
		l.insert(i, fmt.Sprint(i))
		fmt.Printf("insert(%v, fmt.Sprint(%v))\n", i, i)
	}
	for i := 4; i >= 0; i-- {
		e,_ := l.remove(i)
		fmt.Printf("remove(%v): %v\n", i, e)
	}
}

