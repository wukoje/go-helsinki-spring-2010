package main

import "fmt"

func sort(s []int) {
    quicksort(s, 0, len(s)-1)
}

func quicksort(s []int, left int, right int) {
    if right < left { return }

    pi := (left + right) / 2
    pin := partition(s, left, right, pi)
    quicksort(s, left, pin-1)
    quicksort(s, pin+1, right)
}

func partition(s []int, left int, right int, pi int) int {
    p := s[pi]
    s[pi], s[right] = s[right], s[pi]
    si := left
    for i := left; i < right; i++ {
        if s[i] <= p {
            s[i], s[si] = s[si], s[i]
            si++
        }
    }
    s[si], s[right] = s[right], s[si]
    return si
}

func main() {
    x := []int{ 5,0,2,4,9,2,3 }
    sort(x)
    fmt.Println(x)
}
