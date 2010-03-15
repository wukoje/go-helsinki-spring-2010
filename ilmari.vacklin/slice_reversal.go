package main

import "fmt"

func reverse(s []int) {
    l := len(s) - 1
    for i := 0; i < l / 2 + 1; i++ {
        s[l-i], s[i] = s[i], s[l-i]
    }
}

func main() {
    x := []int{ 5, 6, 7, 8, 9 }
    fmt.Println(x)
    reverse(x)
    fmt.Println(x)

    y := []int{ 6, 7, 8, 9 }
    fmt.Println(y)
    reverse(y)
    fmt.Println(y)

    z := []int{ 8, 9 }
    fmt.Println(z)
    reverse(z)
    fmt.Println(z)

    ö := []int{ 9 }
    fmt.Println(ö)
    reverse(ö)
    fmt.Println(ö)
}
