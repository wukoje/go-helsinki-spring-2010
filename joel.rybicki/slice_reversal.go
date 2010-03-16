package main

import fmt "fmt"

// reverse reverses the contents of s in place
func slice_reversal(s []int) {
  // This is a neat trick
  for start, end := 0, len(s)-1; start < end; start, end = start+1, end -1 {
    s[start], s[end] = s[end], s[start]
  }
}

func main() {
  // Output some slices before and after reversal.
  a := [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  fmt.Println( a )
  slice_reversal( a[0:len(a)] )
  fmt.Println( a )
  slice_reversal( a[5:len(a)] )
  fmt.Println( a )
  slice_reversal( a[0:5] )
  fmt.Println( a )
}
