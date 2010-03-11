package main

import "fmt"

func search(s []int, e int) int {
	low := 0;
	high := len(s);

	for low <= high {
		ind := (low + high) / 2;
		if s[ind] == e {
			return ind;
		} else if(s[ind] < e) {
			low = ind + 1;
		} else {
			high = ind - 1;
		}
	}
	return -1;
}

func main() {
	a := []int{1,2,3,4,5};
	fmt.Printf("%d\n",search(a,3));
	a = []int{1,2,3,4,5};
	fmt.Printf("%d\n",search(a,1));
	a = []int{1,2,3,4,5};
	fmt.Printf("%d\n",search(a,5));
	a = []int{1,2,3,4,5,6,7,8};
	fmt.Printf("%d\n",search(a,7));
	a = []int{3,8,12,15,19,20,27,42,53,62};
	fmt.Printf("%d\n",search(a,7));
	a = []int{3,8,12,15,19,20,27,42,53,62};
	fmt.Printf("%d\n",search(a,27));
	a = []int{3,8,12,15,19,20,27,42,53,62};
	fmt.Printf("%d\n",search(a,13));
	a = []int{3,8,12,15,19,20,27,42,53,62};
	fmt.Printf("%d\n",search(a,3));
	a = []int{3,8,12,15,19,20,27,42,53,62};
	fmt.Printf("%d\n",search(a,62));
}
