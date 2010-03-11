package main

import (
	"fmt"
	"strings"
	"hash"
	"crypto/md4"
	"crypto/md5"
	"crypto/sha1"
)

type filter struct {
	table  []bool
	hashes []hash.Hash
}

// Returns a new filter of given size with given hash functions.
// Size can't be smaller than the number of functions given.
func New(size int, funcs []func() hash.Hash) *filter {
	if size < len(funcs) {
		size = len(funcs)
	}
	s := make([]bool, size)
	h := make([]hash.Hash, len(funcs))
	for i, _ := range h {
		h[i] = funcs[i]()
	}
	return &filter{s, h}
}

// Adds a string to the filter.
func (f *filter) Add(e string) {
	b := make([]byte, len(e))
	strings.NewReader(e).Read(b)
	for _, i := range f.getIndices(b) {
		f.table[i] = true
	}
}

// Makes an query to the filter. Returns wheter a strings is,
// or is not part of the filter.
func (f *filter) Query(e string) bool {
	b := make([]byte, len(e))
	strings.NewReader(e).Read(b)
	for _, i := range f.getIndices(b) {
		if !f.table[i] {
			return false
		}
	}
	return true
}

// Returns a slice of indices for given slice of bytes.
// There will be a one index for every hash function in the filter.
func (f *filter) getIndices(b []byte) []int {
	indices := make([]int, len(f.hashes))
	for i, hash := range f.hashes {
		hash.Write(b)
		index := 0
		for _, v := range hash.Sum() {
			index += int(v)
		}
		indices[i] = index % len(f.table)
		hash.Reset()
	}
	return indices
}

func main() {
	f := New(10, []func() hash.Hash{md4.New, md5.New, sha1.New})
	f.add("jee")
	f.add("joo")
	f.add("foo")
	f.add("bar")
	fmt.Println(f.query("jee"))
	fmt.Println(f.query("joo"))
	fmt.Println(f.query("foo"))
	fmt.Println(f.query("bar"))
	fmt.Println()
	fmt.Println(f.query("baz"))
	fmt.Println(f.query("buz"))
	fmt.Println()
	fmt.Println(f.table)
}
