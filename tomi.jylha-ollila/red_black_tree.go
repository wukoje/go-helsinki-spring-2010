package main


import (
	"fmt"
	"rand"
)


const (
	RED = iota
	BLACK
)


type rbnode struct {
	key    int
	colour int
	parent *rbnode
	left   *rbnode
	right  *rbnode
}


func create_nil() *rbnode {
	nil_node := new(rbnode)
	nil_node.colour = BLACK
	nil_node.parent = nil_node
	nil_node.left = nil_node
	nil_node.right = nil_node
	return nil_node
}


func create_node(key int, nil_node *rbnode) *rbnode {
	node := new(rbnode)
	node.colour = RED
	node.parent = nil_node
	node.left = nil_node
	node.right = nil_node
	node.key = key
	return node
}


func (node *rbnode) is_nil() bool { return node == node.left }


func (node *rbnode) print(level int) {
	if !node.is_nil() {
		node.right.print(level + 1)
		if !node.right.is_nil() && node != node.right.parent {
			fmt.Printf("link failure %d -> %d -> %d\n",
				node.key, node.right.key, node.right.parent.key)
		}
		for i := 0; i < level*2; i++ {
			fmt.Printf(" ")
		}
		if node.colour == RED {
			fmt.Printf("R ")
		} else {
			fmt.Printf("B ")
		}
		fmt.Printf("%d\n", node.key)
		if !node.left.is_nil() && node != node.left.parent {
			fmt.Printf("link failure %d -> %d -> %d\n",
				node.key, node.left.key, node.left.parent.key)
		}
		node.left.print(level + 1)
	}
	if level == 0 {
		fmt.Printf("\n")
	}
}


type rbtree struct {
	root     *rbnode
	nil_node *rbnode
}


func create_rbtree() *rbtree {
	tree := new(rbtree)
	tree.nil_node = create_nil()
	tree.root = tree.nil_node
	return tree
}


func (tree *rbtree) rotate_left(node *rbnode) {
	pivot := node.right
	node.right = pivot.left
	if !pivot.left.is_nil() {
		pivot.left.parent = node
	}
	parent := node.parent
	pivot.parent = parent
	if parent.is_nil() {
		tree.root = pivot
	} else if parent.left == node {
		parent.left = pivot
	} else {
		parent.right = pivot
	}
	pivot.left = node
	node.parent = pivot
}


func (tree *rbtree) rotate_right(node *rbnode) {
	pivot := node.left
	node.left = pivot.right
	if !pivot.right.is_nil() {
		pivot.right.parent = node
	}
	parent := node.parent
	pivot.parent = parent
	if parent.is_nil() {
		tree.root = pivot
	} else if parent.right == node {
		parent.right = pivot
	} else {
		parent.left = pivot
	}
	pivot.right = node
	node.parent = pivot
}


func (tree *rbtree) insert_fix(node *rbnode) {
	for node.parent.colour == RED {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle.colour == RED {
				node.parent.colour = BLACK
				uncle.colour = BLACK
				node.parent.parent.colour = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.rotate_left(node)
				}
				node.parent.colour = BLACK
				node.parent.parent.colour = RED
				tree.rotate_right(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if uncle.colour == RED {
				node.parent.colour = BLACK
				uncle.colour = BLACK
				node.parent.parent.colour = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rotate_right(node)
				}
				node.parent.colour = BLACK
				node.parent.parent.colour = RED
				tree.rotate_left(node.parent.parent)
			}
		}
	}
	tree.root.colour = BLACK
}


func (tree *rbtree) insert(key int) *rbnode {
	new_node := create_node(key, tree.nil_node)
	node := tree.root
	prev := tree.nil_node
	for !node.is_nil() {
		prev = node
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	new_node.parent = prev
	if prev.is_nil() {
		tree.root = new_node
	} else if key < prev.key {
		prev.left = new_node
	} else {
		prev.right = new_node
	}
	tree.insert_fix(new_node)
	return new_node
}


func (tree *rbtree) delete_fix(node *rbnode) {
	for node != tree.root && node.colour == BLACK {
		if node == node.parent.left {
			sibling := node.parent.right
			if sibling.colour == RED {
				sibling.colour = BLACK
				node.parent.colour = RED
				tree.rotate_left(node.parent)
				sibling = node.parent.right
			}
			if sibling.left.colour == BLACK && sibling.right.colour == BLACK {
				sibling.colour = RED
				node = node.parent
			} else {
				if sibling.right.colour == BLACK {
					sibling.left.colour = BLACK
					sibling.colour = RED
					tree.rotate_right(sibling)
					sibling = node.parent.right
				}
				sibling.colour = node.parent.colour
				node.parent.colour = BLACK
				sibling.right.colour = BLACK
				tree.rotate_left(node.parent)
				node = tree.root
			}
		} else {
			sibling := node.parent.left
			if sibling.colour == RED {
				sibling.colour = BLACK
				node.parent.colour = RED
				tree.rotate_right(node.parent)
				sibling = node.parent.left
			}
			if sibling.left.colour == BLACK && sibling.right.colour == BLACK {
				sibling.colour = RED
				node = node.parent
			} else {
				if sibling.left.colour == BLACK {
					sibling.right.colour = BLACK
					sibling.colour = RED
					tree.rotate_left(sibling)
					sibling = node.parent.left
				}
				sibling.colour = node.parent.colour
				node.parent.colour = BLACK
				sibling.left.colour = BLACK
				tree.rotate_right(node.parent)
				node = tree.root
			}
		}
	}
	node.colour = BLACK
}


func (tree *rbtree) delete(node *rbnode) {
	if node == nil || node.is_nil() {
		return
	}
	if node.left.is_nil() || node.right.is_nil() {
		child := node.left
		if child.is_nil() {
			child = node.right
		}
		parent := node.parent
		if parent.is_nil() {
			tree.root = child
		} else if node == parent.left {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent
		if node.colour == BLACK {
			tree.delete_fix(child)
		}
		return
	}
	succ := node.right
	for !succ.left.is_nil() {
		succ = succ.left
	}
	node.key, succ.key = succ.key, node.key
	child := succ.right
	parent := succ.parent
	if succ == parent.left {
		parent.left = child
	} else {
		parent.right = child
	}
	child.parent = parent
	if succ.colour == BLACK {
		tree.delete_fix(child)
	}
}


func (tree *rbtree) search(key int) *rbnode {
	node := tree.root
	for !node.is_nil() {
		if key == node.key {
			return node
		} else if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	return nil
}


func (tree *rbtree) print() { tree.root.print(0) }


func main() {
	tree := create_rbtree()
	tree.insert(3)
	tree.print()
	tree.insert(5)
	tree.print()
	tree.insert(6)
	tree.print()
	tree.insert(4)
	tree.print()
	tree.delete(tree.search(5))
	tree.print()
	tree.delete(tree.search(3))
	tree.print()
	tree.delete(tree.search(6))
	tree.print()
	tree.delete(tree.search(4))
	tree.print()
	for i := 0; i < 30; i++ {
		tree.insert(rand.Intn(100))
	}
	for i := 0; i < 10; i++ {
		tree.insert(45 + i)
	}
	for i := 0; i < 10; i++ {
		tree.delete(tree.search(45 + i))
	}
	tree.print()
}
