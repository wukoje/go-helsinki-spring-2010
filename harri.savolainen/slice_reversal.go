package main

import fmt "fmt"

func main() {
        s := []int{1, 2, 3, 4, 5, 6, 7, 8}
        fmt.Println("Original:")
        for _, v := range s {
                fmt.Printf("%d", v)
        }
        fmt.Printf("\n")

        fmt.Println("Reversed:")
        for _, v := range reverse(s) {
                fmt.Printf("%d", v)
        }
        fmt.Printf("\n")
}

func reverse(s []int) []int {

        for i := 0; i < len(s)/2; i++ {

                a := s[i]
                b := s[len(s)-i-1]

                s[i] = b
                s[len(s)-i-1] = a

        }

        return s

}