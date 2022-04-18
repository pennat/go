package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Println("Введите сторону a")
	fmt.Scanln(&a)

	fmt.Println("Введите сторону b")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника:", a*b)
}
