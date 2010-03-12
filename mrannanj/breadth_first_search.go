package main

import (
	"fmt"
	"rand"
	"time"
	"container/list"
)

type tree struct {
	root *node
}

type node struct {
	key         int
	left, right *node
}

func (t *tree) search(k int) bool {
	for pos := t.root; pos != nil; {
		if pos.key == k {
			return true
		} else if k < pos.key {
			pos = pos.left
		} else {
			pos = pos.right
		}
	}
	return false
}

func (t *tree) insert(k int) bool {
	if t.root == nil {
		t.root = &node{k, nil, nil}
		return true
	}
	pos := t.root
	for k != pos.key {
		if k < pos.key {
			if pos.left != nil {
				pos = pos.left
			} else {
				pos.left = &node{k, nil, nil}
				return true
			}
		} else {
			if pos.right != nil {
				pos = pos.right
			} else {
				pos.right = &node{k, nil, nil}
				return true
			}
		}
	}
	return false
}

func (t *tree) breadthSearch(k int) bool {
	if t.root == nil {
		return false
	}

	queue := new(list.List)
	queue.PushBack(t.root)
	for queue.Len() > 0 {
		qelement := queue.Front()
		cur := qelement.Value.(*node)
		queue.Remove(qelement)

		if cur.key == k {
			return true
		}
		if cur.left != nil {
			queue.PushBack(cur.left)
		}
		if cur.right != nil {
			queue.PushBack(cur.right)
		}
	}
	return false
}

func (t *tree) printTree() {
	if t.root == nil {
		return
	}

	var buf string
	queue := new(list.List)
	queue.PushBack(t.root)
	in, out := 0, 1
	for queue.Len() > 0 {
		qelement := queue.Front()
		cur := qelement.Value.(*node)
		queue.Remove(qelement)

		if cur.left != nil {
			queue.PushBack(cur.left)
			in++
			buf += "_"
		}
		buf += fmt.Sprint(cur.key)
		if cur.right != nil {
			queue.PushBack(cur.right)
			in++
			buf += "_"
		}
		buf += " "
		if out--; out <= 0 {
			out, in = in, out
			buf += "\n"
		}
	}
	fmt.Print(buf)
}

func main() {
	rand.Seed(time.Nanoseconds())
	t := &tree{nil}
	for i := 0; i < 30; i++ {
		t.insert(rand.Intn(30))
	}

	fmt.Println("Tree:")
	t.printTree()

	for i := 0; i < 5; i++ {
		n := rand.Intn(30)
		if t.breadthSearch(n) {
			fmt.Println("Number", n, "is in the tree.")
		} else {
			fmt.Println("Number", n, "is not in the tree.")
		}
	}

}
