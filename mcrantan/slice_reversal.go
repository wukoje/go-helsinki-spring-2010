package main

import "rand"
import "fmt"
import "time"

func rndSlice(length int) []int {
//  sl := make([]int, length)
//  for i := 0; i < len(sl); i++ {
//    sl[i] = rand.Intn(1023)
//    sl := rand.Perm(length)
    return rand.Perm(length)
}

func reverse(a []int) []int {
  if len(a) < 2 {
     return a
  }

// I take it that 'in-place' means no temporary variables

  for i := 0; i < len(a) / 2; i++ {
    a[i] ^= a[len(a) - i - 1]
    a[len(a) - i - 1] ^= a[i]
    a[i] ^= a[len(a) - i - 1]
  }

  return a
}

func main() {

  for reps := 0; reps < 4; reps++ {
    rand.Seed(time.Nanoseconds())
    length := rand.Intn(10) + 21
    a := rndSlice(length)
    fmt.Printf("  Random slice, length %d: ", len(a))
    for i := 0; i < len(a); i++ {
      fmt.Printf(" %d", a[i])
    }
    fmt.Printf("\n")

    reverse(a)
    fmt.Printf("Reversed slice, length %d: ", len(a))
    for i := 0; i < len(a); i++ {
      fmt.Printf(" %d", a[i])
    }
    fmt.Printf("\n")
  }
}

