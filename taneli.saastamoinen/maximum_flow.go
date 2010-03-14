package main

import (
	"fmt"
)

type graph struct {
	n		 int
	adj	 []int
	cap  []int
	flow []int
}

func create(i int) *graph {
	g := new(graph)
	g.n = i
	g.adj = make([]int, i*i)
	g.cap = make([]int, i*i)
	g.flow = make([]int, i*i)
	return g
}

func (g *graph) addEdge(i, j, w int) {
	if i >= 0 && j >= 0 && i < g.n && j < g.n {
		g.adj[i*g.n+j] = 1
		g.cap[i*g.n+j] = w
	}
}

func (g *graph) neighbours(i int) []int {
	if i < 0 || i >= g.n {
		return nil
	}
	count := 0
	for j := 0; j < g.n; j++ {
		if g.adj[i*g.n+j] == 1 {
			count++
		}
		if g.adj[j*g.n+i] == 1 {
			count++
		}
	}
	v := make([]int, count)
	c := 0
	for j := 0; j < g.n; j++ {
		if g.adj[i*g.n+j] == 1 {
			v[c] = j
			c++
		}
		if g.adj[j*g.n+i] == 1 {
			v[c] = j
			c++
		}
	}
	return v
}

type pathEdge struct {
	i    int
	j    int
	resi int
}

func makePathEdge(i, j, r int) *pathEdge {
	e := new(pathEdge)
	e.i = i
	e.j = j
	e.resi = r
	return e
}

func (p *pathEdge) equals(q pathEdge) bool {
	return p.i == q.i && p.j == q.j && p.resi == q.resi
}

func (g *graph) findPath(s, t int, path []pathEdge) []pathEdge {
	if s < 0 || t < 0 || s >= g.n || t >= g.n {
		return nil
	}
	if s == t {
		return path
	}
	if ns := g.neighbours(s); ns != nil {
		for _, n := range ns {
			cap := g.cap[s*g.n + n]
			resi := cap - g.flow[s*g.n + n]
			edge := makePathEdge(s, n, resi)
			edgeDone := false
			for _, z := range path {
				if edge.equals(z) {
					edgeDone = true
					break
				}
			}
			if resi > 0 && !edgeDone {
				p := make([]pathEdge, len(path)+1)
				copy(p, path)
				p[len(p)-1] = *edge
				if r := g.findPath(n, t, p); r != nil {
					return r
				}
			}
		}
	}
	return nil
}

func (g *graph) maxFlow(s, t int) int {
	path := g.findPath(s, t, make([]pathEdge, 0))
	for path != nil {
		flow := 1 << 31 - 1
		for _, edge := range path {
			if edge.resi < flow {
				flow = edge.resi
			}
		}
		for _, edge := range path {
			u := edge.i
			v := edge.j
			g.flow[u*g.n + v] += flow
			g.flow[v*g.n + u] -= flow
		}
		path = g.findPath(s, t, make([]pathEdge, 0))
	}
	totalFlow := 0
	for _, e := range g.neighbours(s) {
		totalFlow += g.flow[s*g.n + e]
	}
	return totalFlow
}

func (g *graph) String() string {
	s := "\n"
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			s += fmt.Sprintf("%v ", g.adj[i*g.n+j])
		}
		s += "   "
		for j := 0; j < g.n; j++ {
			s += fmt.Sprintf("%v ", g.cap[i*g.n+j])
		}
		s += "   "
		for j := 0; j < g.n; j++ {
			s += fmt.Sprintf("%v ", g.flow[i*g.n+j])
		}
		s += "\n"
	}
	return s
}

func main() {
	g := create(6)
	g.addEdge(0, 1, 3)
	g.addEdge(0, 2, 3)
	g.addEdge(1, 2, 2)
	g.addEdge(1, 3, 3)
	g.addEdge(2, 4, 2)
	g.addEdge(3, 4, 4)
	g.addEdge(3, 5, 2)
	g.addEdge(4, 5, 3)
	fmt.Printf("(this example taken from http://en.wikipedia.org/wiki/Maximum_flow_problem )\n")
	fmt.Printf("graph - adjacency, capacity, flow: %v\n", g)
	fmt.Printf("running Ford-Fulkerson with DFS...\n")
	f := g.maxFlow(0, 5)
	fmt.Printf("maximum flow from vertex #0 to vertex #5 is %v\n", f)
	fmt.Printf("graph with full flow: %v\n", g)
}

