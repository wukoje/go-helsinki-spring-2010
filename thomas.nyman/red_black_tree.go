// Red-black tree implementation according to
// Introduction to Algorithms Second Edition by
// T.H. Cormen and others
package main

import "fmt"

type color int

const (
	red = iota
	black
)

type Tree struct {
	size int
	root *node // root of tree
}

type node struct {
	key    int   // key value
	color  color // node color
	parent *node // parent node
	left   *node // left child node
	right  *node // right child node
}

// FIFO used for breadth-first traversal.
type fifo struct {
	head  int       // head of queue
	tail  int       // next free index
	queue [](*node) // slice to hold node
}

func main() {
	t := New()

	const testsize = 5

	fmt.Println(t)
	// populate tree
	for i := 0; i <= testsize; i++ {
		t.Insert(i)
		fmt.Println(t)
	}

	// search tree for some keys
	for i := -1; i <= testsize+1; i++ {
		if t.Search(i) {
			fmt.Println("Found", i)
		} else {
			fmt.Println("Did not find", i)
		}
	}

	// empty tree in reverse order
	for i := testsize; i >= 0; i-- {
		t.Delete(i)
		fmt.Println(t)
	}

	// populate tree and empty it in insert order
	for i := 0; i <= testsize; i++ {
		t.Insert(i)
		fmt.Println(t)
	}
	for i := 0; i <= testsize; i++ {
		t.Delete(i)
		fmt.Println(t)
	}
}

// creates new empty Tree
func New() *Tree { return &Tree{0, nil} }

// searches t for key, returns true if found, otherwise false
func (t *Tree) Search(key int) bool {
	if t.treeSearch(t.root, key) != nil {
		return true
	}
	return false
}

// inserts new node with specified key into t
func (t *Tree) Insert(key int) {
	t.insert(&node{key, red, nil, nil, nil})
	t.size++
}

// deletes node with specified key from t (if present)
func (t *Tree) Delete(key int) {
	if n := t.treeSearch(t.root, key); n != nil {
		t.delete(n)
		t.size--
	}
}

// returns string representation or n
func (n *node) String() string {
	s := "(" + fmt.Sprint(n.key)

	if n.color == red {
		s += ":red:"
	} else {
		s += ":black:"
	}

	if n.parent != nil {
		s += fmt.Sprint(n.parent.key) + ":"
	} else {
		s += "<nil>:"
	}

	if n.left != nil {
		s += fmt.Sprint(n.left.key) + ":"
	} else {
		s += "<nil>:"
	}

	if n.right != nil {
		s += fmt.Sprint(n.right.key)
	} else {
		s += "<nil>"
	}
	return s + ")"
}

// returns string representation of t with nodes in level order
func (t *Tree) String() string {
	if t.root == nil {
		return "Tree is empty"
	}
	s := "[ "
	// traverse t in breadth-first fashion
	f := &fifo{0, 0, make([](*node), t.size)}
	f.enqueue(t.root)
	for !f.empty() {
		n := f.dequeue()
		s += fmt.Sprint(n) + " "
		if n.left != nil {
			f.enqueue(n.left)
		}
		if n.right != nil {
			f.enqueue(n.right)
		}
	}
	return s + "]"
}

// adds n to f
func (f *fifo) enqueue(n *node) {
	f.queue[f.tail] = n
	f.tail = (f.tail + 1) % (len(f.queue) + 1)
}

// returns first node from f
func (f *fifo) dequeue() *node {
	n := f.queue[f.head]
	f.head = (f.head + 1) % (len(f.queue) + 1)
	return n
}

// returns true if fifo is empty, otherwise false
func (f *fifo) empty() bool { return f.head == f.tail }

// returns grandparent of n or nil if no grandparent exists
func (n *node) grandparent() *node {
	if n != nil && n.parent != nil {
		return n.parent.parent
	}
	return nil
}

// searches t from n down for key, returns node with key or null
func (t *Tree) treeSearch(n *node, key int) *node {
	if n == nil || key == n.key {
		return n
	}
	if key < n.key {
		return t.treeSearch(n.left, key)
	}
	return t.treeSearch(n.right, key)
}

// inserts n to t
func (t *Tree) insert(n *node) {
	o := t.root
	var m *node
	m = nil

	for o != nil {
		m = o
		if n.key < o.key {
			o = o.left
		} else {
			o = o.right
		}
	}

	n.parent = m

	switch {
	case m == nil:
		t.root = n
	case n.key < m.key:
		m.left = n
	default:
		m.right = n
	}

	n.left, n.right, n.color = nil, nil, red
	t.insertFixup(n)
}

