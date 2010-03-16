/**

1 fizzbuzz.go

Output the numbers from 1 to 100, one per line. However, if the number is divisible by 3 output 'Fizz', if the number is ivisible by 5 output 'Buzz' and if the number is divisible by both 3 and 5 output 'FizzBuzz'

*/

package main

import (
       "fmt"
)

func main() {
       for i := 1; i < 100; i++ {
               switch {
               case i%3 == 0 && i%5 == 0:
                       fmt.Printf("FizzBuzz\n")
               case i%3 == 0:
                       fmt.Printf("Fizz\n")
               case i%5 == 0:
                       fmt.Printf("Buzz\n")
               default:
                       fmt.Printf("%d\n", i)
               }
       }
}
