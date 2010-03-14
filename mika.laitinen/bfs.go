package main

import "fmt"
import . "container/vector"

type t [][]int;

func bfs(arr t, startNode int, endNode int) int {
	var vecNodes IntVector
	var vecPrices IntVector
	var allowedNodes IntVector

	allowedNodes.Resize(len(arr),len(arr));

	vecNodes.Push(startNode);
	vecPrices.Push(0);

	for i := 0; i < len(allowedNodes); i++ {
		allowedNodes[i] = 1;
	}

	for ind := 0; ind < len(vecNodes); ind++ {
		for i := 0; i < len(arr[vecNodes[ind]]); i++ {
			if arr[vecNodes.At(ind)][i] > 0 && allowedNodes.At(i) > 0 {
				if i == endNode {
					return vecPrices.At(ind) + 1;
				}
				allowedNodes[i] = 0;
				vecNodes.Push(i);
				vecPrices.Push(vecPrices.At(ind) + 1);
			}
		}
	}

	return -1;
}

var graphA = [][]int {
	[]int{0,1,0,0,0},
	[]int{0,0,1,0,0},
	[]int{0,0,0,1,0},
	[]int{0,0,0,0,1},
	[]int{0,0,0,0,0}};

var graphB = [][]int{
	[]int{0,0,0,0,0,0,1,0,0,0,0,0},
	[]int{0,0,1,0,0,1,0,1,0,0,0,0},
	[]int{0,1,0,0,0,0,1,0,0,0,0,0},
	[]int{0,0,0,0,0,0,0,0,0,0,0,0},
	[]int{0,0,0,0,0,0,0,0,0,0,1,1},
	[]int{0,1,0,0,0,0,0,0,0,1,0,0},
	[]int{1,0,1,0,0,0,0,0,0,0,0,0},
	[]int{0,1,0,0,0,0,0,0,1,0,0,0},
	[]int{0,0,0,0,0,0,0,1,0,0,1,0},
	[]int{0,0,0,0,0,1,0,0,0,0,1,0},
	[]int{0,0,0,0,1,0,0,0,1,1,0,0},
	[]int{0,0,0,0,1,0,0,0,0,0,0,0}};

func main() {
	fmt.Printf("Route length: %d\n",bfs(graphA,0,4));
	fmt.Printf("Route length: %d\n",bfs(graphB,0,11));
}
