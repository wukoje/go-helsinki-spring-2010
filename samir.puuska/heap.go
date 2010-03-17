
package main
import (
    "fmt"
    //"strings"
 //   "container/list"
)

type heap struct {

  back []int
}

func new(s int) *heap {
 z := make([]int,s)
 z = z[0:0]
 return &heap{z}
  
}

func (h *heap) insert(i int){

  
  free := len(h.back)
  h.back = h.back[0 :free+1]
  h.back[free] = i
  h.upHeap(free)
}


func (h *heap) upHeap(u int){
  //fmt.Println(h.back[u])
  if u ==0{
      return
  }
  
  parent := (u-1)/2
 if h.back[u] >= h.back[parent]{
  return 
 }
  h.back[u],h.back[parent] = h.back[parent],h.back[u]
  h.upHeap(parent)
 
  
}
func (h *heap) pop() int {
  pop := h.back[0]
  h.back[0] = h.back[len(h.back)-1]
  h.back = h.back[0:len(h.back)-1]
 return pop
}
func (h *heap) downHeap(d int){
  left := 2*d
  right := 2*d+1
  heaplen := len(h.back)
  smallest :=0
  if (left >= heaplen) && (h.back[left] > h.back[d]){
    smallest = left
  }else {
   smallest = d 
  }
  
  if  (right >= heaplen) && (h.back[right] > h.back[smallest]){
  smallest = right
  }
  if smallest < d {
   h.back[d],h.back[smallest] = h.back[smallest],h.back[d]
   h.downHeap(smallest)
  }
  
}
func main () {
  fmt.Print("Initialized\n")
  s := new(40)

  s.insert(7)
  s.insert(21)
  s.insert(14)
  fmt.Println(s.pop())
  fmt.Println(s.pop())
  fmt.Println(s.pop())

}
