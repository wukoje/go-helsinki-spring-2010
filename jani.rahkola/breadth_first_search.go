package main

import (
	"fmt"
	"container/list"
)

// Graph as an array of lists. Indices are the nodes of the
// graph, and the list in index i lists all the nodes that i
// is connected to.
type graph struct {
	nodes []*list.List
}

func New() *graph {
	slice := make([]*list.List, 100)
	for i, _ := range slice {
		slice[i] = list.New()
	}
	return &graph{slice}
}

// Checks wheter an edge between x and y exists.
func (g *graph) adjecent(x int, y int) bool {
	if x > len(g.nodes) || y > len(g.nodes) {
		return false
	}
	if x < 0 || y < 0 {
		return false
	}
	if g.nodes[x].Len() == 0 || g.nodes[y].Len() == 0 {
		return false
	}

	for n := range g.nodes[x].Iter() {
		if n.(int) == y {
			return true
		}
	}
	return false
}

// Adds an new edge to the graph
// Negative values are not added.
func (g *graph) add(x int, y int) {
	if g.adjecent(x, y) { // edge already exists
		return
	}
	if x == y {
		return
	}
	if x < 0 || y < 0 {
		return
	}
	if x > len(g.nodes) || y > len(g.nodes) {
		bigger := x
		if y > x {
			bigger = y
		}
		newSlice := make([]*list.List, bigger*2)
		for i, _ := range newSlice {
			newSlice[i] = list.New()
		}
		for i, e := range g.nodes {
			newSlice[i] = e
		}
		g.nodes = newSlice
	}
	g.nodes[x].PushBack(y)
	g.nodes[y].PushBack(x)
}

// A breadth first search for a value.
func (g *graph) search(e int) bool {
	// first the root of the graph needs to be found
	i := 0
	for g.nodes[i].Len() == 0 {
		i++
	}
	// is root the searched element?
	if e == i {
		return true
	}
	queue := list.New()
	queue.PushBackList(g.nodes[i])
	visited := make([]bool, len(g.nodes))
	visited[i] = true
	for queue.Len() != 0 {
		v := queue.Front().Value.(int)
		if v == e {
			return true
		}
		queue.Remove(queue.Front())
		visited[v] = true
		for n := range g.nodes[v].Iter() {
			if !visited[n.(int)] {
				queue.PushBack(n.(int))
			}
		}
	}
	return false
}

func main() {
	graph := New()
	graph.add(6, 4)
	graph.add(4, 5)
	graph.add(4, 3)
	graph.add(5, 1)
	graph.add(5, 2)
	graph.add(3, 2)
	for i := 1; i <= 6; i++ {
		fmt.Println(graph.search(i))
	}
	fmt.Println(graph.search(20))

	graph.add(1, 200)
	fmt.Println(graph.search(200))
	fmt.Println(graph.search(2000))
}
