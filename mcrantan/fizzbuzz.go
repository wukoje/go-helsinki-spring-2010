package main

import "fmt"

func main() {

// Declaring j in 'if' makes j visible only within the 'if'

  var j int = 0

  for i:= 1; i <= 100; i++ {

// If there was tertiary operator ( ? and : from C), this
// could be one-liner

    if j = i % 3; j == 0 {
      fmt.Printf("Fizz")
    }

    if k := i % 5; k == 0 {
      fmt.Printf("Buzz\n")
      continue
    }

    if j != 0 {
      fmt.Printf("%d", i)
    }

    fmt.Printf("\n")
  } 
}
