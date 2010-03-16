package main

import (
	"fmt"
)

type node struct {
	key    int
	red    bool
	parent *node
	left   *node
	right  *node
}

type tree struct {
	root *node
	nil  *node // dummy node for leaves
}

func (n *node) isNil() bool { return n == n.left }

func initNode(key int, nil *node) *node {
	n := new(node)
	n.red = true
	n.parent = nil
	n.left = nil
	n.right = nil
	n.key = key
	return n
}

func initTree() *tree {
	t := &tree{}

	nil := &node{}
	nil.red = false
	nil.parent = nil
	nil.left = nil
	nil.right = nil
	t.nil = nil

	t.root = t.nil

	return t
}

func (t *tree) search(key int) *node {
	n := t.root
	for !n.isNil() {
		if key == n.key {
			return n
		} else if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return nil
}

func (t *tree) rotateLeft(n *node) {
	p := n.right
	n.right = p.left
	if !p.left.isNil() {
		p.left.parent = n
	}
	parent := n.parent
	p.parent = parent
	if parent.isNil() {
		t.root = p
	} else if parent.left == n {
		parent.left = p
	} else {
		parent.right = p
	}
	p.left = n
	n.parent = p
}


func (t *tree) rotateRight(n *node) {
	p := n.left
	n.left = p.right
	if !p.right.isNil() {
		p.right.parent = n
	}
	parent := n.parent
	p.parent = parent
	if parent.isNil() {
		t.root = p
	} else if parent.right == n {
		parent.right = p
	} else {
		parent.left = p
	}
	p.right = n
	n.parent = p
}

func (t *tree) insertBalance(n *node) {
	for n.parent.red == true {
		if n.parent == n.parent.parent.left {
			uncle := n.parent.parent.right
			if uncle.red == true {
				n.parent.red = false
				uncle.red = false
				n.parent.parent.red = true
				n = n.parent.parent
			} else {
				if n == n.parent.right {
					n = n.parent
					t.rotateLeft(n)
				}
				n.parent.red = false
				n.parent.parent.red = true
				t.rotateRight(n.parent.parent)
			}
		} else {
			uncle := n.parent.parent.left
			if uncle.red == true {
				n.parent.red = false
				uncle.red = false
				n.parent.parent.red = true
				n = n.parent.parent
			} else {
				if n == n.parent.left {
					n = n.parent
					t.rotateRight(n)
				}
				n.parent.red = false
				n.parent.parent.red = true
				t.rotateLeft(n.parent.parent)
			}
		}
	}
	t.root.red = false
}

func (t *tree) insert(key int) *node {
	n := t.root
	prev := t.nil
	for !n.isNil() {
		prev = n
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	newNode := initNode(key, t.nil)
	newNode.parent = prev
	if prev.isNil() {
		t.root = newNode
	} else if key < prev.key {
		prev.left = newNode
	} else {
		prev.right = newNode
	}
	t.insertBalance(newNode)
	return newNode
}

func (t *tree) deleteBalance(n *node) {
	for n != t.root && n.red == false {
		if n == n.parent.left {
			sib := n.parent.right
			if sib.red == true {
				n.parent.red = true
				sib.red = false
				t.rotateLeft(n.parent)
				sib = n.parent.right
			}
			if sib.left.red == false && sib.right.red == false {
				n = n.parent
				sib.red = true
			} else {
				if sib.right.red == false {
					sib.red = true
					sib.left.red = false
					t.rotateRight(sib)
					sib = n.parent.right
				}
				sib.right.red = false
				sib.red = n.parent.red
				n.parent.red = false
				t.rotateLeft(n.parent)
				n = t.root
			}
		} else {
			sib := n.parent.left
			if sib.red == true {
				n.parent.red = true
				sib.red = false
				t.rotateRight(n.parent)
				sib = n.parent.left
			}
			if sib.left.red == false && sib.right.red == false {
				sib.red = true
				n = n.parent
			} else {
				if sib.left.red == false {
					sib.red = true
					sib.right.red = false
					t.rotateLeft(sib)
					sib = n.parent.left
				}
				sib.left.red = false
				sib.red = n.parent.red
				n.parent.red = false
				t.rotateRight(n.parent)
				n = t.root
			}
		}
	}
	n.red = false
}

func (t *tree) delete(n *node) {
	if n == nil || n.isNil() {
		return
	}

	if n.left.isNil() || n.right.isNil() {
		child := n.left
		if child.isNil() {
			child = n.right
		}
		parent := n.parent
		if parent.isNil() {
			t.root = child
		} else if n == parent.left {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent
		if n.red == false {
			t.deleteBalance(child)
		}
		return
	}

	succ := n.right
	for !succ.left.isNil() {
		succ = succ.left
	}
	n.key, succ.key = succ.key, n.key

	child := succ.right
	parent := succ.parent
	if succ == parent.left {
		parent.left = child
	} else {
		parent.right = child
	}
	child.parent = parent
	if succ.red == false {
		t.deleteBalance(child)
	}
}

func (n *node) print(depth int) {
	if !n.isNil() {
		n.right.print(depth + 1)

		for i := 0; i < depth; i++ {
			fmt.Print("    ")
		}

		if n.red == true {
			fmt.Print("R ")
		} else {
			fmt.Print("B ")
		}
		fmt.Printf("%d\n", n.key)

		n.left.print(depth + 1)
	}
	if depth == 0 {
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	tree := initTree()
	tree.insert(7)
	tree.root.print(0)
	tree.insert(3)
	tree.root.print(0)
	tree.insert(8)
	tree.root.print(0)
	tree.insert(2)
	tree.root.print(0)
	tree.delete(tree.search(3))
	tree.root.print(0)
	tree.delete(tree.search(7))
	tree.root.print(0)
	tree.delete(tree.search(8))
	tree.root.print(0)
	tree.delete(tree.search(2))
	tree.root.print(0)
}
