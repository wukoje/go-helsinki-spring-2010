package main

import fmt "fmt"

func main() {
	list := make([]int, 10)
	list[0] = 3
	list[1] = 5
	list[2] = 6
	list[3] = 8
	list[4] = 9
	list[5] = 11
	list[6] = 14
	list[7] = 16
	list[8] = 18
	list[9] = 20

	fmt.Println(list)
	fmt.Println("Searching 4, found at index: ", binary_search(list, 4))
	fmt.Println("Searching 18, found at index: ", binary_search(list, 18))


        list = make([]int, 10)
        list[0] = 7
        list[1] = 16
        list[2] = 18
        list[3] = 25
        list[4] = 29
        list[5] = 31
        list[6] = 45
        list[7] = 67
        list[8] = 89
        list[9] = 93

        fmt.Println(list)
        fmt.Println("Searching 4, found at index: ", binary_search(list, 4))
        fmt.Println("Searching 18, found at index: ", binary_search(list, 18))
        fmt.Println("Searching 30, found at index: ", binary_search(list, 30))
        fmt.Println("Searching 45, found at index: ", binary_search(list, 45))
        fmt.Println("Searching 88, found at index: ", binary_search(list, 88))
        fmt.Println("Searching 93, found at index: ", binary_search(list, 93))


}


func binary_search(list []int, e int) int {
	min := 0
	max := len(list) - 1

	for min <= max {
		mid := min + ((max - min) / 2)
		switch {

			case e > list[mid]:
				min = mid + 1

			case e < list[mid]:
				max = mid - 1

			case e == list[mid]:
				return mid
		} 
	} 
	return -1
}
