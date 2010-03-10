package main

import "fmt"

// pivot with last element (save one swap!)
// return where it went
func partition(s []int) int {
	last := len(s)-1
	v := s[last]
	
	place := 0
	for i:= 0; i < len(s); i++ {
		if s[i] < v {
			s[i], s[place] = s[place], s[i]
			place++;
		}
	}
	s[place], s[last] = s[last], s[place]
	return place
}

// sort modifies the slice s so that the integers are sorted in
// place using quicksort
func sort(s []int) {
	if len(s) <= 1 {
		return
	}	
	split := partition(s)
	sort(s[0:split])
	sort(s[split+1:])
}

func main() {
	s := []int{9,4,11,44,3}
	fmt.Println(s)
	sort(s)
	fmt.Println(s)
}
