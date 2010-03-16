package main

import "rand"
import "fmt"
import "time"

const (
	Naccept  = 10
	Npropose = 10
	Nlearn   = 10
	Nproc    = Naccept + Nlearn + Npropose
	Nbuf     = 100

	Tprep = iota
	Tpromise
	Taccept
	Taccepted
)

type msg struct {
	dst  int
	src  int
	kind int
	n    int
	nv   int
	val  *string
}

// routing table
var acceptid [Naccept]int
var proposeid [Npropose]int
var learnid [Nlearn]int

var ctab [Nproc]*chan *msg
var router *chan *msg

func accept(id int, input *chan *msg) {
	var val *string = nil
	n, nv := -1, -1
	for {
		m := <-*input
		switch m.kind {
		case Tprep:
			if m.n > n {
				n = m.n
				nm := &msg{m.src, id, Tpromise, n, nv, val}
				*router <- nm
			}
		case Taccept:
			if m.n == n {
				nv = m.nv
				val = m.val
				nm := &msg{m.src, id, Taccepted, n, nv, val}
				*router <- nm
				for _, dst := range learnid {
					nm := &msg{dst, id, Taccepted, n, nv, val}
					*router <- nm
				}
			} else {
			}
		}
	}
}

func propose(id int, xc *chan int, input *chan *msg, pval *string) {
	var val *string = nil
	pid := id
	nprom := 0
	accsent := 0
	naccept := 0
	nv := -1
	ticker := time.NewTicker(100000000 + int64(rand.Intn(1000000)))
out: for {
		for _, dst := range acceptid {
			m := &msg{dst, id, Tprep, pid, 0, nil}
			*router <- m
		}
	loop: for {
			select {
			case m := <-*input:
				switch m.kind {
				case Tpromise:
					if m.nv > nv {
						nv = m.nv
						val = m.val
					}
					for nprom++; nprom > Naccept/2 && accsent < nprom; accsent++ {
						if val == nil {
							val = pval
							nv = pid
						}
						nm := &msg{m.src, id, Taccept, pid, nv, val}
						*router <- nm
					}
				case Taccepted:
					naccept++
					if naccept > Naccept/2 {
						break out
					}
				default:
					break loop
				}
			case <-ticker.C:
				break loop
			}
		}
		naccept, nprom, accsent = 0, 0, 0
		nv = -1
		val = nil
		pid += Npropose

	}
	fmt.Printf("%d proposed %s got %s\n", id, *pval, *val)
	ticker.Stop()
	*xc <- 0
}

func learn(id int, input *chan *msg) {
	for {
		<-*input
		//fmt.Printf("%d learned %d.%s\n", id, m.nv, *m.val)
	}
}

func route(input *chan *msg) {
	var delay [Nproc]*msg
	for {
		m := <-*input
		r := rand.Intn(1000)
		switch {
		case r > 100:
			*ctab[m.dst] <- m
		case r >= 30 && r <= 100:
			if delay[m.dst] != nil {
				*ctab[m.dst] <- delay[m.dst]
				delay[m.dst] = nil
			} else {
				delay[m.dst] = m
			}
		case r < 30:
		}
	}
}

func main() {
	id := 0
	nc := make(chan *msg, Nbuf)
	nc = make(chan *msg, Nbuf)
	router = &nc
	go route(router)
	for i := 0; i < Naccept; i++ {
		nc := make(chan *msg, Nbuf)
		ctab[id] = &nc
		acceptid[i] = id
		go accept(id, &nc)
		id++
	}
	for i := 0; i < Nlearn; i++ {
		nc := make(chan *msg, Nbuf)
		ctab[id] = &nc
		learnid[i] = id
		go learn(id, &nc)
		id++
	}
	xc := make(chan int)
	for i := 0; i < Npropose; i++ {
		nc := make(chan *msg, Nbuf)
		ctab[id] = &nc
		proposeid[i] = id
		val := fmt.Sprintf("prop%d", i)
		go propose(id, &xc, &nc, &val)
		id++
	}
	for i := 0; i < Npropose; i++ {
		<-xc
	}
}
