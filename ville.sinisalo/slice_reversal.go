package main

import ( 
	"fmt"
	"rand"
	"time"
)

func main() {
	/* Generating a random table
	   of 1..15 elements and values
	   in [0,300) */
	rand.Seed(time.Nanoseconds())

	size := rand.Intn(15) + 1
	slice := make([]int, size)
	
	for i:=0; i<size; i++ {
		slice[i] = rand.Intn(300)
	}

	for _, v := range slice {
		fmt.Printf("%d\n", v)
	}
	reverse(slice)
	fmt.Printf("\n")
	for _, v := range slice {
		fmt.Printf("%d\n", v)
	}
}

func reverse(s []int) {
	half := len(s)/2
	l := len(s)-1
	for i:=0; i<half; i++ {
		s[i], s[l-i] = s[l-i], s[i]
	}
}
