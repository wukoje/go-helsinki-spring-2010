// Algorithm refers to "Introduction to Algorithms 2nd Edition" Chapter 13, p238~p261
// Modification: use single rb_rotate operation

package main

import "fmt"

//Define colors type
type colors int
const ( red colors = iota; black; )

//Define braches
type branches int
const ( left branches = iota; right; )

//Define tree structure
type rbTree struct {
	root *tNode
	nil  *tNode
}

//Define node structure
type tNode struct {
	value int
	left *tNode
	right *tNode 
	parent *tNode
	color colors
}

func rb_create_tree() *rbTree {
	myTree := new(rbTree)
	myTree.nil = new(tNode)
	myTree.root = myTree.nil
	myTree.nil.color = black
	return myTree
}

func (myTree *rbTree) rb_output_tree(p *tNode, level int) {
	if p == myTree.nil {
		return
	}
	fmt.Print("|")
	for i := 0; i < 4*level; i++ {
		fmt.Print("_")
	}
	if p.color==black { fmt.Print("BLACK:") }
	if p.color==red { fmt.Print("RED:") }
	
	fmt.Println(p.value)
	myTree.rb_output_tree(p.left, level+1)
	myTree.rb_output_tree(p.right, level+1)
}

func isFound(p *tNode, pTree *rbTree) (b bool) {
	if p != pTree.nil {
		fmt.Println("Node is found!")
		b = true
	} else {
		fmt.Println("Node is not found!")
		b = false
	}
	return b
}

func (myTree *rbTree) rb_rotate(m *tNode, branch branches) {
	switch branch {
		case left:
			n := m.right
			m.right = n.left
			if n.left != myTree.nil {
				n.left.parent = m
			}
			n.parent = m.parent
			switch {
				case m.parent == myTree.nil: myTree.root = n
				case m.parent.left == m: m.parent.left = n
				default: m.parent.right = n
			}
			m.parent, n.left = n, m
		case right:
			n := m.left
			m.left = n.right
			if n.right != myTree.nil {
				n.right.parent = m
			}
			n.parent = m.parent
			switch {
				case m.parent == myTree.nil: myTree.root = n
				case m.parent.right == m: m.parent.right = n
				default: m.parent.left = n
			}
			m.parent, n.right = n, m
	}
}

func (myTree *rbTree) rb_search(value int) (p *tNode) {
	for p=myTree.root; p != myTree.nil && p.value != value; {
		switch {
			case p.value > value: p = p.left
			default: p = p.right
		}
	}
	return p
}

func (myTree *rbTree) rb_insert(value int) {
	var m, n *tNode
	n = myTree.nil
	p := new(tNode)
	p.value = value

	for m = myTree.root; m != myTree.nil; {
		n = m
		if m.value < p.value  {
			m = m.right
		} else {
			m = m.left
		}
	}
	p.parent = n
	if p.parent == myTree.nil {
		myTree.root = p
	} else if n.value < p.value {
		n.right = p
	} else {
		n.left = p
	}
	p.left, p.right = myTree.nil, myTree.nil
	p.color = red
	myTree.rb_insert_fix(p)
}

func (myTree *rbTree) rb_insert_fix(p *tNode) {
	for p.parent.color==red {
		if p.parent == p.parent.parent.left {
			n := p.parent.parent.right
			switch n.color {
				case red:
					p.parent.color = black
					n.color = black
					p.parent.parent.color = red
					p = p.parent.parent
				case black:
					if p == p.parent.right {
						p = p.parent
						myTree.rb_rotate(p, left)
					}
					p.parent.color = black
					p.parent.parent.color = red
					myTree.rb_rotate(p.parent.parent, right)
			}
		} else {
			n := p.parent.parent.left
			switch n.color {
				case red:
					p.parent.color = black
					n.color = black
					p.parent.parent.color = red
					p = p.parent.parent
				case black:
					if p == p.parent.left {
						p = p.parent
						myTree.rb_rotate(p, right)
					}
					p.parent.color = black
					p.parent.parent.color = red
					myTree.rb_rotate(p.parent.parent, left)
			}
		}
	}
	myTree.root.color = black
}

func (myTree *rbTree) rb_delete(p *tNode) *tNode {
	var m, n *tNode
	if p.left == myTree.nil || p.right == myTree.nil {
		n = p
	} else {
		if p.right != myTree.nil {
			for n = n.right; n.left != myTree.nil; n = n.left {}
		} else {
			for n = p.parent; n != myTree.nil && p == n.right; p,n = n,n.parent {}
		}
	}
	if n.left != myTree.nil {
		m = n.left
	} else {
		m = n.right
	}
	m.parent = n.parent
	if n.parent == myTree.nil {
		myTree.root = m
	} else {
		if n == n.parent.left {
			n.parent.left = m
		} else {
			n.parent.right = m
		}
	}
	if n != p {
		p.value = n.value
	}
	if n.color==black {
		myTree.rb_delete_fix(m)
	}
	return n
}

func (myTree *rbTree) rb_delete_fix(m *tNode) {
	var n *tNode
	for m != myTree.root && m.color == black {
		if m == m.parent.left {
			n = m.parent.right
			if n.color==red {
				n.color, m.parent.color = black, red
				myTree.rb_rotate(m.parent, left)
				n = m.parent.right
			}
			if n.left.color==black && n.right.color==black {
				n.color = red
				m = m.parent
			} else {
				if n.right.color==black {
					n.left.color, n.color = black, red
					myTree.rb_rotate(n, right)
					n = m.parent.right
				}
				n.color = m.parent.color
				m.parent.color, n.right.color = black, black
				myTree.rb_rotate(m.parent, left)
				m = myTree.root
			}
		} else {
			n = m.parent.left
			if n.color==red {
				n.color, m.parent.color = black, red
				myTree.rb_rotate(m.parent, right)
				n = m.parent.left
			}
			if n.left.color==black && n.right.color==black {
				n.color = red
				m = m.parent
			} else {
				if n.left.color==black {
					n.right.color, n.color = black, red
					myTree.rb_rotate(n, left)
					n = m.parent.left
				}
				n.color = m.parent.color
				m.parent.color, n.left.color = black, black
				myTree.rb_rotate(m.parent, right)
				m = myTree.root
			}
		}
	}
	m.color = black
}

func main() {
	myTree := rb_create_tree()
	for i := 0; i < 16; i++ { myTree.rb_insert(i) }

	fmt.Println("Tree Structure:")
	myTree.rb_output_tree(myTree.root, 0)

	p := myTree.rb_search(2)
	fmt.Println("Searching 2 ...")
	if isFound(p, myTree) { myTree.rb_delete(p) }
	
	p = myTree.rb_search(99)
	fmt.Println("Searching 99 ...")
	if isFound(p, myTree) { myTree.rb_delete(p) }
}

