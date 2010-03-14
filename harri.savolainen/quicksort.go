package main

import fmt "fmt" 

func main() {

	var s = []int {5, 76, 6, 1, 45,2,3,4,1000,5234,56,35,23423,54676,23,32,34,8}
	
	fmt.Println("Original")
	for _, v := range s {
                fmt.Printf("%d ", v)
        }
	fmt.Printf("\n")
	
	qs(s, 0, len(s) - 1)
	
	fmt.Println("Quicksorted")
        for _, v := range s {
                fmt.Printf("%d ", v)
        }
	fmt.Printf("\n")

	


}

func qs(array []int, left int, right int) {
	if right > left { 
		pivotIndex := (left + right)/2 
		pivotNewIndex := partition(array,left,right,pivotIndex)
		qs(array, left, pivotNewIndex -1)
		qs(array, pivotNewIndex + 1,right)	
	}
		
}

func partition(array []int, left int, right int, pivotIndex int) int {
     pivotValue := array[pivotIndex]
     array[pivotIndex],array[right] = array[right],array[pivotIndex]
     storeIndex := left
     for i:=left; i<right;i++ { 
         if array[i] <= pivotValue { 
             array[i],array[storeIndex]=array[storeIndex],array[i]
             storeIndex++
	}
     }
     array[storeIndex],array[right]=array[right],array[storeIndex]
     return storeIndex
}
