package base

// func main() {
// 	const filename = "abc.txt"
// 	//结果返回两个值byte[] err
// /*	contents, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		fmt.Print(err)
// 	} else {
// 		fmt.Printf("%s\n",contents)
// 	}*/

// 	if contents,err := ioutil.ReadFile(filename); err != nil{
// 		fmt.Println(err)
// 	}else {
// 		fmt.Printf("%s\n",contents)
// 	}
// 	fmt.Println(eval(3,5,"*"))
// }

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}
