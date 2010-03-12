package main

import (
	"container/list"
	"fmt"
)

type Node interface { }

// Directed unweighted graph
type Graph struct {
	adjacent map[Node](*list.List)
}

func New() (g *Graph) {
	return &Graph{make(map[Node](*list.List))}
}

func (g *Graph) AddEdge(from, to Node) {
	if _, present := g.adjacent[from]; !present {
		g.adjacent[from] = list.New()
	}
	g.adjacent[from].PushBack(to)
}

func (g *Graph) GetAdjacents(from Node) (<-chan interface{}) {
	if _, present := g.adjacent[from]; !present {
		return list.New().Iter()
	} else {
		return g.adjacent[from].Iter()
	}
	panic("unreachable")
}

func (g *Graph) BFS(root Node) (result *Graph) {
	result = New()
	checked := make(map[Node](bool))
	checked[root] = true
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		u := queue.Front()
		queue.Remove(u)
		uv := u.Value
		for v := range g.GetAdjacents(uv) {
			if _, present := checked[v]; !present {
				checked[v] = true
				result.AddEdge(uv, v)
				queue.PushBack(v)
			}
		}
	}
	return
}

func (g *Graph) String() (result string) {
	for node, adjacent := range g.adjacent {
		result += fmt.Sprintf("%v: ", node)
		for a := range adjacent.Iter() {
			result += fmt.Sprintf("%v ", a)
		}
		result += fmt.Sprint("\n")
	}
	return
}

func main() {
	g := New()
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(5, 4)
	g.AddEdge(2, 5)
	fmt.Printf("%v", g)
	fmt.Println("Result:")
	t := g.BFS(2)
	fmt.Printf("%v", t)
}
