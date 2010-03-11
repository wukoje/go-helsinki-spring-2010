package main
import fmt "fmt"

// integer array for representating graph
type graph [][]int

// data structure queue with basic properties
type queue []int
func (q* queue) push(e int) {
    *q=(*q)[0:len(*q)+1]
    (*q)[len(*q)-1]=e
}
func (q* queue) pop() int {
    ret:=(*q)[0]
    *q=(*q)[1:]
    return ret
}
func create(s []int) *queue {
    tempslice:=queue(s[0:0])
    return &tempslice
}
func (q* queue) empty() bool {
    if(len(*q)==0) { return true }
	return false
}

// returns the length of shortest path from source to destination in graph g
func bfs(g graph, source int, dest int) int {
    q:=create(make([]int,1000))
    q.push(source)
    visited:=make([]bool,len(g))
    dist:=make([]int,len(g))
    dist[source]=0
    for !q.empty() {
        e := q.pop()
        visited[e]=true
        if(e==dest) {
            return dist[e]
        }
        for i:=0;i<len(g);i++ {
            if(g[e][i]==1 && !visited[i]) {
            	visited[i]=true
            	dist[i]=dist[e]+1
                q.push(i)
            }
        }
    }
    return -1
}


func main() {
	// integer array for representating graph
    g := [][]int{
       []int{0,1,0,0,0,0,0,0},
       []int{0,0,1,0,1,1,0,0},
       []int{0,0,0,1,0,0,1,0},
       []int{0,0,1,0,0,0,0,1},
       []int{1,0,0,0,0,1,0,0},
       []int{0,0,0,0,0,0,1,0},
       []int{0,0,0,0,0,1,0,0},
       []int{0,0,0,1,0,0,1,0}}
    source := 0
    destination := 6
    fmt.Printf("Shortest route from %d to %d is %d steps\n",source,destination,bfs(g,source,destination))
    

}
