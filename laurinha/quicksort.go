// the implementation of quicksort here follows the one in
// Introduction to Algorithms, 2nd edition by
// Cormen, Leiserson, Rivest and Stein
package main

import (
	"fmt"
	"rand"
)

// swaps s[a] and s[b]
func swap(s []int, a, b int) {
	s[a], s[b] = s[b], s[a]
}

// partitions s into two parts, using last element as the splitting element
// returns the new index of the splitting value
func partition(s []int) int {
	end := len(s) - 1
	split := s[end]
	i := -1
	for j := 0; j < end; j++ {
		if s[j] <= split {
			i++
			swap(s, i, j)
		}
	}
	swap(s, i+1, end)
	return i+1
}

// partitions s with a random splitting element
func randomPartition(s []int) int {
	index := rand.Intn(len(s))
	swap(s, index, len(s) - 1)
	return partition(s)
}

// sorts s
func sort(s []int) {
	if len(s) != 0 {
		split := randomPartition(s)
		if split > 0 {
			sort(s[0 : split])
		}
		if split+1 < len(s) {
			sort(s[split+1 : ])
		}
	}
}

func main() {
	slice := []int{1,123,5,23,17}
	fmt.Printf("original: %v\n", slice)
	sort(slice)
	fmt.Printf("sorted: %v\n", slice)

	slice = []int{}
	fmt.Printf("original: %v\n", slice)
	sort(slice)
	fmt.Printf("sorted: %v\n", slice)
}
