// Jussi Kokkala
package main
import fmt "fmt"
func search(s []int, e int) int {
    var lower_limit = 0
    var upper_limit = len(s)-1
    for  iter := 0 ; lower_limit+1<upper_limit ; iter ++ {
        var middle = (upper_limit+lower_limit)/2
        if s[middle] > e {
            upper_limit = middle
        } else {
            lower_limit = middle
        }
    }
    if s[lower_limit] == e {
        return lower_limit
    }
    if s[upper_limit] == e {
        return upper_limit
    }
    return -1
}

func test(s [] int, e int) {
    fmt.Printf("Searching for %d in (",e)
    for i:=0 ; i<len(s); i++ {
        fmt.Printf("%d ",s[i])
    }
    fmt.Printf(")\n")
    var result = search(s,e)
    fmt.Printf("Result: %d\n",result);
}   

func main() {
    test([]int{1,2,3,4,5,6,7,8,9},3)
    test([]int{1,2,3,4,5,6,7,8,9},0)
    test([]int{1,2,3,4,5,6,7,8,9},10)
    test([]int{1,2,3,4,5,6,7,8,9,10},10)
    test([]int{1,2,3,4,5,6,8,12,14,22,299},7)
    test([]int{1,2,2,2,2,2,2,2,2,2,2,2,2,2},4)
}
