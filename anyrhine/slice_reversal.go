package main

import "fmt"

func reverse(s []int) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		j--
		i++
	}
}

func main() {
	tab := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(tab); i++ {
		fmt.Printf("%d ", tab[i])
	}
	fmt.Printf("\n")
	reverse(tab)
	for i := 0; i < len(tab); i++ {
		fmt.Printf("%d ", tab[i])
	}
	fmt.Printf("\n")
}
