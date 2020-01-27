package funcs

import "testing"

func TestPiFuncNullPattern(t *testing.T) {
	got := piFunc("#")
	expected := []int{0}

	assertSliceEqual(got, expected, "on null pattern", t)
}

func TestPiFuncOk(t *testing.T) {
	got := piFunc("aa")
	expected := []int{0, 1}

	assertSliceEqual(got, expected, "on double letters", t)

	got = piFunc("bac")
	expected = []int{0, 0, 0}

	assertSliceEqual(got, expected, "on different letters", t)

	got = piFunc("ababacabaca")
	expected = []int{0, 0, 1, 2, 3, 0, 1, 2, 3, 0, 1}

	assertSliceEqual(got, expected, "on crossed letters", t)
}

func TestPiFuncCyrillic(t *testing.T) {
	got := piFunc("абакаба")
	expected := []int{0, 0, 1, 0, 1, 2, 3}

	assertSliceEqual(got, expected, "on cyrillic letters", t)
}
