package string

import "testing"

func TestIsChineseChar(t *testing.T) {
	s := "wo我是shi"
	println(IsChineseChar(s))
}

func TestReg(t *testing.T) {
	println(Reg("123", "\\d+"))
}
