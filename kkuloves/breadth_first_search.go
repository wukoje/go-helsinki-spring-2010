/*
 * breadth_first_search.go: A breadth-first search that finds the
 * distance (as number of edges) between a pair of nodes. The
 * usage is:
 *
 * 	breadth_first_search -s=SourceNode -t=TargetNode -g=GraphFile
 *
 * The option -d may be used to see the progress of the search.
 *
 * The implementation is somewhat eccentric and uses an adjacency
 * matrix of individual bits to store the edges. The advantage of
 * this approach is that a few simple bit manipulation operations
 * can be used to very quickly establish the set of nodes to visit
 * on each iteration of the search. I plan to extend this program
 * into a Monte Carlo estimator for network reliability; we have
 * an ongoing contest at work to see who produces the fastest
 * two-terminal network reliability estimator for our biological
 * graphs. (This plan also determined the input file format.)
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type bits uint64

const numBits = 64

type AdjacencySet []bits

type Graph struct {
	nodes     []string
	nodeIds   map[string]int
	neighbors []AdjacencySet
}

var (
	graphFile = flag.String("g", "breadth_first_search.graph",
		"Filename of the graph file to load")
	sourceNode = flag.String("s", "",
		"Source node to search from (e.g. root of graph)")
	targetNode = flag.String("t", "",
		"Target node to search to")
	debugPrint = flag.Bool("d", false,
		"Print debug output while searching")
)

// DistanceBetween returns the shortest distance (in number of links)
// between source and target nodes in g. Returns -1 if no path
// exists between source and target. The node ids must be valid.
func (g *Graph) DistanceBetween(source, target int) int {
	var (
		n, depth  int
		exhausted bool
	)
	if source == target {
		return 0
	}
	n = cap(g.nodes) / numBits
	visited := make(AdjacencySet, n, n)
	neighbors := make(AdjacencySet, n, n)
	queue := make(AdjacencySet, n, n)
	visited.Set(source)
	copy(queue, g.neighbors[source])

	for !exhausted {
		depth++
		if *debugPrint {
			fmt.Fprintf(os.Stderr, "At depth %d:\n", depth)
		}

		// Visit the queued nodes
		for i, nodes := range queue {
			visited[i] |= nodes
			for n = i * numBits; nodes != 0; nodes, n = (nodes >> 1), (n + 1) {
				if (nodes & 1) != 0 {
					if *debugPrint {
						fmt.Fprintf(os.Stderr, "\t%s\n", g.NodeName(n))
					}
					if n == target {
						return depth
					}
					for j, adjacent := range g.neighbors[n] {
						neighbors[j] |= adjacent
					}
				}
			}
		}

		// Queue the previously unvisited neighbors
		exhausted = true
		for i, nodes := range neighbors {
			b := nodes &^ visited[i]
			if b != 0 {
				exhausted = false
			}
			queue[i] = b
			neighbors[i] = 0
		}
	}
	return -1
}

// NodeId returns the numeric id of node or -1 if node not in g.
func (g *Graph) NodeId(node string) int {
	id, exists := g.nodeIds[node]
	if !exists {
		return -1
	}
	return id
}

// NodeName returns the name corresponding to nodeId, which must be valid.
func (g *Graph) NodeName(nodeId int) string { return g.nodes[nodeId] }

// EnsureHasEdge ensures that an edge between node1 and node2 is in g.
func (g *Graph) EnsureHasEdge(node1, node2 string) {
	a := g.EnsureHasNode(node1)
	b := g.EnsureHasNode(node2)
	g.neighbors[a].Set(b)
	g.neighbors[b].Set(a)
}

// EnsureHasNode ensures that node is present in g and returns its id.
func (g *Graph) EnsureHasNode(node string) int {
	id, exists := g.nodeIds[node]
	if !exists {
		if len(g.nodes) == cap(g.nodes) || len(g.nodes) != len(g.neighbors) ||
			len(g.nodes) != len(g.nodeIds) {
			fmt.Fprintln(os.Stderr, "Error: Graph capacity exceeded!")
			os.Exit(1)
		}
		id = len(g.nodes)
		g.nodes = g.nodes[0 : id+1]
		g.nodes[id] = node
		g.nodeIds[node] = id
		g.neighbors = g.neighbors[0 : id+1]
		g.neighbors[id] = make(AdjacencySet, 0, cap(g.nodes)/numBits)
	}
	return id
}

// NumNodes returns the number of nodes in g.
func (g *Graph) NumNodes() int { return len(g.nodes) }

// Set sets the bit for node in a to 1.
func (a *AdjacencySet) Set(node int) {
	word, bit := (node / numBits), bits(node%numBits)
	if word >= len(*a) {
		if word >= cap(*a) {
			fmt.Fprintln(os.Stderr, "Error: Adjacency set capacity exceeded!")
			os.Exit(1)
		}
		*a = (*a)[0 : word+1]
	}
	(*a)[word] |= (1 << bit)
}

// Unset sets the bit for node in a to 0.
func (a AdjacencySet) Unset(node int) {
	word, bit := (node / numBits), bits(node%numBits)
	if word < len(a) {
		a[word] &^= (1 << bit)
	}
}

// Returns true iff the bit for node in a is 1.
func (a AdjacencySet) IsSet(node int) bool {
	word, bit := (node / numBits), bits(node%numBits)
	if word < len(a) {
		return (a[word] & (1 << bit)) != 0
	}
	return false
}

func (a AdjacencySet) String() string {
	b := make([]byte, cap(a)*numBits)
	for i := cap(a)*numBits - 1; i >= 0; i-- {
		if a.IsSet(i) {
			b[i] = '1'
		} else {
			b[i] = '0'
		}
	}
	return string(b)
}

// NewGraph allocates a new graph with the given capacity for nodes.
func NewGraph(capacity uint) *Graph {
	g := new(Graph)
	capacity += numBits - (capacity % numBits) // Make divisible by numBits
	g.nodes = make([]string, 0, capacity)
	g.nodeIds = make(map[string]int, capacity)
	g.neighbors = make([]AdjacencySet, 0, capacity)
	return g
}

// ReadGraph reads a graph from input. The first line of input must
// have a sufficient node capacity as a decimal integer. The two first
// space-separated fields on each following line define a pair of
// nodes and an edge connecting them.
func ReadGraph(input io.Reader) (g *Graph) {
	var (
		err  os.Error
		line string
		rd   = bufio.NewReader(input)
	)
	if line, err = rd.ReadString('\n'); err == nil {
		var capacity uint
		capacity, err = strconv.Atoui(line[0 : len(line)-1])
		g = NewGraph(capacity)
	}
	for err == nil {
		if line, err = rd.ReadString('\n'); err == nil {
			fields := strings.Fields(line[0 : len(line)-1])
			if len(fields) >= 2 {
				g.EnsureHasEdge(fields[0], fields[1])
			}
		}
	}
	if err != nil && err != os.EOF {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	return
}

func main() {
	var (
		g        *Graph
		sid, tid int
	)

	flag.Parse()

	if file, err := os.Open(*graphFile, os.O_RDONLY, 0); err == nil {
		fmt.Fprintf(os.Stderr, "Reading graph from \"%s\"...\n", *graphFile)
		g = ReadGraph(file)
		file.Close()
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s\n", *graphFile, err)
	}
	if g == nil || g.NumNodes() < 2 {
		fmt.Fprintln(os.Stderr, "No graph, aborting.\n"+
			"Input format for the undirected graph:\n"+
			"<total number of nodes>\n"+
			"<node1> <node2>[ other fields ignored]\n"+
			"... (one line as above per edge)\n")
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "%d nodes read.\n", g.NumNodes())

	if *sourceNode == "" {
		// Assume the first node as root arbitrarily if no source specified
		sid = 0
		*sourceNode = g.NodeName(sid)
		fmt.Fprintf(os.Stderr, "Notice: No source node specified, using: %s\n",
			*sourceNode)
	} else {
		sid = g.NodeId(*sourceNode)
		if sid < 0 {
			notFound(*sourceNode)
		}
	}

	if *targetNode == "" {
		fmt.Fprintln(os.Stderr, "Error: Target node not specified!\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	tid = g.NodeId(*targetNode)
	if tid < 0 {
		notFound(*targetNode)
	}

	if d := g.DistanceBetween(sid, tid); d >= 0 {
		fmt.Printf("Distance from \"%s\" to \"%s\": %d\n",
			*sourceNode, *targetNode, d)
	} else {
		fmt.Printf("No path between \"%s\" and \"%s\".\n",
			*sourceNode, *targetNode)
	}
}

func notFound(node string) {
	fmt.Fprintf(os.Stderr, "Node \"%s\" does not exist in the graph.\n", node)
	os.Exit(1)
}
