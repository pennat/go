package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Println("Введите первое число:")
	fmt.Scanln(&a)

	fmt.Println("Введите второе число:")
	fmt.Scanln(&b)

	fmt.Println("Введите арифметическую операцию (+, -, *, /, !, ^): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
		if b == 0 {
			fmt.Println("Делить на ноль нельзя")
			os.Exit(1)
		}
	case "!":
		res = factorial(a)
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func factorial(a float64) float64 {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
