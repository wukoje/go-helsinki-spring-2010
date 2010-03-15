package main

import "fmt"

func search(s []int, e int) int {
    l := len(s)

    if l == 0 { return -1 }

    pi := l / 2
    p := s[pi]

    if e == p { return pi }
    if e < p { return search(s[0:pi], e) }
    if e > p { return pi + search(s[pi:],  e) }

    return -1
}

func main() {
    fmt.Println(search([]int{ 5, 6, 7, 8, 9 }, 5)) // should be 0
    fmt.Println(search([]int{ 5, 6, 7, 8, 9 }, 6)) // should be 1
    fmt.Println(search([]int{ 5, 6, 7, 8, 9 }, 7)) // should be 2
    fmt.Println(search([]int{ 5, 6, 7, 8, 9 }, 8)) // should be 3
    fmt.Println(search([]int{ 5, 6, 7, 8, 9 }, 9)) // should be 4
   
    fmt.Println(search([]int{ 5 }, 4)) // should be -1
    fmt.Println(search([]int{ 5 }, 5)) // should be 0

    fmt.Println(search([]int{ 5, 6 }, 6)) // should be 1
}
