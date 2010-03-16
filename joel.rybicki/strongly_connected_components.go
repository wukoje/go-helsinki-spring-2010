package main

import (
  "fmt"
  "flag"
  "os"
  "./graph"
  "container/vector"
)

var s vector.IntVector // stack

// Runs a depth-first traversal from node 'root'. 
// Pushes the traversal into stack s
func dfs(g *graph.Graph, root int) {
  if g.Nodes[root] != 0 {
    return
  }
  g.Nodes[root] = 1 // mark as visited
  for i, val := range g.Adj[root] {
    if val != 0 && g.Nodes[i] == 0 { 
      // go here if there's an edge and the node hasn't been visited
      dfs(g, i)
    }
  }
  s.Push(root)
}

// Prints the strongly connected components
func components(g *graph.Graph) {
  for len(s) < len(g.Nodes) {
    for i, visited := range g.Nodes {
      if visited == 0 {
        dfs(g, i)
        fmt.Println()
      }
    }
  }
  // hack: copy our dfs stack so the following dfs won't mess the stack
  var q vector.IntVector
  for i := 0; i<len(s); i++ {
    q.Insert(i, s[i])
  }

  // now to a dfs with the transpose of g
  t := g.Transpose()
  for len(q) != 0 {
    // clear our dfs stack
    s.Cut(0, len(s))

    // pop a new node to traverse
    v := q.Pop()
    for t.Nodes[v] != 0 && len(q) > 0 { // take out the nodes we've already visited 
      v = q.Pop()
    }

    // dfs on v in the transpose
    dfs(t, v)
    if len(s) != 0 { 
        // the dfs stack now contains a strongly connected component, print it
        fmt.Print("Component: ")
        fmt.Println(s)
    }
  }
}

func main() {
  flag.Parse()
  if flag.NArg() < 1 {
    fmt.Println("Give a path to a graph file.")
    os.Exit(1)
  }

  input_file := flag.Arg(0)

  g := graph.ReadGraph(input_file)
  fmt.Println(g.Adj)

  t := g.Transpose()
  fmt.Println(t.Adj)
  
  components(g)
}
