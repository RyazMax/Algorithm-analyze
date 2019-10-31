package sorts

import "testing"

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInsertionSortSimple(t *testing.T) {
	a := []int{3, 1, 8, -1, 4, 1}
	b := []int{-1, 1, 1, 3, 4, 8}

	InsertSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestInsertionSortEmpty(t *testing.T) {
	a := []int{}
	b := []int{}

	InsertSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestInsertionSortOne(t *testing.T) {
	a := []int{5}
	b := []int{5}

	InsertSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestInsertionSortOrdered(t *testing.T) {
	a := []int{-2, 1, 3, 3, 8}
	b := []int{-2, 1, 3, 3, 8}

	InsertSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestInsertionSortAntiOrdered(t *testing.T) {
	a := []int{8, 3, 3, 1, -2}
	b := []int{-2, 1, 3, 3, 8}

	InsertSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

// MergeSortTests

func TestMergeSortSimple(t *testing.T) {
	a := []int{3, 1, 8, -1, 4, 1}
	b := []int{-1, 1, 1, 3, 4, 8}

	MergeSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestMergeSortEmpty(t *testing.T) {
	a := []int{}
	b := []int{}

	MergeSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestMergeSortOne(t *testing.T) {
	a := []int{5}
	b := []int{5}

	MergeSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestMergeSortOrdered(t *testing.T) {
	a := []int{-2, 1, 3, 3, 8}
	b := []int{-2, 1, 3, 3, 8}

	MergeSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestMergeSortAntiOrdered(t *testing.T) {
	a := []int{8, 3, 3, 1, -2}
	b := []int{-2, 1, 3, 3, 8}

	MergeSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

// Radix sort tests

func TestRadixSortSimple(t *testing.T) {
	a := []int{3, 1, 8, 3, 4, 1}
	b := []int{1, 1, 3, 3, 4, 8}

	RadixSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestRadixSortEmpty(t *testing.T) {
	a := []int{}
	b := []int{}

	RadixSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestRadixSortOne(t *testing.T) {
	a := []int{5}
	b := []int{5}

	RadixSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestRadixSortOrdered(t *testing.T) {
	a := []int{1, 3, 3, 8, 28}
	b := []int{1, 3, 3, 8, 28}

	RadixSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}

func TestRadixSortAntiOrdered(t *testing.T) {
	a := []int{28, 8, 3, 3, 1}
	b := []int{1, 3, 3, 8, 28}

	RadixSort(a)
	if !isEqual(a, b) {
		t.Errorf("\n\tExpected %v\n\tFound %v", b, a)
	}
}
