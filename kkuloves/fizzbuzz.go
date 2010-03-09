/*
 * fizzbuzz.go: For each integer from 1 to 100, output one line that
 * contains either the decimal representation of the integer, or
 * "Fizz" if the integer is divisible by 3, or "Buzz" if the integer is
 * divisible by 5, or "FizzBuzz" if the integer is divisible by both
 * 3 and 5.
 *
 * For Introduction to Go, Spring 2010
 * Kimmo Kulovesi <kkuloves@cs.helsinki.fi>
 */
package main

import (
	"fmt"
)

const linesToPrint = 100

func main() {
	for i := 1; i <= linesToPrint; i++ {
		// It would probably make more sense to just do this
		// with two ifs, but that would not give me the
		// opportunity to play with Go's switch.
		switch 0 {
		default:
			// Not divisible by 3 or 5
			fmt.Println(i)
		case i%3 + i%5:
			// Divisible by 3 and 5. The above case is (here)
			// equivalent to "i % 15" but this is more readable
			// (although it requires non-negative values) and
			// one might _hope_ that the compiler could reuse the
			// results of dividing i by 3 and 5 for the following
			// cases.
			fmt.Println("FizzBuzz")
		case i % 3:
			// Divisible by 3
			fmt.Println("Fizz")
			// Apparently fallthrough can not be made conditional;
			// I tried to do "if (i % 5 == 0) { fallthrough }" here,
			// which, along with small changes, would have eliminated
			// the separate FizzBuzz case (although that's possibly
			// more readable).
		case i % 5:
			// Divisible by 5
			fmt.Println("Buzz")
		}
	}
}
