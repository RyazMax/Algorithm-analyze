package funcs

// BoyerMur finds substring s in string t
func BoyerMur(s, t string) []int {
	if len(s) == 0 {
		return []int{}
	}

	sInRune := []rune(s)
	tInRune := []rune(t)
	shifts := make(map[rune]int)

	for i := 0; i < len(sInRune)-1; i++ {
		shifts[sInRune[i]] = len(sInRune) - i - 1
	}

	res := make([]int, 0)
	pos := len(sInRune) - 1

	for pos < len(tInRune) {
		j := 0
		for ; j < len(sInRune) && sInRune[len(sInRune)-j-1] == tInRune[pos-j]; j++ {
		}
		if j == len(sInRune) {
			res = append(res, pos-len(sInRune)+1)
		}
		if sh, exist := shifts[tInRune[pos]]; exist {
			pos += sh
		} else {
			pos += len(sInRune)
		}
	}

	return res
}
