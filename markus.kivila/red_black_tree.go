/*
2 red_black_tree.go
Implement a red-black tree that support search, insert and delete.

Based on C code at http://en.literateprograms.org/Red-black_tree_(C)
*/
package main

import (
	"fmt"
	"rand"
	"strings"
)

type node struct {
	black bool
	p *node
	rc *node
	lc *node
	data elem
}

type elem interface {
	less(e elem) bool
	equals(e elem) bool
	key() int  // these two could be more generic!?
	val() string
}

type rbtree struct {
	root *node
}

// standard search tree insertion
// returns the new node 
func (t *rbtree) treeInsert(data elem) *node {
	if t.root == nil {
		t.root = &node{true, nil, nil, nil, data}
		fmt.Printf("DBG: Inserted root node\n")
	} else {
		var n *node = t.root
		for {
			if data.less(n.data) {
				if n.lc == nil {
					n.lc = &node{false, n, nil, nil, data}
					fmt.Printf("DBG: Inserted to the left\n")
					return n.lc
				} else {
					n = n.lc
				}
			} else if n.rc == nil {
				n.rc = &node{false, n, nil, nil, data}
				fmt.Printf("DBG: Inserted to the right\n")
				return n.rc
			} else {
				n = n.rc
			}
		}
	}
	return t.root
}

// inserts data to the tree
func (t *rbtree) insert(data elem) {
	n := t.treeInsert(data)
	for n != t.root && !n.p.black {
		u := uncle(n)
		if gp := grandparent(n); gp != nil && n.p == gp.lc {
			if u != nil && !u.black {
				n.p.black = true
				u.black = true
				gp.black = false
				n = gp
			} else {
				if n == n.p.rc {
					n = n.p
					t.leftRotate(n)
				}
				n.p.black = true
				gp.black = false
				t.rightRotate(gp)
			}
		} else {
			if u != nil && !u.black {
				n.p.black = true
				u.black = true
				gp.black = false
				n = gp
			} else {
				if n == n.p.lc {
					n = n.p
					t.rightRotate(n)
				}
				n.p.black = true
				gp.black = false
				t.leftRotate(gp)
			}
		}
	}
	t.root.black = true
}

func (t *rbtree) delete(key elem) {
	n := t.searchNode(key)
	if n == nil {
		return // nothing to delete
	}
	if n.lc != nil && n.rc != nil {
		prev := t.maxNode(n.lc)
		n.data = prev.data
		n = prev // kill prev instead
	}

	var child *node
	if n.rc == nil {
		child = n.lc
	} else {
		child = n.rc
	}
	if n.black {
		if child != nil {
			n.black = child.black
		} else {
			n.black = true
		}
		t.delCase1(n)
	}
	// then replace
	if n.p == nil {
		t.root = child
	} else {
		if n == n.p.lc {
			n.p.lc = child
		} else {
			n.p.rc = child
		}
	}
	if child != nil {
		child.p = n.p
	}
	if n.p == nil && child != nil {
		child.black = true
	}
}

// -- all special cases of delete in separate functions --
// Sure we could do this in-place because of tail recursion, but
// let's try to get this working first...

