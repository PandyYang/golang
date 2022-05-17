package main

import "fmt"

//数组是值类型
func printArray(arr [5]int) {
	//以为是值类型 在方法内部会将各个数组进行拷贝
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

//但是可以使用指针来进行改变为引用传递
func printArray2(arr *[3]int) {
	arr[0] = 1000
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// func main() {

// 	fmt.Println("---------------------普通的值传递-----------------------")
// 	arr11 := [5]int{5,5,5,5,5}
// 	printArray(arr11)
// 	fmt.Println(arr11)//原始值不会被改变 因为方法中拷贝了一份
// 	fmt.Println("--------------------------------------------------------")

// 	fmt.Println("--------------使用指针进行引用传递----------------------")
// 	arr10 := [3]int{3,3,3}
// 	printArray2(&arr10)
// 	fmt.Println(arr10)//原始值会被改变 因为指针指向的具体值已经被改变
// 	fmt.Println("-------------------------------------------------------")

// 	var arr1 [5] int
// 	printArray(arr1)

// 	arr2 := [3]int{1,3,5}
// 	arr3 := [...]int{2,4,6,8,10}//... 会让编译器接管具体的变量数量
// 	printArray(arr3)

// 	//arr2参数不匹配 只有3个int数值 所以无法进行传值
// 	//printArray(arr2)

// 	var grid [4][5]int //4行5列

// 	//遍历数组方法1
// 	for i := 0; i<len(arr3); i++ {
// 		fmt.Println(arr3[i])
// 	}
// 	//遍历数组方法2
// 	for i:= range arr3{
// 		fmt.Println(arr3[i])
// 	}
// 	//使用range打印下标以及数值
// 	for i,v:=range arr3 {
// 		fmt.Println(i,v)
// 	}
// 	//使用range只打印数值
// 	for _,v := range arr3{
// 		fmt.Println(v)
// 	}
// 	//数组循环遍历 判断
// 	numbers := [6]int{1,2,3,4,5,6}
// 	maxi := -1
// 	maxValue := -1
// 	for i, v := range numbers{
// 		if v > maxValue {
// 			maxi, maxValue = i, v
// 		}
// 		fmt.Println(maxi,maxValue)
// 	}
// 	//数组求和
// 	sum := 0
// 	for _, v := range numbers {
// 		sum += v
// 	}

// 	fmt.Println(sum)
// 	fmt.Println(arr1,arr2,arr3)
// 	fmt.Println(grid)

// 	//使用range的理由
// 		//美观 简洁
// 		//意义明确

// }
