/*
 *	find maximum flow through repeated breadth-first search for an
 *	augmenting path (edmonds-karp). some additional inefficiency
 *	due to the adjlist edge representation.
 */

package main

import "fmt"
import "rand"

type edge struct {
	fwd      *node
	back     *node
	next     *edge
	use, cap int
}

type node struct {
	str  string
	adj  *edge
	next *node
	p    *node
	min  int // free capacity from end of p to here.
}

const MAXCAP = 1000
const debug bool = true

var black, white node

func augment(g *node, d *node) *node {
	lpp := &g.next
	lp := g
	for lp != nil {
		for e := lp.adj; e != nil; e = e.next {
			fwd := e.fwd
			if e.use < e.cap && fwd.next == &white {
				fwd.next = nil
				fwd.p = lp
				if lp.min < e.cap-e.use {
					fwd.min = lp.min
				} else {
					fwd.min = e.cap - e.use
				}
				if fwd == d {
					return fwd
				}
				*lpp = fwd
				lpp = &fwd.next
			}
		}
		lp = lp.next
	}
	return nil
}

func maxflow(tab []*node, g *node, b *node) int {
	cap := 0
	for {
		for i, _ := range tab {
			tab[i].min = MAXCAP + 1
			tab[i].next = &white
			tab[i].p = nil
		}
		path := augment(g, b)
		if path == nil {
			break
		}
		if debug {
			fmt.Printf("aug:%d", path.min)
		}
		cap += path.min
		for p := path; p != nil; p = p.p {
			if debug {
				fmt.Printf(" %s", p.str)
			}
			for e := p.adj; e != nil; e = e.next {
				if e.back == p {
					e.use += path.min
				}
			}
		}
		if debug {
			fmt.Printf("\n")
		}
	}
	return cap
}

func link(g *node, b *node, cap int) { g.adj = &edge{b, g, g.adj, 0, cap} }

func zap(tab []*node) {
	for i, _ := range tab {
		tab[i].next = &white
		tab[i].p = nil
		for e := tab[i].adj; e != nil; e = e.next {
			e.use = 0
		}
	}
}

func main() {
	tab := make([]*node, 1000)
	for i, _ := range tab {
		tab[i] = &node{fmt.Sprintf("n%02d", i), nil, nil, nil, 0}
	}
	for i := 0; i < 2000; i++ {
		a := rand.Int() % len(tab)
		b := rand.Int() % (len(tab) - 1)
		if a <= b {
			b++
		}
		link(tab[a], tab[b], rand.Intn(MAXCAP))
	}
	for i := 0; i < len(tab); i++ {
		start := rand.Int() % len(tab)
		flow := maxflow(tab, tab[start], tab[i])
		fmt.Printf("maxflow(%s %s): %d\n", tab[start].str, tab[i].str, flow)
		zap(tab)
	}
}
