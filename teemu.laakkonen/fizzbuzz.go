package main

import fmt "fmt"

func main() {

     var i int  =1

     for i<101 {

	switch {
     	 case i%3==0 && i%5==0 :
	    fmt.Printf("FizzBuzz\n")
	 case i%3==0 :
	   fmt.Printf("Fizz\n")
	 case i%5==0 :
	   fmt.Printf("Buzz\n")
	 default:
		fmt.Println(i)
	 }

	i++

     }

}
