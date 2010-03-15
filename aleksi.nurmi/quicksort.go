
package main

import "fmt"

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func sort(s []int) {
	if len(s) > 1 {
		i := 0
		j := len(s) - 1
		pivot := s[j / 2]
		for i < j {
			for s[i] < pivot {
				i++
			}
			for pivot < s[j] {
				j--
			}
			if i <= j {
				swap(&s[i], &s[j])
				i++
				j--
			}
		}
		if 0 < j {
			sort(s[0:j+1])
		}
		if i < len(s)-1 {
			sort(s[i:])
		}
	}
}

func main() {
	t := []int {0, 3, 6, 1, 8, 9, 0, 1, 3, 5}
	fmt.Println(t)
	sort(t)
	fmt.Println(t)
	
	t = []int {3, 1, 1, 7, 6 }
	fmt.Println(t)
	sort(t)
	fmt.Println(t)

	t = []int {}
	fmt.Println(t)
	sort(t)
	fmt.Println(t)

	t = []int {3}
	fmt.Println(t)
	sort(t)
	fmt.Println(t)
	
	t = []int {3, 1}
	fmt.Println(t)
	sort(t)
	fmt.Println(t)

	t = []int {3, 1, 1}
	fmt.Println(t)
	sort(t)
	fmt.Println(t)
}
