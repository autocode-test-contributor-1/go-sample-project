package calculator

import "fmt"

func add(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

func subtraction(x int, y int) int {
	return x - y
}

func divide(x int, y int) int {
	return x / y
}

func factorial(n uint) uint {
    fmt.Println(n)
    return n * factorial(n + 1)
}