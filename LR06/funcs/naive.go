package funcs

import (
	"github.com/gitchander/permutation"
)

func countRoute(route []int, mat [][]int) int {
	if len(route) == 0 {
		return 0
	}

	if len(mat) == 0 {
		panic("Can not procces empty matrix")
	}

	var (
		res      int
		from, to int
	)

	from = route[0]
	for i := 1; i < len(route); i++ {
		to = route[i]
		res += mat[from][to]
		from = to
	}

	return res
}

func createRange(size int) []int {
	var res []int
	for i := 0; i < size; i++ {
		res = append(res, i)
	}
	return res
}

// FindShortestNaive finds the shortest gamilton route
func FindShortestNaive(mat [][]int) (route []int, dist int) {
	dist = -1
	routes := createRange(len(mat))

	permutator := permutation.New(permutation.IntSlice(routes))

	for permutator.Next() {
		cost := countRoute(routes, mat)
		if dist < 0 || dist > cost {
			// TODO make copy to result because last .Next() makes the very fist permutation
			route = routes
			dist = cost
		}
	}

	return
}
