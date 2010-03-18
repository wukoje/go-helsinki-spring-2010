package main

import (
	"fmt"
)

// numbers must be in numerical order.
func search(s []int, e int) int {
    n := len(s)-1
    low := 0
    high := n
    for low < high {
        mid := low + ((high - low) / 2)
        if (s[mid] < e) {
            low = mid + 1
        } else {
            high = mid
        }
    }
    var i int
    if ((low < n) && (s[low] == e)) {
        i = low
    } else {
        i = -1
    }
    return i
}

func main() {
    s := make([]int, 10)
    for i,_ := range s {
        s[i] = i*2
    }
    fmt.Println(s)
    fmt.Println(search(s, 0))
    fmt.Println(search(s, 2))
    fmt.Println(search(s, 8))
    fmt.Println(search(s, 20))
}
