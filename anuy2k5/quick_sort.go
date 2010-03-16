package main

import(
	"fmt"
	"rand"
)

func sort (s []int) {
 quicksort(s, 0, len(s)-1)
}
func quicksort(s []int, low int, high int){
	low = 0
	high = len(s)-1
	pivot := s[0]

	for low != high {
		if s[high] < pivot{
			s[low] = s[high]
			low++
		} else {
			high--
			continue
		}
		
		for low != high {
			if s[low] >= pivot{
				s[high] = s[low]
				high--
				break
			} else {
				low++
			}
		}
	}

	s[low] = pivot

	if low > 1 {
		sort(s[0:low])
	}
	if high < len(s)-2 {
		sort(s[high+1:len(s)])
	}
}

func main(){
	//Case1	
	s := make([]int, 25)
	s = []int{7,42,86,24,1,9}
	fmt.Print("\ncase 1: \nmade slice: ")
	fmt.Println(s)
	fmt.Print("Quick sorted slice")
	sort(s)
	fmt.Println(s)

	//Case2
	for i, _ := range s {
 		s[i] = rand.Intn(100)
	 }
	fmt.Print("\ncase 2: \nmade slice: ")
	fmt.Println(s)
	fmt.Print("quicksorted slice ")
	sort(s)
	fmt.Println(s)
	fmt.Printf("\n")
}


