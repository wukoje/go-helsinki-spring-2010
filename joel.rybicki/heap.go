package main

import fmt "fmt"

type heap struct {
  arr []int // the slice storing the heap
  length int  // the amount if items in stored in the heap
}

func parent(i int) int {
  return i/2
}

func left(i int) int {
  return i*2 + 1
}

func right(i int) int {
  return i*2 + 2
}

// insert adds element e to the heap
func (h *heap) insert(e int) {
  i := h.length
  h.arr[i] = e

  // now fix the heap
  for i > 0 {
    j := parent(i)
    if h.arr[i] < h.arr[j] {
      h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
      i = j
    } else break;
  }

  h.length++
}


// pop returns the smallest element from the heap
// undefined behaviour if the heap is empty!
func (h *heap) pop() int {
  min := h.arr[0]
  h.length--
  h.arr[0] = h.arr[h.length]

  // fix the heap
  i := 0
  for left(i) < h.length {
    j := i
    if left(i) < h.length && h.arr[left(i)] < h.arr[j] {
      j = left(i)
    }
    if right(i) < h.length && h.arr[right(i)] < h.arr[j] {
      j = right(i)
    }
    if h.arr[j] < h.arr[i] {
        h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
        i = j
    } else {
      break
    }
  }
  return min
}


// create takes the given slice and creates a heap type that uses the slice for storage

func create(s []int) *heap {
  h := new(heap)
  h.arr = s
  h.length = 0
  for i := 0; i<len(s); i++ {
    s[i] = -1
  }
  return h
}

func main() {
  a := [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

  var b [15]int
  h2 := create(b[0:])
  for i := len(a)-1; i>=0; i-- {
    h2.insert(a[i])
    fmt.Printf("Inserting %d", a[i])
    fmt.Printf("%d\n", b[0:])
  }
  fmt.Printf("Now popping\n")
  for h2.length != 0 {
    fmt.Printf("Popped %d. Remaining %d\n", h2.pop(), b[0:])
  }

  fmt.Printf("Testing again")

  h := create(a[2:10])
  fmt.Println(a[0:])
  h.insert(5)
  h.insert(8)
  h.insert(666)
  h.insert(7)
  fmt.Println(a[0:])

  for h2.length != 0 {
    fmt.Printf("Popped %d. Remaining %d\n", h2.pop(), b[0:])
  }
}
