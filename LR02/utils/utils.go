package utils

import "fmt"

// PrintMatrixFloat64 prints matrix into strings
func PrintMatrixFloat64(a [][]float64) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%.3f ", a[i][j])
		}
		fmt.Println()
	}
}

// ScanMatrixFloat64 scans matrix from stdin
func ScanMatrixFloat64(n, m int) [][]float64 {
	res := make([][]float64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]float64, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%f", &res[i][j])
		}
	}
	return res
}
