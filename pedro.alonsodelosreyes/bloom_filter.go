package main

import (
	"fmt"
	//"math"
)

func main() {

	a := make([]int, 255)
	b := &bloom{3, a, 0}
	fmt.Println(b.storage)
	fmt.Println("Add some words like: cat, dog, seal, human, apples, oranges, book, bottle, pencil, vodka.")
	b.add("cat")
	//s:="cat"
	//f:="dog"
	//fmt.Println(s[0], s[0]+s[1]+s[2])
	//fmt.Println(f[0], f[0]+f[1]+f[2])
	b.add("dog")
	b.add("seal")
	b.add("human")
	b.add("apples")
	b.add("oranges")
	b.add("book")
	b.add("bottle")
	b.add("pencil")
	b.add("vodka")
	fmt.Println(b.storage)
	fmt.Println("Search for the following: human, seal, vodka and weather")
	fmt.Println("human = ", b.query("human"))
	fmt.Println("seal = ", b.query("seal"))
	fmt.Println("vodka = ", b.query("vodka"))
	fmt.Println("weather = ", b.query("weather"))
}

type bloom struct {
	k       int   //The k were supposed to be the number of hashes but due to the lillte time i have left I'll just use 3
	storage []int // Basically to represent the concept of a bloom filter
	n       int
}

func (b *bloom) hash(num int) (int, int, int) {

	return ((num*(num+3))%len(b.storage) - 1), (num % (len(b.storage) - 1)), ((num * (num + 20)) % (len(b.storage) - 1))

}

func (b *bloom) add(e string) {

	n := 0
	for i := 0; i < len(e); i++ {

		n += (int)(e[i])

	}
	v, j, i := b.hash(n)
	fmt.Println(v, j, i)
	b.storage[i], b.storage[j], b.storage[v] = 1, 1, 1
	b.n++
}


func (b *bloom) query(e string) bool {

	n := 0
	for i := 0; i < len(e); i++ {

		n += (int)(e[i])

	}
	i, j, v := b.hash(n)
	found := true
	if b.storage[i] == 0 {

		found = false

	} else if b.storage[j] == 0 {

		found = false

	} else if b.storage[v] == 0 {

		found = false
	}
	return found
}
