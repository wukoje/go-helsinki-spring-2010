package main

import (
	"fmt"
	"rand"
	"time"
	)

func partition (array []int, left, right, pivotIndex int) int {
	pivotValue := array[pivotIndex]
	array[pivotIndex], array[right] = array[right], array[pivotIndex]
	storeIndex := left
	for i := left; i < right; i++ {
		if array[i] <= pivotValue {
			array[i], array[storeIndex] = array[storeIndex], array[i]
			storeIndex = storeIndex + 1
		}
	}
	array[storeIndex], array[right] = array[right], array[storeIndex]
	return storeIndex
}

func quicksort(array []int, left, right int) {
	if right > left {
		pivotIndex := (left + right) / 2
		pivotNewIndex := partition(array, left, right, pivotIndex)
		quicksort(array, left, pivotNewIndex - 1)
		quicksort(array, pivotNewIndex + 1, right)
	}
}

func main() {
	rand.Seed(time.Nanoseconds())
	array := make([]int, 1000)
	for i,_ := range array {
		array[i] = rand.Intn(len(array))
	}
	fmt.Printf("unordered:\t%v\n", array)
	quicksort(array, 0, len(array) - 1)
	fmt.Printf("ordered:\t%v\n", array)
}
