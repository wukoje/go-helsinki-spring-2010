package main

import "fmt"


// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
  start := 0
  end := len(s) - 1
  middle := len(s) / 2

  ret_index := -1

  for start <= end {
    middle = (start + end) / 2

    if s[middle] == e {
      ret_index = middle
      break
    } else if e < s[middle] {
      end = middle -1
    } else {
      start = middle + 1
    } 
    
  }
  return ret_index
}

func main() {
  slice1 := []int{1,2,3,4,7,11,14,18}
  slice2 := []int{-23,-12,0,22,23,66,69,78,99}
 
  e1 := 7
  e2 := 23
  e3 := -12
  
  fmt.Println("First slice:")
  fmt.Println(slice1)
  fmt.Printf("Searching for %d\n", e1)
  fmt.Printf("Index for %d is: %d\n", e1, search(slice1, e1))
  fmt.Printf("Searching for %d\n", e2)
  fmt.Printf("Index for %d is: %d\n", e2, search(slice1, e2))
  fmt.Println("Second slice:")
  fmt.Println(slice2)
  fmt.Printf("Searching for %d\n", e1)
  fmt.Printf("Index for %d is: %d\n", e1, search(slice2, e1))
  fmt.Printf("Searching for %d\n", e2)
  fmt.Printf("Index for %d is: %d\n", e2, search(slice2, e2))
  fmt.Printf("Searching for %d\n", e3)
  fmt.Printf("Index for %d is: %d\n", e3, search(slice2, e3))
  
  
}
