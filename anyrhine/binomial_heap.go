package main

import "fmt"

type bnode struct {
	key    int
	rank   int8
	parent *bnode
	next   *bnode
	clist  *bnode
}

func link(p *bnode, q *bnode) *bnode {
	if q.key < p.key {
		p, q = q, p
	}
	q.parent = p
	q.next = p.clist
	p.clist = q
	p.next = nil
	p.parent = nil
	p.rank++
	return p
}

func reverse(p *bnode) *bnode {
	var prev *bnode = nil
	for p != nil {
		t := p.next
		p.next = prev
		prev = p
		p = t
	}
	return prev
}

func merge(rp **bnode, a *bnode, b *bnode) {
	var rank int8
	var trail *bnode
	switch {
	case a == nil:
		*rp = b
		return
	case b == nil:
		*rp = a
		return
	case a.rank < b.rank:
		rank = a.rank
	case a.rank >= b.rank:
		rank = b.rank
	}
	sump := rp
	ptab := [3]*bnode{a, b, nil}
	nilcount := 1
	trail = nil
	for ; nilcount < 2; rank++ {
		var tab [3]*bnode
		i := 0
		for j, e := range ptab {
			if e != nil && e.rank == rank {
				tab[i] = e
				ptab[j] = e.next
				i++
			}
		}
		switch i {
		case 3:
			ptab[2] = link(tab[0], tab[1])
			*sump = tab[2]
			sump = &tab[2].next
		case 2:
			ptab[2] = link(tab[0], tab[1])
		case 1:
			*sump = tab[0]
			sump = &tab[0].next
		}
		nilcount = 0
		trail = nil
		for _, e := range ptab {
			if e == nil {
				nilcount++
			} else {
				trail = e
			}
		}
	}
	*sump = trail
}

func insert(pp **bnode, key int) *bnode {
	n := new(bnode)
	n.key = key
	merge(pp, *pp, n)
	return n
}

func delmin(rpp **bnode) (int, bool) {
	if *rpp == nil {
		return -1, false
	}
	minp := rpp
	for pp := rpp; (*pp) != nil; pp = &(*pp).next {
		if (*pp).key < (*minp).key {
			minp = pp
		}
	}
	min := *minp
	*minp = min.next
	merge(rpp, *rpp, reverse(min.clist))
	return min.key, true
}

func decrease(p *bnode, nkey int) *bnode {
	if p.key < nkey {
		return p
	}
	for p.parent != nil && p.parent.key > nkey {
		p.key = p.parent.key
		p = p.parent
	}
	p.key = nkey
	return p
}

const N = 100000

func main() {
	var b, p *bnode
	x := 128941
	for i := 0; i < N; i++ {
		x = x<<10 + x>>6 + i
		p = insert(&b, x)
		decrease(p, p.key-1000)
	}
	for i := 0; i < N; i++ {
		x, ok := delmin(&b)
		x = x
		//fmt.Printf("%d\n", x)
		if !ok {
			fmt.Print("katastrof\n")
		}
	}
}
