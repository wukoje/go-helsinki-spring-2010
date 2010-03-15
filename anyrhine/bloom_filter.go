/*
 *	bloom filter.
 *
 *	test with something like
 *	sed 20000q broken_telephone.words | 8.bloom_filter | grep '^match' | wc -l
 */

package main

import (
	"fmt"
	"bufio"
	"os"
)

type Bloom []uint32

var seed = []uint32{
	1208219249, 1418167501, 1629598171, 1842339523,
	2056243181, 2271187571, 2487098177, 2703848461}

var nfunc = 5 // must be less than len(seed).

func hash(h uint32, s string) uint32 {
	for _, c := range s {
		h += uint32(c)
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return h
}

func (b Bloom) add(s string) {
	tab := []uint32(b)
	bbits := uint32(len(tab)) * 32
	for i := 0; i < nfunc; i++ {
		x := hash(seed[i], s) % bbits
		tab[x/32] |= 1 << (x % 32)
	}
}

func (b Bloom) query(s string) bool {
	tab := []uint32(b)
	bbits := uint32(len(tab)) * 32
	for i := 0; i < nfunc; i++ {
		x := hash(seed[i], s) % bbits
		if tab[x/32]&(1<<(x%32)) == 0 {
			return false
		}
	}
	return true
}

func main() {
	b := make(Bloom, 10000)
	r := bufio.NewReader(os.Stdin)
	for {
		if s, err := r.ReadString('\n'); err == nil {
			s = s[0 : len(s)-1]
			if b.query(s) {
				fmt.Printf("match '%s'\n", s)
			} else {
				fmt.Printf("no match '%s'\n", s)
				b.add(s)
			}
		} else {
			break
		}
	}
}
