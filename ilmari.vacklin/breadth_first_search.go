package main

import "fmt"
import "container/list"

type node struct {
    val string
    in, out *list.List
}

func create(val string) (*node) {
    n := new(node)
    n.val = val
    n.out = list.New()
    n.in = list.New()
    return n
}

func bdf(n *node, e string) (*node) {
    if n.val == e { return n }    

    q := list.New()

    for {
        for l := n.out.Front(); l != nil; l = l.Next() {
            q.PushBack(l.Value)
        }

        if q.Len() == 0 { break }

        l := q.Front()
        q.Remove(l)
        n = l.Value.(*node)
        
        if n.val == e { return n }
    }

    return nil
}

func main() {
    x := create("foo")
    y := create("bar")
    z := create("quux")
    x.out.PushFront(z)
    x.out.PushFront(y)
    fmt.Println(bdf(x, "quux"))
    fmt.Println(bdf(x, "froob")) // should be nil
}
