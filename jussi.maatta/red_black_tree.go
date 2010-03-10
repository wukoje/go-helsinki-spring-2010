// Red-Black Tree.
// This implementation is based on the sample algorithms presented in
// Cormen, Leiserson, Rivest, Stein: Introduction to Algorithms, 2nd ed. (2001)

package main

import (
	"fmt"
	"rand"
	"time"
)

type rbtree struct {
	root *rbnode
	nil  *rbnode // sentinel
}

type rbnode struct {
	left, right, p *rbnode
	black          bool
	val            int
}

func rbtreeNew() *rbtree {
	t := &rbtree{}
	t.nil = &rbnode{}
	t.nil.black = true
	t.nil.val = -99999 // makes debugging easier
	t.root = t.nil
	return t
}

func (t *rbtree) search(val int) *rbnode {
	x := t.root
	for x != t.nil && val != x.val {
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (t *rbtree) succ(x *rbnode) *rbnode {
	if x.right != t.nil {
		return t.min(x.right)
	}
	y := x.p
	for y != t.nil && x == y.right {
		x = y
		y = y.p
	}
	return y
}

func (t *rbtree) min(x *rbnode) *rbnode {
	for x.left != t.nil {
		x = x.left
	}
	return x
}

func (t *rbtree) leftRotate(x *rbnode) {
	y := x.right
	x.right = y.left
	if y.left != t.nil {
		y.left.p = x
	}
	y.p = x.p
	if x.p == t.nil {
		t.root = y
	} else if x.p.left == x {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func (t *rbtree) rightRotate(x *rbnode) {
	y := x.left
	x.left = y.right
	if y.right != t.nil {
		y.right.p = x
	}
	y.p = x.p
	if x.p == t.nil {
		t.root = y
	} else if x.p.right == x {
		x.p.right = y
	} else {
		x.p.left = y
	}
	y.right = x
	x.p = y
}

func (t *rbtree) insert(val int) {
	var x, y *rbnode
	z := &rbnode{}
	z.val = val
	y = t.nil
	x = t.root
	for x != t.nil {
		y = x
		if z.val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == t.nil {
		t.root = z
	} else if z.val < y.val {
		y.left = z
	} else {
		y.right = z
	}
	z.left = t.nil
	z.right = t.nil
	z.black = false
	t.insertFixup(z)
}

func (t *rbtree) insertFixup(z *rbnode) {
	for !z.p.black {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if !y.black {
				z.p.black = true
				y.black = true
				z.p.p.black = false
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					t.leftRotate(z)
				}
				z.p.black = true
				z.p.p.black = false
				t.rightRotate(z.p.p)
			}
		} else {
			y := z.p.p.left
			if !y.black {
				z.p.black = true
				y.black = true
				z.p.p.black = false
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					t.rightRotate(z)
				}
				z.p.black = true
				z.p.p.black = false
				t.leftRotate(z.p.p)
			}
		}
	}
	t.root.black = true
}

func (t *rbtree) delete(z *rbnode) *rbnode {
	var x, y *rbnode
	if z.left == t.nil || z.right == t.nil {
		y = z
	} else {
		y = t.succ(z)
	}
	if y.left != t.nil {
		x = y.left
	} else {
		x = y.right
	}
	x.p = y.p
	if y.p == t.nil {
		t.root = x
	} else {
		if y == y.p.left {
			y.p.left = x
		} else {
			y.p.right = x
		}
	}
	if y != z {
		z.val = y.val
	}
	if y.black {
		t.deleteFixup(x)
	}
	return y
}

func (t *rbtree) deleteFixup(x *rbnode) {
	var w *rbnode
	for x != t.root && x.black {
		if x == x.p.left {
			w = x.p.right
			if !w.black {
				w.black = true
				x.p.black = false
				t.leftRotate(x.p)
				w = x.p.right
			}
			if w.left.black && w.right.black {
				w.black = false
				x = x.p
			} else {
				if w.right.black {
					w.left.black = true
					w.black = false
					t.rightRotate(w)
					w = x.p.right
				}
				w.black = x.p.black
				x.p.black = true
				w.right.black = true
				t.leftRotate(x.p)
				x = t.root
			}
		} else {
			w = x.p.left
			if !w.black {
				w.black = true
				x.p.black = false
				t.rightRotate(x.p)
				w = x.p.left
			}
			if w.left.black && w.right.black {
				w.black = false
				x = x.p
			} else {
				if w.left.black {
					w.right.black = true
					w.black = false
					t.leftRotate(w)
					w = x.p.left
				}
				w.black = x.p.black
				x.p.black = true
				w.left.black = true
				t.rightRotate(x.p)
				x = t.root
			}
		}
	}
	x.black = true
}

func (t *rbtree) magickPrint() { t.magickPrint2(t.root, 0) }

func (t *rbtree) magickPrint2(p *rbnode, level int) {
	if p == t.nil {
		return
	}
	for i := 0; i < 4*level; i++ {
		fmt.Print(" ")
	}
	switch p.black {
	case true:
		fmt.Print("b_")
	case false:
		fmt.Print("r_")
	}
	fmt.Println(p.val)
	t.magickPrint2(p.left, level+1)
	t.magickPrint2(p.right, level+1)
}

func main() {
	seed := time.Nanoseconds()
	rand.Seed(seed)
	fmt.Println("seed:", seed)

	t := rbtreeNew()

	fmt.Println("filling a red-black tree with 20 random integers from the range [0,40]")
	for i := 0; i < 20; i++ {
		t.insert(rand.Intn(41))
	}

	fmt.Println("this is what the tree looks like:")
	t.magickPrint()

	fmt.Println("generating 10 randints from the range, searching and deleting if found")
	for i := 0; i < 10; i++ {
		x := rand.Intn(41)
		p := t.search(x)
		if p != t.nil {
			fmt.Println("\t", x, "found, deleting")
			t.delete(p)
		} else {
			fmt.Println("\t", x, "not found")
		}
	}

	fmt.Println("now the tree looks like this:")
	t.magickPrint()
}
