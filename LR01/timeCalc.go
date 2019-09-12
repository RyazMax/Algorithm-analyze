package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage go run timeCalc.go strLen -key")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Incorrect number")
		return
	}

	var finder distanceFinder

	switch os.Args[2] {
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

	var duration time.Duration
	for i := 0; i < 10; i++ {
		str1 := randStringRunes(n)
		str2 := randStringRunes(n)
		t := time.Now()
		_ = finder(nil, str1, str2)
		duration += time.Since(t)
	}

	fmt.Println(duration / 10)

}
