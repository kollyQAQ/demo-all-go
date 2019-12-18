package main

import (
	"fmt"
	"unsafe"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/9
*/
func main() {

	// map 类型实际上是一个指针类型
	var m map[int]int
	var p uintptr
	fmt.Println(unsafe.Sizeof(m), unsafe.Sizeof(p)) // 8 8

	// slice、struct 是引用类型
	// 引用类型之所以可以引用，是因为我们创建引用类型的变量，其实是一个标头值，
	// 标头值里包含一个指针，指向底层的数据结构，当我们在函数中传递引用类型时，其实传递的是这个标头值的副本
	var sli []int
	var ch chan string
	var fn func(i int) int
	var inter interface {
		Do() string
	}
	fmt.Println(unsafe.Sizeof(sli), unsafe.Sizeof(ch), unsafe.Sizeof(fn), unsafe.Sizeof(inter), unsafe.Sizeof(p)) // 24 8

	fmt.Println("========================================================")

	dict := map[string]int{"key1": 1, "key2": 2, "key3": 3, "key4": 4}
	fmt.Printf("%v, %p, %p\n", dict, dict, &dict) //0xc42007a1b0, 0xc42000c030
	ChangeMap(dict)
	fmt.Printf("%v, %p, %p\n", dict, dict, &dict) //0xc42007a1b0, 0xc42000c030
}

func ChangeMap(dict map[string]int) {
	dict["key1"] = 100
	fmt.Printf("%v, %p, %p\n", dict, dict, &dict) //0xc42007a1b0, 0xc42000c038
}
