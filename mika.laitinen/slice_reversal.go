package main

import "fmt"

func reverse(s []int) {
	for i := 0; i < len(s) / 2; i++ {
		tmp := s[i];
		s[i] = s[len(s) - 1 - i];
		s[len(s) - 1 - i] = tmp;
	}
}

func printArr(a []int) {
	for i := 0; i < len(a); i++ {
		if i == len(a) - 1 {
			fmt.Printf("%d\n",a[i]);
		} else {
			fmt.Printf("%d ",a[i]);
		}
	}
}

func main() {
	a := []int{1,2,3,4,5,6,7,8};
	printArr(a);
	reverse(a);
	printArr(a);
	reverse(a[3:6]);
	printArr(a);
	b := []int{10,2,3,4,7,4,5,3};
	printArr(b);
	reverse(b);
	printArr(b);
	reverse(b[0:5]);
	printArr(b);
}
