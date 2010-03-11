package main

import (
	"fmt"
)

func main() {

	testList := new(list)
	testList.firstNode = new(node)
	testList.firstNode.element = "Test"

	printList(testList)
	testList.insert(1, "Insert test")
	printList(testList)
	testList.insert(1, "Second test")
	printList(testList)
	testList.insert(4, "Third test")
	printList(testList)
	fmt.Println("Length of the list:", len(testList))
	testList.delete(0)
	printList(testList)
	fmt.Println("Length of the list:", len(testList))
	testList.delete(0)
	testList.delete(0)
	printList(testList)
	fmt.Println("Length of the list:", len(testList))
	testList.delete(0)
	printList(testList)
	fmt.Println("Length of the list:", len(testList))
}

type list struct {
	firstNode *node
}

type node struct {
	element  string
	previous *node
	next     *node
}

// insert adds the element e at index i in the list l
// if index is over the length of the list, then the element will be appended to the list
func (l *list) insert(i int, e string) {
	tempNode := l.firstNode
	var prevNode *node

	for j := 0; tempNode != nil; j, tempNode = j+1, tempNode.next {
		if j == i {
			break
		}
		if tempNode != nil {
			prevNode = tempNode
		}

	}

	if tempNode == nil {
		tempNode = new(node)
		tempNode.element = e
		tempNode.previous = prevNode
		prevNode.next = tempNode
	} else {
		newNode := new(node)
		newNode.element = e
		newNode.previous = prevNode
		newNode.next = tempNode
		prevNode.next = newNode
		tempNode.previous = newNode
	}
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	tempNode := l.firstNode
	var prevNode *node

	for j := 0; tempNode.next != nil; j, tempNode = j+1, tempNode.next {
		if j == i {
			break
		}
		if tempNode != nil {
			prevNode = tempNode
		}
	}

	if tempNode.next != nil {
		if tempNode.previous == nil {
			l.firstNode = tempNode.next
			tempNode.next.previous = nil
		} else {
			prevNode.next = tempNode.next
			tempNode.next.previous = prevNode
		}
	} else {
		if tempNode.previous == nil {
			l.firstNode = nil
		} else {
			prevNode.next = nil
		}
	}
}

func printList(l *list) {
	fmt.Printf("doubly-linked-list: -> ")
	for tempNode := l.firstNode; tempNode != nil; tempNode = tempNode.next {
		fmt.Printf("%s -> ", tempNode.element)
	}
	fmt.Println("nil")
}

func len(l *list) (count int) {
	for tempNode := l.firstNode; tempNode != nil; tempNode, count = tempNode.next, count+1 {
	}
	return
}
