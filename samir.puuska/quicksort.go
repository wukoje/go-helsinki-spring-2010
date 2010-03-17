package main
import (
      "fmt"
      )


func partition(array []int,left int,right int, pivot int) int{
  var pivotValue = array[pivot]
  array[pivot],array[right] = array[right],array[pivot]
  var store = left
  for i := left; i < right;i++ {

    if array[i] <= pivotValue {
     array[i],array[store] = array[store],array[i]
     store++;
    }
  }
  array[store],array[right] = array[right],array[store]
  return store
  
  
}

func sort(s []int, left int , right int){
  if right > left {
  pi := (left+right)/2
  pin := partition(s,left,right,pi)
  sort(s,left,pin -1)
  sort(s,pin +1,right)
  }
  
  
}


func main () {
    var s = []int{666,4,-4,4,2,7,3,776,4354354,4,6}
    fmt.Println(s)
    sort(s,0,len(s)-1)
    fmt.Println(s)
}
