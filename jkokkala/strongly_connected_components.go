// Kosaraju's algorithm for finding strongly connected components
package main
import fmt "fmt"

// Graph is an array of strings as it is easy to read from file
type graph []string

// Does DFS for graph and returns an array of nodes in DF order

// Returns array of strings representing the strongly connected components
type SCC_solver struct {
    visited []bool
    places []int
    g graph
}

func(s* SCC_solver) addplace(e int) {
    s.places=s.places[0:len(s.places)+1]
    s.places[len(s.places)-1]=e
}

func(s* SCC_solver) dfs(e int, direction bool) {
	s.visited[e]=direction
	for  i:=0 ; i<len(s.g) ; i++ {
        edge := false
        if(direction && s.is_edge(e,i)) { edge = true }
        if(!direction && s.is_edge(i,e)) { edge = true }
        if(s.visited[i]!=direction && edge) {
            s.dfs(i,direction)
        }
    }
    if(direction) { s.addplace(e) }
}

// Returns an array of strongly connected components (represented as integer arrays) in given graph g
func (s* SCC_solver) solve(g graph) [][]int {
	s.visited=make([]bool,len(g))
	s.places = make([]int,len(g))[0:0]
	s.g=g
    var ret [][]int = make([][]int,len(g))[0:0]

    // Travel all nodes normally and save the order
    s.dfs(0,true)

    k := len(s.places)-1
    used := make([]bool,len(s.places))
    // Travel the nodes in the previous order with reversed graph and find strongly connected components 
    for k>=0 {
        for k>=0 {
            if(!used[k]) { break }
            k--
        }
        if(k<0) { break }
        ret = ret[0:len(ret)+1]
        ret[len(ret)-1] = make([]int,len(s.g))[0:0]
        s.dfs(s.places[k],false)
        for i:=0 ; i < len(s.places) ; i++ {
            if(!used[i] && !s.visited[s.places[i]]) {
                ret[len(ret)-1] = ret[len(ret)-1][0:len(ret[len(ret)-1])+1]
                ret[len(ret)-1][len(ret[len(ret)-1])-1] = s.places[i]
                used[i] = true
            }
        }
    }
    return ret
}

// Function for checking if there exists an edge from node a to node b
func (s* SCC_solver) is_edge(a int,b int) bool {
	if(s.g[a][b]=='1') { return true }
	return false
}

func main() {
	var testgraph graph
	var solver SCC_solver

    // Graph in http://en.wikipedia.org/wiki/File:Scc.png
	testgraph = []string{
		"01000000",
		"00101100",
		"00010010",
		"00100001",
		"10000100",
		"00000010",
		"00000100",
		"00010010" }

    // The solution 
	SCC := solver.solve(testgraph)

    // Output the components
	for i:=0 ; i< len(SCC) ; i++ {
        for j:=0 ; j< len(SCC[i]) ; j++ {
            fmt.Printf("%d ",SCC[i][j])
        }
        fmt.Printf("\n")
    }
}
