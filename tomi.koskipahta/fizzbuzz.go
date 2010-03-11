package main

import fmt "fmt"

func main() {
	
	for i := 1; i <= 100; i++ {
		
		//divisible by 3 and not by 5
		if i % 3 == 0 && i % 5 != 0 {
			fmt.Println("Fizz")
		
		//divisible by 5 and not by 3	
		} else if i % 5 == 0 && i % 3 != 0 {
			fmt.Println("Buzz")
		
		//divisible by both 3 and 5	
		} else if i % 5 == 0 && i % 3 == 0 {
			fmt.Println("FizzBuzz")
		
		//if not divisible by 3 or 5, print the number	
		} else {
			fmt.Println(i)
			
		}
		
	}
}
