package main


import (
	"fmt"
	"rand"
)


type bintree struct {
	key      int
	parent   *bintree
	children []*bintree
}


func create_bintree(key int) *bintree {
	bt := new(bintree)
	bt.key = key
	bt.children = make([]*bintree, 0, 4)
	return bt
}


func (bt *bintree) merge(other *bintree) *bintree {
	new_bt := bt
	if other.key < bt.key {
		new_bt = other
		other = bt
	}
	if len(new_bt.children) == cap(new_bt.children) {
		children := make([]*bintree, len(new_bt.children)*2)
		for i, e := range new_bt.children {
			children[i] = e
		}
		new_bt.children = children[0:len(new_bt.children)]
	}
	new_bt.children = new_bt.children[0 : len(new_bt.children)+1]
	new_bt.children[len(new_bt.children)-1] = other
	other.parent = new_bt
	return new_bt
}


type binheap struct {
	trees []*bintree
}


func create_binheap() *binheap {
	bh := new(binheap)
	bh.trees = make([]*bintree, 48)
	return bh
}


func (bh *binheap) merge(other *binheap) {
	new_bh := create_binheap()
	for i, _ := range bh.trees {
		switch {
		case bh.trees[i] == nil && other.trees[i] != nil:
			if new_bh.trees[i] != nil {
				new_bh.trees[i+1] = new_bh.trees[i].merge(other.trees[i])
				new_bh.trees[i] = nil
			} else {
				new_bh.trees[i] = other.trees[i]
			}
		case other.trees[i] == nil && bh.trees[i] != nil:
			if new_bh.trees[i] != nil {
				new_bh.trees[i+1] = new_bh.trees[i].merge(bh.trees[i])
				new_bh.trees[i] = nil
			} else {
				new_bh.trees[i] = bh.trees[i]
			}
		case bh.trees[i] != nil && other.trees[i] != nil:
			new_bh.trees[i+1] = bh.trees[i].merge(other.trees[i])
		}
	}
	bh.trees = new_bh.trees
}


func (bh *binheap) insert(item int) *bintree {
	heap := create_binheap()
	tree := create_bintree(item)
	heap.trees[0] = tree
	bh.merge(heap)
	return tree
}


func (bh *binheap) pop() int {
	var smallest_key int
	smallest_index := -1
	for i, tree := range bh.trees {
		if tree != nil {
			if smallest_index < 0 || tree.key < smallest_key {
				smallest_key = tree.key
				smallest_index = i
			}
		}
	}
	if smallest_index < 0 {
		return -1
	}
	orphans := bh.trees[smallest_index].children
	bh.trees[smallest_index] = nil
	for _, orphan := range orphans {
		if orphan != nil {
			heap := create_binheap()
			heap.trees[len(orphan.children)] = orphan
			bh.merge(heap)
		}
	}
	return smallest_key
}


func (bh *binheap) decrease_key(node *bintree, key int) {
	if node.key < key {
		return
	}
	node.key = key
	parent := node.parent
	for parent != nil && parent.key > node.key {
		node.key, parent.key = parent.key, node.key
		node = parent
		parent = node.parent
	}
}


func main() {
	bh := create_binheap()
	fmt.Printf("Inserting")
	for i := 0; i < 29; i++ {
		num := rand.Intn(100)
		fmt.Printf(" %d", num)
		bh.insert(num)
	}
	node := bh.insert(100)
	fmt.Printf(" %d", node.key)
	fmt.Printf("\ndecrease %d to -1", node.key)
	bh.decrease_key(node, -1)
	fmt.Printf("\nPopping")
	for i := 0; i < 30; i++ {
		num := bh.pop()
		fmt.Printf(" %d", num)
	}
	fmt.Printf("\n")
}
