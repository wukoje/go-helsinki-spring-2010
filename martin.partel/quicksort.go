package main

import (
	"rand"
	"fmt"
	"time"
)

func quicksort(a []int, done chan<- bool) {
	if len(a) > 1 {
		pivotIndex := len(a) / 2 //rand.Intn(len(a))
		pivotIndex = partition(a, pivotIndex)

		left := a[0:pivotIndex]
		right := a[pivotIndex+1 : len(a)]

		wait := make(chan bool)

		go quicksort(left, wait)
		go quicksort(right, wait)
		<-wait
		<-wait
	}
	done <- true
}

func partition(a []int, pivotIndex int) int {
	rightIndex := len(a) - 1
	pivotValue := a[pivotIndex]

	tmp := a[rightIndex]
	a[rightIndex] = a[pivotIndex]
	a[pivotIndex] = tmp

	storeIndex := 0
	for i := 0; i < rightIndex; i++ {
		x := a[i]
		if x <= pivotValue {
			tmp = a[i]
			a[i] = a[storeIndex]
			a[storeIndex] = tmp
			storeIndex++
		}
	}

	tmp = a[rightIndex]
	a[rightIndex] = a[storeIndex]
	a[storeIndex] = tmp

	return storeIndex
}


func createRandomArray(size int) []int {
	a := make([]int, size)
	for i, _ := range a {
		a[i] = rand.Intn(size * 10)
	}
	return a
}

func printArray(a []int) {
	fmt.Print("[")
	for i, v := range a {
		fmt.Printf("%d", v)
		if i < len(a)-1 {
			fmt.Print(",")
		}
	}
	fmt.Println("]")
}

func inOrder(a []int) bool {
	if len(a) > 0 {
		prev := a[0]
		for i := 1; i < len(a); i++ {
			if a[i] < prev {
				return false
			}
			prev = a[i]
		}
	}

	return true
}

func main() {
	rand.Seed(time.Seconds())

	for i := 0; i < 5; i++ {
		a := createRandomArray(10 * i)
		printArray(a)

		ch := make(chan bool)
		go quicksort(a, ch)
		<-ch

		if !inOrder(a) {
			fmt.Println("oh darn")
		}

		printArray(a)
		fmt.Println()
	}
}
