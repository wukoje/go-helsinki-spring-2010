package main

import fmt "fmt" 

func main() { 
	s := []int{10,20,30,36,40,45,65,77,87}
	e := 77
	fmt.Printf("Searching for %d, elements are: ",e);
	for i, v := range s { fmt.Printf("%d[%d] ",i,v) }
 	fmt.Printf("\n")
	
	fmt.Printf("Search result, element in slice is: %d\n",search(s, e))
	
} 

func search(s []int, e int) int {

	return bs(s,e,0,len(s))
}

func bs(s []int, value int, low int, high int) int {

	fmt.Printf("Examining slice: ")	
	for _, v := range s[low:high] { fmt.Printf("[%d] ",v) }	
	fmt.Printf("\n")
	
	if high < low {
		return -1 // not found
		
	} 
	
	mid := low + ((high - low) / 2) 
	
		if s[mid] > value {
			return bs(s, value, low, mid-1)
		} else { 
			if s[mid] < value {
				return bs(s, value, mid+1, high)	
			} else {
				return mid // found
			}
		}
	
	panic("unreachable");
   }

