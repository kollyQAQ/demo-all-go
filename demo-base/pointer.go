package main

import (
	"fmt"
	"testing"
	"unsafe"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/9

& 是取地址符号，即取得某个变量的地址，如；&a
* 是指针运算符，可以表示一个变量是指针类型，也可以表示一个指针变量所指向的存储单元，也就是这个地址所存储的值

经常会见到: p . *p , &p 三个符号
p 是一个指针变量的名字，表示此指针变量指向的内存地址
*p 表示此指针指向的内存地址中存放的内容
&p 就是取指针 p 的地址(指针 p 同时也是个变量，既然是变量，编译器肯定要为其分配内存地址)
*/
func main() {

	type Rect struct {
		width  float64
		height float64
	}

	// & 是取地址符号，取到 Rect 类型对象的地址
	rect := Rect{100, 100}
	fmt.Println(&rect) // &{100 100}
	// * 可以表示一个变量是指针类型 (r 是一个指针变量)
	var r *Rect = &Rect{100, 100}
	fmt.Println(r) // &{100 100}
	// * 也可以表示指针类型变量所指向的存储单元，也就是这个地址所指向的值
	fmt.Println(*r) // {100 100}
	// & 查看这个指针变量的地址，基本数据类型直接打印地址
	fmt.Println(&r) //0xc00008e018

	i := 100
	fmt.Println(&i)

	fmt.Println("========================================================")

	slice1 := []int{1, 2, 3}
	slice2 := slice1
	println(slice1)
	println(slice2)
	p1 := &slice1
	p2 := &slice2
	fmt.Printf("%p %p\n", p1, &p1)
	fmt.Printf("%p %p\n", p2, &p2)

	fmt.Println("========================================================")

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
	ChangeMap1(dict)
	fmt.Printf("%v, %p, %p\n", dict, dict, &dict) //0xc42007a1b0, 0xc42000c030
}

func ChangeMap1(dict map[string]int) {
	dict["key1"] = 100
	fmt.Printf("%v, %p, %p\n", dict, dict, &dict) //0xc42007a1b0, 0xc42000c038
}

func TestPoint(t *testing.T) {
	slice := []int64{1, 2, 3}
	println("%p", slice)
}
