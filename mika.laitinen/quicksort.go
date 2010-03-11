package main

import "fmt"

func sort(s []int) {
	if len(s) <= 1 {
		return;
	}

	pind := len(s) - 1;
	store := 0;
	for i := store; i <= pind; i++ {
		if s[i] < s[pind] {
			s[i], s[store] = s[store], s[i];
			store++;
		}
	}
	s[store], s[pind] = s[pind], s[store];

	sort(s[0:store]);
	sort(s[store+1:len(s)]);
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
	a := []int{3,5,7,3,7,2,1,7,3,9,4,7,3,5,3,7,8,6,4,3};
	fmt.Printf("Unsorted: ");
	printArr(a);
	sort(a);
	printArr(a);
	a = []int{100,10,1,20,300,400,13,122,2003};
	fmt.Printf("Unsorted: ");
	printArr(a);
	sort(a);
	printArr(a);
	a = []int{100,10,1};
	fmt.Printf("Unsorted: ");
	printArr(a);
	sort(a);
	printArr(a);
	a = []int{10,100,1};
	fmt.Printf("Unsorted: ");
	printArr(a);
	sort(a);
	printArr(a);
	a = []int{1,10,100};
	fmt.Printf("Unsorted: ");
	printArr(a);
	sort(a);
	printArr(a);
}
