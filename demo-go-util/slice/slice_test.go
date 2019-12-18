package slice

import (
	"kolly.com/demo-all-go/demo-go-util/print"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int64{1, 2, 3, 4, 5}
	s = InsertInt64Slice(s, 2, 6)
	print.PrintSlice(s)
}

func TestFindInSlice(t *testing.T) {
	s := []int64{1, 2, 3, 4, 5}
	var a int64 = 1
	println(FindInSlice(a, s))
}
