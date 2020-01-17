package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"./sorts"
)

func fillOrder(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return a
}

func fillAntiOrder(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = n - i
	}
	return a
}

func fillRandom(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Int()
	}
	return a
}

const (
	MIN_SIZE = 1
	MAX_SIZE = 12000
	EXPS     = 6
)

func testCase(a []int, sort func(a []int)) time.Duration {
	b := make([]int, len(a))
	copy(b, a)
	t := time.Now()
	sort(b)
	scince := time.Since(t)
	return scince
}

func writeRow(i int, a, b, c time.Duration, file *os.File) {
	file.WriteString(strconv.FormatInt(int64(i), 10) + " " + strconv.FormatInt(a.Nanoseconds(), 10) + " " +
		strconv.FormatInt(b.Nanoseconds(), 10) + " " + strconv.FormatInt(c.Nanoseconds(), 10) + "\n")
}

func main() {

	fileOrdered, _ := os.OpenFile("ordered.log", os.O_WRONLY|os.O_CREATE, 0644)
	fileAntiOrdered, _ := os.OpenFile("antiordered.log", os.O_WRONLY|os.O_CREATE, 0644)
	fileRandom, _ := os.OpenFile("random.log", os.O_WRONLY|os.O_CREATE, 0644)
	defer fileOrdered.Close()
	defer fileAntiOrdered.Close()
	defer fileRandom.Close()

	for i := MIN_SIZE; i <= MAX_SIZE; i++ {
		var insertTime, mergeTime, radixTime time.Duration

		// Ordered
		for j := 0; j < EXPS; j++ {
			a := fillOrder(i)
			insertTime += testCase(a, sorts.InsertSort)
			mergeTime += testCase(a, sorts.MergeSort)
			radixTime += testCase(a, sorts.RadixSort)
		}
		insertTime = insertTime / EXPS
		mergeTime = mergeTime / EXPS
		radixTime = radixTime / EXPS
		writeRow(i, insertTime, mergeTime, radixTime, fileOrdered)
		insertTime = 0
		mergeTime = 0
		radixTime = 0

		// AntiOrdered
		for j := 0; j < EXPS; j++ {
			a := fillAntiOrder(i)
			insertTime += testCase(a, sorts.InsertSort)
			mergeTime += testCase(a, sorts.MergeSort)
			radixTime += testCase(a, sorts.RadixSort)
		}
		insertTime = insertTime / EXPS
		mergeTime = mergeTime / EXPS
		radixTime = radixTime / EXPS
		writeRow(i, insertTime, mergeTime, radixTime, fileAntiOrdered)
		insertTime = 0
		mergeTime = 0
		radixTime = 0

		// Random
		for j := 0; j < EXPS; j++ {
			a := fillRandom(i)
			insertTime += testCase(a, sorts.InsertSort)
			mergeTime += testCase(a, sorts.MergeSort)
			radixTime += testCase(a, sorts.RadixSort)
		}

		insertTime = insertTime / EXPS
		mergeTime = mergeTime / EXPS
		radixTime = radixTime / EXPS
		writeRow(i, insertTime, mergeTime, radixTime, fileRandom)

		if i%100 == 0 {
			log.Println(i, insertTime, mergeTime, radixTime)
		}
	}
}
