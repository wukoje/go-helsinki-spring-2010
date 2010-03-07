package main

import "fmt"

// swaps s[a] and s[b]
func swap(s []int, a int, b int) {
    s[a], s[b] = s[b], s[a]
}

// reverses the slice given as a parameter
func reverse(s []int) {
    for b, e := 0, len(s) - 1; b < e; b, e = b+1, e-1 {
        swap(s, b, e)
    }
}

func main() {
    slice := []int{1,2,3,4,5}
    fmt.Printf("original: %v\n", slice)
    reverse(slice)
    fmt.Printf("reversed: %v\n", slice)

    slice = []int{}
    fmt.Printf("original: %v\n", slice)
    reverse(slice)
    fmt.Printf("reversed: %v\n", slice)
}
