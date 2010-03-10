package main

import (
	heap "container/heap"
	vector "container/vector"
	"scanner"
	"fmt"
	"os"
	"strconv"
)

type edge struct {
	to     int
	weight int
	num    int
}
type node []edge

func (e edge) Less(f interface{}) bool {
	return e.weight < f.(edge).weight
}

// Calculates the minimum spanning tree using Prim's algorithm
// Returns cost of the tree and indices of its edges
func minimumSpanningTree(graph []node) (int, []int) {
	if len(graph) == 0 {
		return 0, nil
	}
	pq := new(vector.Vector)
	heap.Init(pq)

	used := make([]bool, len(graph))
	used[0] = true
	for _, e := range graph[0] {
		heap.Push(pq, e)
	}
	total := 0
	res := make([]int, 0, len(graph)-1)
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(edge)
		if used[cur.to] {
			continue
		}
		used[cur.to] = true
		total += cur.weight
		res = res[0 : 1+len(res)]
		res[len(res)-1] = cur.num

		for _, e := range graph[cur.to] {
			heap.Push(pq, e)
		}
	}
	return total, res
}

// Test code

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
func main() {
	n, m := readInt(), readInt()
	arrs := make([]vector.Vector, n)
	for i := 0; i < m; i++ {
		a, b, w := readInt(), readInt(), readInt()
		arrs[a].Push(edge{b, w, i})
		arrs[b].Push(edge{a, w, i})
	}
	graph := make([]node, n)
	for i, v := range arrs {
		graph[i] = make([]edge, v.Len())
		for j, e := range v {
			graph[i][j] = e.(edge)
		}
	}
	r, edges := minimumSpanningTree(graph)
	fmt.Println("MST cost:", r)
	fmt.Println("Chosen edges:", edges)
}
