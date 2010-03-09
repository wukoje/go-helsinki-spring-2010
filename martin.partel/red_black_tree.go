package main

/*
 * Rotation madness!
 * Implementation based on old lecture notes, which were based on Cormen et al.
 *
 * There is also a main() with test code, a fairly complete sanity check
 * method used by main() and a nifty graphviz printer.
 */

import "os"
import "fmt"

type color int

const (
	red   color = color(iota)
	black color = color(iota)
)

func (c color) String() string {
	switch c {
	case red:
		return "red"
	case black:
		return "black"
	}
	return "bad-color"
}

type RBTree struct {
	root    *rbNode
	nilNode *rbNode // A sentinel node to avoid a bunch of null checks
	// The links on nil may be changed in edge cases
	size uint
}

type rbNode struct {
	color  color
	parent *rbNode
	left   *rbNode
	right  *rbNode
	key    int
}

func (n *rbNode) grandparent() *rbNode { return n.parent.parent }

func NewRBTree() *RBTree {
	nilNode := new(rbNode)
	nilNode.color = black
	nilNode.parent = nilNode
	nilNode.left = nilNode
	nilNode.right = nilNode
	return &RBTree{nilNode, nilNode, 0}
}

func (tree *RBTree) Size() uint { return tree.size }

func (tree *RBTree) Insert(key int) {
	parent := tree.nilNode
	current := tree.root
	for current != tree.nilNode {
		parent = current
		if key < current.key {
			current = current.left
		} else {
			current = current.right
		}
	}

	newNode := &rbNode{red, parent, tree.nilNode, tree.nilNode, key}

	if parent == tree.nilNode {
		tree.root = newNode
	} else {
		if key < parent.key {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}

	tree.insertFixup(newNode)
	tree.size++
}

func (tree *RBTree) insertFixup(node *rbNode) {
	// Fixup loop invariant: the only error in the tree is that
	// node and its parent are both red
	for node.parent.color == red {

		if node.parent == node.grandparent().left {
			if uncle := node.grandparent().right; uncle.color == red {
				// Case 1: the parent and uncle are red
				// => move redness up to grandparent
				node.parent.color = black
				uncle.color = black
				node.grandparent().color = red
				node = node.grandparent()
			} else {
				if node == node.parent.right {
					// Case 2: the parent is red, the uncle is black and
					// node == grandparent.left.right
					// => convert to case 3
					node = node.parent
					tree.leftRotate(node)
				}
				// Case 3: the parent is red, the uncle is black and
				// node == grandparent.left.left
				// => move redness from parent to grand parent and
				// rotate it to the right.
				node.parent.color = black
				node.grandparent().color = red
				tree.rightRotate(node.grandparent())
			}

		} else {

			// As above with left <-> right
			if uncle := node.grandparent().left; uncle.color == red {
				// Case 1
				node.parent.color = black
				uncle.color = black
				node.grandparent().color = red
				node = node.grandparent()
			} else {
				if node == node.parent.left {
					// Case 2
					node = node.parent
					tree.rightRotate(node)
				}
				// Case 3
				node.parent.color = black
				node.grandparent().color = red
				tree.leftRotate(node.grandparent())
			}
		}
	}

	tree.root.color = black // In case the redness was moved to the root
}

func (tree *RBTree) Delete(key int) {
	node := tree.findNode(key)
	if node != nil {
		tree.deleteNode(node)
	}
}

func (tree *RBTree) findNode(key int) *rbNode {
	node := tree.root
	for node != tree.nilNode {
		if key == node.key {
			return node
		} else {
			if key < node.key {
				node = node.left
			} else {
				node = node.right
			}
		}
	}
	return nil
}

func (tree *RBTree) deleteNode(node *rbNode) {
	var nodeToSplice *rbNode
	if node.left == tree.nilNode || node.right == tree.nilNode {
		nodeToSplice = node
	} else {
		nodeToSplice = tree.successor(node)
	}
	// nodeToSplice can have at most one child

	var child *rbNode
	if nodeToSplice.left != tree.nilNode {
		child = nodeToSplice.left
	} else {
		child = nodeToSplice.right
	}

	child.parent = nodeToSplice.parent

	if nodeToSplice == tree.root {
		tree.root = child
	} else {
		if nodeToSplice == nodeToSplice.parent.left {
			nodeToSplice.parent.left = child
		} else {
			nodeToSplice.parent.right = child
		}
	}

	if nodeToSplice != node {
		node.key = nodeToSplice.key
	}

	if nodeToSplice.color == black {
		tree.deleteFixup(child) // child may be nil, but child.parent is set
	}

	tree.size--
}

func (tree *RBTree) successor(node *rbNode) *rbNode {
	if node.right != tree.nilNode {
		return tree.subtreeMinimum(node.right)
	}

	parent := node.parent
	for parent != tree.nilNode && node == parent.right {
		node = parent
		parent = node.parent
	}
	return parent
}

func (tree *RBTree) subtreeMinimum(node *rbNode) *rbNode {
	for node.left != tree.nilNode {
		node = node.left
	}
	return node
}

func (tree *RBTree) deleteFixup(node *rbNode) {
	for node != tree.root && node.color == black {
		if node == node.parent.left {

			sibling := node.parent.right
			if sibling.color == red {
				// Case 1: red sibling
				// => convert to one of the cases with black sibling
				sibling.color = black
				node.parent.color = red
				tree.leftRotate(node.parent)
				sibling = node.parent.right
			}

			if sibling.left.color == black && sibling.right.color == black {
				// Case 2: both nephews are black
				// => make sibling red, need extra black in parent
				// instead of only left subtree
				sibling.color = red
				node = node.parent
			} else {
				if sibling.right.color == black {
					// Case 3: only right nephew is black
					// => transform into case 4
					sibling.left.color = black
					sibling.color = red
					tree.rightRotate(sibling)
					sibling = node.parent.right
				}

				// Case 4: right nephew is red
				// => magic that kills the extra black we carry around
				sibling.color = node.parent.color
				node.parent.color = black
				sibling.right.color = black
				tree.leftRotate(node.parent)
				node = tree.root
			}

		} else {
			// As above with left <-> right

			sibling := node.parent.left
			if sibling.color == red {
				// Case 1
				sibling.color = black
				node.parent.color = red
				tree.rightRotate(node.parent)
				sibling = node.parent.left
			}

			if sibling.right.color == black && sibling.left.color == black {
				// Case 2
				sibling.color = red
				node = node.parent
			} else {
				if sibling.left.color == black {
					// Case 3
					sibling.right.color = black
					sibling.color = red
					tree.leftRotate(sibling)
					sibling = node.parent.left
				}

				// Case 4
				sibling.color = node.parent.color
				node.parent.color = black
				sibling.left.color = black
				tree.rightRotate(node.parent)
				node = tree.root
			}

		}
	}

	node.color = black
}

func (tree *RBTree) leftRotate(node *rbNode) {
	oldRight := node.right
	oldParent := node.parent

	node.right = oldRight.left
	if oldRight.left != tree.nilNode {
		oldRight.left.parent = node
	}

	oldRight.left = node
	oldRight.parent = oldParent
	node.parent = oldRight

	if oldParent == tree.nilNode {
		tree.root = oldRight
	} else {
		if node == oldParent.left {
			oldParent.left = oldRight
		} else {
			oldParent.right = oldRight
		}
	}
}

func (tree *RBTree) rightRotate(node *rbNode) {
	oldLeft := node.left
	oldParent := node.parent

	node.left = oldLeft.right
	if oldLeft.right != tree.nilNode {
		oldLeft.right.parent = node
	}

	oldLeft.right = node
	oldLeft.parent = oldParent
	node.parent = oldLeft

	if oldParent == tree.nilNode {
		tree.root = oldLeft
	} else {
		if node == oldParent.right {
			oldParent.right = oldLeft
		} else {
			oldParent.left = oldLeft
		}
	}
}


func (tree *RBTree) Search(key int) bool {
	node := tree.findNode(key)
	return node != nil
}


func (tree *RBTree) Verify() os.Error {
	switch {
	case tree.root.color != black:
		return os.NewError("Root not black")
	case tree.nilNode.color != black:
		return os.NewError("nilNode.color modified")
	}

	if tree.root != tree.nilNode {
		_, _, bstError := tree.checkBinarySearchCondition(tree.root)
		if bstError != nil {
			return bstError
		}
	}

	var linkError os.Error = nil
	tree.inorderWalk(tree.root, func(n *rbNode) {
		if n.left != tree.nilNode && n.left.parent != n {
			linkError = os.NewError(fmt.Sprint(n.left.key) + " has incorrect parent link")
		}
		if n.right != tree.nilNode && n.right.parent != n {
			linkError = os.NewError(fmt.Sprint(n.right.key) + " has incorrect parent link")
		}
	})
	if linkError != nil {
		return linkError
	}

	nodeCount := uint(0)
	tree.inorderWalk(tree.root, func(*rbNode) { nodeCount++ })
	if nodeCount != tree.size {
		return os.NewError(fmt.Sprintf("Expected size %d but found %d nodes", tree.size, nodeCount))
	}

	var colorError os.Error = nil
	tree.inorderWalk(tree.root, func(n *rbNode) {
		if n.color == red {
			switch red {
			case n.left.color:
				colorError = os.NewError("Left child of red was red")
			case n.right.color:
				colorError = os.NewError("Right child of red was red")
			}
		}
	})
	if colorError != nil {
		return colorError
	}

	_, blackError := tree.checkBlackCounts(tree.root)

	if blackError != nil {
		return blackError
	}

	return nil
}

func (tree *RBTree) inorderWalk(node *rbNode, f func(*rbNode)) {
	if node != tree.nilNode {
		tree.inorderWalk(node.left, f)
		f(node)
		tree.inorderWalk(node.right, f)
	}
}

func (tree *RBTree) checkBinarySearchCondition(node *rbNode) (*int, *int, os.Error) {
	var min, max *int
	var err os.Error

	if node.left != tree.nilNode {
		_, max, err = tree.checkBinarySearchCondition(node.left)
		if err != nil {
			return nil, nil, err
		}
		if max != nil && *max > node.key {
			return nil, nil, os.NewError(fmt.Sprintf("%d found under %d", max, node.key))
		}
	} else {
		max = &node.key
	}
	if node.right != tree.nilNode {
		min, _, err = tree.checkBinarySearchCondition(node.right)
		if err != nil {
			return nil, nil, err
		}
		if min != nil && *min < node.key {
			return nil, nil, os.NewError(fmt.Sprintf("%d found under %d", min, node.key))
		}
	} else {
		min = &node.key
	}

	return min, max, nil
}

func (tree *RBTree) checkBlackCounts(node *rbNode) (int, os.Error) {
	if node != tree.nilNode {
		leftCount, leftErr := tree.checkBlackCounts(node.left)
		rightCount, rightErr := tree.checkBlackCounts(node.right)

		var err os.Error = nil
		if leftErr != nil {
			err = leftErr
		} else if rightErr != nil {
			err = rightErr
		}

		total := leftCount + rightCount
		if node.color == black {
			total++
		}

		return total, err
	}
	// Weirdly enough I couldn't do "} else { return ... }" here.
	// "function ends without a return statement" it said.

	return 1, nil
}


func main() {
	tree := NewRBTree()

	testValues := []int{201, 360, 698, 636, 996, 571, 640, 750, 935, 928, 791, 582, 273, 531, 400, 448, 559, 900, 863, 796, 732, 433, 542, 376, 392, 589, 942, 125, 785, 211, 496, 697, 283, 16, 945, 301, 254, 86, 568, 496, 959, 677, 918, 689, 342, 435, 263, 751, 365, 574, 346, 274, 15, 516, 852, 694, 412, 898, 642, 360, 789, 599, 361, 548, 299, 6, 124, 996, 884, 242, 923, 436, 905, 707, 986, 10, 284, 124, 640, 853, 487, 73, 191, 678, 202, 412, 467, 159, 994, 760, 385, 197, 858, 501, 879, 508, 68, 633, 428, 309}
	outsideValues := []int{266, 205, 972, 72, 650, 485, 45, 944, 55, 871, 79, 521, 422, 291, 973, 257, 182, 463, 893, 644, 773, 278, 630, 966, 688, 511, 661, 598, 150, 189, 779, 554, 815, 334, 482, 602, 566, 539, 28, 820, 788, 127, 476, 951, 595, 763, 240, 268, 175, 432}
	valuesInDeleteOrder := []int{750, 159, 707, 284, 412, 599, 201, 400, 16, 254, 365, 900, 789, 508, 385, 791, 879, 898, 642, 582, 273, 568, 342, 448, 996, 86, 918, 863, 574, 124, 760, 412, 942, 309, 531, 694, 346, 274, 698, 283, 697, 124, 301, 548, 361, 436, 299, 392, 211, 15, 905, 853, 125, 559, 487, 191, 10, 928, 68, 852, 732, 242, 640, 360, 677, 994, 689, 986, 202, 73, 785, 935, 197, 360, 496, 996, 923, 589, 796, 501, 959, 435, 376, 542, 571, 6, 858, 516, 945, 636, 467, 263, 633, 884, 640, 751, 428, 678, 496, 433}
	// duplicates: []int{360, 996, 640, 496, 496, 412, 360, 124, 996, 124, 640, 412}

	drawAfterEachInsert := false
	drawAfterEachDelete := false
	drawFullTree := false

	for i, e := range testValues {
		fmt.Printf("Adding %d\n", e)
		tree.Insert(e)

		if drawAfterEachInsert {
			tree.DrawGraphviz(fmt.Sprintf("rbtree_insert_%d.png", i))
		}

		err := tree.Verify()
		if err != nil {
			fmt.Printf("Darn: %s\n", err.String())
			bail(tree)
		}
	}

	if drawFullTree {
		err := tree.DrawGraphviz("rbtree.png")
		if err != nil {
			fmt.Printf("failed to draw tree: %s\n", err.String())
		}
	}

	for _, e := range testValues {
		fmt.Printf("Searching for %d...", e)
		if tree.Search(e) {
			fmt.Printf("OK, found\n")
		} else {
			fmt.Printf("Drat, not found\n")
			bail(tree)
		}
	}

	for _, e := range outsideValues {
		fmt.Printf("Searching for %d...", e)
		if !tree.Search(e) {
			fmt.Printf("OK, not found\n")
		} else {
			fmt.Printf("Drat, found, shouldn't have\n")
			bail(tree)
		}
	}

	for i, e := range valuesInDeleteOrder {
		fmt.Printf("Deleting %d\n", e)
		tree.Delete(e)

		if drawAfterEachDelete {
			tree.DrawGraphviz(fmt.Sprintf("rbtree_delete_%d.png", i))
		}

		err := tree.Verify()
		if err != nil {
			fmt.Printf("Darn: %s\n", err.String())
			bail(tree)
		}
	}

	if tree.Size() != 0 {
		fmt.Printf("Oh shoot! The tree is not empty.\n")
		bail(tree)
	}

	fmt.Printf("All done.\n")
}

func bail(tree *RBTree) {
	tree.DrawGraphviz("rbtree_epic_fail.png")
	os.Exit(1)
}

func (tree *RBTree) DrawGraphviz(filename string) os.Error {
	read, write, err := os.Pipe()
	if err != nil {
		return err
	}

	// $PATH lookup doesn't work so I hard-coded it :/
	argv := []string{"/usr/bin/dot", "-Tpng", "-o" + filename}
	fds := []*os.File{read, os.Stdout, os.Stderr}
	pid, err := os.ForkExec("/usr/bin/dot", argv, os.Environ(), "", fds)

	(*os.File).Close(read)

	if err != nil {
		(*os.File).Close(read)
		(*os.File).Close(write)
		return err
	}

	write.WriteString(tree.GraphvizInput())
	(*os.File).Close(write)

	os.Wait(pid, 0)

	return nil
}

func (tree *RBTree) GraphvizInput() string {
	nodeDecls := ""
	nextNodeNum := 0
	createNodeId := func(node *rbNode) string {
		nodeId := fmt.Sprintf("node%d_%d", nextNodeNum, node.key)
		nextNodeNum++
		attrs := ""
		attrs += "fillcolor=" + node.color.String()
		if node == tree.nilNode {
			attrs += ",shape=rectangle,label=NIL"
		} else {
			attrs += ",label=" + fmt.Sprint(node.key)
		}
		nodeDecls += nodeId + " [" + attrs + "]"
		return nodeId
	}

	nodeIds := map[*rbNode]string{}
	nodeId := func(node *rbNode) string {
		if node != tree.nilNode {
			id, ok := nodeIds[node]
			if !ok {
				id = createNodeId(node)
				nodeIds[node] = id
			}

			return id
		}
		return createNodeId(node)
	}

	edgeDecls := ""
	tree.inorderWalk(tree.root, func(node *rbNode) {
		edgeDecls += fmt.Sprintf("%s -- %s\n", nodeId(node), nodeId(node.left))
		edgeDecls += fmt.Sprintf("%s -- %s\n", nodeId(node), nodeId(node.right))
	})

	result := "graph {\n"
	result += "graph [ordering=out]\n"
	result += "node [shape=circle,style=filled,fontcolor=white]\n"
	result += nodeDecls
	result += edgeDecls
	result += "}\n"
	return result
}
