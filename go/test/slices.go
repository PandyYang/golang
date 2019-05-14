package main

import "fmt"

//使用切片就可以不使用指针 直接改变底层数组的值
//起到一个视图的作用
func updateslice(arr []int){
	arr[0] = 100
}

func main() {
	arr2 := [...]int{8,8,8,8,8}
	fmt.Println(arr2)
	updateslice(arr2[:3])
	fmt.Println(arr2)
	fmt.Println(arr2)

	arr := [...]int{7,8,9,55,6,2,3,54,5,6,25,6,2}
	fmt.Println("arr[2:6]",arr[2:6]) //左开右闭
	fmt.Println("arr[2:]",arr[2:])//从第二个元素开始所有元素
	fmt.Println("arr[:6]",arr[:6])//第六个元素之前 不含第六个元素
	fmt.Println("arr[:]",arr[:])//所有的元素

	arr3 := [...]int{0,1,2,3,4,5,6,7,8}
	s1 := arr3[2:6]
	s2 := s1[3:5]
	fmt.Println("s1----------------------------")
	fmt.Println(s1) //2 3 4 5
	fmt.Println("s2----------------------------")
	fmt.Println(s2)// 5 6

	//slice是可以向后扩展的 不能向后扩展 s[i]无法超越len(s)  向后扩展不能超越cap
	fmt.Printf("s1=%v,s1.len = %d,s1.cap = %d\n",s1,len(s1),cap(s1)); //s1=[2 3 4 5],s1.len = 4,s1.cap = 7
	fmt.Printf("s2=%v,s2.len = %d,s2.cap = %d\n",s2,len(s2),cap(s2)); //s2=[5 6],s2.len = 2,s2.cap = 4

	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)

	//s4 and s5 no longer arr
	fmt.Println(s1,s2,s3,s4,s5,arr3) // arr3现在没有12???

}
