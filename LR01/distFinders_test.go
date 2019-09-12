package main

import (
	"bytes"
	"testing"
)

const (
	emptyStringMatrixOutput  = "0 \n1 \n"
	equalStringsMatrixOutput = `0 1 2 3 4 5 
1 0 1 2 3 4 
2 1 0 1 2 3 
3 2 1 0 1 2 
4 3 2 1 0 1 
5 4 3 2 1 0 
`
	kirillicStringsMatrixOutput = `0 1 2 3 
1 1 2 3 
2 1 2 3 
3 2 2 3 
4 3 3 2 
`
	transitionStringsMatrixOutput = `0 1 2 3 
1 1 2 3 
2 2 2 2 
3 3 2 2 
`
)

func TestEditorDistanceAllEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := editorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found", res)
	}

	if outstream.String() != "0 \n" {
		t.Error("Expected:\n0 \n \nFound:\n", outstream.String())
	}
}
func TestEditorDistanceEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := editorDistance(outstream, "a", "")
	if res != 1 {
		t.Error("Expected 1, found ", res)
	}

	if emptyStringMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", emptyStringMatrixOutput, "Found:\n", outstream.String())
	}

	outstream = bytes.NewBufferString("")
	res = editorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}

	if "0 \n" != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", "0 \n", "Found:\n", outstream.String())
	}
}

func TestEditorDistanceEqual(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := editorDistance(outstream, "Hello", "Hello")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}

	if equalStringsMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", equalStringsMatrixOutput, "Found:\n", outstream.String())
	}
}

func TestEditorDistanceKirillic(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := editorDistance(outstream, "скат", "кот")
	if res != 2 {
		t.Error("Expected 2, found ", res)
	}

	if kirillicStringsMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", kirillicStringsMatrixOutput, "Found:\n", outstream.String())
	}
}

func TestDomerauEditorDistanceAllEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauEditorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found", res)
	}

	if outstream.String() != "0 \n" {
		t.Error("Expected:\n0 \n \nFound:\n", outstream.String())
	}
}

func TestDomerautEditorDistanceEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauEditorDistance(outstream, "a", "")
	if res != 1 {
		t.Error("Expected 1, found ", res)
	}

	if emptyStringMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", emptyStringMatrixOutput, "Found:\n", outstream.String())
	}

	outstream = bytes.NewBufferString("")
	res = domerauEditorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}

	if "0 \n" != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", "0 \n", "Found:\n", outstream.String())
	}

}

func TestDomerauEditorDistanceEqual(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauEditorDistance(outstream, "Hello", "Hello")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}

	if equalStringsMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", equalStringsMatrixOutput, "Found:\n", outstream.String())
	}
}

func TestDomerauEditorDistanceKirillic(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauEditorDistance(outstream, "скат", "кот")
	if res != 2 {
		t.Error("Expected 2, found ", res)
	}

	if kirillicStringsMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", kirillicStringsMatrixOutput, "Found:\n", outstream.String())
	}
}

func TestDomerauEditorDistanceTransition(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauEditorDistance(outstream, "bca", "dac")
	if res != 2 {
		t.Error("Expected 2, found ", res)
	}

	if transitionStringsMatrixOutput != outstream.String() {
		t.Error("Wrong matrix!\nExpected:\n", transitionStringsMatrixOutput, "Found:\n", outstream.String())
	}
}

func TestDomerauRecursiveEditorDistanceAllEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauRecursiveEditorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found", res)
	}
}

func TestDomerauRecursiveEditorDistanceEmpty(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauRecursiveEditorDistance(outstream, "a", "")
	if res != 1 {
		t.Error("Expected 1, found ", res)
	}

	outstream = bytes.NewBufferString("")
	res = domerauRecursiveEditorDistance(outstream, "", "")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}
}

func TestDomerauRecursiveEditorDistanceEqual(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauRecursiveEditorDistance(outstream, "Hello", "Hello")
	if res != 0 {
		t.Error("Expected 0, found ", res)
	}
}

func TestDomerauRecursiveEditorDistanceKirillic(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauRecursiveEditorDistance(outstream, "скат", "кот")
	if res != 2 {
		t.Error("Expected 2, found ", res)
	}
}

func TestDomerauRecursiveEditorDistanceTransition(t *testing.T) {
	outstream := bytes.NewBufferString("")
	res := domerauRecursiveEditorDistance(outstream, "bca", "dac")
	if res != 2 {
		t.Error("Expected 2, found ", res)
	}
}
