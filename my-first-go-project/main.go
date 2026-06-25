package main

import (
	"fmt"
	"my-first-go-project/pacote"
)

func main() {
	fmt.Println("hello, world!")
	pacote.GetFoo()
	sayHello()
	fmt.Println(sum(10, 10))

	a, b := swap(10, 20)
	fmt.Println(a, b)

	res, rem := divide(10, 3)
	fmt.Println(res, rem)

	fmt.Println(sumMany(1, 2, 3, 4, 5))
}

func sayHello() {
	fmt.Println("Hello")
}

func sum(a int, b int) int {
	return a + b
}

func sumMany(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}

func swap(a, b int) (int, int) {
	return a, b
}

func divide(a int, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// go run main.go
// go build main.go -> Executable
