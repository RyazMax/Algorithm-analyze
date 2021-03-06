package funcs

import "testing"

func TestKmpNullPattern(t *testing.T) {
	got := Kmp("", "Aloha")
	expected := []int{}

	assertSliceEqual(got, expected, "on null pattern", t)
	got = Kmp("", "")
	expected = []int{}

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

func TestKmpCyrillic(t *testing.T) {
	got := Kmp("Дом", "домЭтонеДом")
	expected := []int{8}

	assertSliceEqual(got, expected, "on cyrillic letters", t)

	got = Kmp("аба", "аабакаба")
	expected = []int{1, 5}

	assertSliceEqual(got, expected, "on cyrillic letters", t)
}

func TestKmpOk(t *testing.T) {
	got := Kmp("a", "abAba")
	expected := []int{0, 4}

	assertSliceEqual(got, expected, "on ok str", t)

	got = Kmp("ab", "abacbaaba")
	expected = []int{0, 6}

	assertSliceEqual(got, expected, "on ok str", t)

	got = Kmp("abaa", "aabaaa")
	expected = []int{1}

	assertSliceEqual(got, expected, "on ok str", t)
}
