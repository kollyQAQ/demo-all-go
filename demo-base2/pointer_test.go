package test

import (
	"fmt"
	"testing"
)

func TestPoint1(t *testing.T) {
	a := 1
	println(&a) // 0xc00008e1a8
	var b *int
	b = &a
	var c *int
	c = &a
	println(b)  // 0xc00008e1a8
	println(&b) // 0xc000086020
	println(*b) // 1
	println(c)  // 0xc00008e1a8
	println(&c) // 0xc000086028
	println(*c) // 1
}

/*
https://studygolang.com/topics/3682
	切片是基于数组的。可以把切片简单的理解为数组中的一段(快照)，它保留了指向数组中某个元素（可以不是第一个元素）的指针，以及它本身的长度（也就是切片的大小）。
	var s=make([]int,0,10)这条语句创建了一个长度为10的数组，并在该数组之上建立了一个切片s，该切片保留指向数组第一个元素的指针，以及它的长度（0）。
	t = append(s,1,2,3,4) 这条语句向 s 指向的数组中追加了 4 个元素（注意这里 s 并没有变化，它的长度仍然是0，仅仅是底层数组变了）并返回新的切片（指向数组的第一个元素，长度为4）赋值给 t 。
	现在 s 和 t 都指向同一个数组的第一个元素，因此他们的地址是相同的。而 s 的长度是0，t 的长度是4. 可以参考go的文档：go doc builtin.make 以及 go doc builtin.append
*/
func TestPoint2(t *testing.T) {
	slice1 := []int64{1, 2, 3}
	slice2 := slice1
	println(slice1)            // [3/3]0xc000018220
	println(slice2)            // [3/3]0xc000018220，slice1和slice2指向同一块内存区域
	slice2 = append(slice2, 4) // append 后发生扩容，slice2指向新分配的内存区域，但是 append 不一定导致扩容
	println(slice2)            // [3/3]0xc00001a1b0
	p1 := &slice1
	p2 := &slice2
	fmt.Printf("%p %p %v %p\n", p1, *p1, *p1, &p1) // 指针p1指向的地址，指针p1指向的地址的内容，指针p1本身的地址
	fmt.Printf("%p %p %v %p\n", p2, *p2, *p2, &p2)
}
