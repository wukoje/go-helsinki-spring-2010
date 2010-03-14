/*
Ex 7
Define a graph data structure
 and implement a function that performs a breadth first search
 for a certain node.

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
	"container/list"
	"strconv"
)

/*
graph represented by adjacency list
each index of the array represents a node
each vector of the node contains all the connected nodes
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

//add a new edge to graph
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
	g.nodes[y].PushBack(x)
}

/*
data structure used in search
the array contains all the nodes involved in breadth_first introval in FIFO order
cur_idx is the current index of element that is checked
tail is the next space available
*/
type checklist struct {
	checked_nodes []int
	cur_idx int
	tail int
}

// check if e is already in checklist
func (ch *checklist) contain(e int) bool {
	for i:=0; i<ch.tail; i++ {
		if ch.checked_nodes[i] == e {
			return true
		}
	}
	return false
}

// if checklist is full, return false
// else add new node to checklist
func (ch *checklist) add(e int) bool {
	if ch.tail == len(ch.checked_nodes) { return false }
	if !ch.contain(e) {
		ch.checked_nodes[ch.tail] = e
		ch.tail++
	}
	return true
}

func (ch *checklist) String() string {
	s := ""
	if ch != nil {
		for i:=0; i<ch.tail; i++ {
			s += strconv.Itoa(ch.checked_nodes[i]) + "\t"
		}
	}
	return s
}

//search the given element and return index
//if element does not exist, return -1
func (g *graph) search(e int) (int, *checklist) {
	if len(g.nodes) > 0 {
		chk := &checklist{make([]int,len(g.nodes)), -1, 0}

		for i:=0; i<len(g.nodes); i++ {	//find first unchecked nodes
			if i == e { return i,chk }
			if !chk.add(i) { return -1,chk }

			//check nodes that connect to current node
			for chk.cur_idx < (chk.tail-1) {
				chk.cur_idx++
				cur_list := g.nodes[chk.checked_nodes[chk.cur_idx]]			
				if cur_list != nil {	
					for n := range cur_list.Iter() {
						if n.(int) == e { return n.(int),chk }
						if !chk.add(n.(int)) { return -1,chk }
					}
				}
			}
		}
		return -1,chk			
	}
	return -1,nil
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

func main(){
	fmt.Println("#Test case 1: all nodes are connected")
	graph := newGraph(6)
	graph.addEdge(0,1)
	graph.addEdge(0,3)
	graph.addEdge(1,3)
	graph.addEdge(1,2)
	graph.addEdge(1,4)
	graph.addEdge(3,4)
	graph.addEdge(4,5)
	fmt.Println(graph)
	i,chk := graph.search(2)
	fmt.Printf("search 2\n searching path:\t%v\n result:%d\n",chk,i)
	i,chk = graph.search(6)
	fmt.Printf("search 6\n searching path:\t%v\n result:%d\n",chk,i)

	fmt.Println("\n#Test case 2: there exist single nodes")
	graph = newGraph(8)
	graph.addEdge(0,2)
	graph.addEdge(0,5)
	graph.addEdge(1,5)
	graph.addEdge(5,3)
	graph.addEdge(5,6)
	graph.addEdge(5,2)
	graph.addEdge(2,6)
	fmt.Println(graph)
	i,chk = graph.search(7)
	fmt.Printf("search 7\n searching path:\t%v\n result:%d\n",chk,i)
	i,chk = graph.search(8)
	fmt.Printf("search 8\n searching path:\t%v\n result:%d\n",chk,i)
}
