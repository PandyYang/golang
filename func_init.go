package main

import "fmt"

// init is a special function, always run ahead main func
func init() {
	println("init...")
}

func main() {
	fmt.Println("main...")
}
