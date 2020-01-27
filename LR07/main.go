package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"./funcs"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give at least 1 argument")
		fmt.Println("Usagee go run main.go <func num>")
		fmt.Println("Func nums:")
		fmt.Println("1. KMP")
		fmt.Println("2. KMP online")
		fmt.Println("3. BoyerMoor")
		os.Exit(1)
	}
	funcNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Incorrect argument")
		os.Exit(1)
	}

	fmt.Println("Input substring and string")
	var s, t string
	fmt.Scan(&s, &t)

	switch funcNum {
	case 1:
		fmt.Println("ANS:", funcs.Kmp(s, t))
	case 2:
		fmt.Println("ANS:", funcs.KmpOnline(s, strings.NewReader(t)))
	case 3:
		fmt.Println("ANS:", funcs.BoyerMur(s, t))
	}
}
