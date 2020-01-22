package funcs

func piFunc(s string) []int {
	sInRune := []rune(s)

	pi := make([]int, len(sInRune))
	for i := 1; i < len(sInRune); i++ {
		j := pi[i-1]
		for j > 0 && sInRune[i] != sInRune[j] {
			j = pi[j-1]
		}
		if sInRune[i] == sInRune[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}
