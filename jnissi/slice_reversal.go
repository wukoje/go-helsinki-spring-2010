package main

import (
	"fmt"
	"rand"
)

func reverse(a []int) {
	l := len(a)
	for i, _ := range a[0 : l/2] {
		a[i], a[l-1-i] = a[l-1-i], a[i]
	}
}

func printReversePrint(a []int) {
	fmt.Printf("Original:\t%v\n", a)
	reverse(a)
	fmt.Printf("Reversed:\t%v\n", a)
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := array[2:9]
	printReversePrint(a)
	a = make([]int, 20)
	for i, _ := range a {
		a[i] = rand.Int()%99 + 1
	}
	printReversePrint(a)
}
