package main

import (
	"fmt"
)

type node struct {
    prev *node
    next *node
    data string
}

type list struct {
    head *node
    tail *node
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
    newNode := new(node)
    newNode.data = e
    currentNode := l.head
    var prevNode *node
    prevNode = nil
    for n:=0; n<i; n++ {
        if currentNode == nil {
            break
        }
        prevNode = currentNode
        currentNode = currentNode.next
    }
    if prevNode != nil {
        prevNode.next = newNode
        newNode.prev = prevNode
    }
    if i == 0 || l.head == nil {
        l.head = newNode
    }
    if currentNode == nil {
        l.tail = newNode
    } else {
        newNode.next = currentNode
        currentNode.prev = newNode
    }
}


// delete removes the element at index i in the list l
func (l *list) delete(i int) {
    currentNode := l.head
    for n:=0; n<i; n++ {
        if currentNode == nil {
            return
        }
        currentNode = currentNode.next
    }
    if currentNode == nil {
        return
    }
    if i == 0 {
        l.head = currentNode.next
        if currentNode.next != nil {
            currentNode.next.prev = nil
        }
    } else {
        currentNode.prev.next = currentNode.next
        currentNode.next.prev = currentNode.prev
    }
    if currentNode.next == nil {
        l.tail = currentNode.prev
    }
}

func printNodes(list *list) {
    fmt.Println("All nodes:")
    current := list.head
    for current != nil {
        fmt.Println(current.data)
        current = current.next
    }
}

func main() {
    list := new(list)
    list.head = nil
    list.tail = nil
    list.insert(0, "foo")
    list.insert(0, "bar")
    list.insert(2, "foobar")
    list.insert(3, "foobaz")
    fmt.Println(list.head.data)
    printNodes(list)
    list.delete(2)
    printNodes(list)
    list.delete(0)
    list.delete(0)
    list.delete(0)
    printNodes(list)
    list.delete(0)
    printNodes(list)
    list.insert(0, "foo")
    printNodes(list)
}
