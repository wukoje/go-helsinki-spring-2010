package main

import "fmt"

func reverse(s []int) {
  start := 0
  end := len(s) - 1

  for start < end {
    s[start], s[end] = s[end], s[start]
    start += 1
    end -= 1  
  }
}

func main() {
  example := []int{1, 2, 3, 5, 18, 22}
  
  fmt.Println("Initial slice:")
  fmt.Println(example)
  
  reverse(example)
  
  fmt.Println("Reversed slice:")
  fmt.Println(example)
}
