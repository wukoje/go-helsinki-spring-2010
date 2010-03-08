package main

import (
	"fmt"
	"flag"
	"strconv"
)

func main() {
	const SIZE int = 9 // default array size
	size := SIZE

	flag.Parse() // try to get array size from command line
	if flag.NArg()>0 {
		tmp_size, err := strconv.Atoi(flag.Arg(0))
		if err!=nil {
			fmt.Printf ("cannot parse command line parameter, using default array size (%d)\n", SIZE)
			size = SIZE
		} else {
			size = tmp_size
		}
	}

	if size<1 {
		fmt.Printf ("Array size (%d) too small, using default size (%d) instead\n", size, SIZE)
		size = SIZE
	}

	slice := make([]int, size) // create a new slice and initialize it
	for i, _ := range slice {
		slice[i]=i;
	}

	fmt.Printf ("order before slice reversal:\n")
	print_slice(slice)

	reverse (slice)

	fmt.Printf ("\norder after slice reversal:\n")
	print_slice(slice)


}

func reverse(s []int) {
	last_index := len(s)-1
	loop_length := len(s)/2
	tmp := 0
	for i := 0; i<loop_length; i++ {
		tmp = s[i];
		s[i] = s[last_index-i]
		s[last_index-i] = tmp
	}
}

func print_slice(slice []int) {
	for _, c := range slice {
		fmt.Printf ("%v\n", c)
	}	
}
