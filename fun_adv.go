package main

import (
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("hello, %s", name)
}

func test(name string, f func(string)) {
	f(name)
}

func main() {
	test("pandy", sayHello)
}
