package funcs

import "testing"

func TestBMNullPattern(t *testing.T) {
	got := BoyerMur("", "Aloha")
	expected := []int{}

	assertSliceEqual(got, expected, "on null pattern", t)
	got = BoyerMur("", "")
	expected = []int{}

	assertSliceEqual(got, expected, "on null pattern", t)
}

func TestBMNullStr(t *testing.T) {
	got := BoyerMur("A", "")
	expected := []int{}

	assertSliceEqual(got, expected, "on null str", t)

	got = BoyerMur("abacaba", "")
	expected = []int{}

	assertSliceEqual(got, expected, "on null str", t)
}

func TestBMCyrillic(t *testing.T) {
	got := BoyerMur("Дом", "домЭтонеДом")
	expected := []int{8}

	assertSliceEqual(got, expected, "on cyrillic letters", t)

	got = BoyerMur("аба", "аабакаба")
	expected = []int{1, 5}

	assertSliceEqual(got, expected, "on cyrillic letters", t)
}

func TestBMOk(t *testing.T) {
	got := BoyerMur("a", "abAba")
	expected := []int{0, 4}

	assertSliceEqual(got, expected, "on ok str", t)

	got = BoyerMur("ab", "abacbaaba")
	expected = []int{0, 6}

	assertSliceEqual(got, expected, "on ok str", t)

	got = BoyerMur("abaa", "aabaaa")
	expected = []int{1}

	assertSliceEqual(got, expected, "on ok str", t)
}