// C1: n is the new root
func (t *rbtree) delCase1(n *node) {
	fmt.Printf("DBG: del case 1\n")
	if n.p == nil {
		return
	}
	t.delCase2(n)
}
// C2: not root, but has red sibling
func (t *rbtree) delCase2(n *node) {
	fmt.Printf("DBG: del case 2\n")
	if s := sibling(n); s != nil && !s.black {
		n.p.black = false
		s.black = true
		if n == n.p.lc {
			t.leftRotate(n.p)
		} else {
			t.rightRotate(n.p)
		}
	}
	t.delCase3(n)
}
// C3: black relatives
func (t *rbtree) delCase3(n *node) {
	fmt.Printf("DBG: del case 3\n")
	if s := sibling(n); s != nil {
		if n.p.black && (s.black &&
						(s.lc == nil || s.lc.black) &&
						(s.rc == nil || s.rc.black)) {
			s.black = false
			t.delCase1(n.p)
			return
		}
	}
	t.delCase4(n)
}
// C4: red parent, other relatives black
func (t *rbtree) delCase4(n *node) {
	fmt.Printf("DBG: del case 4\n")
	if s := sibling(n); s != nil {
		if !n.p.black &&  s == nil || (s.black &&
								(s.lc == nil || s.lc.black) &&
							    (s.rc == nil || s.rc.black)) {
			s.black = false
			n.p.black = true
			return
		}
	}
	t.delCase5(n)
}
// C5: sibling has red child, prepares for case 6
func (t *rbtree) delCase5(n *node) {
	fmt.Printf("DBG: del case 5\n")
	s := sibling(n)
	if n == n.p.lc && s.black && (s.lc != nil && !s.lc.black) &&
								(s.rc == nil || s.rc.black) {
		s.black = false
		s.lc.black = true
		t.rightRotate(s)
	} else if n == n.p.rc && s.black &&
				(s.lc == nil || s.lc.black) &&
				(s.rc != nil && !s.rc.black) {
		s.black = false
		s.rc.black = true
		t.leftRotate(s)
	}
	t.delCase6(n)
}
// C6: finishing up
func (t *rbtree) delCase6(n *node) {
	fmt.Printf("DBG: del case 6\n")
	s := sibling(n)
	s.black = n.p.black
	if n == n.p.lc {
		if s.rc != nil {
			s.rc.black = true
		}
		t.leftRotate(n.p)
	} else {
		if s.lc != nil {
			s.lc.black = true
		}
		t.rightRotate(n.p)
	}
}

func (t *rbtree) maxNode(n *node) *node {
	for n.rc != nil {
		n = n.rc
	}
	return n
}

func (t *rbtree) search(key elem) elem {
	if n := t.searchNode(key); n != nil {
		return n.data
	}
	return nil
}

func (t *rbtree) searchNode(key elem) *node {
	if t.root == nil {
		return nil
	}
	n := t.root
	for n != nil {
		if n.data.equals(key) {
			return n
		}
		if key.less(n.data) {
			n = n.lc
		} else {
			n = n.rc
		}
	}
	return nil
}


/* Moves the given node to left child of its right subtree
   Does not alter coloring
   For example:
      
	  n                 m
	 / \               / \
	lc  m   --->      n  m.rc
	   / \           / \
	m.lc m.rc       lc m.lc
*/
func (t *rbtree) leftRotate(n *node) {
	fmt.Printf("DBG: left rotate\n")
	if m := n.rc; m != nil {
		n.rc = m.lc
		if m.lc != nil {
			m.lc.p = n
		}

		m.p = n.p
		if n.p == nil {
			t.root = m
		} else if n.p.rc == n {
			n.p.rc = m
		} else {
			n.p.lc = m
		}
		m.lc = n
		n.p = m
	}
}

/* Moves the given node right child of its left subtree
   Does not alter coloring
   For example:
      
	  n                 m
	 / \               / \
    m   rc  --->    m.lc  n
   / \                   / \
m.lc m.rc              m.rc rc
*/
func (t *rbtree) rightRotate(n *node) {
	fmt.Printf("DBG: right rotate\n")
	if m := n.lc; m != nil {
		n.lc = m.rc
		if m.rc != nil {
			m.rc.p = n
		}

		m.p = n.p
		if n.p == nil {
			t.root = m
		} else if n.p.lc == n {
			n.p.lc = m
		} else {
			n.p.rc = m
		}
		m.rc = n
		n.p = m
	}
}

func sibling(n *node) *node {
	if n.p == nil {
		return nil
	} else if n == n.p.lc {
		return n.p.rc
	}
	return n.p.lc
}

func grandparent(n *node) *node {
	if n != nil && n.p != nil {
		return n.p.p
	}
	return nil
}

func uncle(n *node) *node {
	if gp := grandparent(n); gp != nil {
		if n.p == gp.rc {
			return gp.lc
		} else {
			return gp.rc
		}
	}
	return nil
}

