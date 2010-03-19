// Red-black tree from Cormen 3rd ed
package main

import (
	"fmt"
	"strconv"
	"container/list"
)

const (
	black = iota
	red
)

type tree struct {
	root     *node
	sentinel *node
}

type node struct {
	parent, left, right *node
	color               int
	key                 int
}

func (t *tree) rotate_right(n *node) {
	y := n.left
	n.left = y.right
	if y.right != t.sentinel {
		y.right.parent = n
	}
	y.parent = n.parent
	switch {
	case n.parent == t.sentinel:
		t.root = y
	case n.parent.right == n:
		n.parent.right = y
	default:
		n.parent.left = y
	}
	y.right = n
	n.parent = y
}

func (t *tree) rotate_left(n *node) {
	y := n.right
	n.right = y.left
	if y.left != t.sentinel {
		y.left.parent = n
	}
	y.parent = n.parent
	switch {
	case n.parent == t.sentinel:
		t.root = y
	case n.parent.left == n:
		n.parent.left = y
	default:
		n.parent.right = y
	}
	y.left = n
	n.parent = y
}

func (t *tree) insert(k int) {
	n := &node{t.sentinel, t.sentinel, t.sentinel, red, k}
	p := t.sentinel
	for x := t.root; x != t.sentinel; {
		p = x
		if n.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	n.parent = p
	switch {
	case p == t.sentinel:
		t.root = n
	case n.key < p.key:
		p.left = n
	default:
		p.right = n
	}
	t.insert_fixup(n)
}

func (t *tree) insert_fixup(n *node) {
	for n.parent.color == red {
		if n.parent == n.parent.parent.left {
			y := n.parent.parent.right
			if y.color == red {
				n.parent.color = black
				y.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					t.rotate_left(n)
				}
				n.parent.color = black
				n.parent.parent.color = red
				t.rotate_right(n.parent.parent)
			}
		} else {
			y := n.parent.parent.left
			if y.color == red {
				n.parent.color = black
				y.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rotate_right(n)
				}
				n.parent.color = black
				n.parent.parent.color = red
				t.rotate_left(n.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *tree) transplant(x *node, y *node) {
	switch {
	case x.parent == t.sentinel:
		t.root = y
	case x == x.parent.left:
		x.parent.left = y
	default:
		x.parent.right = y
	}
	y.parent = x.parent
}

func (t *tree) min(n *node) *node {
	if n == t.sentinel {
		return n
	}
	for n.left != t.sentinel {
		n = n.left
	}
	return n
}

func (t *tree) delete(k int) bool {
	found, n := t.search(k)
	if !found {
		return false
	}
	y := n
	ycolor := y.color
	var x *node
	switch {
	case n.left == t.sentinel:
		x = n.right
		t.transplant(n, n.right)
	case n.right == t.sentinel:
		x = n.left
		t.transplant(n, n.left)
	default:
		y = t.min(n.right)
		ycolor = y.color
		x = y.right
		if y.parent == n {
			n.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
		y.color = n.color
	}
	if ycolor == black {
		t.delete_fixup(x)
	}
	return true
}

func (t *tree) delete_fixup(n *node) {
	for n != t.root && n.color == black {
		if n == n.parent.left {
			w := n.parent.right
			if w.color == red {
				w.color = black
				n.parent.color = red
				t.rotate_left(n.parent)
				w = n.parent.right
			}
			if w.left.color == black &&
				w.right.color == black {
				w.color = red
				n = n.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					t.rotate_right(w)
					w = n.parent.right
				}
				w.color = n.parent.color
				n.parent.color = black
				w.right.color = black
				t.rotate_left(n.parent)
				n = t.root
			}
		} else {
			w := n.parent.left
			if w.color == red {
				w.color = black
				n.parent.color = red
				t.rotate_right(n.parent)
				w = n.parent.left
			}
			if w.right.color == black &&
				w.left.color == black {
				w.color = red
				n = n.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					t.rotate_left(w)
					w = n.parent.left
				}
				w.color = n.parent.color
				n.parent.color = black
				w.left.color = black
				t.rotate_right(n.parent)
				n = t.root
			}
		}
	}
	n.color = black
}

func (t *tree) search(k int) (ok bool, pos *node) {
	for pos = t.root; pos != t.sentinel; {
		if pos.key == k {
			return true, pos
		} else if k < pos.key {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
	return false, t.sentinel
}

func (t *tree) String() (buf string) {
	if t.root == t.sentinel {
		return
	}
	queue := new(list.List)
	queue.PushBack(t.root)
	in, out := 0, 1
	for queue.Len() > 0 {
		qelement := queue.Front()
		cur := qelement.Value.(*node)
		queue.Remove(qelement)

		if cur.left != t.sentinel {
			queue.PushBack(cur.left)
			in++
			buf += "_"
		}

		if cur.color == black {
			buf += "b"
		} else {
			buf += "r"
		}
		buf += strconv.Itoa(cur.key)

		if cur.right != t.sentinel {
			queue.PushBack(cur.right)
			in++
			buf += "_"
		}

		buf += " "

		if out--; out <= 0 {
			out, in = in, out
			buf += "\n"
		}
	}
	return
}

func new_tree() (t *tree) {
	sentinel := &node{nil, nil, nil, black, 0}
	t = &tree{sentinel, sentinel}
	return
}

func main() {
	t := new_tree()

	for i := 0; i <= 100; i++ {
		t.insert(i)
	}
	fmt.Println("Tree after inserting 0...100")
	fmt.Print(t)

	for i := 0; i <= 50; i++ {
		t.delete(i)
	}
	fmt.Println("Tree after deleting 0...50")
	fmt.Print(t)
}

