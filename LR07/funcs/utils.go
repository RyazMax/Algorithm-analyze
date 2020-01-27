package funcs

import "testing"

func assertSliceEqual(got, expected []int, msg string, t *testing.T) {
	pass := true
	if len(got) != len(expected) {
		pass = false
	}

	if pass {
		for i := range got {
			if got[i] != expected[i] {
				pass = false
			}
		}
	}

	if !pass {
		t.Error("WA:", msg, " EXPECTED:", expected, " GOT:", got)
	}
}
