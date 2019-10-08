package main

import (
	"fmt"
	"unsafe"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/4
*/

func main() {
	// 值类型：基础数据类型(int、float、bool和string)、数组类型、结构体类型
	i := 100
	j := i
	fmt.Printf("i=%v,j=%v,&i=%v,&j=%v\n", i, j, &i, &j)
	arr1 := [4]int64{1, 2, 3, 4}
	arr2 := arr1
	fmt.Printf("arr1=%v,arr2=%v,&arr1=%v,&arr2=%v\n", arr1, arr2, unsafe.Pointer(&arr1[0]), unsafe.Pointer(&arr2[0]))
	s1 := struct {
		age  int
		name string
	}{}
	s2 := s1
	fmt.Printf("s1=%v,s2=%v,&s1=%v,&s2=%v\n", s1, s2, unsafe.Pointer(&s1), unsafe.Pointer(&s2))
	// 引用类型：切片类型(slice)、字典类型(map)、通道类型(chan)、函数类型(func)
	sli1 := []int{1, 2, 3, 4}
	sli2 := sli1
	fmt.Printf("sli1=%v,sli2=%v,&sli1=%v,&sli2=%v", sli1, sli2, unsafe.Pointer(&sli1[0]), unsafe.Pointer(&sli2[0]))
}
