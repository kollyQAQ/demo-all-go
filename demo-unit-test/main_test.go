package demo_unit_test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/11
*/

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

func TestAdd(t *testing.T) {

	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}
	t.Log("done")

}

func Add(a, b int) int {

	return a + b
}

func TestAAA(t *testing.T) {
	a := "zzzyy"
	_, err := strconv.Atoi(a)
	if err != nil {
		println(1)
	}
	println(2)
}
