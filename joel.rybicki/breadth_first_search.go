package main

import (
  "fmt"
  "flag"
  "os"
  "./graph"
  "container/vector"
  "strconv"
)

// Runs a breadth-first search from node 'root'.
// Returns true if the node 'sought' is encountered during the search
// and false otherwise.
func bfs(g *graph.Graph, root int, sought int) bool {
  var queue vector.IntVector
  queue.Push(root)
  
  for len(queue) != 0 {
    // dequeue
    n := queue.At(0)
    queue.Delete(0)
    fmt.Printf("Arrived at node %d\n", n)
    if n == sought {
      return true
    } 
    if g.Nodes[n] == 0 {
      g.Nodes[n] = 1 // discovered, ignore later
      for i, val := range g.Adj[n] {
        if val != 0 && g.Nodes[i] == 0 { // has an edge to i and it hasn't been visited
          queue.Push(i)
        }
      }
    }
  }
  return false
}

func main() {
  flag.Parse()
  if flag.NArg() < 3 {
    fmt.Println("Give the path to a graph file as the first parameter, root node as the second and destination as third.")
    os.Exit(1)
  }

  input_file := flag.Arg(0)
  root,e := strconv.Atoi(flag.Arg(1))
  dest,e := strconv.Atoi(flag.Arg(2))

  if e == nil {
    fmt.Printf("Could not parse command line arguments %s", e)
  }

  g := graph.ReadGraph(input_file)
  fmt.Println(g)

  fmt.Println("Starting BFS.")
  found := bfs(g, root, dest)
  
  if found {
    fmt.Println("Found the node!")
  } else {
    fmt.Println("Could not reach the destination.")
  }
}

