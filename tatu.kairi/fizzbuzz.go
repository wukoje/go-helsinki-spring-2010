/*
1 fizzbuzz.go

Output the numbers from 1 to 100, one per line. However, if the number is divisible by 3 output 'Fizz', if the number is divisible output 'Buzz' and if the number is divisible by both 3 and 5 output 'FizzBuzz'

Author: tatu.kairi@gmail.com
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {

		result := ""

		if i % 3 == 0 {
			result += "Fizz"
		}

		if i % 5 == 0 {
			result += "Buzz"
		}

		if result == "" {
			result += strconv.Itoa(i)
		}

		fmt.Printf("%s\n", result)

	}
}    

