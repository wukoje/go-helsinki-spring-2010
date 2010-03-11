package main

import "fmt"

type binomialTree struct {
	parent *binomialTree
	next   *binomialTree
	child  *binomialTree
	rank   int
	value  int
}

type binomialHeap struct {
	head *binomialTree
}

// merges the rootlists given as pointers and returns a pointer to the new
// rootlists head
func mergeRootLists(head1 *binomialTree, head2 *binomialTree) *binomialTree {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head *binomialTree
	if head1.rank <= head2.rank {
		head = head1
	} else {
		head = head2
	}

	for head1 != nil && head2 != nil {
		next1, next2 := head1.next, head2.next
		if head1.rank <= head2.rank {
			head1.next = head2
			head2.next = next1
		} else {
			head2.next = head1
			head1.next = next2
		}
		head1, head2 = next1, next2
	}

	return head
}

// links two binomial trees of equal rank, making t1 a child of t2
func binomialLink(t1,t2 *binomialTree) {
	t1.parent = t2
	t1.next = t2.child
	t2.child = t1
	t2.rank++
}

// adds contents of h2 into h1, destroying h2 in the process
func (h1 *binomialHeap) union(h2 *binomialHeap) {
	h1.head = mergeRootLists(h1.head, h2.head)
	if h1.head == nil {
		return
	}

	var prev *binomialTree = nil
	curr, next := h1.head, h1.head.next

	for next != nil {
		if curr.rank != next.rank ||
			(next.next != nil && next.rank == next.next.rank) {
			prev = curr
			curr = next
		} else if curr.value <= next.value {
			curr.next = next.next
			binomialLink(next, curr)
		} else {
			if prev == nil {
				h1.head = next
			} else {
				prev.next = next
			}
			binomialLink(curr, next)
			curr = next
		}
		next = curr.next
	}
	h2.head = nil //otherwise h2's head would point somewhere inside h1
}

// inserts i into heap h, returns a pointer to the node containing i for use
// with decrease key
func (h *binomialHeap) insert(i int) *binomialTree {
	head := &binomialTree{value: i, rank: 0}
	heap := &binomialHeap{head}
	h.union(heap)
	return head
}

// sets the parent of each node current's next-list to nil
func nilParents(current *binomialTree) {
	for current != nil {
		current.parent = nil
		current = current.next
	}
}

// reverses the next-pointers starting from current and returns the new head
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
	if h.head == nil {
		return 0, false
	}
	var prev, minPrev *binomialTree = nil, nil
	current, min := h.head, h.head

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
		h.head = min.next
	}

	minValue := min.value

	nilParents(min.child)
	head := reversePointers(min.child)
	heap := &binomialHeap{head}
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
	fmt.Println("inserting 5, 3, 7 and 15")
	heap.insert(5)
	heap.insert(3)
	heap.insert(7)
	n := heap.insert(15)
	ret, _ := heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ok := heap.decreaseKey(n, 1)
	fmt.Printf("decreaseKey 15->1, returned %v\n", ok)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
	ret, _ = heap.pop()
	fmt.Printf("pop: %v\n", ret)
}
