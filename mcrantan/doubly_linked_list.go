package main

import(
  "os"
  "fmt"
  "rand"
  "time"
  "flag"
)

type node struct {
  value string
  next *node
  prev *node
}

type list struct {
  size int
  root *node
}

func main() {

  var SIZE int
  flag.IntVar(&SIZE, "s", 4, " Initial heap size (default: 4) ")
  flag.Parse()

  rand.Seed(time.Nanoseconds())

  mylist := &list{0,nil}

  fmt.Printf("List length: %d.\n", mylist.size)
  fmt.Printf("Inserting %d nodes to list.\n", SIZE)

  for i := 0; i < SIZE; i++ {
    mylist.insert(0, rndstring(8))
  }
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Inserting %d nodes to list at index %d.\n", 1, mylist.size / 2)
  mylist.insert(mylist.size / 2, rndstring(8))
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Inserting %d nodes to list at index %d.\n", 1, mylist.size - 1)
  mylist.insert(mylist.size - 1, rndstring(8))
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Deleting %d nodes from list at index %d.\n", 1, mylist.size / 2)
  mylist.delete(mylist.size / 2)
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Deleting %d nodes from list at index %d.\n", 1, 0)
  mylist.delete(0)
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Deleting %d nodes from list at index %d.\n", 1, mylist.size - 1)
  mylist.delete(mylist.size - 1)
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("\nCreating new list with 0 elements.\n")
  mylist = &list{0,nil}
  fmt.Printf("List contains currently %d nodes with values:\n  %v.\n", mylist.size, mylist)

  fmt.Printf("Deleting %d nodes from list at index %d.\n", 1, 0)
  mylist.delete(0)

  fmt.Printf("Deleting %d nodes from list at index %d.\n", 1, -1)
  mylist.delete(-1)

  fmt.Printf("Inserting %d nodes to list at index %d.\n", 1, -1)
  mylist.insert(-1, rndstring(8))

  fmt.Printf("Inserting %d nodes to list at index %d.\n", 1, 42)
  mylist.insert(42, rndstring(8))

  os.Exit(0)
}

func (l *list) String() string {

  if l.size == 0 {
    return "[]"
  }

  s := ""
  curnode := l.root

  for i := 0; i < l.size; i++ {
    s += " " + curnode.value + " "
    curnode = curnode.next
  }

  return s
}

func rndstring(length int) string {

//  picks := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//  picks := map[string]int{"0":0,"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"a":10,"b":11,"c":12,"d":13,"e":14,"f":15}
  picks := map[int]string{0:"0",1:"1",2:"2",3:"3",4:"4",5:"5",6:"6",7:"7",8:"8",9:"9",10:"a",11:"b",12:"c",13:"d",14:"e",15:"f"}
  plen := len(picks)
  s := ""

  for i := 0; i < length; i++ {
    s += picks[rand.Intn(plen)]
  }

  return s
}

// insert adds the element e at index i in the list l
func (l *list) insert(i int, e string) {

  if i < 0 || i > l.size {
    fmt.Printf("\nCan't insert to list at index %d (list length: %d).\n", i, l.size)
    return
  }

  newnode := &node{e, nil, nil}

  if l.size == 0 {
    l.root = newnode
    (l.root).value = e
    l.size++
    return
  }
    
  curnode := l.root

  for j := 0; j < i; j++ {
    curnode = curnode.next
  }

  if i == 0 {                           // First (=root) node?
    (l.root).prev = newnode
    newnode.next = l.root
    l.root = newnode
  } else if i == l.size - 1 {           // Last node?
    newnode.prev = curnode.prev
    newnode.next = curnode
    (curnode.prev).next = newnode
    curnode.prev = newnode
    
  } else {
    (curnode.prev).next = newnode       // Node somewhere between
    newnode.prev = curnode.prev
    newnode.next = curnode
    curnode.prev = newnode
  }

  l.size++
}

// delete removes the element at index i in the list l
func (l *list) delete(i int) {

  if l.size <= i || i < 0 {
    fmt.Printf("\nCan't delete from list at index %d (list length: %d).\n", i, l.size)
    return
  }

  curnode := l.root

  for j := 0; j < i; j++ {
    curnode = curnode.next
  }

  if i == 0 {                           // First (=root) node?
    l.root = curnode.next
  } else if i == l.size - 1 {           // Last node?
    (curnode.prev).next = nil
  } else {
    (curnode.prev).next = curnode.next
    (curnode.next).prev = curnode.prev
  }

  curnode.value = ""
  curnode.prev = nil            // Redundant if this was root node
  curnode.next = nil            // Redundant if this was last node
  l.size--
}


