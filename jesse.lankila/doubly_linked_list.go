package main

import "fmt"

type list struct {
  head *elem
  tail *elem
  size int
}

type elem struct {
  payload string
  prev *elem
  next *elem
}

// search for linked list - returns the element in index i
// if i is negative, returns the first element; if i >= list.size, returns last element
func (l *list) search(i int) (ret_elem *elem) {
  
  if i < 0 {
    ret_elem = l.head
  } else if i >= l.size {
    ret_elem = l.tail
  } else {
  
    curr_elem := l.head
    j := 0
    for j < i {
      curr_elem = curr_elem.next
      j++
    }
    ret_elem = curr_elem
  }
  return ret_elem
}  

// insert adds the element e at index i in the list l
// if i is negative, inserts as first element; if i >= list.size, as last element
func (l *list) insert(i int, e string) {

  var new_elem *elem = new(elem)
  new_elem.payload = e
  
  // insert to head
  if i <= 0 {
    // empty list
    if l.head == nil {
      l.head = new_elem
      l.tail = new_elem
    } else {
      new_elem.next, l.head.prev = l.head, new_elem
      l.head = new_elem
    }
    
  // insert to tail
  } else if i >= l.size {
    l.tail.next, new_elem.prev = new_elem, l.tail
    l.tail = new_elem
    
  // search in the middle
  } else {
    var curr_elem *elem = l.search(i)
    // we replace current node with with new one
    curr_elem.prev.next, new_elem.prev = new_elem, curr_elem.prev
    curr_elem.prev, new_elem.next = new_elem, curr_elem
  }
  
  // finally, increase size by one
  l.size++

}

// delete removes the element at index i in the list l
func (l *list) delete(i int) string {
  
  var curr_elem *elem = l.search(i)
  
  if i <= 0 {
    l.head = curr_elem.next
    l.head.prev = nil
  } else if i >= l.size-1 {
    l.tail = curr_elem.prev
    l.tail.next = nil
  } else {
  curr_elem.prev.next, curr_elem.next.prev = curr_elem.next, curr_elem.prev
  }
  l.size--
  return curr_elem.payload
}

func main() {

  example := new(list)
  example.insert(0, "plaa")
  example.insert(1, "hyrk")
  example.insert(1, "zing")
 
  for i:=0; i < example.size; i++ {
    fmt.Println(example.search(i).payload)
  }
  
  str := example.delete(2)
  
  fmt.Printf("Deleted element: %s\n", str)
  for i:=0; i < example.size; i++ {
    fmt.Println(example.search(i).payload)
  }
} 
