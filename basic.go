package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内变量 外面不能写 :=

var (
	sdsd = 1
	dw   = "ss"
)

var aa = 3
var ss = "ddd"
var nn = true

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValie() {
	var a int = 4
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "def"
	var s = "abx"
	fmt.Println(a, b, c, d, s)
}

func variableShorter() {
	a, b, c, d := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}

// func main() {
// 	fmt.Println("Hello World")
// 	variableZeroValue()
// 	variableInitialValie()
// 	variableTypeDeduction()
// 	variableShorter()
// 	euler()
// 	tiangle()
// 	consts()
// 	enums()
// }

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	//e的i Pi次方+1 前面是底数 后面是指数
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	//以下可以输出0
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func tiangle() {
	var a, b int = 3, 4
	var c int
	//类型不一致时 必须进行强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//定义常量
func consts() {
	const filename = "abx.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举类型
func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	//使用简便方式 iota  输出0 1 2
	const (
		x = iota
		y
		z
	)
	//使用iota作为种子进行一些逻辑运算
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(x, y, z)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
