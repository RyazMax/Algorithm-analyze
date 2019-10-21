package matrix

import "testing"

var (
	aInt [][]int
	bInt [][]int
	cInt [][]int
	dInt [][]int
	eInt [][]int
	fInt [][]int
)

func init() {
	// Первый кейс
	aInt = InitMatrixInt(2, 2)
	bInt = InitMatrixInt(2, 2)
	cInt = InitMatrixInt(2, 2)
	aInt[0][0], aInt[0][1], aInt[1][0], aInt[1][1] = 1, 1, 1, 1
	bInt[0][0], bInt[0][1], bInt[1][0], bInt[1][1] = 1, 2, 2, 1
	cInt[0][0], cInt[0][1], cInt[1][0], cInt[1][1] = 3, 3, 3, 3

	// Второй кейс
	dInt = InitMatrixInt(2, 3)
	eInt = InitMatrixInt(3, 2)
	fInt = InitMatrixInt(2, 2)
	dInt[0][0], dInt[0][1], dInt[0][2] = 1, 2, 3
	dInt[1][0], dInt[1][1], dInt[1][2] = 2, 3, 4
	eInt[0][0], eInt[0][1], eInt[1][0], eInt[1][1], eInt[2][0], eInt[2][1] = 1, 2, 3, 4, 1, 2
	fInt[0][0], fInt[0][1], fInt[1][0], fInt[1][1] = 10, 16, 15, 24
}

func TestRegMulInt(t *testing.T) {
	if !isEqualInt(RegMulInt(aInt, bInt), cInt) {
		t.Error("Incorrect mul on square matrixes")
	}

	if !isEqualInt(RegMulInt(dInt, eInt), fInt) {
		t.Error("Incorrect answer on diff shapes matrixes")
	}
}

func TestVinMulInt(t *testing.T) {
	if !isEqualInt(VinMulInt(aInt, bInt), cInt) {
		t.Error("Incorrect mul on square matrixes")
	}

	if !isEqualInt(VinMulInt(dInt, eInt), fInt) {
		t.Error("Incorrect answer on diff shapes matrixes")
	}
}

func TestOptVinMulInt(t *testing.T) {
	if !isEqualInt(OptVinMulInt(aInt, bInt), cInt) {
		t.Error("Incorrect mul on square matrixes")
	}

	if !isEqualInt(OptVinMulInt(dInt, eInt), fInt) {
		t.Error("Incorrect answer on diff shapes matrixes")
	}
}
