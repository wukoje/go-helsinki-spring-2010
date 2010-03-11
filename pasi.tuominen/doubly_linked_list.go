package main

import (
  "fmt"
)


type list struct {

	head *node

}

type node struct {
	next *node
	prev *node
	e string
}

func (l *list) print() {
	node := l.head
	for (node != nil) {
		fmt.Printf("%s\n", node.e)
		node = node.next
	}
}

func (l *list) printReversed() {
	node := l.head
	for (node.next != nil) {
		node = node.next
	}
	for (node != nil) {
		fmt.Printf("%s\n", node.e)
		node = node.prev
	}
}

func (l *list) append(e string) {
	last := l.head;
	newNode := createNode(e)
	
	if (last == nil) {
		l.head = newNode
	} else {
		for last.next != nil {
			last = last.next
		}

		last.next = createNode(e)
		last.next.prev = last
	}
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	node := l.head
	newNode := createNode(e)
	
	if (i == 0) {	//HEAD
		if (node == nil) {
			l.head = newNode
		} else {
			newNode.next = node
			node.prev = newNode
			l.head = newNode
		}
	} else {
		
		//traverse to node
		for index := 0; index < i-1; index++ {
		
			if (node.next == nil) {
				
				panic()
			}
			node = node.next
		}

		if (node.next == nil) { //TAIL
			node.next = newNode
			newNode.prev = node
		} else {
			
			newNode.next = node.next
			node.next = newNode
			
			newNode.next.prev = newNode
			newNode.prev = node
			
			
		}
	}
}



// delete removes the element at index i in the list l

func (l *list) delete(i int) {
	
	node := l.head
	if (i==0) {
		if (node != nil) {
			l.head = node.next
			l.head.prev = nil
		} else {
			panic();
		}
	} else {
	
		index := 0
		for (index < i) {
			node = node.next
			index++
			if (node == nil) {
				panic()
			}
		}
		if (node.next == nil) { //TAIL

			node.prev.next = nil
			node.prev = nil
			
		} else {

			node.prev.next = node.next
			node.next.prev = node.prev

		}

	//GC handles the rest
	}
}

func createNode(e string) *node {
	r := new(node);
	r.e = e
	return r
}


func main() {
	ll := new(list)
	ll.append("Node 1");
	ll.append("Node 2");
	ll.append("Node 3");
	ll.append("Node 4");
	ll.append("Node 5");
	
	fmt.Printf(":::5 nodes:\n")
	ll.print()
	fmt.Printf(":::Nodes reversed:\n")
	ll.printReversed()

	
	ll.insert(3,"Inserted node 1");
	ll.insert(0,"Inserted node 2");
	ll.insert(7,"Inserted node 3");  //TAIL

	fmt.Printf(":::After adding 3 nodes:\n")
	ll.print()
	fmt.Printf(":::Nodes reversed:\n")
	ll.printReversed()

	ll.delete(0) 
	ll.delete(3)
	ll.delete(5) //TAIL

	
	fmt.Printf(":::After deletion of 3 nodes:\n")
	ll.print()
	fmt.Printf(":::Nodes reversed:\n")
	ll.printReversed()
}
