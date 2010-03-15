package main

import "fmt"
import "rand"
import "flag"

type graph struct {
	str      string
	adj      []*graph
	next     *graph
	dfi, low int
}

var rndseed = flag.Int64("s", 13, "random seed")

var snil graph

type Tarjan struct {
	dfi   int
	stack *graph
	scc   []*graph
}

func (t *Tarjan) tarjanr(p *graph) {
	p.dfi, p.low = t.dfi, t.dfi
	t.dfi++
	p.next = t.stack
	t.stack = p
	for _, q := range p.adj {
		if q.dfi == 0 {
			t.tarjanr(q)
			if q.low < p.low {
				p.low = q.low
			}
		} else if q.next != &snil && q.low < p.low {
			p.low = q.low
		}
	}
	if p.low == p.dfi {
		scc := t.stack
		t.stack = p.next
		p.next = nil
		t.scc = t.scc[0 : len(t.scc)+1]
		t.scc[len(t.scc)-1] = scc
	}
}

func tarjan(tab []graph) []*graph {
	g := &graph{"root", make([]*graph, len(tab)), nil, 0, 0}
	t := &Tarjan{1, nil, make([]*graph, 0, len(tab))}
	for i, _ := range tab {
		tab[i] = graph{tab[i].str, tab[i].adj, &snil, 0, 0}
		g.adj[i] = &tab[i]
	}
	t.tarjanr(g)
	pp := &t.scc[len(t.scc)-1]
	for *pp != g && *pp != nil {
		pp = &(*pp).next
	}
	if *pp == g {
		*pp = g.next
	}
	return t.scc
}

func link(g *graph, b *graph) {
	nadj := make([]*graph, len(g.adj)+1)
	for i, c := range g.adj {
		nadj[i] = c
	}
	nadj[len(g.adj)] = b
	g.adj = nadj
}

func main() {
	flag.Parse()
	tab := make([]graph, 100)
	for i, _ := range tab {
		tab[i].str = fmt.Sprintf("n%d", i)
	}
	rand.Seed(*rndseed)
	for i := 0; i < 200; i++ {
		a := rand.Intn(len(tab))
		b := rand.Intn(len(tab) - 1)
		if a >= b {
			b++
		}
		link(&tab[a], &tab[b])
	}
	for _, scc := range tarjan(tab) {
		fmt.Print("scc:")
		for q := scc; q != nil; q = q.next {
			fmt.Printf(" %s", q.str)
		}
		fmt.Print("\n")
	}
}
