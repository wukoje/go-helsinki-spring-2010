package main

import "fmt"

type binomialTree struct {
	parent *binomialTree
	next   *binomialTree
	child  *binomialTree
	degree int
	value  int
}

type binomialHeap struct {
	cap *binomialTree
}

// merges the rootlists given as pointers and returns a pointer to the new
// rootlists cap
func mergeRootLists(cap1 *binomialTree, cap2 *binomialTree) *binomialTree {
	if cap1 == nil {
		return cap2
	}
	if cap2 == nil {
		return cap1
	}

	var cap *binomialTree
	if cap1.degree <= cap2.degree {
		cap = cap1
	} else {
		cap = cap2
	}

	for cap1 != nil && cap2 != nil {
		next1, next2 := cap1.next, cap2.next
		if cap1.degree <= cap2.degree {
			cap1.next = cap2
			cap2.next = next1
		} else {
			cap2.next = cap1
			cap1.next = next2
		}
		cap1, cap2 = next1, next2
	}

	return cap
}

// links two binomial trees of equal degree, making t1 a child of t2
func binomialLink(t1, t2 *binomialTree) {
	t1.parent = t2
	t1.next = t2.child
	t2.child = t1
	t2.degree++
}

// adds contents of h2 into h1, destroying h2 in the process
func (h1 *binomialHeap) union(h2 *binomialHeap) {
	h1.cap = mergeRootLists(h1.cap, h2.cap)
	if h1.cap == nil {
		return
	}

	var prev *binomialTree = nil
	curr, next := h1.cap, h1.cap.next

	for next != nil {
		if curr.degree != next.degree ||
			(next.next != nil && next.degree == next.next.degree) {
			prev = curr
			curr = next
		} else if curr.value <= next.value {
			curr.next = next.next
			binomialLink(next, curr)
		} else {
			if prev == nil {
				h1.cap = next
			} else {
				prev.next = next
			}
			binomialLink(curr, next)
			curr = next
		}
		next = curr.next
	}
	h2.cap = nil //otherwise h2's cap would point somewhere inside h1
}

// inserts i into heap h, returns a pointer to the node containing i for use
// with decrease key
func (h *binomialHeap) insert(i int) *binomialTree {
	cap := &binomialTree{value: i, degree: 0}
	heap := &binomialHeap{cap}
	h.union(heap)
	return cap
}

// sets the parent of each node current's next-list to nil
func nilParents(current *binomialTree) {
	for current != nil {
		current.parent = nil
		current = current.next
	}
}

// reverses the next-pointers starting from current and returns the new cap
func reversePointers(current *binomialTree) *binomialTree {
	if current == nil {
		return nil
	}

	var prev *binomialTree = nil

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

// removes the smallest element e from h and returns it
// ok is false if h was empty
func (h *binomialHeap) pop() (e int, ok bool) {
	if h.cap == nil {
		return 0, false
	}
	var prev, minPrev *binomialTree = nil, nil
	current, min := h.cap, h.cap

	for current != nil {
		if current.value < min.value {
			minPrev = prev
			min = current
		}
		prev = current
		current = current.next
	}

	if minPrev != nil {
		minPrev.next = min.next
	} else {
		h.cap = min.next
	}

	minValue := min.value

	nilParents(min.child)
	cap := reversePointers(min.child)
	heap := &binomialHeap{cap}
	h.union(heap)

	return minValue, true
}

// decreases the value of n in h
// returns false if i is larger than the current value of n or n is nil,
// otherwise returns true
func (h *binomialHeap) decreaseKey(n *binomialTree, i int) (ok bool) {
	if n == nil || i > n.value {
		return false
	}
	n.value = i
	for n.parent != nil && n.value < n.parent.value {
		n.value, n.parent.value = n.parent.value, n.value
		n = n.parent
	}
	return true
}

func main() {
	heap := new(binomialHeap)
	fmt.Println("inserting 25, 33, 67 and 75")
	heap.insert(25)
	heap.insert(33)
	heap.insert(67)
	n := heap.insert(75)
	ret, _ := heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ok := heap.decreaseKey(n, 1)
	fmt.Printf("decreaseKey 75->1, returned %v\n", ok)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
}

