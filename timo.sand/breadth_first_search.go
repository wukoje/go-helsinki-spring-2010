package main

import(
	"fmt"
)

var open chan node

func main() {
	
	fmt.Println("\nTesting tree creation and node assignment")
	testTree := new(Tree)
	testTree.Insert(5)
	fmt.Printf("%v", testTree)
	
	fmt.Println("\nTesting insert on tree")
	testTree.Insert(4)
	fmt.Printf("%v", testTree)
	testTree.Insert(8)
	fmt.Printf("%v", testTree)
	testTree.Insert(3)
	fmt.Printf("%v", testTree)
	
	if search(testTree, 3) != nil {
		fmt.Println("Yay")
	}

	if search(testTree, 10) != nil {
		fmt.Println("Yay")
	}
}

func search(t *Tree, value int) *node {
	t.resizeChannel(open)
	fmt.Println("Channel size:", cap(open))
	open <- *t.root
	OK := true
	var tmpNode node
	for OK {
		tmpNode, OK = <- open
		if tmpNode.key == value {
			return &tmpNode
		} else {
			if tmpNode.left != nil	{
				open <- *tmpNode.left
			} else if tmpNode.right != nil {
				open <- *tmpNode.right
			}
		}
	}
	
	fmt.Println("No match found for:", value)
	
	return nil
}

type Tree struct {
	root *node
	height int
	nodeCount int
}

type node struct {
	key int
	parent *node
	left *node
	right *node
}

func (t *Tree) Insert(value int) {
	newNode := new(node)
	newNode.key = value
	if t.root == nil {
		t.root = newNode
	} else {
		tmpNode := t.root
		var tmpParent *node
		for tmpNode != nil {
			tmpParent = tmpNode
			if value < tmpNode.key {
				tmpNode = tmpNode.left
			} else {
				tmpNode = tmpNode.right
			}
		}
		newNode.parent = tmpParent
		if value < tmpParent.key {
			tmpParent.left = newNode
		} else {
			tmpParent.right = newNode
		}
	}
	t.nodeCount++
}

func (t *Tree) String() (result string) {
	result = "Inorder: "
	result += inorderTreeWalk(t.root)
	result += "\nPre-order: "	
	result += preorderTreeWalk(t.root)
	result += "\nPost-order: "	
	result += postorderTreeWalk(t.root)
	result += "\n"
	return
}

func inorderTreeWalk(n *node) (result string) {
	if n != nil {
		result += inorderTreeWalk(n.left)
		result += fmt.Sprint(n.key, " ")
		result += inorderTreeWalk(n.right)
	} 
	return
}

func preorderTreeWalk(n *node) (result string) {
	if n != nil {
		result += fmt.Sprint(n.key, " ")
		result += preorderTreeWalk(n.left)
		result += preorderTreeWalk(n.right)
	}
	return
}

func postorderTreeWalk(n *node) (result string) {
	if n != nil {
		result += postorderTreeWalk(n.left)
		result += postorderTreeWalk(n.right)
		result += fmt.Sprint(n.key, " ")
	} 
	return
}

func (t *Tree) resizeChannel(chan node) {
	channelSize := t.nodeCount/2+1
	open = make(chan node, channelSize)
}


