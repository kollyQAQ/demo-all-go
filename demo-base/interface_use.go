package main

import "fmt"

type Animal interface {
	Move() string
	Eat() string
}

type Bird struct {
}

func (b Bird) Move() string {
	return "Bird fly"
}

func (b Bird) Eat() string {
	return "Bird eat"
}

type Dog struct {
}

func (Dog) Move() string {
	return "Dog Run"
}

func (Dog) Eat() string {
	return "Dog eat"
}

func AnimalPlay(animal Animal) {
	fmt.Println(animal.Move())
	fmt.Println(animal.Eat())
}

func main() {
	bird := Bird{}
	AnimalPlay(bird)
	dog := Dog{}
	AnimalPlay(dog)
}
