// sort modifies the slice s so that the integers are sorted in
// place using quicksort
package main

import (
	"fmt";
	"rand"; 
)

func main() {
	size :=10
	list := make([]int,size)
	orig := make([]int,size)
	for i:=0; i<size; i++ {
		list[i] = rand.Int()
		orig[i] = list[i]
	//	fmt.Println(list[i])
	}
	sort(list,0,size-1)
	fmt.Println("Sorted \t \t original")
	for i:=0; i<size; i++ {
		fmt.Println(list[i],"\t", orig[i])
	}
}

func sort(A []int,l int, r int) {
	if l<r {
		pIndex := (l+r)/2
		newIndex := partition(A,l,r,pIndex)
		sort(A,l,newIndex-1)
		sort(A,newIndex+1,r)
	}
}
func partition(A []int,l int, r int,index int) int{
	pValue := A[index]
	temp := A[index]
	A[index] = A[r]
	A[r] = temp
	sIndex := l
	for i:=l; i<r; i++ {
		if A[i] < pValue {
			temp = A[i]
			A[i] = A[sIndex]
			A[sIndex] = temp
			sIndex++
		}
	}
	temp = A[sIndex] 
	A[sIndex] = A[r]
	A[r] = temp
	return sIndex
}
