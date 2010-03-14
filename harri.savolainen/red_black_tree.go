package main

import fmt "fmt"

func main() {

	var s = []int{5, 76, 6, 1, 45, 2, 3, 4, 1000, 5234, 56, 35, 23423, 54676, 23, 32, 34, 8}

	rbTree := newTree()

	for _, v := range s {
		rbTree.insert(v)
	}

	fmt.Printf("Searching for max:%d\n", rbTree.max(rbTree.root).key)
	fmt.Printf("Searching for min:%d\n", rbTree.min(rbTree.root).key)
	fmt.Printf("Searching for key 23 (0=not found, number=found) : %d\n", rbTree.search(rbTree.root, 23).key)
	fmt.Printf("Deleting node with key 23 via deletekey(key) (0=not found, number=deleted): %d\n", rbTree.deletekey(23).key)
	fmt.Printf("Deleting node with key 23 again (0=not found, number=deleted): %d\n", rbTree.deletekey(23).key)

	node := rbTree.search(rbTree.root, 1000)
	fmt.Printf("Search for key 1000 (0=not found, number=found):%d\n", node.key)
	fmt.Printf("Deleting node with via delete(node) 1000 (0=not found, number=deleted):%d\n", rbTree.delete(node).key)
	fmt.Printf("Searching for key again 1000 (0=not found, number=found) : %d\n", rbTree.search(rbTree.root, 1000).key)

}


type node struct {
	key   int
	left  *node
	right *node
	p     *node
	color string
	tnil  int
}


func newNode(key int) *node {
	n := new(node)
	n.key = key
	n.left = nil
	n.right = nil
	n.p = nil
	n.color = "b"
	n.tnil = 0
	return n
}

type rbtree struct {
	root *node
	tnil *node
}


func newTree() *rbtree {
	rbt := new(rbtree)
	rbt.tnil = newNode(0)
	rbt.tnil.tnil = 1
	rbt.root = rbt.tnil
	return rbt
}

func (rbt *rbtree) insert(k int) {
	z := newNode(k)
	z.left = rbt.tnil
	z.right = rbt.tnil
	x := rbt.root
	y := rbt.tnil
	tnil := rbt.tnil

	for x != tnil {
		y = x
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == rbt.tnil {
		rbt.root = z
	} else {
		if k < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}
	z.color = "r"
	rbt.rbinsertfixup(z)
}

func (rbt *rbtree) rbinsertfixup(z *node) {
	for z.p.color == "r" {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == "r" {
				z.p.color = "b"
				y.color = "b"
				z.p.p.color = "r"
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					rbt.leftrotate(z)
				}
				z.p.color = "b"
				z.p.p.color = "r"
				rbt.rightrotate(z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == "r" {
				z.p.color = "b"
				y.color = "b"
				z.p.p.color = "r"
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					rbt.rightrotate(z)
				}
				z.p.color = "b"
				z.p.p.color = "r"
				rbt.leftrotate(z.p.p)
			}
		}

	}
	rbt.root.color = "b"

}

func (rbt *rbtree) rightrotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != rbt.tnil {
		y.right.p = x
	}
	w := x.p
	y.p = w
	if w == rbt.tnil {
		rbt.root = y
	} else {
		if w.right == x {
			w.right = y
		} else {
			w.left = y
		}
	}
	y.right = x
	x.p = y
}

func (rbt *rbtree) leftrotate(x *node) {
	y := x.right
	x.right = y.left
	if y.left != rbt.tnil {
		y.left.p = x
	}
	w := x.p
	y.p = w
	if w == rbt.tnil {
		rbt.root = y
	} else {
		if w.left == x {
			w.left = y
		} else {
			w.right = y
		}
	}
	y.left = x
	x.p = y

}

func (rbt *rbtree) delete(z *node) *node {

	if z.left == rbt.tnil || z.right == rbt.tnil {

		x := rbt.tnil

		if z.left != rbt.tnil {
			x = z.left
		} else {
			x = z.right
		}
		w := z.p
		if w == rbt.tnil {
			rbt.root = x
		} else {
			if z == w.left {
				w.left = x
			} else {
				w.right = x
			}
		}
		x.p = w
		if z.color == "b" {
			rbt.rbdeletefixup(x)
		}
		return z
	}
	y := rbt.succ(z)
	x := y.right
	w := y.p
	if y == w.left {
		w.left = x
	} else {
		w.right = x
	}
	x.p = w
	z.key = y.key
	if y.color == "b" {
		rbt.rbdeletefixup(x)
	}
	return y
}

func (rbt *rbtree) rbdeletefixup(x *node) {
	for x != rbt.root && x.color == "b" {
		if x == x.p.left {
			w := x.p.right
			if w.color == "r" {
				w.color = "b"
				x.p.color = "r"
				rbt.leftrotate(x.p)
				w = x.p.right
			}
			if w.left.color == "b" && w.right.color == "b" {
				w.color = "r"
				x = x.p
			} else {
				if w.right.color == "b" {
					w.left.color = "b"
					w.color = "r"
					rbt.rightrotate(w)
					w = x.p.right
				}
				w.color = x.p.color
				x.p.color = "b"
				w.right.color = "b"
				rbt.leftrotate(x.p)
				x = rbt.root
			}
		} else {
			w := x.p.left
			if w.color == "r" {
				w.color = "b"
				x.p.color = "r"
				rbt.rightrotate(x.p)
				w = x.p.left
			}
			if w.right.color == "b" && w.left.color == "b" {
				w.color = "r"
				x = x.p
			} else {
				if w.left.color == "b" {
					w.right.color = "b"
					w.color = "r"
					rbt.leftrotate(w)
					w = x.p.left
				}
				w.color = x.p.color
				x.p.color = "b"
				w.left.color = "b"
				rbt.rightrotate(x.p)
				x = rbt.root
			}

		}

	}
	x.color = "b"

}


func (rbt *rbtree) succ(x *node) *node {
	if x.right != rbt.tnil {
		return rbt.min(x.right)
	} else {
		y := x.p
		for y != rbt.tnil && x == y.right {
			x = y
			y = x.p
		}
		return y
	}
	panic("unreachable")
}


func (rbt *rbtree) min(x *node) *node {
	if x.left.tnil == 0 {
		y := rbt.min(x.left)
		return y

	}
	return x
}

func (rbt *rbtree) max(x *node) *node {
	if x.right.tnil == 0 {
		y := rbt.max(x.right)
		return y

	}
	return x
}

func (rbt *rbtree) search(x *node, k int) *node {
	if x == rbt.tnil || x.key == k {
		return x
	}
	if k < x.key {
		return rbt.search(x.left, k)
	} else {
		return rbt.search(x.right, k)
	}
	panic("unreachable")
}

func (rbt *rbtree) deletekey(k int) *node {
	result := rbt.search(rbt.root, k)
	if result.tnil == 0 {
		return rbt.delete(result)
	} else {
		return result
	}
	panic("unreachable")
}
