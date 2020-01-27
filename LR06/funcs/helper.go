package funcs

import "testing"

// AssertSliceEqual checks if a,b are equal int slices
func AssertSliceEqual(got, expect []int, msg string, t *testing.T) {
	if len(got) != (len(expect)) {
		t.Error(msg, "GOT ", got, "EXPECT ", expect)
	}

	isOk := true
	for i, v := range got {
		if v != expect[i] {
			isOk = false
			break
		}
	}

	if !isOk {
		t.Error(msg, "GOT ", got, "EXPECT ", expect)
	}
}
