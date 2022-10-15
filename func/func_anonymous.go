package _func

import (
	"fmt"
	"math"
)

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

	getSqrt := func(a float64) float64 { return math.Sqrt(a) }
	fmt.Print(getSqrt(2))

	fns := []func(x int) int{
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	println(fns[0](100))

	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "Hello World"
		},
	}

	println(d.fn)

	fc := make(chan func() string, 2)

	fc <- func() string {
		return "Hello World"
	}

	println((<-fc)())
}
