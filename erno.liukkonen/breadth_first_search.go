package main

import "fmt"
import "container/list"

func main() {
	graph := new(graph)
	test1 := &node { 1, list.New() }
	test2 := &node { 2, list.New() }
	test3 := &node { 3, list.New() }
	test4 := &node { 4, list.New() }

	graph.base = test1

	test1.e.PushBack(test2)
	test2.e.PushBack(test3)
	test3.e.PushBack(test4)
	test4.e.PushBack(test1)

	fmt.Println(graph.search(2))
}

type node struct {
	num int
	e *list.List
}

type graph struct {
	base *node
}

func (g *graph) search(num int) *node {
	que := list.New()
	
	add_q := func(n *node) {
		que.PushBack(n)
	}

	remove_q := func() *node {
		temp := que.Front()
		if temp == nil {
			return nil
		}
		que.Remove(temp)
		return temp.Value.(*node)
	}
	add_q(g.base)

	temp2 := map[int] bool {}
	for que.Front() != nil {
		nod := remove_q()
		temp2[nod.num] = true
		
		if nod.num == num && nod!= nil {
			return nod
		}

		for edge := range nod.e.Iter() {
			e := edge.(*node)
			if _, b := temp2[e.num]; !b {
				add_q(edge.(*node))
			}
		}
	}
	return nil
}

