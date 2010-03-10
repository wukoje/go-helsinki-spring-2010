package main

import "fmt"
import "rand"

func sort(tab []int) {
	if len(tab) <= 1 {
		return
	}

	// todo: be smarter about piv selection
	piv := tab[0]
	tab[0], tab[len(tab)-1] = tab[len(tab)-1], tab[0]

	i := 0
	j := len(tab) - 2
	for {
		for i < j && tab[i] <= piv {
			i++
		}
		for i < j && tab[j] > piv {
			j--
		}
		if i < j {
			tab[i], tab[j] = tab[j], tab[i]
		} else {
			break
		}
	}
	tab[len(tab)-1] = tab[i]
	tab[i] = piv

	sort(tab[0:i])
	sort(tab[i+1 : len(tab)])
}

func main() {
	tab := make([]int, 2000)
	for i, _ := range tab {
		tab[i] = rand.Int() % 1000
	}
	sort(tab)
	for _, n := range tab {
		fmt.Printf("%d ", n)
	}
	fmt.Printf("\n")
}
