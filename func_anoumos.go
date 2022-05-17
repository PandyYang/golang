package main

import "fmt"

func main() {
	max := func(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	i := max(1, 2)

	fmt.Print(i)
}
