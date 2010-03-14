/**
 * binomial_heap.go: A binomial heap implementation customized from CLRS (ch. 19)
 * http://mitpress.mit.edu/algorithms/
 */

package bheap

import (
	"math"
)

type Bnode struct {
	deg int
	parent *Bnode
	sibling *Bnode
	child *Bnode

	key int
	value interface{}
}

type Bheap struct {
	head *Bnode
}

// Minimum returns the smallest element from heap
func (h *Bheap) Minimum() *Bnode {
	n, _ := minimum(h)
	return n
}

// Insert adds the given node into heap
func (h *Bheap) Insert(n *Bnode) {
	h.head = union(h, n)
}

// Delete removes the given node from heap
func (h *Bheap) Delete(n *Bnode) {
	h.DecreaseKey(n, math.MinInt32)
	h.ExtractMinimum()
}

// DecreaseKey decreases the key on given node
func (h *Bheap) DecreaseKey(n *Bnode, key int) {
	if key > n.key {
		panic("DecreaseKey: New key greater then current key (",
			key, " > ", n.key, ")")
	}

	n.key = key
	this := n
	p := this.parent
	for p != nil && this.key < p.key {
		// swap p & this
		p.key, this.key = this.key, p.key
		p.value, this.value = this.value, p.value

		this = p
		p = this.parent
	}
}

// ExtractMinimum removes and returns the smallest element from heap
func (h *Bheap) ExtractMinimum() *Bnode {
	node, prev := minimum(h)

	if node == nil {
		return nil
	}

	if prev != nil {
		prev.sibling = node.sibling
	} else {
		h.head = node.sibling
	}

	// build a new binomial tree from n's children (in reverse)
	var ntree *Bnode
	child := node.child
	for child != nil {
		next := child.sibling
		child.sibling = ntree
		ntree = child
		child = next
	}

	h.head = union(h, ntree)
	return node
}

// make c a child of p
func link(c *Bnode, p *Bnode) {
	c.parent = p
	c.sibling = p.child
	p.child = c
	p.deg++
}

// merges t to root list of h and returns the new root list
func merge(h *Bheap, t *Bnode) *Bnode {
	var head, this, chosen *Bnode
	n1 := h.head
	n2 := t

	for n1 != nil && n2 != nil {
		if n1.deg < n2.deg {
			chosen = n1
			n1 = n1.sibling
		} else {
			chosen = n2
			n2 = n2.sibling
		}

		if head == nil {
			head = chosen
			this = head
		} else {
			this.sibling = chosen
			this = this.sibling
		}
	}

	if n1 != nil {
		chosen = n1
	} else {
		chosen = n2
	}

	if head == nil {
		head = chosen
	} else {
		this.sibling = chosen
	}

	return head
}

// unites bheap with given tree
func union(h *Bheap, t *Bnode) *Bnode {
	head := merge(h, t)
	if head == nil {
		return nil
	}

	var prev *Bnode
	this := head
	next := this.sibling
	for next != nil {
		if this.deg != next.deg ||
			(next.sibling != nil && next.sibling.deg == this.deg) {
			prev = this
			this = next
		} else {
			if this.key <= next.key {
				this.sibling = next.sibling
				link(next, this)
			} else {
				if prev == nil {
					head = next
				} else {
					prev.sibling = next
				}
				link(this, next)
				this = next
			}
		}
		next = this.sibling
	}
	return head
}


// lookup minimum and it's previous sibling
func minimum(h *Bheap) (node *Bnode, prev *Bnode) {
	if h.head == nil {
		return
	}

	min := math.MaxInt32
	this := h.head
	node = h.head
	var tprev *Bnode
	for this != nil {
		if this.key < min {
			min = this.key
			prev = tprev
			node = this
		}
		tprev = this
		this = this.sibling
	}
	return
}
