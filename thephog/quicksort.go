package main

import "fmt"
import "rand"

func swap(s []int, index1 int, index2 int) { s[index1], s[index2] = s[index2], s[index1] }

func qsort(s []int, left int, right int) {

	if left >= right {
		return
	}

	//swap random number to the end of the array
	//in order to get a good enough pivot
	swap(s, rand.Intn(right-left)+left, right)
	pivot := s[right]

	i, j := left, right-1
	for {

		for s[i] < pivot {
			i++
		}

		for s[j] > pivot && j > left {
			j--
		}

		if i < j {
			swap(s, i, j)
		} else {
			break
		}

	}

	s[right] = s[i]
	s[i] = pivot

	qsort(s, left, i-1)
	qsort(s, i+1, right)
}

func sort(s []int) { qsort(s, 0, len(s)-1) }

func main() {

	slice := [...]int{5, 123, 22, 8, 12, 1, 2, -5, 19, 10000, 125, 88, 1337, 293, 1, 0, 17, 15, -88, -35}[0:]

	fmt.Printf("Unsorted slice: \n")
	for _, v := range slice {
		fmt.Printf("%d , ", v)
	}

	sort(slice)
	fmt.Printf("\nSorted slice: \n")
	for _, v := range slice {
		fmt.Printf("%d , ", v)
	}

	fmt.Printf("\n")
}
