package funcs

import "testing"

func TestKmpNullPattern(t *testing.T) {
	got := Kmp("", "Aloha")
	expected := []int{0, 1, 2, 3, 4}

	assertSliceEqual(got, expected, "on null pattern", t)
	got = Kmp("", "")
	expected = []int{0}

	assertSliceEqual(got, expected, "on null pattern", t)
}

func TestKmpNullStr(t *testing.T) {
	got := Kmp("A", "")
	expected := []int{}

	assertSliceEqual(got, expected, "on null str", t)

	got = Kmp("abacaba", "")
	expected = []int{}

	assertSliceEqual(got, expected, "on null str", t)
}

func TestKmpOk(t *testing.T) {
	got := Kmp("a", "abAba")
	expected := []int{0, 4}

	assertSliceEqual(got, expected, "on null str", t)
}
