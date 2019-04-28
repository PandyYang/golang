package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	for ; n>0; n/=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string){
	file ,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	//省略起始与递增条件
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

//省略其实递增结束条件
func forever() {
	for {
		fmt.Println("112")
	}
}

func main() {
	fmt.Println(
		convertToBin(5), //101
		convertToBin(13), //
	)
	printFile("abc.txt")
	forever()
}
