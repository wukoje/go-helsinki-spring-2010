package main

import (
	"fmt"
	"rand"
	"os"
	"strconv"
	"math"
	"time"
)

// Number of elements to add
const N = 100000
// Largest prime fewer than 2^63
const P = 9223372036854775783

type hashPar struct {
	a, b uint64
}

type Bloom struct {
	size uint64
	array []uint64
	params []hashPar
}

func hash(in string, a, b uint64) uint64 {
	for _,c := range []byte(in) {
		b = (b + uint64(c) * a) % P
	}
	return b
}

// Create new bloom filter with table of size m and k hash functions.
// Optimal value for k is m/n * ln(2) where n is the amount of 
// strings to be stored.
func NewBloom(m uint64, k int) *Bloom {
	if m == 0 { panic("Bloom filter size must be >0") }
	if k == 0 { panic("Bloom filter hash rounds must be >0") }
	ret := new(Bloom)
	ret.size = m
	ret.array = make([]uint64, (m-1)/64+1)
	ret.params = make([]hashPar, k)
	for k--; k>=0; k-- {
		ret.params[k] = hashPar{
			uint64(rand.Int63n(P-1)+1), // a 
			uint64(rand.Int63n(P-1)) }  // b
	}
	return ret
}

func (b *Bloom) Add(s string) {
	for _,p := range b.params {
		h := hash(s, p.a, p.b) % b.size
		b.array[h/64] |= 1 << (h%64)
	}
}

func (b *Bloom) Query(s string) bool {
	for _,p := range b.params {
		h := hash(s, p.a, p.b) % b.size
		if b.array[h/64] & (1 << (h%64)) == 0 { return false }
	}
	return true
}

// Generate a random string of length length.
func randString(length int) string {
	ret := make([]byte, length)
	for i := range ret {
		ret[i] = byte(rand.Intn(256))
	}
	return string(ret)
}

func cmdError() {
	fmt.Printf("Usage: %s load queries table_size\n\n", os.Args[0])
	fmt.Printf("where load is the amount of elements in the set, "+
		   "queries is the amount of test\nqueries to perform and "+
		   "table_size is the size of the lookup table.\nNumber "+
		   "of hash functions is automatically computed to the "+
		   "optimal value of\ntable_size/load*ln(2)\n"+
		   "table_size should be a *prime number* closest (+/-) "+
		   "to 2^n + 2^(n+1) for some n.\n\n"+
		   "Try for example:\n%s 10000 100000 98299\n",os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 4 { cmdError() }
	load, err := strconv.Atoui64(os.Args[1])
	if err != nil || load == 0 { cmdError() }
	quer, err := strconv.Atoui64(os.Args[2])
	if err != nil { cmdError() }
	size, err := strconv.Atoui64(os.Args[3])
	if err != nil || size <= 0{ cmdError() }

	/* Calculate optimal amount of hash rounds */
	hash := int(float(size) / float(load) * math.Ln2 +.5)

	rand.Seed(time.Nanoseconds())
	b := NewBloom(size, hash)

	fmt.Printf("Using structure of size %.0f kB for set of size %d, "+
		   "%d hash rounds.\n", float(size)/(8*1024), load, hash)

	fmt.Printf("Adding .. ");
	for i:=uint64(0); i<load; i++ { 
		b.Add(randString(200+rand.Intn(100))) 
	}
	fmt.Printf("\nQuerying .. ", )
	tot := uint64(0)
	for i:=uint64(0); i<quer; i++ {
		if b.Query(randString(100+rand.Intn(100))) { tot++ }
	}
	fmt.Printf("%.3f%% false positives.\n", 
			100.0*float(tot)/float(quer))
	fmt.Printf("%d/%d\n", tot, quer)
}
