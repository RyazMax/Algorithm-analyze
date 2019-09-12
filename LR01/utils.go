package main

import (
	"fmt"
	"io"
)

func min(a, b int) (res int) {
	res = a
	if b < a {
		res = b
	}
	return
}

func printSlice(outstream io.Writer, a []int) {
	if outstream == nil {
		return
	}

	for _, val := range a {
		fmt.Fprintf(outstream, "%v ", val)
	}
	fmt.Fprintln(outstream)
}
