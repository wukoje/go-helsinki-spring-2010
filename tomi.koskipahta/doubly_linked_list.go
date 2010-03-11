package main

import fmt "fmt"

func main() {

	l := new(list)
	l.first = nil
	l.last = nil
	
	fmt.Println()
	fmt.Println("Inserting index 0 and Data#0")
	l.insert(0, "Data#0")
	l.toString()
	
	fmt.Println()
	fmt.Println("Inserting index 2 and Data#1")
	l.insert(2, "Data#1")
	l.toString()
		
	fmt.Println()
	fmt.Println("Inserting index 4 and Data#2")
	l.insert(4, "Data#2")
	l.toString()
		
	fmt.Println()
	fmt.Println("Inserting index 5 and Data#3")
	l.insert(5, "Data#3")
	l.toString()
		
	fmt.Println()
	fmt.Println("Inserting index 3 and Data#4")
	l.insert(3, "Data#4")
	l.toString()
		
	fmt.Println()
	fmt.Println("Inserting index 1 and Data#5")
	l.insert(1, "Data#5")
	l.toString()
		
	fmt.Println()
	fmt.Println("Inserting index 0 and Data#6")
	l.insert(0, "Data#6")
	l.toString()
	
	fmt.Println()
	fmt.Println("Deleting index 5")
	l.delete(5)
	l.toString()
		
	fmt.Println()
	fmt.Println("Deleting index 7")
	l.delete(7)
	l.toString()
		
	fmt.Println()
	fmt.Println("Deleting index 2")
	l.delete(2)
	l.toString()
			
	fmt.Println()
	fmt.Println("Deleting index 1")
	l.delete(1)
	l.toString()
			
	fmt.Println()
	fmt.Println("Deleting index 0")
	l.delete(0)
	l.toString()
}





type node struct {
	data string
	next *node
	prev *node
}





type list struct {
	first *node   // points to first node of list; null for empty list
    last *node
}






func (l *list) insert(i int, e string) { 	//Insert in given index,
	n := l.first				//If index higher than nodes
	if n == nil {				//data will be inserted last
		x := new(node)
		x.data = e
		l.last = x
		l.first = x
		return
	}
	
		if i == 0 {
			x := new(node)
			x.data = e
			x.next = n
			
			if n.prev == nil {
				l.first = x
				x.prev = nil
				n.prev = x
				
			} else if n.prev != nil{
				x.prev = n.prev
				x.prev.next = x
				x.next.prev = x
			}
			return
		}

	
	
	n = l.first
	for j := 1; j <= i; j++  {
		
		if n.next != nil {
			n = n.next
			
		} else {
			x := new(node)
			x.data = e
			x.next = n
			l.last = x
			x.next = nil
			n.next = x
			x.prev = n
			return
		}
		
		if i == j {
			x := new(node)
			x.data = e
			x.next = n
			
			if n.prev == nil {
				l.first = x
				x.prev = nil
				n.prev = x
				
			} else if n.prev != nil{
				x.prev = n.prev
				x.prev.next = x
				x.next.prev = x
			}
			return
		}
		
	}
}






func (l *list) delete(i int) {
	n := l.first

	if n == nil {
		fmt.Println("Error, list is empty.")
		return
	}
	
	for j := 0; j <= i; j++  {
		if j == i { 
			
			if n.prev != nil && n.next != nil {
				n.prev.next = n.next
				n.next.prev = n.prev
				
			} else if n.prev != nil && n.next == nil {
				n.prev.next = n.next
				l.last = n.prev
				
			} else if n.prev == nil && n.next != nil {
				n.next.prev = n.prev
				l.first = n.next
			}		
			n.next = nil
			n.prev = nil			
			return
		}
		if n.next != nil {
			n = n.next
		} else {
			fmt.Println("Delete-Error. Index ", i , " was not found.")
			return
		}
	}
	fmt.Println("Delete-Error. Index ", i , " was not found.")
}






func (l *list) toString() {
	n := l.first
	for n != nil {
		fmt.Print(n.data, " <-> ")
		n = n.next
	}
	fmt.Println()
}
