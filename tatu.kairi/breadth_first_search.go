/*
1 breadth_first_search.go

Define a graph data structure and implement a function that performs a breadth first search for a certain node.
*/

package main

import (
	"fmt"
)

type node struct {
	successors []*node
	value string
}

type graph struct {
	root *node
	size int
}

func search(g *graph, searchable string) *node {

	queue, discovered := make([]node, g.size), make([]node, g.size)
	queue[0] = *g.root
	queueSize := 1
	discIndex := 0

	for ; queueSize > 0; {
		currentNode := queue[queueSize - 1]
		queueSize--

		if currentNode.value == searchable {
			return &currentNode
		}

		for _,e := range currentNode.successors {
			if isInQueue(discovered, *e) {
				break
			}

			queue[queueSize] = *e
			queueSize++	
			discovered[discIndex] = *e
			discIndex++
		}
	}

	return nil
}

func isInQueue(s []node, e node) bool {
	for _,el := range s {
		if el.value == e.value {
			return true
		}
	}
	return false
}

func main(){

	/* Graph here is like this:

                Kassel - München
               /         Erfut
              /         /  
    Frankfurt - Würzburg 
              \         \
               \         Nürnberg - Stuttgart
                Mannheim - Karlsruhe - Augsburg
	*/

	node1, node2, node3, node4, node5, node6, node7, node8, node9, node10 := new(node), new(node), new(node), new(node), new(node), new(node), new(node), new(node), new(node), new(node)

	node1.successors = []*node{node2, node3, node4}
	node1.value = "Frankfurt"
	fmt.Printf("%v\n", node1)

	node2.successors = []*node{node5}
	node2.value = "Mannheim"

	node3.successors = []*node{node6, node7}
	node3.value = "Würzburg"

	node4.successors = []*node{node8}
	node4.value = "Kassel"

	node5.successors = []*node{node9}
	node5.value = "Karlsruhe"

	node6.successors = []*node{node10}
	node6.value = "Nürnberg"

	node7.value = "Erfut"

	node8.value = "München"

	node9.value = "Augsburg"

	node10.value = "Stuttgart"

	g := new(graph)
	g.root = node1
	g.size = 10

	fmt.Printf("Found '%s'\n", search(g, "Frankfurt").value)
	fmt.Printf("Found '%s'\n", search(g, "Stuttgart").value)
	fmt.Printf("Found '%s'\n", search(g, "München").value)


	n := search(g, "Helsinki")
	if n == nil {
		fmt.Printf("If node is not found, nil is returned\n")
	} else {
		fmt.Printf("Found '%s'\n", n.value)
	}



}

