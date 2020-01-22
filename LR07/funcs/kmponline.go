package funcs

import (
	"strings"
)

// KmpOnline finds substring s in string t using only O(len(s)) memory
func KmpOnline(s string, t *strings.Reader) []int {
	if len(s) == 0 {
		return []int{}
	}

	s = s + "#"
	pi := piFunc(s)
	sInRune := []rune(s)
	lens := len(sInRune)

	res := make([]int, 0)

	lastIdx := 0
	count := 0
	for t.Len() > 0 {
		r, _, _ := t.ReadRune()

		j := lastIdx
		for j > 0 && sInRune[j] != r {
			j = pi[j-1]
		}
		if sInRune[j] == r {
			j++
		}
		lastIdx = j

		if j == lens-1 {
			res = append(res, count-lens+2)
		}
		count++
	}
	return res
}

// KmpOnlineWrapper is wrapper to use t as string instead of strings.Reader
func KmpOnlineWrapper(s, t string) []int {
	return KmpOnline(s, strings.NewReader(t))
}
