package funcs

import "testing"

func TestCreateRange(t *testing.T) {
	a := createRange(0)
	b := make([]int, 0)
	AssertSliceEqual(a, b, "null size:", t)

	a = createRange(1)
	b = []int{0}
	AssertSliceEqual(a, b, "size of 1:", t)

	a = createRange(4)
	b = []int{0, 1, 2, 3}
	AssertSliceEqual(a, b, "size of 4:", t)
}

func TestCountRangeOk(t *testing.T) {
	mat := [][]int{
		[]int{0, 2},
		[]int{1, 0},
	}

	route := []int{0, 1}
	res := countRoute(route, mat)
	if res != 2 {
		t.Error("GOT", res, "EXPECT", 2)
	}

	route = []int{1, 0}
	res = countRoute(route, mat)
	if res != 1 {
		t.Error("GOT", res, "EXPECT", 1)
	}

	mat = [][]int{
		[]int{0, 2, 3},
		[]int{2, 0, 1},
		[]int{3, 4, 0},
	}

	route = []int{1, 2, 0}
	res = countRoute(route, mat)
	if res != 4 {
		t.Error("GOT", res, "EXPECT", 4)
	}

	route = []int{2, 1, 0}
	res = countRoute(route, mat)
	if res != 6 {
		t.Error("GOT", res, "EXPECT", 6)
	}

	route = []int{}
	res = countRoute(route, mat)
	if res != 0 {
		t.Error("GOT", res, "EXPECT", 0)
	}

	mat = [][]int{
		[]int{0},
	}
	route = []int{0}
	res = countRoute(route, mat)
	if res != 0 {
		t.Error("GOT", res, "EXPECT", 0)
	}
}

func TestCountRangePanic(t *testing.T) {
	// TODO: need to find out how to proccess panic throw recover
}

func TestFindShortestNaiveOk(t *testing.T) {
	// Case with only one point
	mat := [][]int{
		[]int{0},
	}
	route, dist := FindShortestNaive(mat)

	AssertSliceEqual(route, []int{0}, "One point matrix:", t)
	if dist != 0 {
		t.Error("GOT", dist, "EXPECTED", 0)
	}

	// Cases with 2x2 matrix
	// One of different
	mat = [][]int{
		[]int{0, 2},
		[]int{1, 0},
	}
	route, dist = FindShortestNaive(mat)

	AssertSliceEqual(route, []int{1, 0}, "2x2 matrix, with different routes", t)
	if dist != 1 {
		t.Error("GOT", dist, "EXPECTED", 1)
	}

	// Equal distance routes
	mat = [][]int{
		[]int{0, 2},
		[]int{2, 0},
	}
	route, dist = FindShortestNaive(mat)

	AssertSliceEqual(route, []int{0, 1}, "2x2 matrix, equal distance routes should return first route", t)
	if dist != 2 {
		t.Error("GOT", dist, "EXPECTED", 2)
	}
}
