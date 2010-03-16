package main

// Output the numbers from 1 to 100, one per line. However, if the number is divisible by 3 output 'Fizz', if the number is divisible by 5 output 'Buzz' and if the number is divisible by both 3 and 5 output 'FizzBuzz'. In case you output any combination of Fizz/Buzz don't output the number.

import (
    "fmt"
)

func main() {
    for number := 1; number <= 100; number++ {
        buzzed := false
        if (number % 3) == 0 {
            fmt.Printf("Fizz")
            buzzed = true
        }
        if (number % 5) == 0 {
            fmt.Printf("Buzz")
            buzzed = true
        }
        if !buzzed {
            fmt.Printf("%d",number)
        }
        fmt.Printf("\n")
    }
}
