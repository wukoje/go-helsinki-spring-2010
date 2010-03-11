package main

import "fmt"

type heap struct {
	H []int
	V int
}// your type here

// insert adds element e to the heap
func (h *heap) insert(e int) {
	h.H[h.V] = e;
	h.bubble_up(h.V);
	h.V++;
}

func (h *heap) bubble_down(e int) {
	next := 2*e + 1;
	if h.V <= next {
		return;
	}
	if h.V == next + 1 {
		if(h.H[e] < h.H[next]) {
			h.H[e], h.H[next] = h.H[next], h.H[e];
		}
		return;
	}
	if h.H[next] > h.H[next+1] {
		if(h.H[e] < h.H[next]) {
			h.H[e], h.H[next] = h.H[next], h.H[e];
			h.bubble_down(next);
		}
	} else {
		if(h.H[e] < h.H[next+1]) {
			h.H[e], h.H[next+1] = h.H[next+1], h.H[e];
			h.bubble_down(next + 1);
		}
	}
}

func (h *heap) bubble_up(e int) {
	next := (e-1)/2;
	if next >= 0 && h.H[e] > h.H[next] {
		h.H[e], h.H[next] = h.H[next], h.H[e];
		h.bubble_up(next);
	}
}

// pop returns the top element from the heap
func (h *heap) pop() int {
	r := h.H[0];
	h.V--;
	h.H[0] = h.H[h.V];
	if h.V > 1 {
		h.bubble_down(0);
	}
	return r;
}

func (h *heap) pr() {
	// prints the heap as a "tree", so that one line represents
	// a depth in the tree; last line may be incomplete (of course)
	ended := false;
	limit := 1;
	ll := 0;
	ind := 0;
	for !ended {
		ll += limit;
		for ; ind < ll; ind++ {
			if ind == h.V {
				ended = true;
				break;
			}
			fmt.Printf("%d ",h.H[ind]);
		}
		fmt.Printf("\n");
		limit *= 2;
	}
}

func (h *heap) empty() bool {
	return h.V == 0;
}

// create takes the given slice and creates a heap type that uses the slice for storage
func create(s []int) *heap {
	var h heap
	h.V = 0;
	h.H = s;
	return &h;
}

func main() {
	var s[30]int
	a := []int{2,3,5,2,6,2,3,5,7,8,4,8,3,8,3,5,2,9,0,4,5,-1};
	h := create(&s);
	for i := 0; i < len(a); i++ {
		h.insert(a[i]);
	}
	h.pr();
	for !h.empty() {
		fmt.Printf("%d\n",h.pop());
	}

	h = create(&s);
	h.insert(3);
	h.insert(4);
	fmt.Printf("%d\n",h.pop());
	fmt.Printf("%d\n",h.pop());
	h.insert(2);
	fmt.Printf("%d\n",h.pop());
	h.insert(100);
	h.insert(10);
	h.insert(10000);
	h.insert(1000);
	h.insert(1000000);
	h.insert(100000);
	fmt.Printf("%d\n",h.pop());
	h.pr();
}
