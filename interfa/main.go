package main

import "fmt"

type Phone interface {
	speak()
	read()
}

type IPhone struct {
	name string
}

func (a *IPhone) read() {
	fmt.Println("read apple")
}

type Oppo struct {
	id int
}

type Mi struct {
	f bool
}

func (a *IPhone) speak() {
	fmt.Println("hi apple")
}

func (a Oppo) speak() {
	fmt.Println("hi vivo")
}

func (a Mi) speak() {
	fmt.Println("hi dami")
}

func show(myPhone Phone) {
	myPhone.speak()
}

func main() {
	show(new(IPhone))
}
