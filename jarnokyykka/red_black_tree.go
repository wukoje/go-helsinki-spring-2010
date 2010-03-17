/* 
   Pyörää ei keksitty uudelleen, pohjana käytettiin ratkaisua mikä löytyy: 
   http://en.literateprograms.org/Red-black_tree_(Java)
*/


package main

import (
	"fmt"
)

type node struct {
	left *node
	right *node	
	parent *node 
	value int
	color string
}

type tree struct {
	root *node
}

func main() {
	t := new(tree)

	t.insert(50)
	t.insert(75)
	t.insert(25)
	t.insert(100)
	t.insert(125)
	t.insert(88)
	t.insert(49)
	t.insert(52)
	t.insert(54)
	t.insert(111)
	t.insert(29)
	t.printTree()

	t.delete(100)
	fmt.Printf("----------\n")
	t.printTree()
}


func (t *tree) insert(v int) {
	newNode := new(node)
	newNode.color = "r"
	newNode.value = v
	if t.root != nil {
		//no duplicates allowed
		if t.search(v) == nil {
			t.root.insertTree(newNode)
			t.insert_case1(newNode)	
		}
	} else {
		t.root = newNode
		t.insert_case1(newNode)
	}
}

func (leaf *node) insertTree(n *node){
	if leaf.value > n.value {
		if leaf.left != nil {
			leaf.left.insertTree(n)
		} else {
			leaf.left = n
			n.parent = leaf
		}
	} else if n.value >= leaf.value {
		if leaf.right != nil {
			leaf.right.insertTree(n)
		} else {
			leaf.right = n
			n.parent = leaf
		}
		
	}
}

func (t *tree) search(v int) *node {
	return t.root.search(v)
}

func (n *node) search(v int) *node {
	if n.value == v {
		return n
	}
	if n.value > v && n.left != nil {
		return n.left.search(v)
	} 
	if n.value < v && n.right != nil {
		return n.right.search(v)
	}
	return nil
}

func (t *tree) delete(v int) {
	n := t.search(v)
	var pred *node	
	var child *node

	if n == nil {
		return
	}
	if n.left != nil && n.right != nil {
		pred = n.left.maximumNode()
		n.value = pred.value
	}

	//child = (n.left == nil) ? n.left : n.right

	if pred.right == nil {
		child = pred.left
	} else {
		child = pred.right
	}

	if n.color == "b" && child != nil {
		n.color = child.color	
		t.delete_case1(n)
	}
	

	replace_node(t, pred, child)
	
	if t.root.color == "r" {
		t.root.color = "b"
	}
}

func (t *tree) delete_case1(n *node) {
	if n.parent == nil {
		return
	} else {
		t.delete_case2(n)
	}
}

func (t *tree) delete_case2(n *node) {
	if n.sibling().color == "r" {
		n.parent.color = "r"
		n.sibling().color = "b"
	}
	if n == n.parent.left {
		rotate_left(t, n.parent)
	} else {
		rotate_right(t, n.parent)
	}


	t.delete_case3(n)
}

func (t *tree) delete_case3(n *node) {
	if n.parent.color == "b" && n.sibling().color == "b" && n.sibling().left.color == "b" && n.sibling().right.color == "b" {
		n.sibling().color = "r"
		t.delete_case1(n.parent)
	} else {
		t.delete_case4(n)
	}

}

func (t *tree) delete_case4(n *node) {
	if n.parent.color == "r" && n.sibling().color == "b" && n.sibling().left.color == "b" && n.sibling().right.color == "b" {
		n.sibling().color = "r"
		n.parent.color = "b"
	} else {
		t.delete_case5(n)
	}


}

func (t *tree) delete_case5(n *node) {
	if n == n.parent.left && n.sibling().color == "b" && n.sibling().left.color == "r" && n.sibling().color == "r" {
		n.sibling().color = "r"
		n.sibling().left.color = "b"
		rotate_right(t, n.sibling())
	} else if n == n.parent.right && n.sibling().color == "b" && n.sibling().right.color == "r" && n.sibling().left.color == "b" {
		n.sibling().color = "r"
		n.sibling().right.color = "b"
		rotate_left(t, n.sibling())
	}

	t.delete_case6(n)
      


}

