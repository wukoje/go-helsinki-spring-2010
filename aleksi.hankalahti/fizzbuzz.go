package main

import "fmt"

const (
	MIN_VAL = 1
	MAX_VAL = 100
)

// Just to make the software more complicated ;)
func getLine(i int) (out string) {
	if i % 3 == 0 {
		out += "Fizz"
	}
	if i % 5 == 0 {
		out += "Buzz"
	}
	if len(out) == 0 {
		out += fmt.Sprintf("%v", i)
	}
	return
}

func main() {
	for i := MIN_VAL; i <= MAX_VAL; i++ {
		fmt.Println(getLine(i))
	}
}

// So long, Pop! I'm off to check my tiger trap! – Bill Watterson
