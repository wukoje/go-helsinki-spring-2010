

package main

import fmt "fmt" 

func main() {
	a := make([]int,100)
	for i := 0; i<100; i++ {
		a[i]=i
	}
	end := search(a,11)
	fmt.Println(end)
}

func search(a []int, x int) int {
	left := 0
	right := len(a)-1
	
	for left < right {
	//	fmt.Println(left,right)
		
		mid := (left+right)/2
		if a[mid] < x {
			left = mid + 1 
		} else if a[mid] > x {
			right = mid - 1
		} else {
			right = mid
			left = mid
			return 1
		}
	}
	return 0
}
