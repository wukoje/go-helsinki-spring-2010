package main

import fmt "fmt"

func main() {
	
	fmt.Println()
	fmt.Println(" The tree looks like this")
	fmt.Println("      3")
	fmt.Println("   A-----B")
	fmt.Println("   2\\    |")
	fmt.Println("     \\   |")
	fmt.Println("      C  | 4")
	fmt.Println("     / \\ |")
	fmt.Println("   1/  3\\|")
	fmt.Println("   E-----D")
	fmt.Println("      5")
	fmt.Println()
	
	//Creation of vertices
	A := new(vertex)
	A.name = ("A")
	A.D = 0
	A.color = ""
	
	B := new(vertex)
	B.name = ("B")
	B.D = 0
	B.color = ""

	C := new(vertex)
	C.name = ("C")
	C.D = 0
	C.color = ""
	
	D := new(vertex)
	D.name = ("D")
	D.D = 0
	D.color = ""
	
	E := new(vertex)
	E.name = ("E")
	E.D = 0
	E.color = ""
	
	//SLice with all vertices on it
	vertices := make([] *vertex, 5)
	vertices[0] = A 
	vertices[1] = B
	vertices[2] = C
	vertices[3] = D
	vertices[4] = E
	
	
	//Creation of new edges
	AB := new(edge)
	AB.v1 = *A
	AB.v2 = *B
	AB.cost = 3
	
	AC := new(edge)
	AC.v1 = *A
	AC.v2 = *C
	AC.cost = 2
	
	BD := new(edge)
	BD.v1 = *B
	BD.v2 = *D
	BD.cost = 4
	
	CE := new(edge)
	CE.v1 = *C
	CE.v2 = *E
	CE.cost = 1
	
	ED := new(edge)
	ED.v1 = *E
	ED.v2 = *D
	ED.cost = 5
	
	CD := new(edge)
	CD.v1 = *C
	CD.v2 = *D
	CD.cost = 3
	
	//Slice with all the edges on it
	edges := make([] *edge, 6)
	edges[0] = AB 
	edges[1] = AC
	edges[2] = BD
	edges[3] = CE
	edges[4] = ED
	edges[5] = CD
	
	
	//New graph, valued by the vertex and edge lists
	graphy := new(graph)
	graphy.v = vertices
	graphy.e = edges
	
	graphy.toString()
	fmt.Println()
	fmt.Println(" After MST the tree looks like this")
	fmt.Println("      3")
	fmt.Println("   A-----B")
	fmt.Println("   2\\")
	fmt.Println("     \\")
	fmt.Println("      C")
	fmt.Println("     / \\")
	fmt.Println("   1/  3\\")
	fmt.Println("   E     D")
	fmt.Println()
	
	asd := prims(graphy)
	asd.toString()
}












type edge struct {
	v1 vertex
	v2 vertex
	cost int
}

type vertex struct {
	name string
	D int
	color string
}

type graph struct {
	v []*vertex
	e []*edge
}












func prims(g *graph) graph {

	if len(g.e) == 0 {
		fmt.Println("Error, graph is empty")
		return *new(graph)
	}

	returnV := make([] *vertex, len(g.v))
	returnE := make([] *edge, len(g.e))
	returnG := new(graph)
	returnG.v = returnV
	returnG.e = returnE
	
	for i := 0; i < len(g.v); i++ {
		g.v[i].D = 999
		g.v[i].color = "WHITE"
	}
	for i := 0; i < len(g.e); i++ {
		g.e[i].v1.D = 999
		g.e[i].v2.D = 999
		g.e[i].v1.color = "WHITE"
		g.e[i].v2.color = "WHITE"
	}
	

	Q := new(heap)
	Q.array = make([] vertex, 30)
	Q.max = len(g.v)
	Q.amount = 0

	startt := g.e[0].v1
	startt.color = "GRAY"
	startt.D = 0
	Q.insert(startt)
	
	for  Q.empty() == false {
		current := Q.pop()
		
		
		for i := 0; i < len(g.v); i++ {
			neigh := *g.v[i]
			for j := 0; j < len(g.e); j++ {
			
				if (neigh.name == g.e[j].v1.name && current.name == g.e[j].v2.name) {
					if g.e[j].v1.color == "WHITE" {
					
						g.e[j].v1.color = "GRAY"
						g.e[j].v1.D = g.e[j].cost
						//g.e[j].v2.D = g.e[j].cost
						Q.insert(g.e[j].v1)
						returnG.add(g.e[j])
							
					} else if g.e[j].v1.color == "GRAY" {
						if g.e[j].v1.D > g.e[j].cost {
							g.e[j].v1.D = g.e[j].cost
							returnG.add(g.e[j])
						}
					}
					current.color = "BLACK"
				}
				if (neigh.name == g.e[j].v2.name && current.name == g.e[j].v1.name) {
					if g.e[j].v2.color == "WHITE" {
				
						g.e[j].v2.color = "GRAY"
						g.e[j].v2.D = g.e[j].cost
						//g.e[j].v1.D = g.e[j].cost
						Q.insert(g.e[j].v2)
						returnG.add(g.e[j])
					
					
					} else if g.e[j].v2.color == "GRAY" {
						if g.e[j].v2.D > g.e[j].cost {
							g.e[j].v2.D = g.e[j].cost
							returnG.add(g.e[j])
						
						}
					}
					current.color = "BLACK"
				}
			}
		}
	}
	
	returnG.dublicates()
	return *returnG
}



