/*
Ex 5
Implement a doubly linked list that supports insert and delete.

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
)


type list struct {
	head *node	//head node of the list
	tail *node	//tail node of the list
	len int		//number of elements in the list
}

type node struct {	//node type that is stored in the list
	value string
	previous *node
	next *node
}

func newList() *list{	//simulate constructor of list
	head := &node{value: "head"}
	tail := &node{value: "tail"}
	len := 0
	head.next = tail
	tail.previous = head

	return &list{head,tail,len}
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string){
	if i > l.len || i < 0{	//index validation
		fmt.Println("Error: index out of bound!")
		return
	}

	n := &node{value:e}

	var cur *node
	if i <= (l.len/2) {	//find position from head
		cur = l.head
		for ;i > 0; i-- {
			cur = cur.next
		}
	}else{			//find position from tail
		cur = l.tail
		for ;i <= l.len; i++ {
			cur = cur.previous
		}
	}

	n.next = cur.next
	n.previous = cur
	cur.next = n
	n.next.previous = n
	l.len++
}


// delete removes the element at index i in the list l
func (l *list) delete(i int){
	if i >= l.len || i < 0{	//index validation
		fmt.Println("Error: index out of bound!")
		return
	}

	var cur *node
	if i <= (l.len/2) {	//find position from head
		cur = l.head
		for ;i >= 0; i-- {
			cur = cur.next
		}
	}else{	//find position from tail
		cur = l.tail
		for ;i < l.len; i++ {
			cur = cur.previous
		}
	}
	cur.previous.next = cur.next
	cur.next.previous = cur.previous
	l.len--
}

// return list as string
func (l *list) String() string{
	s := "List: "

	if l.len == 0{
		s += "empty."
	}else{
		cur := l.head
		s += "["
		for i:=0; i<l.len; i++ {
			cur = cur.next
			s += cur.value + ","
		}
		s = s[0:len(s)-1] + "]"
	}

	return s
}

func main(){
	l := newList()

	fmt.Println("\n#Test case 1: initialize list by inserting")
	fmt.Println("before insert -- ",l)
	l.insert(0,"mon")
	l.insert(1,"tue")
	l.insert(2,"wed")
	l.insert(3,"fri")
	l.insert(4,"sat")
	fmt.Println("after insert -- ",l)

	fmt.Println("\n#Test case 2: insert 'sun' at index 0")
	fmt.Println("before insert -- ",l)
	l.insert(0,"sun")
	fmt.Println("after insert -- ",l)

	fmt.Println("\n#Test case 3: insert 'thu' at index 4")
	fmt.Println("before insert -- ",l)
	l.insert(4,"thu")
	fmt.Println("after insert -- ",l)

	fmt.Println("\n#Test case 4: insert 'day' at index 8 should raise error")
	fmt.Println("before insert -- ",l)
	l.insert(8,"day")

	fmt.Println("\n#Test case 5: delete most of the elements")
	fmt.Println("before delete -- ",l)
	l.delete(0)
	l.delete(0)
	l.delete(0)
	l.delete(0)
	l.delete(0)
	l.delete(0)
	fmt.Println("after delete -- ",l)

	fmt.Println("\n#Test case 6: delete element at index 1 should raise error")
	fmt.Println("before delete -- ",l)
	l.delete(1)

	fmt.Println("\n#Test case 7: delete last element")
	fmt.Println("before delete -- ",l)
	l.delete(0)
	fmt.Println("after delete -- ",l)

	fmt.Println("\n#Test case 8: insert element after emptying list")
	fmt.Println("before insert -- ",l)
	l.insert(0,"new")
	fmt.Println("after insert -- ",l)

}
