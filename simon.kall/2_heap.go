package main
import (
	"fmt";
)
type node struct {
	data int

}
type binHeap struct {
	hp int
	heap []node 
}
func main() {
	myHeap := new(binHeap)
	myHeap.heap = make([]node,100)
	myHeap.hp = -1
	myHeap.addData(5)
	myHeap.addData(1)
	myHeap.addData(31)
	myHeap.addData(3)
		
	myHeap.String()
	myHeap.removeNode(0)
	fmt.Println("after removeNode")
	myHeap.String()
	myHeap.addData(2)
	fmt.Println("after addData")
	myHeap.String()
	
	
}

func (h *binHeap) addData(data int) {
 
	h.hp++
	if h.hp > len(h.heap){
		return
	}
	newNode := new(node)
	newNode.data = data
	child := h.hp
	parent := (child-1)/2
	h.heap[child] = *newNode
	//fmt.Println("new",newNode,"old", &h.heap[0], "newdata",h.heap[child].data,"old data", h.heap[parent].data )
	
	for newNode != &h.heap[0] && h.heap[child].data < h.heap[parent].data {
		h.heap[child],h.heap[parent] = h.heap[parent],h.heap[child] 
		child = parent
		parent = (child-1)/2
	} 
	
}

func (h *binHeap) removeNode(i int){
	
	if i>h.hp || h.heap == nil {
		return
	} else {
		h.heap[i] = h.heap[h.hp]
		h.heap[h.hp] = node{0}
		h.hp--
		left := 2*i+1
		right := 2*i+2
		
		for left <= h.hp && right <= h.hp && h.hp >= 0 {
			if h.heap[left].data < h.heap[right].data {
				h.heap[left],h.heap[i] = h.heap[i],h.heap[left]
				i=2*i+1
			} else {
				h.heap[right],h.heap[i] = h.heap[i],h.heap[right]
				i=2*i+2
			}
			left = 2*i+1
			right = 2*i+2
		}
			
	}
}
func (h *binHeap) String() {
	fmt.Printf("%v",h)
}












