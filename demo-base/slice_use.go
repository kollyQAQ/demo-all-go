package main

import "fmt"

/*
Author: kolly.li@klook.com
Date: 2019/10/9
*/
func main() {

	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]

	newSlice[0] = 10

	fmt.Println(slice)
	fmt.Println(newSlice)

	fmt.Println("========================================================")

	/*
		对于底层数组容量是k的切片slice[i:j]来说
		长度：j-i
		容量:k-i
		此外还有一种 3 个索引的方法，第 3 个用来限定新切片的容量，其用法为 slice[i:j:k]
	*/
	slice1 := []int{1, 2, 3, 4, 5}
	newSlice1 := slice1[1:3]
	newSlice2 := slice[1:2:3]

	fmt.Printf("newSlice1长度:%d,容量:%d\n", len(newSlice1), cap(newSlice1))
	fmt.Printf("newSlice2长度:%d,容量:%d\n", len(newSlice2), cap(newSlice2))

	fmt.Println("========================================================")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d type=%T &s=%p %v\n", len(s), cap(s), s, &s, s)
}
