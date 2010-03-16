// Issue 493041

package main

import (
	"fmt"
)

func main() {

	// fmt.Println("\nTesting tree initialization:")
	// testTree := new(RedBlackTree)
	// testTree.Insert(5)
	// fmt.Printf("%v", testTree)
	//
	// fmt.Println("\nTesting tree population:")
	// fmt.Println("Inserting:", 4)
	// testTree.Insert(4)
	// fmt.Printf("%v\n", testTree)
	// fmt.Println("Inserting:", 3)
	// testTree.Insert(3)
	// fmt.Printf("%v\n", testTree)
	// fmt.Println("Inserting:", 6)
	// testTree.Insert(6)
	// fmt.Printf("%v\n", testTree)
	// fmt.Println("Inserting:", 8)
	// testTree.Insert(8)
	// fmt.Printf("%v\n", testTree)
	// fmt.Println("Inserting:", 10)
	// testTree.Insert(10)
	// fmt.Printf("%v", testTree)

	testTree := new(RedBlackTree)
	for i := 1; i < 15; i++ {
		fmt.Println("Inserting:", i)
		testTree.Insert(i)
		fmt.Printf("%v\n", testTree)
	}

	fmt.Println("Search:", testTree.Search(5))
	fmt.Println("Search:", testTree.Search(15))
	fmt.Println()

	for i := 1; i < 15; i++ {
		fmt.Println("Deleting:", i)
		testTree.Delete(i)
		fmt.Printf("%v\n", testTree)
	}
	fmt.Println("\nDeleting:", 5)
	testTree.Delete(5)
	fmt.Printf("%v\n", testTree)
}


type RedBlackTree struct {
	root *rbNode
}

type rbNode struct {
	red    bool
	key    int
	parent *rbNode
	left   *rbNode
	right  *rbNode
}

func (tree *RedBlackTree) Search(key int) (target *rbNode) {
	tmpNode := tree.root
	return inorderSearch(tmpNode, key)
}

func inorderSearch(root *rbNode, key int) (target *rbNode) {
	if root == nil {
		return
	}
	if key < root.key {
		target = inorderSearch(root.left, key)
	} else if key > root.key {
		target = inorderSearch(root.right, key)
	} else {
		target = root
	}
	return
}

func (tree *RedBlackTree) Delete(key int) {
	deleteNode := tree.Search(key)
	if deleteNode.numberChildren() == 0 {
		deleteNode.deleteChildless()
	} else if deleteNode.numberChildren() == 1 {
		tree.deleteWithOneChild(deleteNode)
	} else {
		tree.deleteWith2Children(deleteNode)
	}
}

func (node *rbNode) deleteChildless() {
	tmpParent := node.parent
	if tmpParent.left == node {
		tmpParent.left = nil
	} else {
		tmpParent.right = nil
	}
}

func (tree *RedBlackTree) deleteWithOneChild(node *rbNode) {
	var child *rbNode
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}
	replaceNode(node, child)
	if !node.red {
		if child.red {
			child.red = false
		} else {
			tree.deleteCase1(child)
		}
	}
}

func (tree *RedBlackTree) deleteWith2Children(node *rbNode) {
	successor := tree.findMin(node.right)
	replaceNode(node, successor)
	tree.deleteCase1(successor)
}

func (tree *RedBlackTree) findMin(node *rbNode) (minNode *rbNode) {
	minNode = node
	for ; minNode.left != nil; minNode = minNode.left {
	}
	return
}

func (tree *RedBlackTree) deleteCase1(node *rbNode) {
	fmt.Println("Case 1")
	if !node.isRoot() {
		tree.deleteCase2(node)
	}
}

func (tree *RedBlackTree) deleteCase2(node *rbNode) {
	fmt.Println("Case 2")
	if node.sibling().red {
		node.parent.red = true
		node.sibling().red = false
		if node == node.parent.left {
			tree.rotateLeft(node.parent)
		} else {
			tree.rotateRight(node.parent)
		}
	}
	tree.deleteCase3(node)

}

