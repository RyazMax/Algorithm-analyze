package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"./matrix"
)

func main() {
	start, _ := strconv.Atoi(os.Args[1])
	stop, _ := strconv.Atoi(os.Args[2])
	step, _ := strconv.Atoi(os.Args[3])
	filename := os.Args[4]

	file, _ := os.OpenFile(filename, os.O_WRONLY, 0644)

	for i := start; i <= stop; i += step {
		a, b := matrix.InitMatrixInt(i, i), matrix.InitMatrixInt(i, i)

		var regTime, vinTime, optVinTime time.Duration

		for j := 0; j < 20; j++ {
			t := time.Now()
			matrix.RegMulInt(a, b)
			regTime += time.Since(t)

			t = time.Now()
			matrix.VinMulInt(a, b)
			vinTime += time.Since(t)

			t = time.Now()
			matrix.OptVinMulInt(a, b)
			optVinTime += time.Since(t)

			log.Println("SIZE: ", i, " IT: ", j, "/20")
		}
		repr := strconv.FormatInt(int64(i), 10) + " " + strconv.FormatInt(regTime.Nanoseconds()/20, 10) + " " +
			strconv.FormatInt(vinTime.Nanoseconds()/20, 10) + " " +
			strconv.FormatInt(optVinTime.Nanoseconds()/20, 10) + "\n"
		file.WriteString(repr)
	}
}
