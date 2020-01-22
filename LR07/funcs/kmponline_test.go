package funcs

import (
	"strings"
	"testing"
)

func TestKmpOnlineNullPattern(t *testing.T) {
	got := KmpOnline("", strings.NewReader("Aloha"))
	expected := []int{}

	assertSliceEqual(got, expected, "on null pattern", t)
	got = KmpOnline("", strings.NewReader(""))
	expected = []int{}

	assertSliceEqual(got, expected, "on null pattern", t)
}

func TestKmpOnlineNullStr(t *testing.T) {
	got := KmpOnline("A", strings.NewReader(""))
	expected := []int{}

	assertSliceEqual(got, expected, "on null str", t)

	got = KmpOnline("abacaba", strings.NewReader(""))
	expected = []int{}

	assertSliceEqual(got, expected, "on null str", t)
}

func TestKmpOnlineCyrillic(t *testing.T) {
	got := KmpOnline("Дом", strings.NewReader("домЭтонеДом"))
	expected := []int{8}

	assertSliceEqual(got, expected, "on cyrillic letters", t)

	got = KmpOnline("аба", strings.NewReader("аабакаба"))
	expected = []int{1, 5}

	assertSliceEqual(got, expected, "on cyrillic letters", t)
}

func TestKmpOnlineOk(t *testing.T) {
	got := KmpOnline("a", strings.NewReader("abAba"))
	expected := []int{0, 4}

	assertSliceEqual(got, expected, "on ok str", t)

	got = KmpOnline("ab", strings.NewReader("abacbaaba"))
	expected = []int{0, 6}

	assertSliceEqual(got, expected, "on ok str", t)

	got = KmpOnline("abaa", strings.NewReader("aabaaa"))
	expected = []int{1}

	assertSliceEqual(got, expected, "on ok str", t)
}
