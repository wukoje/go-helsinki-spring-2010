/**

1 quicksort.go
// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int)

*/

package main

import (
	"fmt"
)

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {

	fmt.Printf("now sorting: %v\n", s)

	l := 1 // how many values have been larger so far (or "days since accident")

	i := 0
	for j := 1; j < len(s); j++ {
		switch {
			case s[i] <= s[j]:
				l++
			case s[i] > s[j]:
				a := s[j]
				k := 1
				for k <= l {
					s[j-k+1] = s[j-k]
					k++
				}
				s[i] = a
				i++
		}
	}

	fmt.Printf("sorted: %v [%v] %v\n\n", s[0:i], s[i], s[i+1:len(s)])

	if len(s[0:i]) > 1 {
		sort(s[0:i])
	}

	if len(s[i+1:len(s)]) > 1 {
		sort(s[i+1:len(s)])
	}
}

func main() {

	// list for testing
	t1 := []int{2,144,233,8,3,89,5,13,1,377,21,1,55,34}

	fmt.Printf("\nfull slice before: %v\n\n", t1)
	sort(t1)
	fmt.Printf("full slice after: %v\n\n", t1)

}
