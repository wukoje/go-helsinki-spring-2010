package main

import fmt "fmt"

func reverse(s []int) {	
	var temp int = 0
	for i, j := 0, len(s)-1; i != j; j,i = j-1, i+1 {
		temp = s[i]
		s[i] = s[j]
		s[j] = temp
    }

	fmt.Printf("Reversed slice...\n")
   	for _,i := range s {
 		fmt.Printf("%d\n",i) // print switched   
	}
}

func main() {

	s := []int{1,2,3}
 	
	fmt.Printf("Actual slice\n")  	
	for _,i := range s {
        fmt.Printf("%d\n",i)
    }
    
	reverse(s)
    fmt.Printf("done!\n")
}

