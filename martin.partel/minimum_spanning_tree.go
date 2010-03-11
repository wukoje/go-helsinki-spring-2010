package main

import "math"
import "container/heap"
import "os"
import "exec"
import "fmt"
import "strconv"
import "scanner"

const MAX_VERTS = 128
const MAX_EDGES = 128

/*
 * Prim's algo as presented in Cormen et al.
 */

type Graph struct {
	verts       []*Vertex
	vertsByName map[string]int
}

func (g *Graph) getOrAddVertex(name string) int {
	index, found := g.vertsByName[name]
	if found {
		return index
	}
	return g.addVertex(name)
}

func (g *Graph) addVertex(name string) int {
	index := len(g.verts)
	g.verts = g.verts[0 : index+1]

	v := new(Vertex)
	v.name = name
	v.adj = make([]int, 0, MAX_EDGES)
	v.weights = make([]int, 0, MAX_EDGES)

	g.vertsByName[name] = index
	g.verts[index] = v

	return index
}

type Vertex struct {
	name    string
	adj     []int
	weights []int
}

type MstNode struct {
	index     int
	heapIndex int
	added     bool
	key       int // Weight of lightest edge that could connect this vertex to the tree
	parent    int
}

type mstHeap struct {
	nodes []*MstNode
}

func (h *mstHeap) Len() int { return len(h.nodes) }

func (h *mstHeap) Less(i int, j int) bool { return h.nodes[i].key < h.nodes[j].key }

func (h *mstHeap) Push(x interface{}) {
	i := len(h.nodes) + 1
	h.nodes = h.nodes[0:i]
	h.nodes[i] = x.(*MstNode)
}

func (h *mstHeap) Pop() interface{} {
	i := len(h.nodes) - 1
	x := h.nodes[i]
	h.nodes = h.nodes[0:i]
	return x
}

func (h *mstHeap) Swap(i int, j int) {
	t := h.nodes[i]
	h.nodes[i] = h.nodes[j]
	h.nodes[j] = t

	h.nodes[i].heapIndex = i
	h.nodes[j].heapIndex = j
}

func (h *mstHeap) decreaseKey(index int, newKey int) {
	h.nodes[index].key = newKey
	for parentIndex := (index - 1) / 2; index > 0 && h.Less(parentIndex, index); index = parentIndex {
		h.Swap(index, parentIndex)
	}
}

func Prim(g *Graph, rootIndex int) []*MstNode {
	nodes := make([]*MstNode, len(g.verts))
	q := &mstHeap{make([]*MstNode, len(g.verts))}

	for i := 0; i < len(g.verts); i++ {
		node := &MstNode{i, i, false, math.MaxInt32, -1}
		nodes[i] = node
		q.nodes[i] = node
	}

	q.nodes[rootIndex].key = 0
	heap.Init(q)

	for q.Len() > 0 {
		n1 := heap.Pop(q).(*MstNode)
		n1.added = true
		v1Index := n1.index
		v1 := g.verts[v1Index]

		for edgeIndex, v2Index := range v1.adj {
			n2 := nodes[v2Index]
			weight := v1.weights[edgeIndex]
			if !n2.added && weight < n2.key {
				n2.parent = n1.index
				q.decreaseKey(n2.heapIndex, weight)
			}
		}
	}

	return nodes
}

func main() {
	drawResult := false

	verts := make([]*Vertex, 0, MAX_VERTS)
	vertsByName := make(map[string]int)
	graph := &Graph{verts, vertsByName}

	var s scanner.Scanner
	s.Init(os.Stdin)

	readEdges(&s, graph)

	if len(graph.verts) == 0 {
		bail(os.NewError("Empty graph given"))
	}

	nodes := Prim(graph, 0)

	for i2, n2 := range nodes {
		i1 := n2.parent
		if i1 > -1 {
			v1 := graph.verts[i1]
			v2 := graph.verts[i2]

			var weight int
			for e, i3 := range v1.adj {
				if i3 == i2 {
					weight = v1.weights[e]
				}
			}

			fmt.Printf("%s -(%d)- %s\n", v1.name, weight, v2.name)
		}
	}

	if drawResult {
		err := drawGraphviz(graph, nodes, "mst.png")
		if err != nil {
			fmt.Printf("trying to execute dot: %s\n", err.String())
		}
	}
}

