package graph

import (
  "os"
  "scanner"
  "strconv"
  "fmt"
)

type Graph struct {
  Adj [][]int // stores the adjacency matrix
  Nodes []int // stores some additional flags (colour?) for vertices if necessary
}

// creates a new Graph object
func CreateGraph(node_count int) *Graph {
  g := new(Graph)
  g.Adj = make([][]int, node_count)
  for i := 0; i<node_count; i++ {
    g.Adj[i] = make([]int, node_count)
  }
  g.Nodes = make([]int, node_count)
  return g
}

// reads the adj matrix from a file
func ReadGraph(filename string) *Graph {
  f, err := os.Open(filename, os.O_RDONLY, 0666)

  if f == nil {
    fmt.Printf("Error opening %s: %s", filename, err)
    os.Exit(1)
  }

  defer f.Close()


  var s scanner.Scanner
  s.Init(f)
  tok := s.Scan()
  node_count, err := strconv.Atoi(s.TokenText()) // the first token is the number of nodes in the graph
  g := CreateGraph(node_count)

  current_node := 0
  current_edge := 0
  
  tok = s.Scan()
  for tok != scanner.EOF && current_node < node_count {
    if tok == '.' {
      current_node++
      current_edge = 0
      tok = s.Scan()
      continue
    }
    weight, err := strconv.Atoi(s.TokenText())

    if err != nil {
      fmt.Printf("Error by Atoi(%s): %s\n", s.TokenText(), err)
      os.Exit(1)
    }

    fmt.Printf("%d -> %d  = %d\n", current_node, current_edge, weight)
    g.Adj[current_node][current_edge] = weight
    current_edge++

    tok = s.Scan()
  }
  return g
}

// returns the transpose of graph g
func (g *Graph) Transpose() *Graph {
  t := CreateGraph(len(g.Nodes))
  for i, arr := range g.Adj {
    for j, _ := range arr  {
      t.Adj[j][i] = g.Adj[i][j]
    }
  }
  return t
}
