package main

import (
	"container/vector"
	"fmt"
	"os"
	"strings"
	"time"
)

type node struct {
	value string
	edges []*node
	visited uint
}

// Search for v in graph with root node n
func (root *node) search(v string) *node {
	if root.value == v { return root }
	root.visited++ // Should check for overflow
	visit := root.visited
	
	q := new(vector.Vector)
	q.Push(root)
	for q.Len() > 0 {
		n := q.Pop().(*node)
		for _,t := range n.edges {
			if t.value == v { return t }
			if t.visited < visit {
				t.visited = visit
				q.Push(t)
			}
		}
	}
	return nil
}

func fatal(s string) {
	fmt.Fprintf(os.Stderr,
		    "Test failed: search(\"%s\") did not return node %s\n",
		    s, strings.ToLower(s))
	os.Exit(1)
}

// main runs a full test set on the structure
// Number of tests to perform
const N_TESTS = 1000
func main() {
	root := node{ "A", make([]*node,0), 0 }
	b := node{ "B", make([]*node,0), 0 }
	c := node{ "C", make([]*node,0), 0 }
	d := node{ "D", make([]*node,0), 0 }
	e := node{ "E", make([]*node,0), 0 }
	f := node{ "F", make([]*node,0), 0 }
	g := node{ "G", make([]*node,0), 0 }
	h := node{ "H", make([]*node,0), 0 }
	i := node{ "I", make([]*node,0), 0 }
	j := node{ "J", make([]*node,0), 0 }
	k := node{ "K", make([]*node,0), 0 }
	l := node{ "L", make([]*node,0), 0 }
	m := node{ "M", make([]*node,0), 0 }
	n := node{ "N", make([]*node,0), 0 }
	o := node{ "O", make([]*node,0), 0 }
		
	/* Create an example graph:
	 .--==A___  
	/   \/ \  \
	|   B   C  D==-.
	|  /|\ /|  |  \ \
	| E F-G H  I==J /
	| |/ \  |      /
	| K   '-L-----'
	\ |    / \
	 'M   N   O */

	root.edges = &[...]*node{&b,&c,&d}
	b.edges = &[...]*node{&root,&e,&f,&g}
	c.edges = &[...]*node{&g,&h}
	d.edges = &[...]*node{&i,&j,&l}
	e.edges = &[...]*node{&k}
	f.edges = &[...]*node{&g,&l}
	h.edges = &[...]*node{&l}
	i.edges = &[...]*node{&j}
	j.edges = &[...]*node{&i}
	k.edges = &[...]*node{&f,&m}
	l.edges = &[...]*node{&n,&o}
	m.edges = &[...]*node{&root}

	nanos := time.Nanoseconds()
	for I:=0; I<N_TESTS; I++ {
		if root.search("A") != &root { fatal("root") }
		if root.search("B") != &b { fatal("B") }
		if root.search("C") != &c { fatal("C") }
		if root.search("D") != &d { fatal("D") }
		if root.search("E") != &e { fatal("E") }
		if root.search("F") != &f { fatal("F") }
		if root.search("G") != &g { fatal("G") }
		if root.search("H") != &h { fatal("H") }
		if root.search("I") != &i { fatal("I") }
		if root.search("J") != &j { fatal("J") }
		if root.search("K") != &k { fatal("K") }
		if root.search("L") != &l { fatal("L") }
		if root.search("M") != &m { fatal("M") }
		if root.search("N") != &n { fatal("N") }
		if root.search("O") != &o { fatal("I") }
		if root.search("X") != nil {
			fmt.Fprintf(os.Stderr, "Search for non-existent "+
					       "value returned a node.\n")
			os.Exit(1)
		}
	}
	nanos = time.Nanoseconds()-nanos

	fmt.Fprintf(os.Stderr, "All tests completed in %d ms\n",
				nanos/1000000)
}
