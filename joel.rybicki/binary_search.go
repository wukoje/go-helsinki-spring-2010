package main

import (
  "fmt"
)

// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
  l, r := 0, len(s)-1
  for l <= r {
    mid := int((l + r) / 2)
    if s[mid] == e {
        return mid
    }
    if s[mid] < e {
        l = mid+1
    } else {
        r = mid-1
    }
  }
  return -1
}

func main() {
  // Find all the values in the array
  a := [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  for i, val := range a {
    if search(a[0:], val) == i {
        fmt.Printf("Found a[%d] contained %s as supposed\n.", i, val)
    } else {
        fmt.Printf("ERROR a[%d] != %d.\n", i, val)
    }
  }
  
  if search(a[0:], 11) == -1 {
    fmt.Printf("OK, does not report non-existing values.")
  } else {
    fmt.Printf("Oh my, found non existing-values.")
  }
}