func (g graph) add(ed *edge) {
	for i := 0; i<len(g.e); i++ {
		if g.e[i] == nil {
			g.e[i] = ed
			return
		}
	}
}


func (g *graph) dublicates() graph{
	for i := 0; i<len(g.e); i++ {
		for j := 0; j<len(g.e); j++ {
			if g.e[i] != nil && g.e[j] != nil {
				if j==i {
					j++
				}
				if j >= len(g.e) {
					return *g
				}
				if g.e[i].v1.name == g.e[j].v1.name && 
				   g.e[i].v2.name == g.e[j].v2.name {
						g.e[i] = nil
				}
			}
		}
	}
	return *g
}







func (g *graph) toString() {
	edges := g.e
	for i := 0; i < len(edges); i++ {
		if edges[i] != nil {
			fmt.Print(edges[i].v1.name, " - ")
			fmt.Print(edges[i].cost, " - ")
			fmt.Println(edges[i].v2.name)
		}
	}

}










type heap struct {
	array []vertex
	//temp int
	max int
	amount int
}


func (h *heap) empty() bool{
	if h.amount != 0 {
		return false
	} else {
		return true
	}
	return false
}


func (h *heap) insert(e vertex) {
	if h.max > h.amount {
		h.array[h.amount] = e
		h.amount++
		
		current := h.amount - 1
		for {
			if current == 0 {
				break
			}
		
			parent := (current - 1) / 2
			
			if h.array[current].D < h.array[parent].D {
				h.array[parent], h.array[current] = h.array[current], h.array[parent]
				current = parent
			} else {
				break
			}
		}
	}
}


func (h *heap) pop() vertex {
	//fmt.Print("POP1 ", h.array[0].name, " ", h.amount, " ")
	popped := h.array[0]
	h.array[0] = h.array[h.amount-1]
	h.array[h.amount-1] = h.array[29]
	h.amount--
	//fmt.Println("POP2 ", h.array[0].name," ", h.amount, " ")

	left := 0
	right := 0
	x := 0 
	
	for {
		parent := h.array[x].D
		left = h.array[x*2+1].D
		right = h.array[x*2+2].D
		
		if left == 0 {
			left = 999
		}
		if right == 0 {
			right = 999
		}
		
		if parent <= left && parent <= right {
			break
			
		} else if parent > left || parent > right {

			if left <= right {
				h.array[x].D, h.array[x*2+1].D = left, parent
				x = x*2+1
			} else {
				h.array[x].D, h.array[x*2+2].D = right, parent
				x = x*2+2
			}	
		}
	}
	return popped	
}

func (h *heap) print() {
	fmt.Println(h.array[0].D)
	
	level := 2
	counter := 0
	
	for i := 1; i <= h.amount-1; i++ {
		fmt.Print(h.array[i].D, " ")
		
		counter++
		if level == counter {
			fmt.Println()
			level = level * 2
			counter = 0
		}
	}
	fmt.Println()
	fmt.Println()
}

