package demo_benchmark

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/1

-benchmem 可以提供每次操作分配内存的次数，以及每次操作分配的字节数

go test -bench=. -run=none
go test -bench=. -benchmem -run=none
*/

const numbers = 100

func TestNormal(b *testing.T) {
	var s string
	for i := 0; i < numbers; i++ {
		s = fmt.Sprintf("%v%v", s, i)
	}
	fmt.Println(s)
}

func BenchmarkSpringf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var sb strings.Builder
		for i := 0; i < numbers; i++ {
			sb.WriteString(strconv.Itoa(i))
		}
		_ = sb.String()
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var buffer bytes.Buffer
		for i := 0; i < numbers; i++ {
			buffer.WriteString(strconv.Itoa(i))
		}
		_ = buffer.String()
	}
	b.StopTimer()
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for idx := 0; idx < b.N; idx++ {
		var s string
		for i := 0; i < numbers; i++ {
			s += strconv.Itoa(i)
		}
	}
	b.StopTimer()
}
