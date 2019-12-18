package test

import (
	"testing"
)

func TestDefer(t *testing.T) {
	println(111)
	defer func() {
		println(222)
	}()
	defer func() {
		println(333)
	}()
	println(444)
}
