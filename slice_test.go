package goutils

import (
	"testing"
)

type testStructure struct {
	idx  int
	name string
}

func TestSliceDeleteByIndex(t *testing.T) {

	var rawSlice = []testStructure{testStructure{
		idx:  1,
		name: "1",
	}, {
		idx:  2,
		name: "2",
	}, {
		idx:  3,
		name: "3",
	}, {
		idx:  4,
		name: "4",
	}, {
		idx:  5,
		name: "5",
	}}
	t.Log("raw slice")
	t.Log(rawSlice)

	Slice_DeleteByIndex(&rawSlice, 3)

	t.Log("after delete index=2 element")
	t.Log(rawSlice)
}
