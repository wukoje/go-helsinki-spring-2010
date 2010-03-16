package main

import "fmt"
import "strconv"
import "container/list"

type HashMethod interface {
     hash(data int32, length uint) uint
}

func firstHash(data int32, length uint) uint {
     	return uint(data) % (length * 8)
}

func secondHash(data int32, length uint) uint {
     	return (2 * uint(data)) % (length * 8)
}

func thirdHash(data int32, length uint) uint {
     	return (3 * uint(data)) % (length * 8)
}

type hashOne func(data int32, length uint) uint 
type hashTwo func(data int32, length uint) uint
type hashThree func(data int32, lenght uint) uint

func (h hashOne) hash(data int32, length uint) uint {
     return firstHash(data, length)
}

func (h hashTwo) hash(data int32, length uint) uint {
     return secondHash(data, length)
}

func (h hashThree) hash(data int32, length uint) uint {
     return thirdHash(data, length)
}

type BloomFilter struct {
	filter []byte
	hashFuncs []HashMethod
}

func (b *BloomFilter) add(data int32) {
     	for _,v := range b.hashFuncs {
	    h := v.hash(data, uint(len(b.filter)))
	    var theByte byte = 0
	    if bitPos := h % 8; bitPos != 0 {
	       theByte = 1 << bitPos
	    }
	    b.filter[h / 8] |= theByte
	}
}

func (b *BloomFilter) check(data int32) bool {
     	for _,v := range b.hashFuncs {
	    h := v.hash(data, uint(len(b.filter)))

	    var theByte byte = 0
	    if bitPos := h % 8; bitPos != 0 {
	       theByte = 1 << bitPos
	    }

	    if b.filter[h / 8] & theByte != theByte {
	       return false
	    }
	}
	return true
}

func (b *BloomFilter) String() string {
	s := ""
	for _, v := range b.filter {
		s += strconv.Itoa(int(v)) + ", "
	}
	return s
}

func NewBloomFilter(size int) *BloomFilter { 
     filter := BloomFilter{make([]byte, size), nil}
     h1 := new(hashOne)
     h2 := new(hashTwo)
     h3 := new(hashThree)
     filter.hashFuncs = []HashMethod {h1, h2, h3}
     return &filter
}

func Find(data int32, l *list.List) bool {
	c := l.Iter()

	for {
		if closed(c) {
			return false
		}

		elem := <-c
		val, _ := elem.(int32)
		if val == data {
			return true
		}
	}
	return false
}

func main() {

	addNums := []int32{5, 10, 15, 20, 25, 30, 35, 8, 16, 24, 32, 999, 1288, 1289, 1334,
		1689, 2555, 8000, 1001, 1337, 1555, 300000, 400000, 500000,
		800000, 666666, 13371337, 5555555, 256246}
	findNums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 888, 999, 1001, 1333, 1337, 1554,
		1555, 1600, 300000, 444444, 500001, 500000, 13371337}

	filter := NewBloomFilter(16)
	l := list.New()

	fmt.Printf("Adding numbers...\n")
	for i, v := range addNums {
		l.PushBack(v)
		filter.add(v)
		fmt.Printf("%d, ", v)
		if i%6 == 0 && i != 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\nBloom filter: %v \n", filter)
	fmt.Printf("\nSearching for numbers...\n")
	for _, v := range findNums {
		fmt.Printf("%d", v)

		if filter.check(v) {
			fmt.Printf(" BLOOM: YES ")
		} else {
			fmt.Printf(" BLOOM: NO ")
		}

		if Find(v, l) {
			fmt.Printf(" LIST: YES ")
		} else {
			fmt.Printf(" LIST: NO")
		}

		fmt.Printf("\n")
	}

}
