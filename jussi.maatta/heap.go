// A heap. Adapted from pseudocode examples presented in
// Kokkarinen & Ala-Mutka: Tietorakenteet ja algoritmit (2000)

package main

import (
	"fmt"
	"os"
)

type heap struct {
	data []int
	size int
}

func (h *heap) String() string {
	return fmt.Sprintf("size=%v data=%v", h.size, h.data[0:h.size])
}

func (h *heap) left(i int) int   { return 2*(i+1) - 1 }
func (h *heap) right(i int) int  { return 2 * (i + 1) }
func (h *heap) parent(i int) int { return (i - 1) / 2 }

func (h *heap) insert(e int) {
	i := h.size
	h.size++
	// climb up until a suitable position is found
	for i > 0 && h.data[h.parent(i)] < e {
		h.data[i] = h.data[h.parent(i)]
		i = h.parent(i)
	}
	h.data[i] = e
}

func (h *heap) pop() int {
	if h.size > 0 {
		x := h.data[0]
		h.data[0] = h.data[h.size-1]
		h.size--
		if h.size > 1 {
			h.heapify(0)
		}
		return x
	}
	// the wrong way to handle errors
	fmt.Fprintf(os.Stderr, "tried to pop from an empty heap\n")
	os.Exit(1)
	return 0 // to make the compiler happy
}

func (h *heap) heapify(i int) {
	j := i
	x := h.data[i]
	for h.left(j) < h.size {
		l := h.left(j)
		if h.right(j) < h.size && h.data[h.right(j)] > h.data[l] {
			l = h.right(j)
		}
		if h.data[l] > x {
			h.data[j] = h.data[l]
			j = l
		} else {
			h.data[j] = x
			return
		}
	}
	// if we get this far, the original root of our subtree will
	// become a leaf
	h.data[j] = x
}

func create(s []int) *heap { return &heap{s, 0} }

func main() {
	h := create(&[10]int{})

	for _, v := range []int{7, 2, 5, 9, 3} {
		fmt.Printf("adding %v to the heap\n", v)
		h.insert(v)
	}
	fmt.Println("popping twice")
	fmt.Println(h.pop())
	fmt.Println(h.pop())
	for _, v := range []int{1, 10, 4} {
		fmt.Printf("adding %v to the heap\n", v)
		h.insert(v)
	}
	fmt.Println("this is what the heap array looks like:")
	fmt.Println(h)
	fmt.Println("popping until there's nothing left...")
	for h.size > 0 {
		fmt.Println(h.pop())
	}
	// h.pop() // for an alternative ending
}
