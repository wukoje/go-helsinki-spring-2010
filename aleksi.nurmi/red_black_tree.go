
package main

import "fmt"

type rbColor bool

const (
	black rbColor = false
	red = true
)

type rbNode struct {
	key int
	color rbColor
	parent, left, right *rbNode
}

var nilNode *rbNode = &rbNode { -1, black, nil, nil, nil }

func (c rbColor) String() string {
	if c == black {
		return "black"
	}
	return "red"
}

func (n *rbNode) String() string {
	if n == nilNode {
		return "NIL"
	}
	return fmt.Sprintf("(%v,%v,%v,%v)", n.key, n.color, n.left, n.right)
}

func rbRightRotate(proot **rbNode, x *rbNode) {
	y := x.left
	x.left = y.right
	if y.right != nilNode {
		x = y.right.parent
	}
	w := x.parent
	y.parent = w
	if w == nilNode {
		*proot = y
	} else {
		if w.left == x {
			w.left = y
		} else {
			w.right = y
		}
	}
	y.right = x
	x.parent = y
}

func rbLeftRotate(proot **rbNode, x *rbNode) {
	y := x.right
	x.right = y.left
	if y.left != nilNode {
		y.left.parent = x
	}
	w := x.parent
	y.parent = w
	if w == nilNode {
		*proot = y
	} else {
		if w.left == x {
			w.left = y
		} else {
			w.right = y
		}
	}
	y.left = x
	x.parent = y
}

func rbMin(x *rbNode) *rbNode {
	for x.left != nilNode {
		x = x.left
	}
	return x
}

func rbSucc(x *rbNode) *rbNode {
	if x.right != nilNode {
		return rbMin(x.right)
	}

	y := x.parent
	for y != nilNode && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

func rbSearch(x *rbNode, k int) *rbNode {
	if x == nilNode || x.key == k {
		return x
	}
	if k < x.key {
		return rbSearch(x.left, k)
	}
	return rbSearch(x.right, k)
}

func rbDeleteFix(proot **rbNode, x *rbNode) {
	for x != *proot && x.color == black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				rbLeftRotate(proot, x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					rbRightRotate(proot, w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				rbLeftRotate(proot, x.parent)
				x = *proot
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				rbRightRotate(proot, x.parent)
				w = x.parent.left
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					rbLeftRotate(proot, w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				rbRightRotate(proot, x.parent)
				x = *proot
			}
		}
	}
	x.color = black
}

func rbDelete(proot **rbNode, z *rbNode) *rbNode {
	var x *rbNode
	if z.left == nilNode || z.right == nilNode {
		if z.left != nilNode {
			x = z.left
		} else {
			x = z.right
		}
		w := z.parent
		if w == nilNode {
			*proot = x
		} else {
			if z == w.left {
				w.left = x
			} else {
				w.right = x
			}
		}
		x.parent = w
		if z.color == black {
			rbDeleteFix(proot, x)
		}
		return z
	}
	y := rbSucc(z)
	x = y.right
	w := y.parent
	if y == w.left {
		w.left = x
	} else {
		w.right = x
	}
	x.parent = w
	z.key = y.key
	if y.color == black {
		rbDeleteFix(proot, x)
	}
	return y
}

func rbInsertFix(proot **rbNode, z *rbNode) {
	for z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					rbLeftRotate(proot, z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				rbRightRotate(proot, z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					rbRightRotate(proot, z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				rbLeftRotate(proot, z.parent.parent)
			}
		}
	}
	(*proot).color = black
}

func rbInsert(proot **rbNode, k int) {
	z := new(rbNode)
	z.key = k
	z.left = nilNode
	z.right = nilNode

	x := *proot
	y := nilNode
	for x != nilNode {
		y = x
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nilNode {
		*proot = z
	} else {
		if k < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}
	z.color = red
	rbInsertFix(proot, z)
}

// in addition, the root must be black
func (n *rbNode) isValid() (valid bool, blacks int) {
	valid = true
	switch {
	// a root is black
	case n.parent == nilNode:
		if n.color != black {
			valid = false
		}
	// a nil node is black
	case n == nilNode:
		if n.color != black {
			println("nil node isn't black")
			valid = false
		}
	// the children of a red node are black
	case n.color == red:
		if !(n.left.color == black && n.right.color == black) {
			println("children of a red node aren't black")
			valid = false
		}
	}
	// all paths from a node to its leaves have an equal number of blacks
	var leftValid, rightValid bool
	var leftBlacks, rightBlacks int
	if n != nilNode {
		leftValid, leftBlacks = n.left.isValid()
		rightValid, rightBlacks = n.right.isValid()
	} else {
		leftValid, rightValid = true, true
	}
	if n.color == black {
		return valid && leftValid && rightValid, leftBlacks + rightBlacks + 1
	}
	return valid && leftValid && rightValid, leftBlacks + rightBlacks
}

func rbNew(s []int) *rbNode {
	root := nilNode
	for _, k := range(s) {
		rbInsert(&root, k)
	}
	return root
}

func main() {
	root := rbNew([]int { 8, 3, 4 })
	valid, blacks := root.isValid()
	valid = valid && root.color == black
	println(valid, blacks)
	fmt.Println(root)

	rbInsert(&root, 3)
	rbInsert(&root, 1)
	rbInsert(&root, 8)
	rbInsert(&root, 4)
	rbInsert(&root, 2)
	valid, blacks = root.isValid()
	valid = valid && root.color == black
	println(valid, blacks)
	fmt.Println(root)

	// -1 is the key of nilNode, but it shouldn't be found
	x := rbSearch(root, nilNode.key)
	println(nilNode.key, "exists", x != nilNode)
	
	x = rbSearch(root, 666)
	println("666 exists", x != nilNode)

	x = rbSearch(root, 2)
	println("2 exists", x != nilNode)
	rbDelete(&root, x)
	fmt.Println(root)
	valid, blacks = root.isValid()
	valid = valid && root.color == black
	println(valid, blacks)

	x = rbSearch(root, 4)
	println("4 exists", x != nilNode)
	rbDelete(&root, x)
	fmt.Println(root)
	valid, blacks = root.isValid()
	valid = valid && root.color == black
	println(valid, blacks)
	
	x = rbSearch(root, 4)
	println("4 exists", x != nilNode)
	rbDelete(&root, x)
	fmt.Println(root)
	valid, blacks = root.isValid()
	valid = valid && root.color == black
	println(valid, blacks)
}

