package main

import "fmt"

// sorts s in place with quicksort
func sort(s []int) {

  if len(s) > 1 {
    pivot := len(s) / 2
    newpivot := partition(s, pivot)
    sort(s[0:newpivot])
    sort(s[newpivot+1:len(s)])
  }
}

// partition function for quicksort
func partition(s []int, pivot int) int {  
  left := 0
  right := len(s) - 1

  pivotval := s[pivot]
  // move pivot out of the way to the end
  s[pivot], s[right] = s[right], s[pivot]

  storeind := left  
  for i := left; i < right; i++ {
    if s[i] <= pivotval {
      s[i], s[storeind] = s[storeind], s[i]
      storeind += 1
    }
  }
  // move pivot back to its place
  s[storeind], s[right] = s[right], s[storeind]  
  return storeind
}

func main() {

  var example = []int { 2,4,5,6,7,8,4,1,22,13,124,35,65,3,21,55 }

  fmt.Println("Example array:")
  fmt.Println(example)
  sort(example)
  fmt.Println("After sorting:")
  fmt.Println(example)
}
