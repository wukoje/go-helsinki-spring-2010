package main

import fmt "fmt"

type list struct {
  first *node
}

type node struct {
  next *node
  prev *node
  elem string
}

func create(prev *node, next *node, e string) *node {
  n := new(node)
  n.prev = prev
  n.next = next
  n.elem = e
  //fmt.Printf("Created %s where", e, prev, next)
  //fmt.Println()
  return n
}

// insert adds the element e at index i in the list l
// if the list is shorter than i then e is placed in to the end of the list
func (l *list) insert(i int, e string) {
  n := l.first

  // if the list is empty, e becomes the first item
  if n == nil {
    l.first = create(nil, nil, e)
    return
  }

  // does this go to the first index?
  if i <= 0 {
    m := create(nil, n, e)
    n.prev = m
    l.first = m
    return
  } 

  // otherwise traverse to the correct position
  j := 1
  for n.next != nil && j < i {
      n = n.next
      j++
  }

  m := create(n, n.next, e)

  // add the element into the list and fix the links
  if (n.next != nil) {
    n.next.prev = m
  }

  n.next = m
}

// delete removes the element at index i in the list l
// if the list is shorter than i then the last element is removed
func (l *list) delete(i int) {
  n := l.first

  j := 0
  for n.next != nil && j < i {
    n = n.next
    j++
  }

  if j == 0 && n != nil {
      l.first = n.next
  }

  tmp := n.next

  if (n.prev != nil) {
    n.prev.next = tmp
  }

  if (n.next != nil) {
    n.next.prev = n.prev
  }

}

func print_list(l *list) {
  n := l.first
  for n != nil {
    fmt.Print("["+n.elem+"] ")
    n = n.next
  }
  fmt.Println()
}

func main() {
  my_list := new(list)
  my_list.insert(1, "nolla")
  print_list(my_list)
  my_list.delete(2)
  print_list(my_list)
  my_list.insert(5, "eka")
  my_list.insert(5, "toka")
  my_list.insert(5, "kolmas")
  my_list.insert(5, "neljäs")
  print_list(my_list)
  my_list.delete(0)
  print_list(my_list)
  my_list.insert(0, "uusi eka")
  print_list(my_list)
  my_list.insert(1, "uusi toka")
  print_list(my_list)
  my_list.insert(10, "vika")
  print_list(my_list)
  my_list.delete(2)
  print_list(my_list)
}
