package main

import (
  "fmt"
  "rand"
  "time"
)

func main() {

  a := []int{42}
  tstSort(a)

  a = []int{3,8}
  tstSort(a)

  a = []int{8,3}
  tstSort(a)
  

  for reps := 0; reps < 4; reps++ {
    rand.Seed(time.Nanoseconds())
    length := rand.Intn(10) + 21
    a := rndSlice(length)
    fmt.Printf("     Random slice, length %d: ", len(a))
    for i := 0; i < len(a); i++ {
      fmt.Printf(" %d", a[i])
    }
    fmt.Printf("\n")

    sort(a)
    fmt.Printf("     Sorted slice, length %d: ", len(a))
    for i := 0; i < len(a); i++ {
      fmt.Printf(" %d", a[i])
    }
    fmt.Printf("\n")
  }
}

func tstSort(a []int) {
  fmt.Printf("  Non-random slice, length %d: ", len(a))
  for i := 0; i < len(a); i++ {
    fmt.Printf(" %d", a[i])
  }
  fmt.Printf("\n")

  sort(a)
  fmt.Printf("      Sorted slice, length %d: ", len(a))
  for i := 0; i < len(a); i++ {
    fmt.Printf(" %d", a[i])
  }
  fmt.Printf("\n")
}

func rndSlice(length int) []int {
    return rand.Perm(length)
}
 

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
// Please see http://golang.org/src/pkg/sort/sort.go for better implementation
// (i.e. use other sort methods for short arrays, choose pivot wiser)

func sort(s []int) {

  if len(s) < 2 {
    return
  }

  if len(s) == 2 {
    if s[0] > s[1] {
      temp := s[1]
      s[1] = s[0]
      s[0] = temp
      return
    }
  }
  left := 0
  right := len(s) - 1
  quicksort(s, left, right)
}

func quicksort(array []int, left int, right int) {

  if right > left {
    pivotIndex := (left + right) / 2
    pivotNewIndex := partition(array, left, right, pivotIndex)
    quicksort(array, left, pivotNewIndex - 1)
    quicksort(array, pivotNewIndex + 1, right)
  }
}

func partition(array []int, left int, right int, pivotIndex int) int {

  var temp int
  pivotValue := array[pivotIndex]
  array[pivotIndex] = array[right]
  array[right] = pivotValue
  storeIndex := left
  for i := left; i <= right - 1; i++ { // left <= i < right
    if array[i] <= pivotValue {
      temp = array[i]
      array[i] = array[storeIndex]
      array[storeIndex] = temp
      storeIndex++
    }
  }
  temp = array[storeIndex]
  array[storeIndex] = array[right]
  array[right] = temp
  return storeIndex
}

