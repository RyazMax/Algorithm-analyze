package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"./funcs"
)

// GeneratePolicy interface for generating input data for substring finders
type GeneratePolicy interface {
	Generate(a, b int) (s, t string)
}

// OnePatternPolicy implement GeneratePolicy that simply repeat given patterns to fill all length
type OnePatternPolicy struct {
	pattern string
	text    string
}

// Generate generates strings
func (p OnePatternPolicy) Generate(a, b int) (s, t string) {
	s = strings.Repeat(p.pattern, a/utf8.RuneCountInString(p.pattern))
	s += p.pattern[:a%utf8.RuneCountInString(p.pattern)]

	t = strings.Repeat(p.text, b/utf8.RuneCountInString(p.text))
	t += p.text[:b%utf8.RuneCountInString(p.text)]
	return
}

// OneAndThen operates like OnePatternPolicy but place patternX in front
type OneAndThen struct {
	patternX string
	patternY string
	text     string
}

// Generate generates
func (p OneAndThen) Generate(a, b int) (s, t string) {
	s, t = OnePatternPolicy{p.patternY, p.text}.Generate(a-1, b)
	s = p.patternX + s
	return
}

// ExperimentInfo holds info about experiment
type ExperimentInfo struct {
	start            int
	stop             int
	step             int
	patternLen       int
	experimentsCount int
}

func carryExperiment(info ExperimentInfo, generator GeneratePolicy, strfinder funcs.Substrfinder, w io.Writer) {
	for i := info.start; i < info.stop; i += info.step {
		s, t := generator.Generate(info.patternLen, i)

		var duration time.Duration
		for j := 0; j < info.experimentsCount; j++ {
			startTime := time.Now()
			strfinder(s, t)
			duration += time.Since(startTime)
		}
		duration = duration / time.Duration(info.experimentsCount)

		log.Println("ITER", i)

		stringToWrite := strconv.Itoa(i) + " " + strconv.FormatInt(duration.Nanoseconds(), 10) + "\n"

		w.Write([]byte(stringToWrite))
	}
}

func main() {
	fileName := os.Args[1]
	funcNum, _ := strconv.Atoi(os.Args[2])

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Cannt open file")
		os.Exit(1)
	}
	defer file.Close()

	info := ExperimentInfo{
		start:            100000,
		stop:             1000000,
		step:             400,
		patternLen:       100,
		experimentsCount: 30,
	}

	var fn12 funcs.Substrfinder

	switch funcNum {
	case 1:
		fn12 = funcs.Kmp
	case 2:
		fn12 = funcs.KmpOnlineWrapper
	case 3:
		fn12 = funcs.BoyerMur
	}

	carryExperiment(info, OnePatternPolicy{"a", "ab"}, fn12, file)
}
