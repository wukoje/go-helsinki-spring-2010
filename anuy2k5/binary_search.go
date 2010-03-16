package main

import "fmt"

func main() {
	l := []int{1,2,3,4,5,6,7,8,9,10}
	a:=80
	b:=6
	c:=5
	
	// Search 80 in the array
	fmt.Printf("find %d",a)
	n := search(l,a)
	fmt.Printf("Pos: %v\n", n)
	
	// Search 6 in the array
	fmt.Printf("find %d",b)
	n = search(l, b)
	fmt.Printf("Pos: %v\n", n)
	
	// Search 5 in the array
	fmt.Printf("find %v",c)
	n = search(l, c)
	fmt.Printf("Pos: %d\n", n)

}

func search(s []int, e int) int {
	n := -1
	beg := 0
	end := len(s) - 1
	
	for beg < end {
		mid := beg + (end - beg)/2
		switch {
		case s[mid] < e:
			beg = mid + 1
		case s[mid] == e:
			beg, end, n = mid, mid, mid
		case s[mid] > e:
			end = mid - 1
		}
	}
	
	return n
}

