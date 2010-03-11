package main

import "fmt"
import "strconv"

type Node struct {
	black               bool
	parent, left, right *Node
	data                int
}

type RedBlackTree struct {
	root *Node
}

func (r *RedBlackTree) pointerFromParent(n *Node) **Node {
	if n.parent == nil {
		return &r.root
	} else if n.parent.left == n {
		return &n.parent.left
	}

	return &n.parent.right
}

func (r *RedBlackTree) rotateLeft(n *Node) {
	pp := r.pointerFromParent(n)
	(*pp) = n.right

	n.right.parent = n.parent
	n.parent = n.right
	n.right = n.parent.left
	if n.right != nil {
		n.right.parent = n
	}
	n.parent.left = n
}

func (r *RedBlackTree) rotateRight(n *Node) {
	pp := r.pointerFromParent(n)
	(*pp) = n.left

	n.left.parent = n.parent
	n.parent = n.left
	n.left = n.parent.right
	if n.left != nil {
		n.left.parent = n
	}
	n.parent.right = n
}

func (r *RedBlackTree) minimum(root *Node) *Node {
	n := root
	for n.left != nil {
		n = n.left
	}
	return n
}

func (r *RedBlackTree) maximum(root *Node) *Node {
	n := root
	for n.right != nil {
		n = n.right
	}
	return n
}

func (r *RedBlackTree) stringify(root *Node, depth int) string {
	s := ""
	if root == nil {
		return s
	}

	s += r.stringify(root.right, depth+1)

	for i := 0; i < depth; i++ {
		s += "\t"
	}

	if root.black {
		s += "(BLACK "
	} else {
		s += "(RED "
	}

	s += strconv.Itoa(root.data) + ")"

	if root.left != nil && root.right != nil {
		s += "<"
	} else if root.left != nil {
		s += "\\"
	} else if root.right != nil {
		s += "/"
	}

	s += "\n"

	s += r.stringify(root.left, depth+1)
	return s
}

func (r *RedBlackTree) findClosest(data int) *Node {
	parent := r.root
	if parent == nil {
		return nil
	}

	for {
		target := parent

		if data > target.data {
			target = target.right
		} else if data < target.data {
			target = target.left
		} else {
			return target
		}

		if target == nil {
			break
		}

		parent = target
	}

	return parent
}

func (r *RedBlackTree) treeInsert(n *Node) {
	if r.root == nil {
		r.root = n
		return
	}

	p := r.findClosest(n.data)
	if p.data == n.data {
		return
	} else if n.data > p.data {
		p.right = n
		n.parent = p
	} else {
		p.left = n
		n.parent = p
	}
}

func (r *RedBlackTree) fixBlackness(n *Node) {
	for n.parent != nil && n.black {

		if n.parent.left == n {

			sibling := n.parent.right
			if sibling != nil && !sibling.black {

				sibling.black = true
				n.parent.black = false
				r.rotateLeft(n.parent)
				sibling = n.parent.right

			}
			if sibling == nil ||
				(sibling.left == nil || sibling.left.black) &&
					(sibling.right == nil || sibling.right.black) {

				if sibling != nil {
					sibling.black = false
				}
				n = n.parent

			} else {

				if sibling == nil || sibling.right == nil || sibling.right.black {
					if sibling != nil {
						if sibling.left != nil {
							sibling.left.black = true
						}
						sibling.black = false
						r.rotateRight(sibling)
						sibling = n.parent.right
					}
				}

				if sibling != nil {
					if sibling.right != nil {
						sibling.right.black = true
					}
					sibling.black = n.parent.black
				}
				n.parent.black = true
				r.rotateLeft(n.parent)
				break

			}

		} else {

			sibling := n.parent.left
			if sibling != nil && !sibling.black {

				sibling.black = true
				n.parent.black = false
				r.rotateRight(n.parent)
				sibling = n.parent.left

			}
			if sibling == nil ||
				(sibling.left == nil || sibling.left.black) &&
					(sibling.right == nil || sibling.right.black) {

				if sibling != nil {
					sibling.black = false
				}
				n = n.parent

			} else {

				if sibling == nil || sibling.left == nil || sibling.left.black {
					if sibling != nil {
						if sibling.right != nil {
							sibling.right.black = true
						}
						sibling.black = false
						r.rotateLeft(sibling)
						sibling = n.parent.left
					}
				}

				if sibling != nil {
					if sibling.left != nil {
						sibling.left.black = true
					}
					sibling.black = n.parent.black
				}
				n.parent.black = true
				r.rotateRight(n.parent)
				break

			}
		}

	}
	n.black = true
}

