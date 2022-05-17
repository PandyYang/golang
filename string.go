package main

// func main(){
// 	s := "PandyYang"//UTF-8 编码
// 	fmt.Printf("%X\n",[]byte(s))
// 	for _, b := range []byte(s){
// 		fmt.Printf("%X ",b)
// 	}
// 	fmt.Println()
// 	for i,ch := range s {
// 		fmt.Printf("(%d %X) ",i, ch)
// 	}
// 	fmt.Println()
// 	fmt.Println("Rune Count:",utf8.RuneCountInString(s))//获得字符数量

// 	bytes := []byte(s)
// 	for len(bytes) > 0 {
// 		ch, size := utf8.DecodeRune(bytes)
// 		bytes = bytes[size:]
// 		fmt.Printf("%c ",ch)
// 	}
// 	fmt.Println()

// 	for i, ch := range []rune(s){
// 		fmt.Printf("(%d %c)",i, ch)
// 	}
// 	fmt.Println()

// }
