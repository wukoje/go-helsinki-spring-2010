package main

import (
        "fmt"
)

func main() {

	s1 := []int{1,2,3};

	s2 := []int{4,5,6,7,8,9}

	s3 := []int{9,9,9,1,1,1,3,3,3,4}


        fmt.Println("s1 = ", s1 )

	fmt.Println("Reversed = ",reverse(s1))

	fmt.Println("s2 = ",s2)

	fmt.Println("Reversed = ",reverse(s2))

	fmt.Println("s3 = ",s3)

        fmt.Println("Reversed = ",reverse(s3))

}

func reverse(s []int) []int {

	
	for i:= 0 ; i < len(s)/2; i++ {

		 a := s[i]
		 b := s[(len(s)-1)-i] 

		s[i] = b
		s[(len(s)-1)-i] = a

	}

	return s
}

