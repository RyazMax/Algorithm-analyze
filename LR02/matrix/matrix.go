package matrix

func isEqualInt(a, b [][]int) bool {
	if len(a) != len(b) || len(a[0]) != len(b[0]) {
		return false
	}

	res := true
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				res = false
				break
			}
		}
	}
	return res
}

func InitMatrixFloat64(a, b int) [][]float64 {
	res := make([][]float64, a)
	for i := 0; i < a; i++ {
		res[i] = make([]float64, b)
	}
	return res
}

func InitMatrixInt(a, b int) [][]int {
	res := make([][]int, a)
	for i := 0; i < a; i++ {
		res[i] = make([]int, b)
	}
	return res
}

// RegMulFloat64 multiply NxM and MxK matrixes with float64 elements
func RegMulFloat64(a, b [][]float64) [][]float64 {
	var (
		n = len(a)
		m = len(b)
		q = len(b[0])
	)
	res := InitMatrixFloat64(n, q) // Инициализация матрицы результата

	for i := 0; i < n; i++ {
		for j := 0; j < q; j++ {
			for k := 0; k < m; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	return res
}

// RegMulInt multiply NxM and MxK matrixes with int elements
func RegMulInt(a, b [][]int) [][]int {
	var (
		n = len(a)
		m = len(b)
		q = len(b[0])
	)
	res := InitMatrixInt(n, q) // Инициализация матрицы результата

	for i := 0; i < n; i++ {
		for j := 0; j < q; j++ {
			for k := 0; k < m; k++ {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return res
}

// VinMulFloat64 multiply NxM and MxK matrixes with float64 elements
// Uses Vinograd algorithm
func VinMulFloat64(a, b [][]float64) [][]float64 {
	var (
		m = len(a)
		n = len(b)
		q = len(b[0])
	)
	res := InitMatrixFloat64(m, q) // Инициализация матрицы результата

	mulH := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			mulH[i] = mulH[i] + a[i][2*j]*a[i][2*j+1]
		}
	}

	mulV := make([]float64, q)
	for i := 0; i < q; i++ {
		for j := 0; j < n/2; j++ {
			mulV[i] = mulV[i] + b[2*j][i]*b[2*j+1][i]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			res[i][j] = -mulH[i] - mulV[j]
			for k := 0; k < n/2; k++ {
				res[i][j] = res[i][j] + (a[i][2*k]+b[2*k+1][j])*(a[i][2*k+1]+b[2*k][j])
			}
		}
	}

	if n%2 == 1 {
		for i := 0; i < m; i++ {
			for j := 0; j < q; j++ {
				res[i][j] = res[i][j] + a[i][n-1]*b[n-1][j]
			}
		}
	}

	return res
}

// VinMulInt multiply NxM and MxK matrixes with int elements
// Uses Vinograd algorithm
func VinMulInt(a, b [][]int) [][]int {
	var (
		m = len(a)
		n = len(b)
		q = len(b[0])
	)
	res := InitMatrixInt(m, q) // Инициализация матрицы результата

	mulH := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			mulH[i] = mulH[i] + a[i][2*j]*a[i][2*j+1]
		}
	}

	mulV := make([]int, q)
	for i := 0; i < q; i++ {
		for j := 0; j < n/2; j++ {
			mulV[i] = mulV[i] + b[2*j][i]*b[2*j+1][i]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			res[i][j] = -mulH[i] - mulV[j]
			for k := 0; k < n/2; k++ {
				res[i][j] = res[i][j] + (a[i][2*k]+b[2*k+1][j])*(a[i][2*k+1]+b[2*k][j])
			}
		}
	}

	if n%2 == 1 {
		for i := 0; i < m; i++ {
			for j := 0; j < q; j++ {
				res[i][j] = res[i][j] + a[i][n-1]*b[n-1][j]
			}
		}
	}

	return res
}

// OptVinMulInt optimized version of VinMulInt
func OptVinMulInt(a, b [][]int) [][]int {
	var (
		m      = len(a)
		n      = len(b)
		q      = len(b[0])
		nMinus = n - 1
	)
	res := InitMatrixInt(m, q) // Инициализация матрицы результата

	mulH := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < nMinus; j += 2 {
			mulH[i] -= a[i][j] * a[i][j+1]
		}
	}

	mulV := make([]int, q)
	for i := 0; i < q; i++ {
		for j := 0; j < nMinus; j += 2 {
			mulV[i] -= b[j][i] * b[j+1][i]
		}
	}

	isOdd := n%2 == 1
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			buf := mulH[i] + mulV[j]
			for k := 0; k < nMinus; k += 2 {
				buf += (a[i][k] + b[k+1][j]) * (a[i][k+1] + b[k][j])
			}
			if isOdd {
				buf += a[i][nMinus] * b[nMinus][j]
			}
			res[i][j] = buf
		}
	}

	return res
}

// OptVinMulFloat optimized version of VinMulFloat
func OptVinMulFloat64(a, b [][]float64) [][]float64 {
	var (
		m      = len(a)
		n      = len(b)
		q      = len(b[0])
		nMinus = n - 1
	)
	res := InitMatrixFloat64(m, q) // Инициализация матрицы результата

	mulH := make([]float64, m)
	for i := 0; i < m; i++ {
		for j := 0; j < nMinus; j += 2 {
			mulH[i] -= a[i][j] * a[i][j+1]
		}
	}

	mulV := make([]float64, q)
	for i := 0; i < q; i++ {
		for j := 0; j < nMinus; j += 2 {
			mulV[i] -= b[j][i] * b[j+1][i]
		}
	}

	isOdd := n%2 == 1
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			buf := mulH[i] + mulV[j]
			for k := 0; k < nMinus; k += 2 {
				buf += (a[i][k] + b[k+1][j]) * (a[i][k+1] + b[k][j])
			}
			if isOdd {
				buf += a[i][nMinus] * b[nMinus][j]
			}
			res[i][j] = buf
		}
	}

	return res
}
