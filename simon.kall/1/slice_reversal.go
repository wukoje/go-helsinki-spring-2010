// reverse reverses the contents of s in place
//Output some slices before and after reversal.

package main

import (
	"fmt"; 
)
func main() {
	o := make([]int,10)
	o[0] = 1
	print(o)
	n := reverse(o)
	print(n)
} 
func print(t []int) {
	fmt.Println()
	for i:=0; i<10;i++ {
		fmt.Println(t[i])
	}
}
func reverse(t []int) []int {
	lenght := len(t) -1
	for i:=0; i<len(t)/2 ;i++ {
		first := t[i]
		t[i] = t[lenght-i]
		t[lenght-i] = first
	}
	return t
}
