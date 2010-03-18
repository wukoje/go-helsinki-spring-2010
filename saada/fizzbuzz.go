package main

import (
	"fmt"
)

func fizzbuzz_logic (i int) string {
	result := ""
	if i%3 == 0 {
		result += "Fizz"
	}
	if i%5 == 0 {
		result += "Buzz"
	}
	return result
}

func fizzbuzz_loop () {
	for i := 1; i <= 100; i++ {
		output := fizzbuzz_logic(i)
		if output == "" {
			fmt.Println(i)
		} else {
			fmt.Println(output)
		}
	}
}

func main() {
	fizzbuzz_loop()
}
