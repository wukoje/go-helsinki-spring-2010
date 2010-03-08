package main

import (
	"fmt"
	"container/list"
)

type node struct {
	id int           // node id
	edges *list.List // adjacency list of edges
}

type graph struct {
	root *node
}

// breadth first search of a node from graph based on node id
func (g *graph) bfs(id int) *node {
	queue := list.New()

	enqueue := func(n *node) {
		queue.PushBack(n)
	}

	dequeue := func() *node {
		f := queue.Front()
		if f == nil {
			return nil
		}
		queue.Remove(f)
		return f.Value.(*node)
	}

	enqueue(g.root)

	visited := map[int] bool {}
	for queue.Front() != nil {
		n := dequeue()
		visited[n.id] = true

		if n != nil && n.id == id {
			return n
		}

		for edge := range n.edges.Iter() {
			en := edge.(*node)
			if _, ok := visited[en.id]; !ok {
				enqueue(edge.(*node))
			}
		}
	}

	return nil
}

func main() {
	g := new(graph)
	n1 := &node { 1, list.New() }
	n2 := &node { 2, list.New() }
	n3 := &node { 3, list.New() }

	g.root = n1

	n1.edges.PushBack(n2)
	n2.edges.PushBack(n3)
	n2.edges.PushBack(n1)
	n3.edges.PushBack(n1)

	fmt.Println(g.bfs(1))
	fmt.Println(g.bfs(2))
	fmt.Println(g.bfs(3))
}
