
package main

import "fmt"

type queueNode struct {
	*graphNode
	next, prev *queueNode
}

type queue struct {
	front, back *queueNode
}

func (q *queue) enqueue(g *graphNode) {
	n := &queueNode { g, nil, q.back }
	if q.back != nil {
		q.back.next = n
	}
	q.back = n
	if q.front == nil {
		q.front = n
	}
}

func (q *queue) dequeue() *graphNode {
	if q.front == nil {
		return nil
	}

	n := q.front

	if q.front.next != nil {
		q.front.next.prev = nil
	}
	q.front = q.front.next

	if q.back == n {
		q.back = nil
	}

	return n.graphNode
}

// also functions as a node of an edge list
type edge struct {
	*graphNode
	next *edge
}

type graphNode struct{
	key string
	*edge
}

// searches for key using BFS starting from the nodes in q
// if the graph is cyclic, diverges if there's no node to be found
// this could be prevented by marking visited nodes
func bfs(q *queue, key string) *graphNode {

	n := q.dequeue()

	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	}

	for e := n.edge; e != nil; e = e.next {
		q.enqueue(e.graphNode)
	}

	return bfs(q, key)
}

func main() {
	// diamond-shaped DAG: g -> g0, g1 -> h
	h := &graphNode { "h", nil }
	g0 := &graphNode{ "g0", &edge { h, nil} }
	g1 := &graphNode{ "g1", &edge { h, nil} }
	g := &graphNode{ "g", &edge{ g0, &edge{ g1, nil } } }

	q := new(queue)
	q.enqueue(g)
	n := bfs(q, "g0")
	fmt.Println(n.key, "found")

	q = new(queue)
	q.enqueue(g)
	n = bfs(q, "h")
	fmt.Println(n.key, "found")

	q = new(queue)
	q.enqueue(g)
	n = bfs(q, "g1")
	fmt.Println(n.key, "found")

	q = new(queue)
	q.enqueue(g)
	n = bfs(q, "foo")
	if n == nil {
		fmt.Println("foo not found")
	}

}

