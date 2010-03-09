/*
Ex1
Output the numbers from 1 to 100, one per line. 
However, if the number is divisible by 3 output 'Fizz', 
if the number is divisible by 5 output 'Buzz' 
and if the number is divisible by both 3 and 5 output 'FizzBuzz'.

Author: Xiaoyan ZHANG
*/

package main

import(
	"fmt"
	"strconv"
)

const(
	divby3 = "Fizz"
	divby5 = "Buzz"
)

func main(){
	for i:=1; i<=100; i++ {
		output := ""

		if i % 3 == 0{
			output += divby3
		}

		if i % 5 == 0{
			output += divby5
		}

		if len(output) == 0{
			output += strconv.Itoa(i)
		}

		fmt.Printf(output + "\n")
	}
}