// Line format: "- from weight to"
// Example: "- a 3 b"
//
func readEdges(s *scanner.Scanner, graph *Graph) {
	tok := s.Scan()
	for tok != scanner.EOF {
		n1 := expectIdent(s, "from")
		w := expectInt(s, "weight")
		n2 := expectIdent(s, "to")

		i1 := graph.getOrAddVertex(n1)
		i2 := graph.getOrAddVertex(n2)
		v1 := graph.verts[i1]
		v2 := graph.verts[i2]

		e1 := len(v1.adj)
		e2 := len(v2.adj)
		v1.adj = v1.adj[0 : e1+1]
		v1.weights = v1.weights[0 : e1+1]
		v2.adj = v2.adj[0 : e2+1]
		v2.weights = v2.weights[0 : e2+1]

		v1.adj[e1] = i2
		v1.weights[e1] = w
		v2.adj[e2] = i1
		v2.weights[e2] = w

		tok = s.Scan()
	}
}

func expectIdent(s *scanner.Scanner, desc string) string {
	tok := s.Scan()
	if tok != scanner.Ident {
		bail(os.NewError(desc + " expected"))
	}
	return s.TokenText()
}

func expectInt(s *scanner.Scanner, desc string) int {
	tok := s.Scan()
	if tok != scanner.Int {
		bail(os.NewError(desc + " expected"))
	}
	n, err := strconv.Atoi(s.TokenText())
	if err != nil {
		bail(err)
	}
	return n
}

func drawGraphviz(graph *Graph, nodes []*MstNode, filename string) os.Error {
	dotPath, err := exec.LookPath("dot")
	if err != nil {
		return err
	}

	read, write, err := os.Pipe()
	if err != nil {
		return err
	}
	defer (*os.File).Close(read)
	defer (*os.File).Close(write)

	argv := []string{dotPath, "-Tpng", "-o" + filename}
	fds := []*os.File{read, os.Stdout, os.Stderr}
	pid, err := os.ForkExec(dotPath, argv, os.Environ(), "", fds)

	if err != nil {
		return err
	}

	write.WriteString(graphvizInput(graph, nodes))
	(*os.File).Close(write)

	os.Wait(pid, 0)

	return nil
}

func graphvizInput(graph *Graph, nodes []*MstNode) string {
	result := "graph {\n"
	result += "node [shape=circle]\n"

	edgesSeen := map[string]bool{}

	for v1Index, n1 := range nodes {
		v1 := graph.verts[v1Index]

		for edgeIndex, v2Index := range v1.adj {
			v2 := graph.verts[v2Index]
			n2 := nodes[v2Index]
			weight := v1.weights[edgeIndex]

			edgeCode := fmt.Sprintf("%d %d %d\n", v1Index, weight, v2Index)
			oppositeEdgeCode := fmt.Sprintf("%d %d %d\n", v2Index, weight, v1Index)
			_, oppositeDrawn := edgesSeen[oppositeEdgeCode]

			if !oppositeDrawn {
				inMst := n1.parent == v2Index || n2.parent == v1Index
				penWidth := "1.0"
				if inMst {
					penWidth = "3.0"
				}

				result += fmt.Sprintf("%s -- %s [label=%d,penwidth=%s]\n", v1.name, v2.name, weight, penWidth)
				edgesSeen[edgeCode] = true
			}
		}
	}

	result += "}\n"

	return result
}

func bail(err os.Error) {
	fmt.Println(err.String())
	os.Exit(1)
}
