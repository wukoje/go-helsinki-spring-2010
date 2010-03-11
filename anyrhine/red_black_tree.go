/*
 *	Left-leaning Red-Black Tree (LLRB)
 *
 *	Closely follows Robert Sedgewick's paper and slides
 *	which were riddled with miniscule bugs
 */

package main

import (
	"fmt"
	"os"
	"bufio"
)

const (
	Red   bool = true
	Black bool = false
)

type Node struct {
	key         string
	val         string
	left, right *Node
	col         bool
}

type Tree struct {
	root *Node
}

func red(p *Node) bool { return p != nil && p.col == Red }

func rotleft(h *Node) *Node {
	x := h.right
	h.col, x.col = Red, h.col
	h.right, x.left = x.left, h
	return x
}

func rotright(h *Node) *Node {
	x := h.left
	h.col, x.col = Red, h.col
	h.left, x.right = x.right, h
	return x
}

func flipcol(h *Node) {
	h.col = !h.col
	h.left.col, h.right.col = !h.left.col, !h.right.col
}

func fixup(h *Node) *Node {
	if red(h.right) && !red(h.left) {
		h = rotleft(h)
	}
	if red(h.left) && red(h.left.left) {
		h = rotright(h)
	}
	if red(h.left) && red(h.right) {
		flipcol(h)
	}
	return h
}

func ins(h *Node, n *Node) *Node {
	if h == nil {
		return n
	}
	if h.key == n.key {
		h.val = n.val
	} else if h.key < n.key {
		h.left = ins(h.left, n)
	} else {
		h.right = ins(h.right, n)
	}
	return fixup(h)
}

func redleft(h *Node) *Node {
	flipcol(h)
	if red(h.right.left) {
		h.right = rotright(h.right)
		h = rotleft(h)
		flipcol(h)
	}
	return h
}

func redright(h *Node) *Node {
	flipcol(h)
	if red(h.left.left) {
		h = rotright(h)
		flipcol(h)
	}
	return h
}

func findmax(h *Node) *Node {
	for h.left != nil {
		h = h.left
	}
	return h
}

func delmax(h *Node) *Node {
	if h.left == nil {
		return nil
	}
	if !red(h.left) && !red(h.left.left) {
		h = redleft(h)
	}
	h.left = delmax(h.left)
	return fixup(h)
}

func del(h *Node, key string) *Node {
	if h == nil {
		return nil
	}
	if h.key < key {
		if h.left != nil && !red(h.left) && !red(h.left.left) {
			h = redleft(h)
		}
		h.left = del(h.left, key)
	} else {
		if red(h.left) {
			h = rotright(h)
		}
		if h.key == key && h.right == nil {
			return nil
		}
		if h.right != nil && !red(h.right) && !red(h.right.left) {
			redright(h)
		}
		if h.key == key {
			m := findmax(h.right)
			h.key = m.key
			h.val = m.val
			h.right = delmax(h.right)
		} else {
			h.right = del(h.right, key)
		}
	}
	return fixup(h)
}

func (t *Tree) lookup(key string) (string, bool) {
	p := t.root
	for p != nil {
		if p.key < key {
			p = p.left
		} else if p.key == key {
			return p.val, true
		} else {
			p = p.right
		}
	}
	return "", false
}

func (t *Tree) insert(key string, val string) {
	n := new(Node)
	n.key = key
	n.val = val
	n.col = Red
	t.root = ins(t.root, n)
}

func (t *Tree) delete(key string) { t.root = del(t.root, key) }

func oktree(p *Node) int {
	if p == nil {
		return 1
	} else {
		if red(p) {
			if red(p.left) || red(p.right) {
				fmt.Print("R")
			}
		}
		hi := oktree(p.left)
		hj := oktree(p.right)
		if (p.left != nil && p.left.key <= p.key) || (p.right != nil && p.right.key >= p.key) {
			fmt.Print("O")
		}
		if hi != 0 && hj != 0 {
			if hi != hj {
				fmt.Print("b")
			}
			if red(p) {
				return hi
			} else {
				return hi + 1
			}
		}
	}
	return 0
}

func (t *Tree) ok() {
	if t.root != nil {
		t.root.col = Black
	}
	oktree(t.root)
}

func main() {
	var t Tree
	n := 0
	if dbfile, err := os.Open("broken_telephone.words", os.O_RDONLY, 0); err == nil {
		r := bufio.NewReader(dbfile)
		for i := 0; ; i++ {
			if s, err := r.ReadString('\n'); err == nil {
				key := i
				val := s[0 : len(s)-1]
				//fmt.Printf("key %d val %s\n", key, val)
				t.insert(fmt.Sprintf("%d", key), val)
				n++
				//t.ok()
			} else {
				break
			}
		}
	} else {
		fmt.Printf("open:%s\n", err)
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		if _, ok := t.lookup(fmt.Sprintf("%d", i)); !ok {
			fmt.Printf("bug %d\n", i)
		}
	}
	for i := 0; i < n; i++ {
		// erase keys in "random" order..
		k := (i * 30011) % n
		t.delete(fmt.Sprintf("%d", k))
		t.ok()
	}
	if t.root != nil {
		fmt.Print("epic fail\n")
	}
}
