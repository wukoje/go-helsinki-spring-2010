package main

import (
        "fmt"
)

func main() {

	
	l1 := &list{nil,nil,0}
        fmt.Println(l1)

	l1.insert(0,"test")

	fmt.Println(l1)
	fmt.Println(l1.print())
	l1.insert(0,"test2")

	fmt.Println(l1)

	fmt.Println(l1.print())

	l1.delete(1)

	fmt.Println(l1)
	fmt.Println(l1.print())

	l1.insert(0,"test2")
	l1.insert(0,"test3")

	l1.insert(0,"test4")

	l1.insert(2,"test5")

	fmt.Println(l1)

	fmt.Println(l1.print())

	l1.delete(1)

	fmt.Println(l1.print())
	fmt.Println(l1)

	




}

type node struct {

	next,prev *node

	 value string

}

type list struct {

	begin,end *node
	size int
 
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {
	
	n := &node{nil,nil,e}
	
    if l.size > 0 {

	if i < l.size-1 && i > 0 {

		a:= l.begin

		for j:= 0 ; j <= i;j++ {

			a = a.next
	
		}

		n.prev = a.prev

		a.prev.next = n

		a.prev = n

		n.next = a

		l.size++
	
	} else if i == 0 {

		a:= l.begin

		a.prev = n
		n.next = a

		l.begin = n

		l.size++

	} else if i > l.size {

		a:= l.end

		n.prev = a
	
		a.next = n

		n.next = nil

		l.end = n

		l.size++

	}
   } else {

	l.begin = n
	l.begin.next = l.end

	l.begin.prev = nil
	
	l.end = n

	l.end.prev = l.begin

	l.end.next = nil

	l.size++

   }




}



// delete removes the element at index i in the list l
func (l *list) delete(i int) {

    if i <= l.size-1 {

	e:= l.begin

	for j:=1;j<i;j++ {

			e = e.next
	}

	if e.prev == nil {

		l.begin = e.next

	} else {
			
		e.prev.next  = e.next
	}

	if e.next == nil {
			
		l.end = e.prev				

	} else {

		e.next.prev = e.prev
			
	}

	e.prev = nil
	e.next = nil

	l.size--

   }		
		

		

}

func (l *list) print() string {

	s:="["
	e := l.begin

	for i:=0;i<l.size;i++ {

		s += (" "+e.value)
		e = e.next 	

	}

	s += " ]"

	return s
}
