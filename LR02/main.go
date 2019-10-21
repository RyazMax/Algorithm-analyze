package main

import (
	"fmt"

	"./matrix"
	"./utils"
)

func main() {
	var n, m, k, algo int

	fmt.Println("Введите размеры матрицы 1: ")
	fmt.Scanf("%d%d", &n, &m)

	fmt.Println("Введите число столбцов матрицы 2: ")
	fmt.Scanf("%d", &k)

	fmt.Println("Введите матрицу 1: ")
	a := utils.ScanMatrixFloat64(n, m)

	fmt.Println("Введите матрицу 2: ")
	b := utils.ScanMatrixFloat64(m, k)

	fmt.Println("Для использования обычного алгоритма введите 0, для алгоритма Винограда 1: ")
	fmt.Scanf("%d", algo)

	var res [][]float64
	if algo == 0 {
		res = matrix.RegMulFloat64(a, b)
	} else if algo == 1 {
		res = matrix.VinMulFloat64(a, b)
	} else {
		res = matrix.OptVinMulFloat64(a, b)
	}

	fmt.Println("Матрица результат: ")
	utils.PrintMatrixFloat64(res)
}
