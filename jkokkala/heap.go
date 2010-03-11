package main
import fmt "fmt"

type heap []int
func create(s [] int) *heap {
    tempslice:=heap(s[0:0])
    return &tempslice
}
// Inserts a new integer on the heap
func ( p *heap) insert(e int) {
	*p = (*p)[0:len(*p)+1]
	s := *p
	var i=len(s)-1
    for( i>0 && s[(i-1)/2]  < e )  {
    	s[i] = s[(i-1)/2]
    	i = (i-1)/2
    }
    s[i]=e
}
// Reorganizes the heap
func (p *heap) heapify(e int) {
    s:=*p
    x:=s[e]
    for l(e) < len(s) {
        t := l(e)
        if(r(e) < len(s)) {
            if(s[l(e)] < s[r(e)]) {
                t = r(e)
            }
        }
        if(s[t] > x) {
            s[e] = s[t]
            e=t
        } else {
        	s[e] = x
        	return
        }
    }
}
// Pops largest element from the heap
func (h *heap) pop() int {
	
    ret := (*h)[0]
	temp := (*h)[len(*h)-1]
	*h = (*h)[0:len(*h)-1]
    if(len(*h) == 0 ) { return ret }
	s := *h
	s[0] = temp
	if(len(s)>1) {
		s.heapify(0)
    }
	return ret

}

func l(e int) int { return e*2+1 }
func r(e int) int { return e*2+2 }

func test_insert(h* heap, e int) {
    fmt.Printf("Inserting %d\n",e)
    (*h).insert(e)
}
func test_pop(h* heap) {
    fmt.Printf("Popping %d\n",(*h).pop())
}
func test() {
    theap := create(make([]int,100))
    test_insert(theap,3)
    test_pop(theap)
    test_insert(theap,4)
    test_insert(theap,9)
    test_insert(theap,7)
    for( len(*theap)>0 ) {
    	test_pop(theap)
    }
}

func main() {
	test()
}
