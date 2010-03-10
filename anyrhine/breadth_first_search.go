package main

import "fmt"
import "rand"

type graph struct {
	str   string
	adj   []*graph
	lnext *graph
}

func (g *graph) search(str string) bool {
	lpp := &g.lnext
	for lp := g; lp != nil; lp = lp.lnext {
		for _, p := range lp.adj {
			if p.lnext == nil && lpp != &p.lnext {
				*lpp = p
				lpp = &p.lnext
			}
		}
		fmt.Printf("%s ", lp.str)
		if lp.str == str {
			return true
		}
	}
	return false
}

func (g *graph) link(b *graph) {
	nadj := make([]*graph, len(g.adj)+1)
	for i, c := range g.adj {
		nadj[i] = c
	}
	nadj[len(g.adj)] = b
	// free g.adj
	g.adj = nadj
}

func main() {
	tab := make([]graph, 100)
	for i, _ := range tab {
		tab[i].str = fmt.Sprintf("n%d", i)
	}
	for i := 0; i < 500; i++ {
		a := rand.Int() % len(tab)
		b := rand.Int() % (len(tab) - 1)
		if a >= b {
			b++
		}
		tab[a].link(&tab[b])
	}
	for i := 0; i < len(tab)+1; i++ {
		start := rand.Int() % len(tab)
		fmt.Printf("search n%d:", i)
		if tab[start].search(fmt.Sprintf("n%d", i)) {
			fmt.Printf("found\n")
		} else {
			fmt.Printf("not found\n")
		}
		for j, _ := range tab {
			tab[j].lnext = nil
		}
	}
}
