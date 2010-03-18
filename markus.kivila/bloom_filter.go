/*
2 bloom_filter.go

Implement a bloom filter that supports add and query
*/
package main

import (
	"fmt"
	"rand"
	"strings"
	"container/vector"
	"hash/adler32"
	"hash/crc32"
)

type bitfield struct {
	bytes []byte
}

func newBitfield(size int) *bitfield {
	return &bitfield{make([]byte, size/8+1)}
}

func (b *bitfield) set(i uint) {
	b.bytes[i/8] |= 0x01 << (i % 8)
}

func (b *bitfield) isSet(i uint) bool {
	return (b.bytes[i/8] & (0x01 << (i % 8))) != 0
}

func (b *bitfield) max() uint {
	return uint(len(b.bytes) * 8)
}

// bitfield marks which elements the set contains
// and the functions are used to map arbitrary 
// data to the bitfield
type bloom struct {
	bf *bitfield
	funcs []func([]byte) uint
}

func newBloom(size int) *bloom {
	b := new(bloom)
	b.bf = newBitfield(size)
	// function setup is a bit too static. 
	// maybe do an iface and user set functions later
	b.funcs = make( []func([]byte) uint, 2)
	b.funcs[0] = func (d []byte) uint {
		return uint(crc32.ChecksumIEEE(d)) % b.bf.max()
	}
	b.funcs[1] = func (d []byte) uint {
		return uint(adler32.Checksum(d)) % b.bf.max()
	}
	return b
}

func (b *bloom) add(data []byte) {
	for i:=0; i<len(b.funcs); i++ {
		b.bf.set( b.funcs[i](data) )
		fmt.Printf("Inserted to slot %d\n", b.funcs[i](data))
	}
}

func (b *bloom) query(data []byte) bool {
	for i:=0; i<len(b.funcs); i++ {
		if slot := b.funcs[i](data); b.bf.isSet(slot) {
			fmt.Printf("Found from slot %d\n", slot)
			return true
		}
	}
	return false
}


// Utility functions for testing...
func (b *bitfield) printHex() {
	for i:=0; i<len(b.bytes); i++ {
		fmt.Printf("%02x ", b.bytes[i])
		if i % 16 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func (b *bitfield) print() {
	mask := uint8( 1<<7 )
	for i:=0; i<len(b.bytes); i++ {
		fmt.Printf(" %d-%d: ", i*8, i*8+7)
		bits := b.bytes[i]
		for j:=0; j<8; j++ {
			if mask & bits != 0 {
				fmt.Printf("1")
			} else {
				fmt.Printf("0")
			}
			bits <<= 1
		}
		if i % 2 == 0 {
			fmt.Printf("\n")
		}
	}
	fmt.Printf("\n")
}

func vec2uint(v interface{}) uint {
	return uint(v.(int))
}

func test_bitfield(size int) {
	bf := newBitfield(size)

	fmt.Printf("Created new bitfield:\n")
	bf.print()
	for i:=0; i<size; i++ {
		if bf.isSet(uint(i)) {
			fmt.Printf("ERR: Byte %d set in empty bitfield\n")
		}
	}

	set := new(vector.Vector)
	for i:=0; i<size/8+1; i++ {
		set.Push(rand.Int() % size)
	}
	for i := range set.Iter() {
		fmt.Printf("Setting bit %d\n", vec2uint(i))
		bf.set(vec2uint(i))
		if !bf.isSet(vec2uint(i)) {
			fmt.Printf("ERR: Byte %d not set\n", i)
		}
	}
	bf.printHex()
	bf.print()
}

func test_bloom() {
	data := []string{ "Foo", "Bar", "Foobar" }
	b := newBloom(1024)
	buf := make([]byte, 1024)

	for i:=0; i<len(data); i++ {
		r := strings.NewReader(data[i])
		if n, err := r.Read(buf); err != nil {
			fmt.Printf("ERR: string.Reader exploded!\n")
		} else {
			fmt.Printf("Feeding bloom with: %v\n", buf[0:n])
			b.add(buf[0:n])
			if !b.query(buf[0:n]) {
				fmt.Printf("ERR: Bloom query failed!!\n")
			}
		}
	}
}

func main() {
	test_bitfield(8)
	fmt.Printf("Testing w/ a larger bf\n")
	test_bitfield(1024)

	test_bloom()
}
