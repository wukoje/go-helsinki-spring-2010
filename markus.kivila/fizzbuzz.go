/*
1 fizzbuzz.go

Output the numbers from 1 to 100, one per line. However, if the number is 
divisible by 3 output 'Fizz', if the number is divisible by 5 output 'Buzz'
and if the number is divisible by both 3 and 5 output 'FizzBuzz'

*/

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i % 3 == 0 {
			fmt.Printf("Fizz\n")
		} else if i % 5 == 0 {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
