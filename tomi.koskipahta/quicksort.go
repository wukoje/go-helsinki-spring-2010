package main

import fmt "fmt"

func main() {
	list := make([]int, 10)
	
	
	list[0] = 12
	list[1] = 57
	list[2] = 62
	list[3] = 23
	list[4] = 38
	list[5] = 98
	list[6] = 54
	list[7] = 21
	list[8] = 11
	list[9] = 44
	
	fmt.Println(list)
	sort(list)
	fmt.Println(list)
	
	list[0] = 76
	list[1] = 2
	list[2] = 24
	list[3] = 97
	list[4] = 87
	list[5] = 42
	list[6] = 32
	list[7] = 9
	list[8] = 62
	list[9] = 55
	
	fmt.Println(list)
	sort(list)
	fmt.Println(list)
}



func sort(s []int) {
	temp := 0
	pivot := len(s)-1
		
	if len(s) <= 1 {
		return
	}
	
	for i := 0; i <= pivot; i++ {
		if s[i] > s[pivot] &&  s[i] != s[pivot-1] {
			temp = s[i]
			s[i] = s[pivot-1]
			s[pivot-1] = s[pivot]
			s[pivot] = temp
			pivot--
			i--
				
		} else if s[i] > s[pivot] &&  s[i] == s[pivot-1] {
			s[i], s[pivot] = s[pivot], s[i]
			pivot--
			i--
		}
	}

	temp1 := s[0:pivot]
	temp2 := s[pivot:]
				
	sort(temp1)
	sort(temp2)
	
	list := make( []int, len(temp1) + len(temp2))
	copy(list, temp1)
	copy(list[len(temp1):], temp2)
	s = list
}
