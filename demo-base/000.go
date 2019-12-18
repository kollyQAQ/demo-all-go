package main

import "fmt"

func main() {
	a := make([]int, 0, 4)
	a = append(a, 7)
	a = append(a, 8)
	a = append(a, 9)
	fmt.Printf("a ===> len: %d cap:%d data:%+v prt:%p\n", len(a), cap(a), a, &a)
	ap(a)
	fmt.Printf("a ===> len: %d cap:%d data:%+v prt:%p\n", len(a), cap(a), a, &a)
	a = a[0:4]
	fmt.Printf("a ===> len: %d cap:%d data:%+v prt:%p\n", len(a), cap(a), a, &a)
}

func ap(aaa []int) {
	fmt.Printf("aaa ===> len: %d cap:%d data:%+v prt:%p\n", len(aaa), cap(aaa), aaa, &aaa)
	aaa[0] = 1
	aaa = append(aaa, 10)
	fmt.Printf("aaa ===> len: %d cap:%d data:%+v prt:%p\n", len(aaa), cap(aaa), aaa, &aaa)
}
