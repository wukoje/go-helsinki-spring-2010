package main

import (
	"fmt"
)

func swap(array []int, x int, y int) {
    a := array[x]
    array[x] = array[y]
    array[y] = a
}

func partition(array []int, left int, right int, pivotIndex int) int {
    pivotValue := array[pivotIndex]
    swap(array, pivotIndex, right)
    storeIndex := left
    for i := left; i < right; i++ {
        if array[i] <= pivotValue {
            swap(array, i, storeIndex)
            storeIndex = storeIndex + 1
        }
    }
    swap(array, storeIndex, right)
    return storeIndex
}

func quicksort(array []int, left int, right int) {
    if right > left {
        pivotIndex := (left+right)/2
        pivotNewIndex := partition(array, left, right, pivotIndex)
        quicksort(array, left, pivotNewIndex - 1)
        quicksort(array, pivotNewIndex + 1, right)
    }
}

func sort(s []int) {
    quicksort(s, 0, len(s)-1)
}

func main() {
    s := make([]int, 20)
    for i,_ := range s {
        s[i] = i+i*2%6
    }
    fmt.Println(s)
    sort(s)
    fmt.Println(s)
}
