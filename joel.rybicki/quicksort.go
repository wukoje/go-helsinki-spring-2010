package main

import fmt "fmt"

func partition(s []int, pivot int) int {
  x := s[pivot] // store the pivot value
  right := len(s)-1
  s[pivot], s[right] = s[right], s[pivot]
  j := 0
  for i := 0; i<right; i++ {
    if s[i] <= x {
        s[i], s[j] = s[j], s[i]
        j++
    }
  }
  s[j], s[right] = s[right], s[j]
  return j
}

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
  n := len(s)
  if (n <= 1) { 
    return 
  }
  pivot := n / 2
  pivot = partition(s, pivot)
  sort(s[0:pivot])
  sort(s[pivot:])
}

func main() {
  a := [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
  fmt.Println( a )
  sort(a[0:])
  fmt.Println( a )
  b := [...]int {5, 2, 14, 3, 58, -4, 123, 8, 9, 7, -11, 12}
  fmt.Println( b )
  sort(b[0:])
  fmt.Println( b )
}
