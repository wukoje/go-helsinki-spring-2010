package main

import "fmt"

// returns true if n divides k, otherwise false
func d(n int,k int) bool {
	return k % n == 0
}

// If i is divisible by 3, returns "Fizz",
// if i is diviseble by 5, returns "Buzz" and
// if i is divisible by both 3 and 5, returns "FizzBuzz".
// Otherwise returns the string representation of i.
func representation(i int) string {
	s := ""
	if d(3,i) {
		s += "Fizz"
	}
	if d(5,i) {
		s += "Buzz"
	}
	if s == "" {
		s += fmt.Sprint(i)
	}
	return s
}

// Prints the FizzBuzz-representation of numbers 1...100
func main() {
    for i := 1; i <= 100; i++ {
        fmt.Println(representation(i))
    }
}
