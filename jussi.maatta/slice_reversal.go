package main

import "fmt"

func reverse(s []int) {
	ls := len(s)
	for i := 0; i < ls/2; i++ {
		s[i], s[ls-i-1] = s[ls-i-1], s[i]
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}

	fmt.Println("The original slice:")
	fmt.Println(a)

	fmt.Println("Let's reverse it:")
	reverse(&a)
	fmt.Println(a)

	fmt.Println("And after reverse(a[1:4]):")
	reverse(a[1:4])
	fmt.Println(a)
}
