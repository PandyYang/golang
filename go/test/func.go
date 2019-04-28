package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//swith可以设置多个满足条件
func eval1(a, b int, op string) (int, error) {
	switch op {
	case  "+":
		return a + b, nil
	case  "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a,b)//可以使用 _ 代替无输入
		return q, nil
	default:
		return 0 ,fmt.Errorf("unsupported operation: %s",op)
	}
}

//函数可以有两个返回值 可以起名 仅适用于简单的函数 对于调用者而言是没有区别的
func div(a, b int)(q, r int){
	return a/b , a%b
}


func apply(op func(int, int) int, a, b int) int{
	//获取当前函数的指针
	p := reflect.ValueOf(op).Pointer()
	//根据指针获取当前运行的函数名
	opName := runtime.FuncForPC(p)
	//打印出相关的函数名 以及两个参数
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval1(5,8,"+"))

	q, r := div(30,4)
	fmt.Println(q, r)
	if result, err := eval1(3,4,"x"); err != nil{
		fmt.Println("error:",err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(apply(pow, 3 ,4))
	//在可以直接定义匿名函数
	fmt.Println(apply(
		func(a int, b int) int {
			return int (math.Pow(
				float64(a),float64(b)))
		}, 3, 4))


	fmt.Println(sum(1, 2, 3, 4, 5))

}
