//func search(s []int, e int) int
//enerate some test data and output a few slices and search results.

package main 

import fmt "fmt"


// search returns the index of the element e in the
// sorted slice s. If e does not exist in s search returns -1
func search(s []int, e int) int {
	
	if s == nil {
		return -1
	}
	
	right := len(s) - 1
	left := 0 
	middle := (left+right) / 2
	for right >= left {
		switch {
		
			case s[middle] < e:
				left = middle + 1
				middle = (left+right) / 2
			case s[middle] > e:
				right = middle - 1
				middle = (left+right) / 2				
			case s[middle] == e:
				return middle
			default:
				return -1 
		}	
	}		
	return -1;
}

func main() {

	table := [10]int {1,2,3,4,5,6,7,8,9,10}
	table2 := [20]int{12,13,23,24,25,26,27,28,30,40,43,50,60,65,75,77,81,91,93,100}


	for i,_ := range table {
		fmt.Printf("Search %d from table 1, result index %d\n", i+2,search(&table, i+2))		
	}

	fmt.Printf("Done with table1")
	for i,_ := range table2 {
		fmt.Printf("Search %d from table 2, result index %d\n", i*7,search(&table2, i*7))		
	}
}
