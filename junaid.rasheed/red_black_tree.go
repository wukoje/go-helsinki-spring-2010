package main

import "fmt"


const (
	Red   bool = true
	Black bool = false
)

type Red_Black_Node struct {
	data        int
	left, right *Red_Black_Node
	color       bool
}

type Red_Black_Tree struct {
	root *Red_Black_Node
}

func (tree *Red_Black_Tree) insert(data int) {
	node := new(Red_Black_Node)
	node.data = data
	node.color = Red
	tree.root = balanceInsert(tree.root, node)
}

func balanceInsert(root *Red_Black_Node, node *Red_Black_Node) *Red_Black_Node {
	if root == nil {
		return node
	}
	if root.data < node.data {
		root.left = balanceInsert(root.left, node)
	} else {
		root.right = balanceInsert(root.right, node)
	}
	return fixUp(root)
}

func (tree *Red_Black_Tree) delete(data int) { tree.root = balanceDelete(tree.root, data) }

func balanceDelete(node *Red_Black_Node, data int) *Red_Black_Node {
	if node == nil {
		return nil
	}
	if node.data < data {
		if node.left != nil && !isRed(node.left) && !isRed(node.left.left) {
			node = flipcolorRotateLeft(node)
		}
		node.left = balanceDelete(node.left, data)
	} else {
		if isRed(node.left) {
			node = rotateRight(node)
		}
		if node.data == data && node.right == nil {
			return nil
		}
		if node.right != nil && !isRed(node.right) && !isRed(node.right.left) {
			changecolor(node)
			if isRed(node.left.left) {
				node = rotateRight(node)
				changecolor(node)
			}
		}
		if node.data == data {
			for node.right.left != nil {
				node.right = node.right.left
			}
			node.data = node.right.data
			node.right = deleteMax(node.right)
		} else {
			node.right = balanceDelete(node.right, data)
		}
	}
	return fixUp(node)
}

func deleteMax(node *Red_Black_Node) *Red_Black_Node {
	if node.left == nil {
		return nil
	}
	if !isRed(node.left) && !isRed(node.left.left) {
		node = flipcolorRotateLeft(node)
	}
	node.left = deleteMax(node.left)
	return fixUp(node)
}

func isRed(node *Red_Black_Node) bool { return node != nil && node.color == Red }

func changecolor(node *Red_Black_Node) {
	node.color = !node.color
	node.left.color, node.right.color = !node.left.color, !node.right.color
}

func fixUp(node *Red_Black_Node) *Red_Black_Node {
	if node.left != nil && node.right != nil {
		if isRed(node.right) && !isRed(node.left) {
			node = rotateLeft(node)
		}
		if isRed(node.left) && isRed(node.left.left) {
			node = rotateRight(node)
		}
		if isRed(node.left) && isRed(node.right) {
			changecolor(node)
		}
	}
	return node
}

func rotateLeft(node *Red_Black_Node) *Red_Black_Node {

	node.color, node.right.color = Red, node.color
	node.right, node.right.left = node.right.left, node
	return node.right
}

func rotateRight(node *Red_Black_Node) *Red_Black_Node {
	node.color, node.left.color = Red, node.color
	node.left, node.left.right = node.right, node
	return node.left
}

func flipcolorRotateLeft(node *Red_Black_Node) *Red_Black_Node {
	changecolor(node)
	if isRed(node.right.left) {
		node.right = rotateRight(node.right)
		node = rotateLeft(node)
		changecolor(node)
	}
	return node
}


func (tree *Red_Black_Tree) search(data int) (int, bool) {
	root := tree.root
	for root != nil {
		if root.data < data {
			root = root.left
		} else if root.data == data {
			return root.data, true
		} else {
			root = root.right
		}
	}
	return 0, false
}

func printTree(node *Red_Black_Node) {
	if node != nil {
		color := ""
		if isRed(node) {
			color = "Red"
		} else {
			color = "Black"
		}
		printTree(node.left)
		fmt.Printf("Node: %v Color: %v\n", node.data, color)
		printTree(node.right)
	}
}


func main() {
	var t Red_Black_Tree
	t.insert(10)
	t.insert(1)
	t.insert(11)
	t.insert(2)
	t.insert(6)
	t.insert(9)
	t.insert(15)

	record, _ := t.search(4)

	if record != 0 {
		fmt.Printf("Record found\n")
	} else {
		fmt.Printf("No such record\n")
	}
	color := ""
	if isRed(t.root) {
		color = "Red"
	} else {
		color = "Black"
	}

	fmt.Printf("Root: %v Color: %v\n", t.root.data, color)
	printTree(t.root)

	t.delete(9)
	fmt.Printf("\nPrint after deletion\n")
	fmt.Printf("Root: %v Color: %v\n", t.root.data, color)
	printTree(t.root)
}

