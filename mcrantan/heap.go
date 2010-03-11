package main

import (
  "fmt"
  "os"
  "rand"
  "time"
  "flag"
)

func main() {

  var SIZE int
  flag.IntVar(&SIZE, "s", 47, "Initial heap size (default: 47)")
  flag.Parse()

  rand.Seed(time.Nanoseconds())
  intarr := make([]int, SIZE)
  myheap := create(intarr)
  fmt.Printf("Created heap for %d ints.\n", SIZE)

  fmt.Printf("Inserting value '42'.\n")
  myheap.insert(42)

  foo := myheap.pop()
  fmt.Printf("Popped value: %d.\n", foo)

  fmt.Printf("Inserting %d random values in range 0..1023 into heap:", SIZE)
  for i := 0; i < SIZE; i++ {
    j := rand.Intn(1024)
    myheap.insert(j)
    fmt.Printf(" %d", j)
  }
  fmt.Printf("\n")

  fmt.Printf("\nAttempting to insert one more value(over heap capacity): ")
  myheap.insert(42)

  fmt.Printf("\nPopping all values from heap:")
  for i :=0; i < SIZE + 1; i++ {
    fmt.Printf(" %d", myheap.pop())
  }
}

type heap struct {
  size int
  array []int
  tail int
}

func create(s []int) *heap {

  if len(s) < 2 {
    fmt.Printf("\nA heap smaller than two items? Makes sense, none.\n\n")
//    return nil
    os.Exit(0)
  }

    return &heap{len(s), s, -1}
} 

func (h *heap) insert(e int) {

  h.tail += 1
  if h.tail == h.size {
    fmt.Printf("Heap is full: can't insert value %d into heap.\n", e)
    h.tail -= 1
    return
  }

  h.array[h.tail] = e
  heapifyUp(h, h.tail)
}


func (h *heap) pop() int {

  if h.tail < 0 {
    fmt.Printf("\nHeap is empty, exiting.\n")
    os.Exit(0)
  }

  retval := h.array[0]

  if h.tail > 0 {
    h.array[0] = h.array[h.tail]
    heapifyDown(h, 0)
  }

  h.tail -= 1
  return retval
}

func heapifyUp (h *heap, i int) {

  if h.array[(i - 1) / 2] > h.array[i] {
    h.array[(i - 1) / 2], h.array[i] = h.array[i], h.array[(i - 1) / 2]
    i = (i - 1) / 2
    heapifyUp(h, i)
  }
}

func heapifyDown (h *heap, i int) {

  left := 2 * i + 1
  right := 2 * i + 2
  smallest := 0
 
  if left <= h.tail && h.array[left] < h.array[i] {
    smallest = left
  } else {
    smallest = i
  }

  if right <= h.tail && h.array[right] < h.array[smallest] {
    smallest = right
  }

  if smallest != i {
    h.array[i], h.array[smallest] = h.array[smallest], h.array[i]
    heapifyDown(h, smallest)
  }
}

// For debugging, if you please.
func (h *heap) String() {

  fmt.Printf("heap.tail=%d, heap.size=%d, len(heap.array)=%d\n", h.tail, h.size, len(h.array))
  for i := 0; i < h.size; i++ {
    fmt.Printf(" %d", h.array[i])
  }
}
