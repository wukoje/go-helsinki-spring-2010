package main

import "fmt"

// an array-based implementation of Minimum Binary Heap
type heap struct {
  storage []int
  size int
}


// insert adds element e to the heap
// returns true if insert succeeded, false if not enough space
func (h *heap) insert(e int) bool {

  // are we out of space?
  if h.size == len(h.storage) {
    // get some more from the underlying array
    if len(h.storage) < cap(h.storage) {
      h.storage = h.storage[0:len(h.storage)+1]
    // oops, the array is full too!
    } else {
      return false
    }
  }
  // add new element to the end
  h.storage[h.size] = e
  curr_ind := h.size
  h.size++
  
  // let's find the correct place for our new element
  parent := get_parent(curr_ind)
  for h.storage[curr_ind] < h.storage[parent] {
    h.storage[curr_ind], h.storage[parent] = h.storage[parent], h.storage[curr_ind]
    curr_ind = parent
    parent = get_parent(curr_ind)
  }
  return true
}

// pop returns the smallest element from the heap
func (h *heap) pop() int {

  // pop the root
  popped := h.storage[0]
  h.size--
  
  // move the last element to root
  h.storage[0] = h.storage[h.size]
  curr_ind := 0  
  var smaller_child int
  if h.storage[1] <= h.storage[2] {
    smaller_child = 1
  } else {
    smaller_child = 2
  }

  // sift new root down
  for h.storage[curr_ind] > h.storage[smaller_child] {
    h.storage [curr_ind], h.storage[smaller_child] = h.storage[smaller_child], h.storage [curr_ind]
    curr_ind = smaller_child
    // element has two children, must compare
    if (2*curr_ind+2) < h.size {
      if h.storage[2*curr_ind + 1] < h.storage[2*curr_ind + 2] {
        smaller_child = 2*curr_ind + 1
      } else {
        smaller_child = 2*curr_ind + 2
      }
    // elem has only left child
    } else if (2*curr_ind+2) == h.size {
      smaller_child = 2*curr_ind + 1
    // elem is leaf
    } else {
      break
    }
  }
  return popped
}

func get_parent(index int)  int {
  return (index - 1) / 2
}

// create takes the given slice and creates a heap type that uses the slice for storage

func create(s []int) *heap {
  
  var new_heap *heap = new(heap)
  new_heap.storage = s
  new_heap.size = 0
  return new_heap
}


func main() {
  var example [20]int
  ex_heap := create(example[0:16])
  var succ bool
  for i := 1; i <= 21; i++ {
    succ = ex_heap.insert(i)
    if !succ {
      fmt.Println("heap insert failed")
    } else {
      fmt.Printf("inserted %d\n", i)
    }
  }
  fmt.Printf("Current heap size: %d\n", ex_heap.size)
  for i := 0; i < 10; i++ {
    popped := ex_heap.pop()
    fmt.Printf("Popped: %d; heap size: %d\n", popped, ex_heap.size)
  }   
}
