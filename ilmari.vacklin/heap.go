package main

import "fmt"

type heap struct {
    a []int
    l int
}

func (h *heap) insert(e int) {
    i := h.l
    h.a[i] = e

    pi := int((i-1)/2)
    p := h.a[pi]

    for p > e {
        h.a[pi], h.a[i] = h.a[i], h.a[pi]
        pi = (pi-1)/2
        p = h.a[pi]
    }

    h.l++
}

func (h *heap) pop() int {
    r := h.a[0]
    i := h.l-1
    v := h.a[i]
    h.l = i

    i = 0
    h.a[i] = v
    h.a[h.l] = 0

    for {
        li := 2*i+1

        if li >= h.l { break }

        l := h.a[li]
        ri := 2*i+2
        r := h.a[ri]

        if v < l && v < r { break }

        if l < r || ri >= h.l {
            h.a[i], h.a[li] = h.a[li], h.a[i]
            i = li
        } else if ri <= h.l {
            h.a[i], h.a[ri] = h.a[ri], h.a[i]
            i = ri
        }
    }
    
    return r
}

func create(l []int) *heap {
    return &heap{ l, 0 }
}

func main() {
    x := create(new([10]int))
    fmt.Println(x)
    x.insert(5)
    fmt.Println(x)
    x.insert(3)
    fmt.Println(x)
    x.insert(6)
    fmt.Println(x)

    y := x.pop()
    fmt.Println(y)
    fmt.Println(x)
}
