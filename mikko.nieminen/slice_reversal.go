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

	fmt.Printf ("order before slice reversal:\n%v\n", slice)

	reverse (slice)

	fmt.Printf ("\norder after slice reversal:\n%v\n", slice)

}

func reverse(s []int) {
	last_index := len(s)-1
	loop_length := len(s)/2
	for i := 0; i<loop_length; i++ {
		s[i], s[last_index-i] = s[last_index-i], s[i] 
	}
}
