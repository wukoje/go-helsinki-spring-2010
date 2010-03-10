package main

import fmt "fmt"
// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
	if s != nil {
		quicksort(s, 0, len(s)-1)
	}	
}

func quicksort(s []int, low int, high int) {
		i := low
		j := high

		pivot := s[(low + high) / 2]

		for i <= j {
			for s[i] < pivot {
				i++
			}

			for s[j] > pivot {
				j--
			}

			if i <= j {
				s[i], s[j] = s[j], s[i];
				i++
				j--
			}
		}		

		// Rekursio
		if low < j {
			quicksort(s, low, j)
		}
		if i < high {
			quicksort(s, i, high)
		}
}


func main() {
	s := []int{6,12,5,2,34,64,7,4,3,43,35,32}
	
	for _,j := range s{
		fmt.Printf("%d,", j)	
	}
	fmt.Printf("\n")	
	sort(s)
	fmt.Printf("\n")
	for _,j := range s{
		fmt.Printf("%d,", j)
	
	}
	fmt.Printf("\n")
}
