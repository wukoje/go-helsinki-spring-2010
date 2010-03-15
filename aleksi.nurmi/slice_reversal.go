
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
}

func main() {
	empty := []int {}
	fmt.Print(empty, "-> ")
	reverse(empty)
	fmt.Println(empty)

	one := []int { 1 }
	fmt.Print(one, "-> ");
	reverse(one)
	fmt.Println(one)

	two := []int { 1, 2 }
	fmt.Print(two, "-> ")
	reverse(two)
	fmt.Println(two)

	three := []int { 1, 2, 3 }
	fmt.Print(three, "-> ")
	reverse(three)
	fmt.Println(three)

	four := []int { 1, 2, 3, 4 }
	fmt.Print(four, "-> ")
	reverse(four)
	fmt.Println(four)
}
