//Quicksort

package main

import "fmt"

func main(){
	s := []int {37, 2, 51, 7, 44, 8, 9, 34, 65, 1, 98, 15, 22, 11, 46}

	fmt.Println(s)

	sort(s, 0, len(s)-1)
	fmt.Println(s)
	
}

func sort(s []int, vasen int, oikea int){

	fmt.Println(s, vasen, oikea)

	if vasen < oikea{
	
		i := partition(s, vasen, oikea)
		sort(s, vasen, i-1)
		sort(s, i+1, oikea)
	}	
}

func partition(s []int, vasen int, oikea int) int{

	i := vasen
	k := oikea

	pivot := s[vasen]

	for k>i {

		for s[i] <= pivot && i<= oikea && k>i{
			i++
		}
		for s[k]>pivot && k>=vasen && k>=i{
			k--
		}

		if k>i {
			s[i], s[k] = s[k], s[i]
		}


	}

	s[vasen], s[k] = s[k], s[vasen]

	return k

}


















