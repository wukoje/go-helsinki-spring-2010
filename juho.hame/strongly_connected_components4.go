/**
2 strongly_connected_components.go
Implement a function that given a directed graph finds the strongly connected components. You get to define your own graph representation.
*/

package main

import "fmt"


type graph struct {
	nodes []*node	// pointer(s) to (every) node in graph ("source node for all nodes in graph?")
}

type node struct {
	cons []*node	// cons, connections to neighbors
}


func add(ns []*node, n *node) {
	for i,m := range ns {
		if m == nil {
			ns[i] = n
			break
		}
	}
}

 // returns true if n NOT IN ns
func (n *node) notIn(ns []*node) bool {
	for _, v := range ns {
		if v == n {
			return false	// found!
		}
	}
	return true		// list ended, not found
}

// returns true if path a, b exists, false else (v for visited)
func isPath(a *node, b *node, v []*node) bool {
	add(v, a)
	is := false

	if b.notIn(a.cons) {
		for _, n := range a.cons {
			if n.notIn(v) {
				is = isPath(n, b, v)
			}
		}
	} else {
		is = true	// !
	}

	return is
}

// returns index of n in ns
func (n *node) indexIn (ns []*node) int {
	for i, v := range ns {
		if v == n {
			return i
		}
	}
	return -1	// not found
}

/**
 *	For each node (v) in graph (g), see what type of connection (one-way
 *	path, two-way path, no path) there is to each other node (w) in
 *	graph (g). (Nodes already proven to be in a strongly connected
 *	component (listed in ss) are skipped.)
 */
func (g *graph) findStrong () {
	ss := make([]*node, len(g.nodes))			// ss for strongs (slice of nodes in strong components)
	for i,v := range g.nodes {
		if v.notIn(ss) {						// skip values already in a strong component
			strong := true						// nodes with no connections are strong!
			cis := make([]*node, len(g.nodes))	// indexes of nodes in component, len(g.nodes) just in case
			cis[0] = v
			ix := 1								// index for keeping track of next available... index in cis
			for j,w := range g.nodes {			// searches all links that work both ways
				if i != j && w.notIn(ss) {		// (again) skip values already in a strong component
					switch {
						// path v > w exists, but path v < w does not: v is not in a strong component
						case isPath(v, w, make([]*node, len(g.nodes))) && !isPath(w, v, make([]*node, len(g.nodes))):
							fmt.Printf("%v  > %v (fail)\n", i, j)
							strong = false
							break
						// path v > w does not exist, but path v < w exists! v is not in a strong component
						case !isPath(v, w, make([]*node, len(g.nodes))) && isPath(w, v, make([]*node, len(g.nodes))):
							fmt.Printf("%v <  %v (fail)\n", i, j)
							strong = false
							break
						// neither v > w or v < w exists, which is ok
						case !isPath(v, w, make([]*node, len(g.nodes))) && !isPath(w, v, make([]*node, len(g.nodes))):
							fmt.Printf("%v ?? %v\n", i, j)
						// because this case is none of the above, path v > w must exist as well as v < w. So we add w to alleged strongly connected component
						default:
							fmt.Printf("%v <> %v\n", i, j)
							cis[ix] = w
							ix++
					}
				}
			}

			// when the algorithm reaches this point it has either a list of nodes in a strong component (cis) OR it has found that v is not in a strong component (!strong). In the case of the former, we print the indexes and add all nodes in component to a list (ss) containing all nodes in any strongly connected component. In the case of the latter, just move on to next node in g (used variables are overwritten)
			if strong {
				fmt.Printf("Indexes of nodes in strong component:")
				ix--	// ix is now index of last added node
				for ix > -1 {	// ix == -1: list ended
					fmt.Printf(" %v", cis[ix].indexIn(g.nodes))
					add(ss, cis[ix])
					ix--
				}
				fmt.Printf("\n")
			}
		}
	}
}


func main() {

	n0 := new(node)
	n1 := new(node)
	n2 := new(node)

	n3 := new(node)
	n4 := new(node)
	n5 := new(node)
	n6 := new(node)

	n7 := new(node)
	n8 := new(node)
	n9 := new(node)

	n0.cons = []*node{n1}
	n1.cons = []*node{n0, n2}
	n2.cons = []*node{n1}
	// n0 <> n1 <> n2

	n3.cons = []*node{n4}
	n4.cons = []*node{n5}
	n5.cons = []*node{n6}
	n6.cons = []*node{n3}
	// n3 > n4 > n5 > n6 > n3 > ...

	n7.cons = []*node{n8}
	n8.cons = []*node{n7}
	n9.cons = []*node{n8}
	// n7 <> n8 < n9

	g := new(graph)
	g.nodes = []*node{n0, n1, n2, n3, n4, n5, n6, n7, n8, n9}

	g.findStrong()

}

