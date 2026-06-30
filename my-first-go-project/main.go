package main

import (
	"fmt"
	// "math"
	// "my-first-go-project/pacote"
)

func main() {
	// Funções
	// fmt.Println("hello, world!")
	// pacote.GetFoo()
	// sayHello()
	// fmt.Println(sum(10, 10))

	// a, b := swap(10, 20)
	// fmt.Println(a, b)

	// res, rem := divide(10, 3)
	// fmt.Println(res, rem)

	// fmt.Println(sumMany(1, 2, 3, 4, 5))

	// Variáveis
	// var firstName, lastName string = "João", "Oliveira"
	// fmt.Println(firstName, lastName)

	// x := 10
	// fmt.Println(x)

	// Tipos
	//
	// bool
	//
	// int int8, int16, int32, int64
	// uint uint8, uint16, uint32, uint64 uintptr
	//
	// byte
	//
	// rune
	//
	// float32, float64
	//
	// complex64, complex128
	//
	// string

	// Array
	// const x = 10
	// arr := [x]int{5: 100, 9: 200} // [0 0 0 0 0 100 0 0 0 200]
	// fmt.Println(arr)

	// Loops

	// var i int
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for i, elem := range arr {
	// 	fmt.Println(i, elem)
	// }
	// for _, elem := range arr {
	// 	fmt.Println(elem)
	// }

	// for range 10 {
	// 	fmt.Println("*")
	// }

	// for i, elem := range arr {
	// 	fmt.Println(&i, &elem)
	// }

	// Ifs
	// if x := math.Sqrt(4); x < 5 { // Declaração de variável + If
	// 	fmt.Println("x is less than 5")
	// } else if x > 1 {
	// 	fmt.Println("x is greater than 1")
	// } else {
	// 	fmt.Println("x is greater than 5")
	// }
	// Switch
	do(1)
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

func do(x int) {
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}

// go run main.go
// go build main.go -> Executable