func (tree *RedBlackTree) deleteCase3(node *rbNode) {
	fmt.Println("Case 3")
	if !node.parent.red && !node.sibling().red && !node.sibling().left.red && !node.sibling().right.red {
		node.sibling().red = true
		tree.deleteCase1(node.parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RedBlackTree) deleteCase4(node *rbNode) {
	fmt.Println("Case 4")
	if node.parent.red && !node.sibling().red && !node.sibling().left.red && !node.sibling().right.red {
		node.sibling().red = true
		node.parent.red = false
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RedBlackTree) deleteCase5(node *rbNode) {
	fmt.Println("Case 5")
	if !node.sibling().red {
		if node == node.parent.left && !node.sibling().right.red && node.sibling().left.red {
			node.sibling().red = true
			node.sibling().left.red = false
			tree.rotateRight(node.sibling())
		} else if node == node.parent.right && !node.sibling().left.red && node.sibling().right.red {
			node.sibling().red = true
			node.sibling().right.red = false
			tree.rotateLeft(node.sibling())
		}
	}
	tree.deleteCase6(node)
}

func (tree *RedBlackTree) deleteCase6(node *rbNode) {
	fmt.Println("Case 6")
	node.sibling().red = node.parent.red
	node.parent.red = false

	if node == node.parent.left {
		node.sibling().right.red = false
		tree.rotateLeft(node.parent)
	} else {
		node.sibling().left.red = false
		tree.rotateRight(node.parent)
	}
}

func replaceNode(parent *rbNode, child *rbNode) {
	if parent == parent.parent.left {
		parent.parent.left = child
	} else {
		parent.parent.right = child
	}
	child.parent = parent.parent
}

func (tree *RedBlackTree) Insert(keyToInsert int) {
	newNode := new(rbNode)
	newNode.key = keyToInsert
	newNode.red = true

	if tree.root == nil {
		tree.root = newNode
	} else {
		tmpNode := tree.root
		var tmpParent *rbNode
		for tmpNode != nil {
			tmpParent = tmpNode
			if keyToInsert < tmpNode.key {
				tmpNode = tmpNode.left
			} else {
				tmpNode = tmpNode.right
			}
		}
		newNode.parent = tmpParent
		if keyToInsert < tmpParent.key {
			tmpParent.left = newNode
		} else {
			tmpParent.right = newNode
		}
	}
	tree.insertCase1(newNode)
}

func (tree *RedBlackTree) insertCase1(node *rbNode) {
	if node.isRoot() {
		node.red = false
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlackTree) insertCase2(node *rbNode) {
	if !node.parent.red {
		return
	} else {
		tree.insertCase3(node)
	}
}

func (tree *RedBlackTree) insertCase3(node *rbNode) {
	if node.uncle() != nil && node.uncle().red {
		node.parent.red = false
		node.uncle().red = false
		node.grandparent().red = true
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlackTree) insertCase4(node *rbNode) {
	if node == node.parent.right && node.parent == node.grandparent().left {
		tree.rotateLeft(node.parent)
		node = node.left
	} else if node == node.parent.left && node.parent == node.grandparent().right {
		tree.rotateRight(node.parent)
		node = node.right
	}
	tree.insertCase5(node)
}

func (tree *RedBlackTree) insertCase5(node *rbNode) {
	node.parent.red = false
	node.grandparent().red = true
	if node == node.parent.left && node.parent == node.grandparent().left {
		tree.rotateRight(node.grandparent())
	} else {
		tree.rotateLeft(node.grandparent())
	}
}

func (node *rbNode) numberChildren() (count int) {
	if node.left != nil {
		count++
	}
	if node.right != nil {
		count++
	}
	return
}

func (node *rbNode) uncle() (uncle *rbNode) {
	if node.parent == node.grandparent().left {
		uncle = node.grandparent().right
	} else {
		uncle = node.grandparent().left
	}
	return
}

func (node *rbNode) sibling() (sibling *rbNode) {
	if node == node.parent.left {
		sibling = node.parent.right
	} else {
		sibling = node.parent.left
	}
	return
}

func (node *rbNode) grandparent() (grandparent *rbNode) {
	return node.parent.parent
}

func (node *rbNode) isRoot() bool { return node.parent == nil }

func (tree *RedBlackTree) rotateLeft(root *rbNode) {
	rotateNode := root.right
	if root.isRoot() {
		tree.root = rotateNode
	} else if root == root.parent.left {
		root.parent.left = rotateNode
	} else {
		root.parent.right = rotateNode
	}
	rotateNode.parent = root.parent
	root.right, rotateNode.left = rotateNode.left, root
	root.parent = rotateNode
}

func (tree *RedBlackTree) rotateRight(root *rbNode) {
	rotateNode := root.left
	if root.isRoot() {
		tree.root = rotateNode
	} else if root == root.parent.left {
		root.parent.left = rotateNode
	} else {
		root.parent.right = rotateNode
	}
	rotateNode.parent = root.parent
	root.left, rotateNode.right = rotateNode.right, root
	root.parent = rotateNode
}

func (tree *RedBlackTree) doubleRotateRight(root *rbNode) {
	tree.rotateLeft(root.left)
	tree.rotateRight(root)
}

func (tree *RedBlackTree) doubleRotateLeft(root *rbNode) {
	tree.rotateRight(root.right)
	tree.rotateLeft(root)
}

func (tree *RedBlackTree) String() (result string) {
	result = "Inorder: "
	result += inorderTreeWalk(tree.root)
	result += "\nPre-order: "
	result += preorderTreeWalk(tree.root)
	result += "\nPost-order: "
	result += postorderTreeWalk(tree.root)
	result += "\n"
	return
}

func inorderTreeWalk(n *rbNode) (result string) {
	if n != nil {
		var color string
		if n.red {
			color = "r"
		} else {
			color = "b"
		}
		result += inorderTreeWalk(n.left)
		result += fmt.Sprintf("%d%s ", n.key, color)
		result += inorderTreeWalk(n.right)
	}
	return
}

func preorderTreeWalk(n *rbNode) (result string) {
	if n != nil {
		var color string
		if n.red {
			color = "r"
		} else {
			color = "b"
		}
		result += fmt.Sprintf("%d%s ", n.key, color)
		result += preorderTreeWalk(n.left)
		result += preorderTreeWalk(n.right)
	}
	return
}

func postorderTreeWalk(n *rbNode) (result string) {
	if n != nil {
		var color string
		if n.red {
			color = "r"
		} else {
			color = "b"
		}
		result += postorderTreeWalk(n.left)
		result += postorderTreeWalk(n.right)
		result += fmt.Sprintf("%d%s ", n.key, color)
	}
	return
}
