package main

import (
	"fmt"
	"os"

	"./funcs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You should give 1 parametr - type of algorithm")
		fmt.Println("n - naive, a - aunt algorithm")
		os.Exit(1)
	}

	var n int
	fmt.Println("Input size of matrix")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Size should be positive")
		os.Exit(1)
	}

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		fmt.Println("Input", i, "line")
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	var (
		route []int
		cost  int
	)
	if os.Args[1] == "a" {
	} else {
		route, cost = funcs.FindShortestNaive(mat)
	}

	fmt.Println("The shortest route is")
	for _, v := range route {
		fmt.Print(v, " ")
	}
	fmt.Println("It's cost is ", cost)
}
