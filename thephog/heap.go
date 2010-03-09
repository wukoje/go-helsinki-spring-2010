package main

import "fmt"
import "os"

func swap(s []int, index1 int, index2 int) { 
     s[index1], s[index2] = s[index2], s[index1] 
}

type heap struct {
	data []int
	size int
}

func (h *heap) left(index int) int { return (index+1)*2 - 1 }

func (h *heap) right(index int) int { return (index + 1) * 2 }

func (h *heap) parent(index int) int { return (index+1)/2 - 1 }

func (h *heap) printSubTree(index int, depth int) {
	if index >= h.size {
		return
	}

	h.printSubTree(h.left(index), depth+1)

	for i := 0; i < depth; i++ {
		fmt.Printf("\t")
	}
	fmt.Printf("%d \n", h.data[index])

	h.printSubTree(h.right(index), depth+1)

}

func (h *heap) grow() {
	newData := make([]int, 2*len(h.data))

	for i, v := range h.data {
		newData[i] = v
	}

	h.data = newData

}

func (h *heap) heapify(index int) {
	for {

		left := h.left(index)
		right := h.right(index)
		target := index

		if left < h.size && h.data[left] > h.data[target] {
			target = left
		}

		if right < h.size && h.data[right] > h.data[target] {
			target = right
		}

		if index == target {
			break
		}

		swap(h.data, index, target)
		index = target
	}

}

func (h *heap) insert(e int) {
	if h.size <= len(h.data) {
		h.grow()
	}

	i := h.size
	for i > 0 && h.data[h.parent(i)] < e {
		h.data[i] = h.data[h.parent(i)]
		i = h.parent(i)
	}

	h.data[i] = e
	h.size++

}

func (h *heap) pop() int {
	if h.size == 0 {
	   	 fmt.Printf("CANNOT POP: HEAP ALREADY EMPTY\n")
	   	 os.Exit(0)
	}

	h.size--
	max := h.data[0]
	h.data[0] = h.data[h.size]
	h.heapify(0)

	return max
}

func main() {
	h := heap{make([]int, 4), 0}
	unordered := []int{8, 1337, 5, 100, 66, 2, 32, 19, 89, 10, 11, 17, 27, 44, 88}

	fmt.Printf("Inserting numbers...\n")
	for _, value := range unordered {
		fmt.Printf("%d, ", value)
		h.insert(value)
	}

	h.printSubTree(0, 0)

	fmt.Printf("\nPopping numbers..\n")
	for {
		value := h.pop()
		fmt.Printf("%d, ", value)
	}

}
