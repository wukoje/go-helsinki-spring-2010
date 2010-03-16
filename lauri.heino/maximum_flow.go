// maximum_flow.go
//
// The maximum flow using the Edmunds-Karp algorithm.

package main

import (
	"fmt"
	"container/list"
	"math"
	)

type Node struct {
	id string
}

type Graph struct {
	nodes *list.List
	edges map[*Node] (map[*Node] float32)
	// ^^ (src->trgt->maximum capacity)
}

// initGraph builds the graph used as an example in the Edmonds-Karp 
// algorithm article on Wikipedia.
func initGraph() *Graph {
	graph := new(Graph)
	graph.nodes = list.New()
	graph.edges = map[*Node] (map[*Node] float32){}
	var nodes [7]*Node

	for i:= 0; i < len(nodes); i++ {
		nodes[i] = new(Node)
		nodes[i].id = string(i + 'A')
		graph.nodes.PushBack(nodes[i])
		graph.edges[nodes[i]] = map[*Node] float32{}
	}

	graph.edges[nodes[0]][nodes[1]] = 3.0
	graph.edges[nodes[0]][nodes[3]] = 3.0
	graph.edges[nodes[1]][nodes[2]] = 4.0
	graph.edges[nodes[2]][nodes[0]] = 3.0
	graph.edges[nodes[2]][nodes[3]] = 1.0
	graph.edges[nodes[2]][nodes[4]] = 2.0
	graph.edges[nodes[3]][nodes[4]] = 2.0
	graph.edges[nodes[3]][nodes[5]] = 6.0
	graph.edges[nodes[4]][nodes[1]] = 1.0
	graph.edges[nodes[4]][nodes[6]] = 1.0
	graph.edges[nodes[5]][nodes[6]] = 9.0
	
	return graph
}

// min returns the minimum of a and b.
func min(a, b float32) float32 {
	if a < b {
		return a
	} 
	return b
}

// maximumFlow computes the maximum flow from src to sink in a graph.
// The algorithm used is the Edmonds-Karp algorithm. 
func maximumFlow(graph *Graph, src *Node, sink *Node) float32 {
	var flow float32
	resCapacity := map[*Node] (map[*Node] float32){}
	// ^^ "src->target->residual capacity"

	//build and fill the resCapacity map
	for elem1 := (graph.nodes).Front(); elem1 != nil; elem1 = elem1.Next() {
		node1 := (elem1.Value).(*Node)
		resCapacity[node1] = map[*Node] float32{}
		for elem2 := (graph.nodes).Front(); elem2 != nil; elem2 = elem2.Next() {
			node2 := (elem2.Value).(*Node)
			if _,ok := graph.edges[node1][node2]; ok {
				resCapacity[node1][node2] = graph.edges[node1][node2]
			} else {
				resCapacity[node1][node2] = 0.0
			}
		}
	}

	flow = 0.0

	for {
		m, parent := BFS(graph, src, sink, resCapacity)

		if m == 0.0 {
			break
		}

		flow += m
		v := sink

		for v != src {
			u:= parent[v]
			resCapacity[u][v] -= m
			resCapacity[v][u] += m
			v = u
		}
	}

	return flow
}

// printCapacityMap prints a capacity map.
func printCapacityMap(capacity map[*Node] (map[*Node] float32)) {
	for node1, _ := range capacity {
		for node2, _ := range capacity[node1] {
			fmt.Printf("*** From %v to %v: %v\n", node1.id, node2.id, capacity[node1][node2])
		}
	}
}

// BFS performs a breadth-first search to find a path from src to sink.
// Returns the path in (flow, parentlist) format.
func BFS(graph *Graph, src *Node, sink *Node, resCapacity map[*Node] (map[*Node] float32)) (float32, map[*Node]*Node) {

	var color = map[*Node] int{}
	var minCapacity=  map[*Node] float32{}
	var parent =  map[*Node] *Node{}
	queue := list.New()

	for elem := (graph.nodes).Front(); elem != nil; elem = elem.Next() {
		node := (elem.Value).(*Node)
		color[node] = 0
		minCapacity[node] = math.MaxFloat32
	}

	color[src] = 1
	queue.PushBack(src)
	
	for queue.Len() > 0 {

		first := queue.Front()
		queue.Remove(first)
		v := (first.Value).(*Node)

		for elem := (graph.nodes).Front(); elem != nil; elem = elem.Next() {

			u := (elem.Value).(*Node)
			
			if color[u] == 0 && resCapacity[v][u] > 0.0 {

				minCapacity[u] = min(minCapacity[v], resCapacity[v][u])
				color[u] = 1
				parent[u] = v

				if u == sink {
					return minCapacity[sink], parent
				}

				queue.PushBack(u)
			}
		}
	}

	return 0.0, parent
} 

func main() {
	fmt.Printf("*** Maximum Flow ***\n")
	g := initGraph()
	f := maximumFlow(g, (g.nodes.Front().Value).(*Node),  (g.nodes.Back().Value).(*Node))
	fmt.Printf("Flow: %v\n", f)
}
