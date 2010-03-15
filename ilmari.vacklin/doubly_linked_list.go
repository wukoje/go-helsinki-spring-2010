package main

import "fmt"

type list struct {
    e string
    next, prev *list
}

func (l *list) String() string {
    s := "\"" + l.e + "\""

    if l.next != nil {
        s += ", " + l.next.String()
    }

    return s
}

func (l *list) insert(i int, e string) {
    if i == 0 && l.next == nil && l.prev == nil { // singleton
        l.e = e
        return
    }

    for i > 0 {
        if l.next == nil { 
            l.next = new(list) 
            l.next.prev = l
        }
        l = l.next
        i = i - 1
    }

    m := new(list)
    m.next = l
    m.prev = l.prev
    m.next.prev = m
    m.prev.next = m
    m.e = e
}

func (l *list) delete(i int) {
    for i > 0 {
        if l.next == nil { return }
        l = l.next
        i = i - 1
    }

    l.prev.next = l.next
    if l.next != nil { l.next.prev = l.prev }
}

func main() {
    x := new(list)
    fmt.Println(x)
    x.insert(0, "foo")
    fmt.Println(x)
    x.insert(1, "bar")
    fmt.Println(x)
    x.insert(3, "quux")
    fmt.Println(x)
    x.insert(2, "duck")
    fmt.Println(x)

    x.delete(1)
    fmt.Println(x)
}