// Checks that all invariants apply to this tree
// no need to check these (implicit)
// 		a) all nodes are red or black
// 		b) all leaves are black
func (t *rbtree) check() {
	t.inv1()
	t.inv2(t.root)
	t.inv3()
	fmt.Printf("DBG: Tree checked\n")
}

func (t *rbtree) inv1() {
	if t.root != nil && !t.root.black {
		fmt.Printf("ERR: Inv1 - Root node not black\n")
	}
}

// red nodes have 2 black children and black parent
func (t *rbtree) inv2(n *node) {
	if n == nil {
		return
	}
	if !n.black {
		if n.lc != nil && !n.lc.black ||
		   n.rc != nil && !n.rc.black ||
		   n.p != nil && !n.p.black {
			fmt.Printf("ERR: Inv2 - color mismatch\n")
		}
	}
	t.inv2(n.lc)
	t.inv2(n.rc)
}

// all paths to leafs contain same number of black nodes
func (t *rbtree) inv3() {
	pathsum = -1
	t.inv3walk(t.root, 0)
}

var (
	pathsum int
)

func (t *rbtree) inv3walk(n *node, sum int) {
	if n == nil || n.black {
		sum++
	}
	if n == nil {
		if pathsum == -1 {
			pathsum = sum
		} else if sum != pathsum {
			fmt.Printf("ERR: Inv3 - Black count failed\n")
		}
		return
	}
	t.inv3walk(n.lc, sum)
	t.inv3walk(n.rc, sum)
}

// outputs a nice text representation of the tree
func (t *rbtree) printOrder() {
	t.printWalker(t.root, 0)
	fmt.Printf("\n")
}

func (t *rbtree) printWalker(n *node, level int) {
	if n == nil {
		fmt.Printf("nil")
		return
	}
	if n.rc != nil {
		t.printWalker(n.rc, level+1)
	}
	for i:=0; i<level; i++ {
		fmt.Printf("..")
	}
	if n.black {
		fmt.Printf("B%d: %s\n", n.data.key(), n.data.val())
	} else {
		fmt.Printf("R%d: %s\n", n.data.key(), n.data.val())
	}
	if n.lc != nil {
		t.printWalker(n.lc, level+1)
	}
}

// wrapper for int,string pair that can be used as tree element
type dummy struct {
	k int
	v string
}

func (d *dummy) less(e elem) bool {
	if d.k < e.key() {
		return true
	}
	return false
}

func (d *dummy) equals(e elem) bool {
	if d.k == e.key() {
		return true
	}
	return false
}

func (d *dummy) key() int {
	return d.k
}

func (d *dummy) val() string {
	return d.v
}

func wrap(k int, v string) *dummy {
	return &dummy{k, v}
}

func test() {
	donnie := "Any inaccuracies in this index may be explained by the fact that" +
			" it has been sorted with the help of a computer."
	a := strings.Split(donnie, " ", 0)

	// add quote to tree in order
	t := rbtree{nil}
	for i:=0; i<len(a); i++ {
		t.insert(wrap(i, a[i]))
	}
	t.printOrder()
	t.check()

	fmt.Printf("Ok, let's scramble the order...\n")
	t2 := rbtree{nil}
	idx := make([]int, len(a))
	for i:=0; i<len(a); i++ {
		idx[i] = i
	}
	for i:=0; i<100; i++ {
		r1 := rand.Int() % len(a)
		r2 := rand.Int() % len(a)
		idx[r1], idx[r2] = idx[r2], idx[r1]
	}
	for i:=0; i<len(a); i++ {
		t2.insert(wrap(idx[i], a[idx[i]]))
	}
	t2.printOrder()
	t2.check()

	fmt.Printf("Testing search...\n")
	for i:=0; i<len(a); i++ {
		e := t2.search(wrap(i,""))
		fmt.Printf("Key %d -> %s\n", i, e.val())
	}

	fmt.Printf("Testing delete...\n")
	for i:=0; i<len(a); i++ {
		fmt.Printf("Deleting key %d\n", idx[i])
		t2.delete(wrap(idx[i],""))
		t2.printOrder()
		t2.check()
	}
}

func main() {
	test()
}
