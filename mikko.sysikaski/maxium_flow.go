// Calculates maximum-flow using push-relabel-algorithm.
// Time complexity is O(n^3), where n is the number of vertices in graph.
package main

import (
	"fmt"
	lst "container/list"
	"scanner"
	"os"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Graph type definitions

type edge struct {
	to   *node
	cap  int
	rcap *int
}

type node struct {
	edges []*edge

	exceed int
	height int
}

// The flow code

func (v *node) push(i int) {
	e := v.edges[i]
	add := min(v.exceed, e.cap)
	v.exceed -= add
	e.to.exceed += add
	e.cap -= add
	*e.rcap += add
}
func (v *node) relabel() {
	h := 1 << 30
	for _, e := range v.edges {
		if e.cap > 0 {
			h = min(h, e.to.height)
		}
	}
	v.height = 1 + h
}
func (v *node) discharge() {
	i := 0
	for v.exceed > 0 {
		if i == len(v.edges) {
			v.relabel()
			i = 0
		} else {
			e := v.edges[i]
			if e.cap > 0 && e.to.height < v.height {
				v.push(i)
			}
			i++
		}
	}
}

func maxflow(nodes []node, start, end int) int {
	list := lst.New()
	for i, _ := range nodes {
		if i != start && i != end {
			list.PushBack(&nodes[i])
			nodes[i].height = 1
		}
	}

	s := &nodes[start]
	s.exceed = 0
	for _, e := range s.edges {
		s.exceed += e.cap
	}
	s.height = len(nodes)
	s.discharge()

	for i := list.Front(); i != nil; i = i.Next() {
		v, _ := i.Value.(*node)
		h := v.height
		v.discharge()
		if v.height != h {
			list.Remove(i)
			list.PushFront(v)
			i = list.Front()
		}
	}
	return nodes[end].exceed
}

// Graph constructing functions

func append(edges []*edge, e *edge) []*edge {
	if cap(edges) == len(edges) {
		s := make([]*edge, len(edges), 1+2*len(edges))
		for i, x := range edges {
			s[i] = x
		}
		edges = s
	}
	edges = edges[0 : 1+len(edges)]
	edges[len(edges)-1] = e
	return edges
}
func join(from, to *node, cap int) {
	ea := &edge{to, cap, nil}
	eb := &edge{from, 0, &ea.cap}
	ea.rcap = &eb.cap
	from.edges = append(from.edges, ea)
	to.edges = append(to.edges, eb)
}

// Testing functions

var scan scanner.Scanner

func init() {
	scan.Init(os.Stdin)
	scan.Mode = scanner.ScanInts
}
func readInt() int {
	scan.Scan()
	r, _ := strconv.Atoi(scan.TokenText())
	return r
}

/* Reads flow network from stdin for testing in following format:
 * The input starts with four integers:
 * - N : number of vertices
 * - M : number of edges
 * - S : source index (0-based)
 * - T : sink index (0-based)
 * This is followed by m lines, each containing 3 integers: From,To and Cap.
 * Each line defines an edge from index From to index To (both 0-based) with
 * capasity Cap.
 */
func main() {
	n := readInt()
	m := readInt()
	start, end := readInt(), readInt()
	graph := make([]node, n)
	for i := 0; i < m; i++ {
		from, to, cap := readInt(), readInt(), readInt()
		join(&graph[from], &graph[to], cap)
	}
	fmt.Printf("Maxflow: %d\n", maxflow(graph, start, end))
}
