package _struct

import "testing"

type (
	A struct {
		Age  int
		Name string
	}

	B struct {
		A
		Address string
	}
)

func (b *B) String() {
	println(b.Age, b.Name, b.Address)
}

func TestStructCopy(t *testing.T) {
	a := &A{
		Age:  11,
		Name: "zhangsan",
	}
	b := new(B)
	CopyStruct(a, b)
	b.Address = "shenzhen"
	b.String()
}
