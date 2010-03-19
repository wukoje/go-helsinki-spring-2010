
package main
import (
      "fmt"
      
)

func reverse(s []int) []int{
  for i :=0 ; i < (len(s))/2 ; i++{
   s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
 }
 return s
}


func main () {
  fmt.Print("Initialized\n")
    var s = []int{1,2,3,4,5,6,7,8,9,10,11,12,455,6666}
    fmt.Println(s)
    s = reverse(s)
    fmt.Println(s)
}
