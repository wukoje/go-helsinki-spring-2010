/*
 * In hindsight, the JSON-based format was probably overkill for this
 * exercise, but I wanted to try it out. See graph.json for an example file.
 */


package main


import (
	"fmt"
	"io/ioutil"
	"json"
	"os"
	"reflect"
)


func append(s []int, n int) []int {
	if len(s) == cap(s) {
		if cap(s) == 0 {
			s = make([]int, 0, 4)
		} else {
			new_s := make([]int, len(s), cap(s)*2)
			for i, e := range s {
				new_s[i] = e
			}
			s = new_s
		}
	}
	s = s[0 : len(s)+1]
	s[len(s)-1] = n
	return s
}


type stack struct {
	elems []int
}


func create_stack() *stack {
	st := new(stack)
	st.elems = make([]int, 4)
	return st
}


func (st *stack) push(n int) { st.elems = append(st.elems, n) }


func (st *stack) pop() int {
	n := st.elems[len(st.elems)-1]
	st.elems = st.elems[0 : len(st.elems)-1]
	return n
}


type vertex struct {
	index      int
	lowlink    int
	in_stack   bool
	neighbours []int
}


type graph struct {
	vertices []*vertex
}


type edge [2]int


func read_edges(data interface{}) ([]edge, string) {
	var edges []edge
	switch t := reflect.NewValue(data).(type) {
	case *reflect.SliceValue:
		edges = make([]edge, t.Len())
		for i := 0; i < t.Len(); i++ {
			elem := t.Elem(i).(*reflect.InterfaceValue).Elem()
			if reflect.Typeof(elem).String() != "*reflect.SliceValue" {
				return nil, fmt.Sprintf("Edge is not a list")
			}
			edge_pair := elem.(*reflect.SliceValue)
			if edge_pair.Len() != 2 {
				return nil, fmt.Sprintf("Edge is not a pair")
			}
			var e edge
			for k := 0; k < 2; k++ {
				vertex := edge_pair.Elem(k).(*reflect.InterfaceValue).Elem()
				if reflect.Typeof(vertex).String() != "*reflect.Float64Value" {
					return nil, fmt.Sprintf("Edge contains non-numbers")
				}
				e[k] = int(vertex.(*reflect.Float64Value).Get())
			}
			edges[i] = e
		}
	default:
		return nil, fmt.Sprintf("The edges field is not a list")
	}
	return edges, "foo"
}


func create_graph(str string) (*graph, string) {
	obj, err := json.Decode(str)
	if err != nil {
		return nil, fmt.Sprintf("%v", err)
	}
	gr := new(graph)
	var edges []edge
	switch t := reflect.NewValue(obj).(type) {
	case *reflect.MapValue:
		for _, keyi := range t.Keys() {
			key := keyi.Interface().(string)
			switch key {
			case "edges":
				var err_str string
				edges, err_str = read_edges(t.Elem(keyi).Interface())
				if edges == nil {
					return nil, err_str
				}
			case "vertex_count":
				vertex_count := int(t.Elem(keyi).Interface().(float64))
				gr.vertices = make([]*vertex, vertex_count)
				for i, _ := range gr.vertices {
					gr.vertices[i] = new(vertex)
				}
			default:
				return nil, fmt.Sprintf("Unknown key: %s", key)
			}
		}
	default:
		return nil, fmt.Sprintf("The input doesn't contain" +
			" a dictionary")
	}
	if gr.vertices == nil {
		return nil, fmt.Sprintf("No vertex count specified")
	}
	if edges == nil {
		return nil, fmt.Sprintf("No edges found")
	}
	for _, edge := range edges {
		if edge[0] < 0 || edge[0] >= len(gr.vertices) ||
			edge[1] < 0 || edge[1] >= len(gr.vertices) {
			return nil, fmt.Sprintf("Invalid vertex number\n")
		}
		gr.vertices[edge[0]].neighbours = append(gr.vertices[edge[0]].neighbours,
			edge[1])
	}
	return gr, "lol"
}


func (gr *graph) scc() {
	fmt.Printf("Strongly connected components:\n")
	for _, vertex := range gr.vertices {
		vertex.index = -1
		vertex.lowlink = -1
		vertex.in_stack = false
	}
	index := 0
	st := create_stack()
	for i, vertex := range gr.vertices {
		if vertex.index < 0 {
			gr.tarjan(i, &index, st)
		}
	}
}


func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}


func (gr *graph) tarjan(v int, index *int, st *stack) {
	vr := gr.vertices[v]
	vr.index = *index
	vr.lowlink = *index
	*index = *index + 1
	st.push(v)
	vr.in_stack = true
	for _, i := range vr.neighbours {
		succ := gr.vertices[i]
		if succ.index < 0 {
			gr.tarjan(i, index, st)
			vr.lowlink = min(vr.lowlink, succ.lowlink)
		} else if succ.in_stack {
			vr.lowlink = min(vr.lowlink, succ.lowlink)
		}
	}
	if vr.lowlink == vr.index {
		first := true
		index := -1
		for index != v {
			index = st.pop()
			gr.vertices[index].in_stack = false
			if !first {
				fmt.Printf(" - ")
			}
			first = false
			fmt.Printf("%d", index)
		}
		fmt.Printf("\n")
	}
}


func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <input_graph.json>\n", os.Args[0])
		os.Exit(1)
	}
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	gr, err_str := create_graph(string(bytes))
	if gr == nil {
		fmt.Printf("%v\n", err_str)
		os.Exit(1)
	}
	gr.scc()
}
