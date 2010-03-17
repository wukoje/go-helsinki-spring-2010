package main

import fmt "fmt"

func main() {
	list := make([]int, 10)
	list[0] = 0
	list[1] = 1
	list[2] = 2
	list[3] = 3
	list[4] = 4
	list[5] = 5
	list[6] = 6
	list[7] = 7
	list[8] = 8
	list[9] = 9

	fmt.Println(list)
	list = reverse(list)
	fmt.Println(list)

    list = make([]int, 10)
	list[0] = 30
	list[1] = 31
	list[2] = 32
	list[3] = 33
	list[4] = 34
	list[5] = 35
	list[6] = 36
	list[7] = 37
	list[8] = 38
	list[9] = 39

    fmt.Println(list)
	list = reverse(list)
	fmt.Println(list)

        


}


func reverse(s []int) []int {
	helper := 0
	for i := 0; i < ( len(s) / 2); i++ {
		helper = s[i]	
		s[i] = s[ (len(s)-1) - i ]
		s[(len(s)-1)-i] = helper 
	}
	return s
}
