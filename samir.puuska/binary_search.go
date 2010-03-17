package main
import (
      "fmt"
      )

func bs(s []int, haku int) int {
 var min = 1
 var max = cap(s)
 var mid = 0
 for true {
    mid = (min + max)/2
      if haku > s[mid]{
      min = mid +1
      }else{
      max = mid -1
      }
    if s[mid] == haku {
      return mid
    } else if min > max {
      return -1
    }
  }



 return mid 
}


func main () {
    var s = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
    var lol = bs(s, 15)
    fmt.Println("The index is: ",lol)
}
