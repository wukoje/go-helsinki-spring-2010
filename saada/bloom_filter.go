package main

import (
	"fmt"
)

type BloomFilter struct {
	bits []bool
	k uint
}

func create (s int) *BloomFilter {
	bf := new(BloomFilter)
	bf.bits = make([]bool, s)
	bf.k = 3
	return bf
}

func (bf *BloomFilter) add (s string) {
	for i:= uint(1); i<= bf.k; i++ {
		key := hash(s, uint(len(bf.bits)), i)
		bf.bits[key] = true
	}
}

func (bf *BloomFilter) query (s string) bool {
	result := true
	for i:= uint(1); i<= bf.k; i++ {
		key := hash(s, uint(len(bf.bits)), i)
		result = result && bf.bits[key]
	}
	return result
}

func hash (s string, n uint, k uint) uint {
	data := uint(0)
	for _,x := range []byte(s) {
		data += uint(x)
	}
	h1 := data % n
	h2 := 1 + (data % (n-1))
	return (h1 + k * h2) % n
}

func main() {
	fmt.Println("Bloom Filter")
	bf := create(123)
	fmt.Println("should not find foo:", bf.query("foo"))
	fmt.Println("adding foo")
	bf.add("foo")
	fmt.Println("should find foo:", bf.query("foo"))
	fmt.Println("should not find bar:", bf.query("bar"))
	fmt.Println("adding bar")
	bf.add("bar")
	fmt.Println("should find bar:", bf.query("bar"))
}
