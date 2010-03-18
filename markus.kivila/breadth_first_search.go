/*
1 breadth_first_search.go

Define a graph data structure and implement a function that performs a 
breadth first search for a certain node.
*/
package main


import (
	"fmt"
	"container/heap"
	"container/vector"
)

type edge struct {
	from *node
	to *node
//	weigth int
}

// node structure stores also all search related info. 
// This results in additional cleanup phase after the search
// (assuming that same graph will be used in another search)
type node struct {
	edges vector.Vector
	name int
	visited bool
	num int
	prev *node
}

// vector takes care of the heap interface
// we only need to add Less function
type nodeQueue struct {
	vector.Vector
}

func (n *nodeQueue) Less(i, j int) bool {
	return n.At(i).(*node).num < n.At(j).(*node).num
}

func newNode(name int) *node {
	fmt.Printf("Creating node %v\n", name)
	return &node{vector.Vector{}, name, false, -1, nil}
}

func (n *node) addEdge(to *node) {
	fmt.Printf("Created edge %v --> %v\n", n.name, to.name)
	n.edges.Push(&edge{n, to})
}

// search function marks the results to the given nodes
func bfs(root, target *node) *vector.Vector {
	queue := &nodeQueue{vector.Vector{}}
	root.num = 0
	heap.Push(queue, root)

	for queue.Len() > 0 {
		n := heap.Pop(queue).(*node)
		n.visited = true

		fmt.Printf("Visited %v\n", n.name)

		if n == target {
			path := vector.Vector{}
			path.Push(n)
			for n.prev != nil {
				n = n.prev
				path.Push(n)
			}
			return &path
		}
		for ei := range n.edges.Iter() {
			e := ei.(*edge)
			if !e.to.visited {
				e.to.prev = n
				e.to.num = n.num + 1
				heap.Push(queue, e.to)
			}
		}
	}
	return nil // never reached
}

// search without side effects
// prints the found path and cleans up the graph
func searchPath(nodes vector.Vector, from, to int) {
	path := bfs(nodes[from].(*node), nodes[to].(*node))

	for path.Len() > 0 {
		fmt.Printf("-> %v ", path.Pop().(*node).name)
	}
	fmt.Printf("\n -- Done --\nCleaning up..\n")
	clearNodes(nodes)
}

func clearNodes(nodes vector.Vector) {
	for ni := range nodes.Iter() {
		n := ni.(*node)
		n.prev = nil
		n.visited = false
		n.num = -1
	}
}

func main() {
	fmt.Printf("Creating graph...\n")
	/*
			0
		   / \
		  1   2
		 /\  / \ 
        3  4 5  6
		   |
		   7
	*/
	nodes := vector.Vector{}
	for i:=0; i<8; i++ {
		nodes.Push(newNode(i))
	}
	for i:=0; i<8; i++ {
		n := nodes[i].(*node)
		switch i {
			case 0:
				n.addEdge(nodes[1].(*node))
				n.addEdge(nodes[2].(*node))
			case 1:
				n.addEdge(nodes[0].(*node))
				n.addEdge(nodes[3].(*node))
				n.addEdge(nodes[4].(*node))
			case 2:
				n.addEdge(nodes[0].(*node))
				n.addEdge(nodes[5].(*node))
				n.addEdge(nodes[6].(*node))
			case 3:
				n.addEdge(nodes[1].(*node))
			case 4:
				n.addEdge(nodes[1].(*node))
				n.addEdge(nodes[7].(*node))
			case 5:
				n.addEdge(nodes[2].(*node))
			case 6:
				n.addEdge(nodes[2].(*node))
			case 7:
				n.addEdge(nodes[4].(*node))
		}
	}
	fmt.Printf("Path from 0 to 7:\n")
	searchPath(nodes, 0, 7)
	fmt.Printf("Path from 0 to 0:\n")
	searchPath(nodes, 0, 0)
	fmt.Printf("Path from 6 to 7:\n")
	searchPath(nodes, 6, 7)
	fmt.Printf("Path from 1 to 2:\n")
	searchPath(nodes, 1, 2)
}
