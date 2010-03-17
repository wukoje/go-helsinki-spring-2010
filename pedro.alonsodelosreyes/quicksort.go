package main

import (
        "fmt"
)

func main() {

	s1 := []int{2,4,6,7,1,3}

	s2 := []int {10,2,4,15,13,21,234,6,7,1,3}
	s3 := []int {0,12,213,45,78,7892,4,6,7,1,3,8,34,23,8}
	

	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	fmt.Println("s3 = ", s3)
	sort(s1)
	sort(s2)
	sort(s3)
	
        fmt.Println(" Sorted = ",s1 )
	fmt.Println(" Sorted = ",s2 )
	fmt.Println("Sorted = ",s3 )
}

func sort(s []int) {

	if len(s) < 1 {

	} else {

		pivot := s[(len(s)-1)]
		pos := len(s)-1
	
		for i:=0;i< len(s);i++ {
		
			if s[i] > pivot && pos > i {
							
				s[pos] = s[i]
				s[i] = pivot
				pos = i					

			} else if s[i] < pivot && pos < i {

				s[pos] = s[i]
				s[i] = s[pos+1]
				s[pos+1] = pivot
				pos++

			}
		}

	
			sort(s[0:pos])
			sort(s[pos+1:])
		
		}

		


}
