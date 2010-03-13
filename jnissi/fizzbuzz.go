package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		var out string = ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out != "" {
			fmt.Printf(out + "\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
