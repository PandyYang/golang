package main

import "fmt"

func main() {

	// anonymous function
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	i := max(1, 2)

	fmt.Print(i)

	// run self
	func(a int, b int) {
		max := 0
		if a > b {
			max = a
		} else {
			max = b
		}
		fmt.Printf("max is %v\n", max)
	}(1, 2)
}
