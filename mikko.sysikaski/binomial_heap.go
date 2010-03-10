package main

import (
	"fmt"
)

// Public API

type LessInterface interface {
	Less(x interface{}) bool
}
// Binomial heap
type BinHeap struct {
	head *binTree
}
// Reference to an element in heap; used as arguments to DecreaseKey
type HeapRef **binTree

// Adds a new element to heap
func (h *BinHeap) Insert(x LessInterface) HeapRef {
	add := &binTree{nil, x, 0, nil, nil, nil}
	add.selfRef = &add
	h.union(BinHeap{add})
	return add.selfRef
}
// Removes and returns the smallest element of the heap
func (h *BinHeap) Pop() LessInterface {
	// find the smallest element
	cur := h.head
	small := cur
	for cur = cur.next; cur != nil; cur = cur.next {
		if cur.key.Less(small.key) {
			small = cur
		}
	}
	if small == h.head {
		h.head = small.next
	} else {
		pred := h.head
		for pred.next != small {
		}
		pred.next = small.next
	}

	// reverse child-list of the smallest element
	var px *binTree
	for x := small.child; x != nil; {
		x.parent = nil
		n := x.next
		x.next = px
		px = x
		x = n
	}
	// merge child list with root list
	h.union(BinHeap{px})

	return small.key
}
// Decreases the value of the element r to the new value x
func (h *BinHeap) DecreaseKey(r HeapRef, x LessInterface) {
	cur := *r
	cur.key = x
	for cur.parent != nil && cur.key.Less(cur.parent.key) {
		p := cur.parent
		cur.key, p.key = p.key, cur.key
		*cur.selfRef = p
		*p.selfRef = cur
		cur.selfRef, p.selfRef = p.selfRef, cur.selfRef
		cur = p
	}
}

// Internal types and methods

type binTree struct {
	parent *binTree
	key    LessInterface
	degree int
	child  *binTree
	next   *binTree

	// Used to maintain references to nodes during DecreaseKey operations.
	selfRef **binTree
}

func (t *binTree) setParent(p *binTree) {
	t.parent = p
	t.next = p.child
	p.child = t
	p.degree++
}
func merge(a, b *binTree) *binTree {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	if a.degree < b.degree {
		a.next = merge(a.next, b)
		return a
	}
	b.next = merge(a, b.next)
	return b
}

func (h *BinHeap) union(b BinHeap) {
	h.head = merge(h.head, b.head)
	if h.head == nil {
		return
	}

	cur := h.head
	var prev *binTree
	for cur.next != nil {
		next := cur.next
		if cur.degree < next.degree || (next.next != nil && next.degree == next.next.degree) {
			prev = cur
			cur = next
		} else {
			if cur.key.Less(next.key) {
				cur.next = next.next
				next.setParent(cur)
			} else {
				if prev == nil {
					h.head = next
				} else {
					prev.next = next
				}
				cur.setParent(next)
				cur = next
			}
		}
	}
}


// Test code

func printTree(t *binTree, d int) {
	for ; t != nil; t = t.next {
		for i := 0; i < d; i++ {
			fmt.Print("\t")
		}
		fmt.Println(t.key)
		printTree(t.child, d+1)
	}
}

type Int int

func (x Int) Less(y interface{}) bool { return int(x) < int(y.(Int)) }
func printTop(h *BinHeap) {
	fmt.Println("Popping min element:", h.Pop())
	fmt.Println("Heap after pop:")
	printTree(h.head, 0)
}
func add(h *BinHeap, x int) HeapRef {
	fmt.Println("Adding", x)
	r := h.Insert(Int(x))
	fmt.Println("Heap after add:")
	printTree(h.head, 0)
	return r
}
func degKey(h *BinHeap, r HeapRef, x int) {
	fmt.Printf("Decreasing key from %d to %d\n", (*r).key, x)
	h.DecreaseKey(r, Int(x))
	fmt.Println("Heap after decrease:")
	printTree(h.head, 0)
}

func main() {
	h := new(BinHeap)
	r := add(h, 2)
	r2 := add(h, 5)
	add(h, 1)
	add(h, 4)
	printTop(h)
	degKey(h, r2, 1)
	degKey(h, r, 0)
	printTop(h)
	printTop(h)
	printTop(h)
}
