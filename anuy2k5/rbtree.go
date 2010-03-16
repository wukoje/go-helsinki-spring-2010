package main

import "fmt"

type Red_Black_Node struct {
 value int
 left, right *Red_Black_Node
 color bool
}

type Tree struct {
 root *Red_Black_Node
}
const (
 Red bool = true
 Black bool = false
)

func (tree *Tree) insert(value int) {
 node := new(Red_Black_Node)
 node.value = value
 node.color = Red
 tree.root = nodeInsert(tree.root, node)
}

func nodeInsert(root *Red_Black_Node, node *Red_Black_Node) *Red_Black_Node {
 if root == nil {
 return node
 }
 if root.value < node.value {
 root.left = nodeInsert(root.left, node)
 } else {
 root.right = nodeInsert(root.right, node)
 }
 return fixUp(root)
}

func (tree *Tree) delete(value int) { tree.root = nodeDelete(tree.root, value) }

func nodeDelete(node *Red_Black_Node, value int) *Red_Black_Node {
 if node == nil {
 return nil
 }
 if node.value < value {
 if node.left != nil && !isRed(node.left) && !isRed(node.left.left) {
 node = flip_color_rotate_left(node)
 }
 node.left = nodeDelete(node.left, value)
 } else {
 if isRed(node.left) {
 node = rotateRight(node)
 }
 if node.value == value && node.right == nil {
 return nil
 }
 if node.right != nil && !isRed(node.right) && !isRed(node.right.left) {
 changecolor(node)
 if isRed(node.left.left) {
 node = rotateRight(node)
 changecolor(node)
 }
 }
 if node.value == value {
 for node.right.left != nil {
 node.right = node.right.left
 }
 node.value = node.right.value
 node.right = deleteMax(node.right)
 } else {
 node.right = nodeDelete(node.right, value)
 }
 }
 return fixUp(node)
}

func deleteMax(node *Red_Black_Node) *Red_Black_Node {
 if node.left == nil {
 return nil
 }
 if !isRed(node.left) && !isRed(node.left.left) {
 node = flip_color_rotate_left(node)
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

func flip_color_rotate_left(node *Red_Black_Node) *Red_Black_Node {
 changecolor(node)
 if isRed(node.right.left) {
 node.right = rotateRight(node.right)
 node = rotateLeft(node)
 changecolor(node)
 }
 return node
}


func (tree *Tree) search(value int) (int, bool) {
 root := tree.root
 for root != nil {
 if root.value < value {
 root = root.left
 } else if root.value == value {
 return root.value, true
 } else {
 root = root.right
 }
 }
 return 0, false
}

func output_Tree(node *Red_Black_Node) {
 if node != nil {
 color := ""
 if isRed(node) {
 color = "Red"
 } else {
 color = "Black"
 }
 output_Tree(node.left)
 fmt.Printf("Node: %v Color  >> %v\n", node.value, color)
 output_Tree(node.right)
 }
}


func main() {
 var rb Tree
 fmt.Printf(" Inserting 2,8,3,6,9,7,11 \n")
 rb.insert(2) 
 rb.insert(8)
 rb.insert(3)
 rb.insert(6)
 rb.insert(9)
 rb.insert(7)
 rb.insert(11)


 record, _ := rb.search(6)

 if record != 0 {
 fmt.Printf("Record Exists\n")
 } else {
 fmt.Printf("Record doesnt exist\n")
 }
 color := ""
 if isRed(rb.root) {
 color = "Red"
 } else {
 color = "Black"
 }

 fmt.Printf("Root: %v Color: %v\n", rb.root.value, color)
 output_Tree(rb.root)

 rb.delete(1)
 fmt.Printf("\nPrint after deletion\n")
 fmt.Printf("Root: %v Color: %v\n", rb.root.value, color)
 output_Tree(rb.root)
}
 
