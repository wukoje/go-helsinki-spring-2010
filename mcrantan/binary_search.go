package main

import (
  "rand"
  "fmt"
  "time"
  "sort"
)

func rndSlice(length int) []int {
    return rand.Perm(length)
}

func search(s []int, e int) int {

  var pos, chval, j, x int
  pos = (len(s) - 1) / 2
  chval = pos / 2 
  j = 1 
  x = 1

// There is no power or logarithm function, so figure out log
  for j = 1; x < len(s); j++ {
    x = x << 1
    if x > len(s) {
      break
    }
  }

  for i := 0; i <= j; i++ {
    if s[pos] < e {
      pos += chval
    }

    if s[pos] > e {
      pos -= chval
    }

    if s[pos] == e {
      return pos
    }

    chval /= 2

    if chval == 0 {
      chval +=1      // chval might become 0
    }

// fmt.Printf("DEBUG: pos=%d, chval=%d\n", pos, chval)
  }

  return -1;
}

func main() {

  for reps := 0; reps < 4; reps++ {
    rand.Seed(time.Nanoseconds())
    length := rand.Intn(100000)
    a := rndSlice(length)

    fmt.Printf("  Initialized random slice, length %d\n", len(a))

    b := sort.IntArray(a)
    sort.Sort(b)

// From the tutorial
    if !sort.IsSorted(b) {
      panic()
    }

    valtofind := rand.Intn(length)  // Small chance of choosing too large value
    if retval := search(b, valtofind); retval < 0 {
      fmt.Printf("Value %d not found.\n", valtofind)
    } else {
      fmt.Printf("Value %d found on position %d.\n", valtofind, retval)
    }

    valtofind = length + 1

    if retval := search(b, valtofind); retval < 0 {
      fmt.Printf("Value %d not found.\n", valtofind)
    } else {
      fmt.Printf("Value %d found on position %d.\n", valtofind, retval)
    }

    fmt.Printf("\n")
  }
}

