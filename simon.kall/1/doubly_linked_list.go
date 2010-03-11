package main

import (
	"fmt";
 
)

func main() {
	list := new(list)
	for i:=0; i<5; i++ {
		list.insert(0,"hej")
	}
//	list.insert(0,"hoj")
//	list.insert(1,"hipp")
	list.print()
	fmt.Println("listan")
	list.delete(0)
	list.print()
}
type list struct {
	first *node
	last *node
}
type node struct {
	data string
	next *node
	prev *node
}
func (l *list) findNode(i int) *node{
	
	index := 0
	searchNode := l.first
	for searchNode != nil {
		if index == i {
		//	fmt.Println("nodefound")
			return searchNode
		}
		index++
		searchNode = searchNode.next
	}
//	fmt.Println("node not found")
	
	return nil
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	newNode := new(node)
	newNode.data = e
	
	if l.first == nil {
		l.first = newNode
		l.last = newNode
		newNode.prev = nil
		newNode.next = nil
	} else {
		prevNode := l.findNode(i)
		if prevNode == nil {
			fmt.Println("node index to high or low")
			return
		}
		newNode.prev = prevNode.prev
		newNode.next = prevNode
		
		if prevNode.prev == nil {
			l.first = newNode
		} else {
			prevNode.prev.next = newNode
		} 
		prevNode.prev = newNode
	}
}


// delete removes the element at index i in the list l

func (l *list) delete(i int) {
	delNode := l.findNode(i)
	if delNode == nil {
		
	} else {
		if delNode.prev == nil {
			l.first = delNode.next
		} else {
			delNode.prev.next = delNode.next
		}
		
		if delNode.next == nil {
			l.last = delNode.prev
		} else {
			delNode.next.prev = delNode.prev
		}
		
		delNode = nil
	} 
}
func (l *list) print() {
	currentNode := l.first
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	} 
}