func (r *RedBlackTree) deleteNode(n *Node) {
	pp := r.pointerFromParent(n)
	if n.left == nil && n.right == nil {

		if n.black {
			r.fixBlackness(n)
		}
		(*pp) = nil

	} else if n.left != nil && n.right != nil {

		next := r.minimum(n.right)
		next.data, n.data = n.data, next.data
		r.deleteNode(next)

	} else {

		if n.left != nil {
			n.left.parent = n.parent
			(*pp) = n.left
		} else {
			n.right.parent = n.parent
			(*pp) = n.right
		}

		if n.black {
			r.fixBlackness(*pp)
		}

	}
}

func (r *RedBlackTree) String() string { return r.stringify(r.root, 0) }

func (r *RedBlackTree) Insert(data int) {
	n := &Node{false, nil, nil, nil, data}
	r.treeInsert(n)

	for n.parent != nil && n.parent.parent != nil && !n.parent.black {

		if n.parent.parent.left == n.parent {

			if uncle := n.parent.parent.right; uncle != nil && !uncle.black {
				n.parent.black = true
				uncle.black = true
				n.parent.parent.black = false
				n = n.parent.parent
			} else {
				if n.parent.right == n {
					r.rotateLeft(n.parent)
					n = n.left
				}

				n.parent.parent.black = false
				n.parent.black = true
				r.rotateRight(n.parent.parent)
				break
			}

		} else {

			if uncle := n.parent.parent.left; uncle != nil && !uncle.black {
				n.parent.black = true
				uncle.black = true
				n.parent.parent.black = false
				n = n.parent.parent
			} else {
				if n.parent.left == n {
					r.rotateRight(n.parent)
					n = n.right
				}

				n.parent.parent.black = false
				n.parent.black = true
				r.rotateLeft(n.parent.parent)
				break
			}

		}

	}

	r.root.black = true
}

func (r *RedBlackTree) Search(data int) bool {
	n := r.findClosest(data)
	return n.data == data
}

func (r *RedBlackTree) Delete(data int) {
	n := r.findClosest(data)
	if n.data != data {
		return
	}

	r.deleteNode(n)
}

func main() {
	insertKeys := []int{55, 612, 12, 33, 9, 15, 12333, 28, 56, 1232, 9929, 99, 223, 562, 8, 1235, 1337, 88}
	searchKeys := []int{1, 2, 3, 4, 5, 6, 15, 1337, 8, 19, 23, 28}
	deleteKeys := []int{28, 133, 15, 9, 9929, 88, 12333, 9, 12, 87}
	tree := &RedBlackTree{nil}

	for _, v := range insertKeys {
		fmt.Printf("\nInserting %d ... \n", v)
		tree.Insert(v)
	}
	fmt.Printf("%v \n", tree)

	for _, v := range searchKeys {
		fmt.Printf("\nSearching for %d ...\n", v)
		if tree.Search(v) {
			fmt.Printf("FOUND %d!\n", v)
		}
	}

	for _, v := range deleteKeys {
		fmt.Printf("\nDeleting %d ... \n", v)
		tree.Delete(v)
	}
	fmt.Printf("%v \n", tree)

}
