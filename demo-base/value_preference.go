package main

import (
	"fmt"
)

/*
Author: kolly.li@klook.com
Date: 2019/10/4
*/

func main() {
	// 值类型：基础数据类型(int、float、bool和string)、数组类型、结构体类型
	i := 100
	j := i
	fmt.Printf("基础数据类型：i=%v,j=%v,&i=%p,&j=%p\n", i, j, &i, &j)
	arr1 := [4]int64{1, 2, 3, 4}
	arr2 := arr1
	fmt.Printf("数组类型：arr1=%v,arr2=%v,&arr1=%p,&arr2=%p,&arr1[0]=%p,&arr2[0]=%p\n", arr1, arr2, &arr1, &arr2, &arr1[0], &arr2[0])
	s1 := struct {
		age   int
		name  string
		score []int
	}{10, "kolly", []int{1, 2, 3}}
	s2 := s1
	fmt.Printf("结构体类型：s1=%v,s2=%v,&s1=%p,&s2=%p,&s1.score=%p,&s2.score=%p,&s1.score[0]=%p,&s2.score[0]=%p\n",
		s1, s2, &s1, &s2, &(s1.score), &(s2.score), &(s1.score[0]), &(s2.score[0]))
	fmt.Println("============================================")
	// 引用类型：切片类型(slice)、字典类型(map)、通道类型(chan)、函数类型(func)
	sli1 := []int{1, 2, 3, 4}
	sli2 := sli1
	fmt.Printf("切片类型：sli1=%v,sli2=%v,&sli1=%p,&sli2=%p,&sli1[0]=%p,&sli2[0]=%p", sli1, sli2, &sli1, &sli2, &sli1[0], &sli2[0])
}
