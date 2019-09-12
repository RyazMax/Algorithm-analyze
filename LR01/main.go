package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var finder distanceFinder
	if len(os.Args) < 4 {
		fmt.Println("Usage go main.go string string [...]")
		return
	}

	switch os.Args[3] {
	case "-l":
		finder = editorDistance
	case "-d":
		finder = domerauEditorDistance
	case "-r":
		finder = domerauRecursiveEditorDistance
	default:
		fmt.Println("Unknown distance function")
		return
	}

	t := time.Now()
	res := finder(os.Stdout, os.Args[1], os.Args[2])
	fmt.Println("Answer is: ", res)
	fmt.Println("Work took: ", time.Since(t))
}
