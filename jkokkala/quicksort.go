package main
import fmt "fmt"

// Sorts array partially
func quicksort(s [] int, low int, high int) {
    i:=low
    j:= high
    x := s[(low+high)/2]
    for i <= j {
        for s[i]<x { i++ }
        for s[j]>x { j-- }
        if i<=j {
            s[i],s[j]=s[j],s[i]
            i++
            j--
        }
    }
    if(low<j) { quicksort(s,low,j) }
    if(high>i) { quicksort(s,i,high) }
}
// Sorts the array
func sort(s [] int) {
    quicksort(s,0,len(s)-1)
}

func main() {
     s := []int{34,1,333,4,3,-1,-14444,99999,9999999,33,4,44,85,0}
     sort(s)
     for i:=0 ; i<len(s); i++ { fmt.Printf("%d ",s[i]) }
     fmt.Printf("\n")

}
