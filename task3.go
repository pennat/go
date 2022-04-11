package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&num)

	var a int = num % 10
	fmt.Println("Количество единиц в числе:", a)

	var b int = (num / 10) % 10
	fmt.Println("Количество десяков в числе:", b)

	var c int = (num / 100) % 10
	fmt.Println("Количество сотен в числе:", c)
}
