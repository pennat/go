package main

import "fmt"

func main() {
	var num int

	fmt.Println("Введите порядковый номер: ")
	fmt.Scanln(&num)

	fmt.Println("Число Фибоначчи: ", fibnum(num))
}

func fibnum(num int) int {
	if num == 1 {
		return 0
	} else if num == 2 || num == 3 {
		return 1
	}
	return fibnum(num-1) + fibnum(num-2)
}
