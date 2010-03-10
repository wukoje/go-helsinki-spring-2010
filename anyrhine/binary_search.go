package main

import "fmt"

func search(s []int, e int) int {
	i := 0
	j := len(s)
	for i < j {
		m := int(uint(i + j) / 2)
		if s[m] < e {
			i = m + 1
		} else if s[m] > e {
			j = m
		} else {
			return m
		}
	}
	return -1

}

func main() {
	tab := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}

	fmt.Printf("array: ")
	for i := 0; i < len(tab); i++ {
		tab[i] *= 2
		fmt.Printf("%d ", tab[i])
	}
	fmt.Printf("\n")

	for i := 1; i < 2*(len(tab)+1); i++ {
		t := search(tab, i)
		if t != -1 {
			fmt.Printf("%d found at %d (%d)\n", i, t, tab[t])
		} else {
			fmt.Printf("%d not found\n", i)
		}
	}
}
