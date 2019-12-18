package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{"张三", 20}
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)
	fmt.Println(v)
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)
	fmt.Println("=======================")
	u1 := v.Interface().(User)
	fmt.Println(u1)
	fmt.Println("=======================")
	fmt.Println(t.Kind()) //获取类型底层类型
	fmt.Println("=======================")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}
	fmt.Println("=======================")
	x := 2
	v1 := reflect.ValueOf(&x)
	v1.Elem().SetInt(100)
	fmt.Println(x)
}

type User struct {
	Name string
	Age  int
}

func (u User) Login() {
	fmt.Printf("login in")
}
