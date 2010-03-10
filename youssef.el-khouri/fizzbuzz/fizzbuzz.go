package main

import fmt "fmt"

func main() { 
  for i := 1; i <= 100; i++ {
    if i % 3 == 0 || i % 5 == 0 {
      if i%3 == 0 {
	fmt.Printf("fizz")
      }
      if i%5 == 0 {
	fmt.Printf("buzz")
      }
    } else {
      fmt.Printf("%d", i)
    }
    fmt.Printf("\n")
  }
}
