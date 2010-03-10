package main

//Implement a doubly linked list that supports insert and delete.
import fmt "fmt"

type node struct {
	next, prev *node
	data string
}

type list struct {
	head *node
}


// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	if l != nil {
		if l.head == nil {
			l.head = &node{nil,nil,""}				
			l.head.next, l.head.prev = l.head, l.head		
		}

		if i == 0 {
			if l.head.next == l.head {
				l.head.data = e
			} else {
				n := &node{l.head, l.head.prev, e}
				l.head.prev, l.head = n, n
				n = l.head.prev
				n.next = l.head
			}
		} else {		
			n := l.head
			//find right index - 1
			for j := 0; j < i-1; j++ {
				if n.next == l.head {
					n.next = &node{l.head,n,""}
					l.head.prev = n.next 
				}
				n = n.next
			}
			temp := n.next			
			n.next = &node{n.next,n,e}
			temp.prev = n.next
		}
	}
}


// delete removes the element at index i in the list l
func (l *list) delete(i int) {
	if l != nil {
		if  l.head != nil {
			n := l.head							
			if i == 0 {
				temp := n.prev
				temp.next, n = n.next, n.next
				n.prev, l.head = temp, temp.next
			} else if l.head.next != nil {
				for j := 0; j < i-1 && l.head != n.next; j++ {
					n = n.next;
				}
				temp := n.next
				n.next, temp =  temp.next, temp.next
				temp.prev = n			
			}
		}
	}
}


func main(){
	//n := &node{nil,nil,""}
	l := &list{nil}
	l.insert(0, "1")
	l.insert(1, "3")
	l.insert(2, "4")
	l.insert(3, "5")
	l.insert(4, "6")
	l.insert(5, "17")
	l.insert(6, "18")
	l.insert(7, "19")
	l.insert(8, "20")
	fmt.Printf("done\n")

	n := l.head
	fmt.Printf("%s,\n", n.data)
	n = n.next	
	//print whole list
	for n != l.head {
		if n == nil {
			n = l.head
		}		
		fmt.Printf("%s, \n", n.data)
		n = n.next
	}

	// reversed
	n = l.head
	fmt.Printf("\n\n%s,\n", n.data)
	n = n.prev	
	//print whole list
	for n != l.head {
		if n == nil {
			n = l.head
		}		
		fmt.Printf("%s, \n", n.data)
		n = n.prev
	}


	//delete test and outputs

	fmt.Printf("\nDelete index 2\n")
	l.delete(2)
	n = l.head
	fmt.Printf("%s,\n", n.data)
	n = n.next	
	for n != l.head {
		fmt.Printf("%s, \n", n.data)
		n = n.next
	}

	
	fmt.Printf("\nDelete index 5\n")
	l.delete(5)
	n = l.head
	fmt.Printf("%s,\n", n.data)
	n = n.next	
	for n != l.head {
		fmt.Printf("%s, \n", n.data)
		n = n.next
	}

	fmt.Printf("\nDelete index 0\n")
	l.delete(0)
	n = l.head
	fmt.Printf("%s,\n", n.data)
	n = n.next	
	for n != l.head {
		fmt.Printf("%s, \n", n.data)
		n = n.next
	}
}
