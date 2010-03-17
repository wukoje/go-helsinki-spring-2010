
package main
import (
    "fmt"
)



type list struct {
    data string
    prev *list
    next *list
}


func (n *list) insert( i int, s string) {
  newnode:= new(list)
  newnode.data = s
  list := n
  for z:=0;z<=i && n !=nil;n, z = n.next, z+1 {
   list = n 
  }
  n = list
  newnode.prev = list
  newnode.next = list.next
  if (list.next !=nil){
      list.next.prev = newnode
  }
  n.next = newnode  
  //fmt.Println(newnode.prev)
}
func (n *list) delete(d int ){
 d++;
  list := n
  for z:=0;z<=d && n !=nil;n, z = n.next, z+1 {
   list = n 
  }
  n = list
  
  
if n != nil {
   if n.prev.next !=nil{
    n.prev.next = n.next
   }
      if n.next !=nil {
	n.next.prev = n.prev
      }
    }
}

func (n *list) traverse() {
  for n =n.next; n != nil; n = n.next {
   fmt.Println(n) 
  }
}

func main () {
   fmt.Print("Initialized\n")
   lista := new(list)
   lista.insert(0,"Suomen turku")
   lista.insert(1,"Zerg Rush")
   lista.insert(2,"poni")
   lista.insert(3,"puuhamaa")
   lista.insert(6,"lelu")
   lista.insert(4,"Koira")
   lista.insert(5,"hippi")
   lista.traverse()
   lista.delete(3)
   fmt.Println("------delete 3------")
   lista.traverse()
   fmt.Println("------delete 0------")
   lista.delete(0)
   lista.traverse()
}
