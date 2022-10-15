package base

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%d,len=%d,cap=%d\n", s, len(s), cap(s))
}

// func main() {
// 	var s [] int //zero value for slice is nil
// 	for i := 0; i<10; i++ {
// 		//printSlice(s)
// 		s = append(s, i*2)
// 		//printSlice(s)
// 		//fmt.Println(s)
// 	}

// 	s1 := []int{2,4,6,8}
// 	printSlice(s1)

// 	s2 := make([]int,16)//长度
// 	s3 := make([]int,10,32)//长度和容量
// 	printSlice(s2)

// 	printSlice(s3)

// 	copy(s2,s1)//将s1中的元素copy到s2中
// 	printSlice(s2)

// 	//进行元素的删除
// 	s2 = append(s2[:3],s2[4:]...)
// 	printSlice(s2)

// 	front := s2[0] //首元素
// 	s2 = s2[1:]//将第一个元素截取

// 	tail := s2[len(s2)-1]//最后一个元素
// 	s2 = s2[:len(s2) - 1]
// 	fmt.Println(front,tail)

// 	//截取首尾的字母字后的元素
// 	printSlice(s2)

// }
