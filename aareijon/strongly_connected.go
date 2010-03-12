package main

import (
"fmt"
"scanner"
"strconv"
"os"
"container/list"
)

type graph_matrix [][]bool

// format:
// 2
// 1 0
// 1 1
// ^^ 2 nodes, links between 0->0, 1->0, 1->1
func read_graph(filename string) graph_matrix {
	if file, err := os.Open(filename, os.O_RDONLY, 0); err == nil {
		var s scanner.Scanner
		s.Init(file)
		s.Mode = scanner.ScanInts
		
		s.Scan()
		size,_ := strconv.Atoi(s.TokenText())

		// assume file is in right format etc
		graph := make([][]bool, size)
		for i := range(graph) {
			row := make([]bool, size)
			for j := range(row) {
				s.Scan()
				if s.TokenText() == "0" {
					row[j] = false
				} else {
					row[j] = true
				}				
			}
			graph[i] = row
		}
		return graph
	}
	return nil
}

func scc(graph graph_matrix) *list.List {
	count := len(graph)
	stack := make([]int, count)[0:0]
	components := list.New()
	component := list.New()
	transpose := false
	visited := make([]bool, count) /* all false by default */
	
	var traverse func (node int)
	// too lazy to pass more parameters, use variables from
	// from closure to make code less clear
	traverse = func(node int) { 
		if visited[node] {
			return
		}
		if transpose {
			component.PushBack(node)
		}

		visited[node] = true
		for i := 0; i < count; i++ {
			if transpose {
				if graph[i][node] {
					traverse(i) 
				}
			} else {
				if graph[node][i] {
					traverse(i)
				}
			}
		}
		if !transpose { // push to stack
			stack = stack[0:len(stack)+1]
			stack[len(stack)-1] = node
		}
	}

	for i := 0; i < count; i++ {
		traverse(i)
	}

	// nicer way to do this? (something else than make(...))
	for i := range(visited) { 
		visited[i] = false
	}
	
	transpose = true
	for i := len(stack)-1; i >= 0; i-- {
		traverse(stack[i])
		if component.Len() > 0 {
			components.PushBack(component)
			component = list.New()
		}
	}
	return components
}

func print_components(components *list.List) {
	fmt.Println("Components: ")
	for comps := components.Front(); comps != nil; comps = comps.Next() {
		fmt.Print("- ")
		com := comps.Value.(*list.List)
		for node := com.Front(); node != nil; node = node.Next() {
			fmt.Printf("%d ", node.Value.(int))
		}
		fmt.Println()
	}
}

func main() {
	graph := read_graph("input-1.graph")
	print_components(scc(graph))
	graph = read_graph("input-2.graph")
	print_components(scc(graph))
}
