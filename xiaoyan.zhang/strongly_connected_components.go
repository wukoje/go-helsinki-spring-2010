/*
Ex 8
Implement a function that given a directed graph
 finds the strongly connected components.

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
	"strconv"
	"container/list"
)

/*
directed graph represented by adjacency list
each index of the array represents a node
each list of the node contains all the connected nodes
length of the list represents how many edges it has
*/
type graph struct {
	nodes []*list.List
}

//constructor of graph
func newGraph(cap int) *graph {
	slice := make([]*list.List,cap)

	for i:=0; i<cap; i++ {
		list := list.New()
		slice[i] = list
	}
	return &graph{slice}
}

//add a new edge x->y to graph
func (g *graph) addEdge(x,y int) {
	if x < 0 || y < 0  || (x >= len(g.nodes)) || (y >= len(g.nodes)) {
		fmt.Println("Error: index out of bound")
		return
	}
	if x == y {
		fmt.Println("Error: not an edge!")
		return
	}

	for n := range g.nodes[x].Iter() {
		if n.(int) == y { return }
	}

	g.nodes[x].PushBack(y)
}

//reverse the directed graph, i.e. x->y => y->x
// and return the new graph
func (g *graph) reverse() *graph {
	if g==nil || len(g.nodes)==0 { return nil }
	
	g_rev := newGraph(len(g.nodes))
	for i:=0; i<len(g.nodes); i++ {
		list := g.nodes[i]
		for j := range list.Iter() {
			g_rev.addEdge(j.(int),i)
		}
	}

	return g_rev
}

//return string to show graph
func (g *graph) String() string{
	s := "Graph: \n"
	for i,list := range g.nodes {
		s += "Node " + strconv.Itoa(i) + ":\t"
		for n := range list.Iter() {
			s += strconv.Itoa(n.(int)) + "\t"
		}
		s += "\n"
	}
	return s
}

/*
data structure used in finding strongly connected components
index of nodes_count represents the node
value of nodes_count represents the count number used in traversal
*/
type checklist struct {
	nodes_count []int
	count int
}

//depth first search
func (g *graph) dfs(chk *checklist, n int) {
	list := g.nodes[n]
	chk.nodes_count[n] = -1

	for i := range list.Iter() {
		if chk.nodes_count[i.(int)] == 0 {
			g.dfs(chk, i.(int))
		}
	}
	chk.count++
	chk.nodes_count[n] = chk.count	
}

//find the max value in the checklist
func (g *graph) find_max_value(chk *checklist) int {
	current := 0
	idx := -1
	for i,c := range chk.nodes_count {
		if c > current {
			idx = i
			current = c
		}
	}
	if idx >= 0 { chk.nodes_count[idx] = -1 }
	return idx
}

/*
find strongly connected components
with the help of dfs both graph and reverse_graph
*/

func (g *graph) find_str_conn_component() {
	if len(g.nodes) <= 0 { return }

	chk := &checklist{make([]int,len(g.nodes)),0}
	for i,_ := range g.nodes {
		if chk.nodes_count[i] == 0 {
			g.dfs(chk,i)
		}
	}

	g_rev := g.reverse()
	chk_rev := &checklist{make([]int,len(g.nodes)),0}
	component := list.New()
	num_checked := 0

	for chk_rev.count <= len(g.nodes) {
		idx := g_rev.find_max_value(chk)
		
		if idx < 0 { break }

		if chk_rev.nodes_count[idx] == 0 {
			g_rev.dfs(chk_rev,idx)
		} else { continue }

		list := list.New()
		for i,c := range chk_rev.nodes_count {
			if c > num_checked && c <= chk_rev.count {
				list.PushBack(i)
			}
		}
		num_checked = chk_rev.count
		component.PushBack(list)
	}

	//Print out strongly connected components
	i := 0
	for l := range component.Iter() {
		i++
		fmt.Printf("strongly connected component %d:\n", i)
		for sub := range l.(*list.List).Iter() {
			fmt.Printf("%d\t", sub.(int))
		}
		fmt.Println("\n")
	}

}


func main() {
	fmt.Println("#Test case 1:")
	g := newGraph(4)
	g.addEdge(0,1)
	g.addEdge(0,2)
	g.addEdge(2,0)
	g.addEdge(2,3)
	g.addEdge(3,0)
	fmt.Println(g)
	g.find_str_conn_component()

	fmt.Println("#Test case 2:")
	g = newGraph(8)
	g.addEdge(0,1)
	g.addEdge(1,2)
	g.addEdge(1,4)
	g.addEdge(1,5)
	g.addEdge(2,3)
	g.addEdge(2,6)
	g.addEdge(3,2)
	g.addEdge(3,7)
	g.addEdge(4,0)
	g.addEdge(4,5)
	g.addEdge(5,6)
	g.addEdge(6,5)
	g.addEdge(7,6)
	g.addEdge(7,3)
	fmt.Println(g)
	g.find_str_conn_component()
}
