/*
	1 doubly_linked_list.go

	Implement a doubly linked list that supports insert and delete.

	type list struct {
		// your code here
	}

	// insert adds the element e at index i in the list l
	func (l *list) insert(i int, e string)

	// delete removes the element at index i in the list l
	func (l *list) delete(i int)
*/

package main

import(
	"os"
	"fmt"
)

type node struct {
	value string
	next *node
	prev *node
}

type list struct {
	count int
	root *node
}

// pretty print list
func (l *list) String() string {

	s := ""

	if l.count == 0 {
		return "[]"
	} else if l.count == 1 {
		return "[" + l.root.value + "]"
	}

	for i, current := 1, l.root; i <= l.count; i, current = i+1, current.next {
		tempNode := *current
		s += "[" + tempNode.value + "] "
	}

	return s
}

// insert adds the element e at index i in the list l. return index or error
func (l *list) insert(i int, e string) (n int, err os.Error){


	if i > l.count || i < 0 {
		return -1, os.NewError("index out of bounds")
	}

	var currentNode node

	pntrCurrent := l.root

	if l.count == 0 { // empty list

		newNode := node{e, nil, nil}
		l.root = &newNode

	} else if i == l.count  { // node is inserted in last index

		for ; pntrCurrent.next != nil; pntrCurrent = currentNode.next {
			currentNode = *pntrCurrent
		}

		newNode := node{e, nil, pntrCurrent}
		pntrCurrent.next = &newNode

	} else { // node is inserted between nodes
		for j := 0; j <= i; j, pntrCurrent = j + 1, currentNode.next {
			currentNode = *pntrCurrent
		}

		newNode := node{e, &currentNode, currentNode.prev}
		currentNode.prev.next = &newNode
		currentNode.prev = &newNode

	}

	l.count += 1


	return i, nil

}

// delete removes the element at index i in the list l. returns the deleted element or error
func (l *list) delete(i int) (s string, err os.Error){

	if i > l.count || i < 0 {
		return "D:", os.NewError("index out of bounds")
	}

	if l.count == 0 { 
		return "D:", os.NewError("list is empty")
	} 

	if l.count == 1 { // shortcut if only one element in the list
		l.count = 0
		temp := *l.root
		l.root = nil
		return temp.value, nil	
	}

	var currentNode node
	pntrCurrent := l.root

  // fetch correct node as currentNode
	for j := 0; j <= i; j, pntrCurrent = j + 1, currentNode.next {
		currentNode = *pntrCurrent
	}

	if i == l.count - 1 { // deleting last item
		currentNode.prev.next = nil
  } else if i == 0 { // deleting first item
    currentNode.next.prev = nil
    l.root = currentNode.next
	} else {
	  currentNode.next.prev = currentNode.prev
		currentNode.prev.next = currentNode.next
	}

	// clean dangling references
  ret := currentNode.value
	currentNode.next = nil
	currentNode.prev = nil

	l.count -= 1

	return ret, nil
}

func main(){

	l := &list{0, nil}
	l.insert(0, "foo")
	l.insert(1, "bar");
	l.insert(2, "asdsad")

  fmt.Printf("Let's start with:\n%v\n---\nadding to index 1:\n", l)	

	l.insert(1, "qwe")

	fmt.Printf("%v\n---\nadding to index 10 causes error:\n", l)	

	_, err := l.insert(10, "overflow!")
	
	if err != nil {
		fmt.Printf("'%v' is the error list returned\n\n", err)
	}

	fmt.Printf("And nothing is changed in the list:\n%v\n---\n", l)	

	_, err = l.delete(10)

	if err != nil {
		fmt.Printf("Deleting from index 10 is also an error:\n'%v' is the error list returns\n---\n", err)
	}


  fmt.Printf("So, the array is:\n%v\n\nDeleting last item:\n", l)
	l.delete(3)
	fmt.Printf("%v\n---\n", l)
  fmt.Printf("deleting from the middle:\n")
  l.delete(1)
	fmt.Printf("%v\n---\n", l)
  fmt.Printf("Deleting first item:\n")
  l.delete(0)
	fmt.Printf("%v\n---\n", l)
  l.delete(0)
  fmt.Printf("Deleting from empty list ('%v') is error:\n", l)

  _, err = l.delete(0)
  
  if err != nil {
    fmt.Printf("'%v' is the error list returned\nDone!\n", err)
  }
}


