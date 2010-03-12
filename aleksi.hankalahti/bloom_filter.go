package main

import "fmt"

type Bloom struct {
	filter []bool
	k      uint
}

func New(size int) *Bloom {
	k := uint(float(size)*0.1*0.7) // by Stetson and Harrison
	return &Bloom{make([]bool, size), k}
}

func (b *Bloom) Add(s string) {
	for i := uint(1); i <= b.k; i++ {
		b.filter[hash(s, uint(len(b.filter)), i)] = true
	}
}

func (b *Bloom) Query(s string) (found bool) {
	found = true
	for i := uint(1); i <= b.k; i++ {
		found = found && b.filter[hash(s, uint(len(b.filter)), i)]
	}
	return
}

func hash(s string, n, k uint) uint {
	// Shamelessly stolen from CLRS. Doesn't work too well if n is not a prime.
	var data uint = 0
	for _, b := range []byte(s) {
		data += uint(b)
	}
	h1 := data % n
	h2 := 1 + (data % (n-1))
	return (h1 + k * h2) % n
}

func main() {
	var b *Bloom = New(101) // FIXME: proper way to use this?
	b.Add("foo")
	b.Add("bar")
	b.Add("baz")
	fmt.Printf("Query 'foo': %v\n", b.Query("foo"))
	fmt.Printf("Query 'blaa': %v\n", b.Query("sblaia"))
}
