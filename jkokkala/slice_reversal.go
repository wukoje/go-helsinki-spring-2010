package main
import fmt "fmt"
func reverse(s []int) {
	var j = len(s)
    for i := 0 ; i< len(s)/2 ; i++ {
        j--
        var temp = s[j]
        s[j] =s[i]
        s[i] = temp


    }
    

}

func test(s [] int) {
    
    for i:=0 ; i<len(s); i++ {
    	fmt.Printf("%d ",s[i])
    }
    fmt.Printf(" -> ")
    reverse(s)
    for i:=0 ; i<len(s); i++ {
    	fmt.Printf("%d ",s[i])
    }
    fmt.Printf("\n")
}

func main() {
    test([]int{1,2,3,4,5,6,7,8,9})
    test([]int{1,2,3,4,5,6,7,8,9,10})
    test([]int{1,2,3,4,5,6,5,4,3,2,1})
    test([]int{1,2,1,2,1,2,1,2,1,2,1,2,1,2})
    
}
