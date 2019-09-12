package main

import (
	"io"
	"math"
	"unicode/utf8"
)

type distanceFinder func(outstream io.Writer, a, b string) int

func editorDistance(outstream io.Writer, a, b string) (dist int) {
	var (
		aLen   = utf8.RuneCountInString(a)
		bLen   = utf8.RuneCountInString(b)
		dpPrev = make([]int, bLen+1)
		dpCurr = make([]int, bLen+1)
		aRunes = []rune(a)
		bRunes = []rune(b)
	)

	for i := 0; i < bLen+1; i++ {
		dpPrev[i] = i
	}
	printSlice(outstream, dpPrev)

	for i := 1; i < aLen+1; i++ {
		dpCurr[0] = i

		for j := 1; j < bLen+1; j++ {
			dpCurr[j] = dpPrev[j-1]
			if aRunes[i-1] != bRunes[j-1] {
				dpCurr[j]++
			}
			dpCurr[j] = min(dpCurr[j], min(dpCurr[j-1]+1, dpPrev[j]+1))
		}

		printSlice(outstream, dpCurr)

		dpPrev, dpCurr = dpCurr, dpPrev
	}

	return dpPrev[bLen]
}

func domerauEditorDistance(outstream io.Writer, a, b string) (dist int) {
	var (
		aLen      = utf8.RuneCountInString(a)
		bLen      = utf8.RuneCountInString(b)
		dpPrePrev = make([]int, bLen+1)
		dpPrev    = make([]int, bLen+1)
		dpCurr    = make([]int, bLen+1)
		aRunes    = []rune(a)
		bRunes    = []rune(b)
	)

	for i := 0; i < bLen+1; i++ {
		dpPrev[i] = i
	}
	printSlice(outstream, dpPrev)

	for i := 1; i < aLen+1; i++ {
		dpCurr[0] = i

		for j := 1; j < bLen+1; j++ {
			dpCurr[j] = dpPrev[j-1]
			if aRunes[i-1] != bRunes[j-1] {
				dpCurr[j]++
			}

			if i > 1 && j > 1 {
				if aRunes[i-2] == bRunes[j-1] && aRunes[i-1] == bRunes[j-2] {
					dpCurr[j] = min(dpCurr[j], dpPrePrev[j-2]+1)
				}
			}
			dpCurr[j] = min(dpCurr[j], min(dpCurr[j-1]+1, dpPrev[j]+1))
		}
		printSlice(outstream, dpCurr)

		dpPrePrev, dpPrev, dpCurr = dpPrev, dpCurr, dpPrePrev
	}

	return dpPrev[bLen]
}

func domerauRecursiveEditorDistance(outstream io.Writer, a, b string) (dist int) {
	return _domerauRecursive([]rune(a), []rune(b), utf8.RuneCountInString(a), utf8.RuneCountInString(b))
}

func _domerauRecursive(a, b []rune, i, j int) (dist int) {
	if i == 0 {
		return j
	}
	if j == 0 {
		return i
	}

	transition := math.MaxInt64
	if i > 2 && j > 2 && a[i-1] == b[j-2] && a[i-2] == b[j-1] {
		transition = _domerauRecursive(a, b, i-2, j-2) + 1
	}
	replace := _domerauRecursive(a, b, i-1, j-1)
	if a[i-1] != b[j-1] {
		replace++
	}
	return min(transition, min(replace, min(_domerauRecursive(a, b, i, j-1)+1, _domerauRecursive(a, b, i-1, j)+1)))
}