// deletes n from t
func (t *Tree) delete(n *node) {
	if n != nil {
		var m *node
		var o *node

		if n.left == nil || n.right == nil {
			m = n
		} else {
			m = t.treeSuccessor(n)
		}

		if m.left != nil {
			o = m.left
		} else {
			o = m.right
		}

		if o != nil {
			o.parent = m.parent
		}

		switch {
		case m.parent == nil:
			t.root = o
		case m == m.parent.left:
			m.parent.left = o
		default:
			m.parent.right = o
		}

		if m != n {
			n.key = m.key
		}
		if m.color == black {
			t.deleteFixup(o)
		}
	}
}

// returns successor to n
func (t *Tree) treeSuccessor(n *node) *node {
	if n.right != nil {
		return t.treeMin(n.right)
	}
	m := n.parent
	for m != nil && n == m.right {
		n, m = m, m.parent
	}
	return m
}

// returns minimum value from n's subtree
func (t *Tree) treeMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}

// rotates t leftwise at n
func (t *Tree) leftRotate(n *node) {
	m := n.right
	n.right = m.left // turn m's left subtree into n's right subtree
	if m.left != nil {
		m.left.parent = n
	}

	m.parent = n.parent

	switch {
	case n.parent == nil:
		t.root = m
	case n == n.parent.left:
		n.parent.left = m
	default:
		n.parent.right = m
	}

	m.left = n
	n.parent = m
}

// rotates t rightwise at n
func (t *Tree) rightRotate(n *node) {
	m := n.left
	n.left = m.right // turn m's right subtree into n's left subtree
	if m.right != nil {
		m.right.parent = n
	}

	m.parent = n.parent

	switch {
	case n.parent == nil:
		(*t).root = m
	case n == n.parent.right:
		n.parent.right = m
	default:
		n.parent.left = m
	}

	m.right = n
	n.parent = m
}

// repairs tree properties of nodes in t if broken during insert
func (t *Tree) insertFixup(n *node) {
	for n.parent != nil && n.parent.color == red {
		if n.grandparent() != nil && n.parent == n.grandparent().left {
			m := n.grandparent().right

			// case 1
			if m != nil && m.color == red {
				n.parent.color = black
				m.color = black
				n.grandparent().color = red
				n = n.grandparent()
			} else {
				// case 2
				if n == n.parent.right {
					n = n.parent
					t.leftRotate(n)
				}
				// case 3
				n.parent.color = black
				n.grandparent().color = red
				t.rightRotate(n.grandparent())
			}
		} else {
			m := n.grandparent().left
			// case 4
			if m != nil && m.color == red {
				n.parent.color = black
				m.color = black
				n.grandparent().color = red
				n = n.grandparent()
			} else {
				// case 5
				if n == n.parent.left {
					n = n.parent
					t.rightRotate(n)
				}
				// case 6
				n.parent.color = black
				n.grandparent().color = red
				t.leftRotate(n.grandparent())
			}
		}
	}
	t.root.color = black
}

// repairs tree properties of nodes in t if broken during delete
func (t *Tree) deleteFixup(n *node) {
	if n != nil {
		for n != t.root && n.color == black {
			if n == n.parent.left {
				m := n.parent.right
				// case 1
				if m.color == red {
					m.color = black
					n.parent.color = red
					t.leftRotate(n.parent)
					m = n.parent.right
				}
				// case 2
				if m.left.color == black && m.right.color == black {
					m.color = red
					n = n.parent
				} else {
					// case 3
					if m.right.color == black {
						m.left.color = black
						m.color = red
						t.rightRotate(m)
						m = n.parent.right
					}
					// case 4
					m.color = n.parent.color
					n.parent.color = black
					m.right.color = black
					t.leftRotate(n.parent)
					n = t.root
				}
			} else {
				m := n.parent.left
				// case 5
				if m.color == red {
					m.color = black
					n.parent.color = red
					t.rightRotate(n.parent)
					m = n.parent.left
				}
				// case 6
				if m.right.color == black && m.left.color == black {
					m.color = red
					n = n.parent
				} else {
					// case 7
					if m.left.color == black {
						m.right.color = black
						m.color = red
						t.leftRotate(m)
						m = n.parent.left
					}
					// case 8
					m.color = n.parent.color
					n.parent.color = black
					m.left.color = black
					t.rightRotate(n.parent)
					n = t.root
				}
			}
		}
		n.color = black
	}
}
