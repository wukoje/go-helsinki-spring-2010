package main

import  ( 
	"fmt"
)

func main() {

	for i := 1; i < 101; i++ {
		switch {
			case i%5==0 && i%3==0:
				fmt.Printf("FizzBuzz")
			case i%3==0:
				fmt.Printf("Fizz")
			case i%5==0:
				fmt.Printf("Buzz")
			default:
				fmt.Printf("%d", i)				
		}
		fmt.Printf("\n")
	}
}