func (t *tree) delete_case6(n *node) {
	n.sibling().color = n.parent.color
	n.parent.color = "b"
	if n == n.parent.left {
		n.sibling().right.color = "b"
		rotate_left(t, n.parent)
	} else {
		n.sibling().left.color = "b"
		rotate_right(t, n.parent)
	}

}


func (n *node) maximumNode() *node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) grandparent() *node {
	if (n != nil) && (n.parent != nil) && (n.parent.parent != nil) {
		return n.parent.parent
	}
	return nil
}

func (n *node) uncle() *node {
	g := n.grandparent()
	if g == nil || n == nil {
		return nil
	}
	if n.parent == g.left {
		return g.right
	} else {
 		return g.left
	}
	return nil
}

func (n *node) sibling() *node {
	if n.parent == nil {
		return nil
	}		
	if n == n.parent.left {
		return n.parent.right
	} else {
		return n.parent.right
	}
	return nil
}

func rotate_left(t *tree, n *node) {
	r := n.right
	replace_node(t, n, r)
	n.right = r.left
	if r.left != nil {
		r.left.parent = n
	}
	r.left = n
	n.parent = r
}

func rotate_right(t *tree, n *node) {
	l := n.left
	replace_node(t, n, l)
	n.left = l.right
	if(l.right != nil) {
		l.right.parent = n
	}
	l.right = n
	n.parent = l
}

func replace_node(t *tree, oldn *node, newn *node) {
	if oldn.parent == nil {
		t.root = newn
	} else {
		if oldn == oldn.parent.left {
			oldn.parent.left = newn
		} else {
			oldn.parent.right = newn
		}
	}
	if newn != nil {
		newn.parent = oldn.parent
	}
}

func (t *tree) insert_case1(n *node) {
	if n.parent == nil {
		n.color = "b"
	} else {
		t.insert_case2(n)
	}
}

func (t *tree) insert_case2(n *node) {
	if n.parent.color == "b" {
		//tree ok
		return
	} else {
		t.insert_case3(n)
	}
}

func (t *tree) insert_case3(n *node) {
	if n.uncle() != nil {
		if n.uncle().color == "r" {
			n.parent.color = "b"
			n.uncle().color = "b"
			n.grandparent().color = "r"
			t.insert_case1(n.grandparent())
		} else {
			t.insert_case4(n)
		}
	} else {
		t.insert_case4(n)
	}
}

func (t *tree) foo(a int) {
	fmt.Printf("---- %d -----\n", a)
	t.printTree()
	fmt.Printf("-----------.\n")
}

func (t *tree) insert_case4(n *node) {
	if n == n.parent.right && n.parent == n.grandparent().left {
		rotate_left(t, n.parent)
		n = n.left
	} else if n == n.parent.left && n.parent == n.grandparent().right {
		rotate_right(t, n.parent)
		n = n.right
	}
	t.insert_case5(n)
}


func (t *tree) insert_case5(n *node) {
	n.parent.color = "b"
	n.grandparent().color = "r"
	
	if n == n.parent.left && n.parent == n.grandparent().left {
		rotate_right(t, n.grandparent())
	} else {
		rotate_left(t, n.grandparent()) 
	}
}

//prints tree sideways
func (t *tree) printTree() {
	fmt.Printf("-----\n")
	printNode(t.root, 0)
	fmt.Printf("-----\n")
}

func printNode(n *node, indent int) {
	if n == nil {
		fmt.Printf("no values")
		return
	}
	if n.right != nil {
		printNode(n.right, indent + 4)
	}
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	if n.color == "r" {
		fmt.Printf("r%d\n", n.value)
	} else {
		fmt.Printf("b%d\n", n.value)
	} 
	if n.left != nil {
		printNode(n.left, indent + 4)
	}
}
