package main

// 闭包是指一个函数和与其相关的引用环境而组合的实体
// 闭包 = 函数 + 引用环境
func add2() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//func main() {
//	var f = add2()
//	println(f(10))
//	println(f(10))
//}
