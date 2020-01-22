package funcs

import (
	"unicode/utf8"
)

// Kmp finds substring s in string s
func Kmp(s, t string) []int {
	if len(s) == 0 {
		return []int{}
	}

	pi := piFunc(s + "#" + t)

	res := make([]int, 0)
	lens := utf8.RuneCountInString(s)
	for i, v := range pi {
		if v == lens {
			res = append(res, i-2*lens)
		}
	}
	return res
}
