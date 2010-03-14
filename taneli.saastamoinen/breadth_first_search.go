package main

import (
	"fmt"
	"container/vector"
)

type graph struct {
	n int
	m []int
}

func create(i int) *graph {
	g := new(graph)
	g.n = i
	g.m = make([]int, i*i)
	return g
}

func (g *graph) addEdge(i, j int) {
	if i >= 0 && j >= 0 && i < g.n && j < g.n {
		g.m[i*g.n+j] = 1
		g.m[j*g.n+i] = 1
	}
}

func (g *graph) neighbours(i int) *vector.IntVector {
	if i >= 0 && i < g.n {
		v := new(vector.IntVector)
		for j := 0; j < g.n; j++ {
			if g.m[i*g.n+j] == 1 {
				v.Push(j)
			}
		}
		return v
	}
	return nil
}

func (g *graph) String() string {
	s := "\n"
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			s += fmt.Sprintf("%v ", g.m[i*g.n+j])
		}
		s += "\n"
	}
	return s
}

func (g *graph) bfs(i int) ([]int, []int) {
	if i < 0 || i >= g.n {
		return nil, nil
	}
	colour := make([]int, g.n)
	distance := make([]int, g.n)
	parent := make([]int, g.n)
	colour[i] = 1
	distance[i] = 0
	parent[i] = -1
	v := new(vector.IntVector)
	v.Push(i)
	for v.Len() > 0 {
		n := v.At(0)
		v.Delete(0)
		for _, e := range g.neighbours(n).Data() {
			if colour[e] == 0 {
				colour[e] = 1
				distance[e] = distance[n] + 1
				parent[e] = n
				v.Push(e)
			}
		}
		colour[n] = 2
	}
	return distance, parent
}

func main() {
	g := create(5)
	g.addEdge(0, 1)
	g.addEdge(0, 4)
	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 3)
	g.addEdge(3, 4)
	fmt.Printf("graph: %v\n", g)
	fmt.Printf("breadth-first search from node #4...\n")
	d, p := g.bfs(4)
	fmt.Printf("distances: %v, parents: %v\n", d, p)
	fmt.Printf("breadth-first search from node #1...\n")
	d, p = g.bfs(1)
	fmt.Printf("distances: %v, parents: %v\n", d, p)
	fmt.Printf("breadth-first search from node #2...\n")
	d, p = g.bfs(2)
	fmt.Printf("distances: %v, parents: %v\n", d, p)
}

