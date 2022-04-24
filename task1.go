package main

import "fmt"

func main() {
	var num int

	fmt.Println("Введите порядковый номер: ")
	fmt.Scanln(&num)

	fmt.Println(fibnum(num))
}

func fibnum(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 || num == 2 {
		return 1
	}
	return fibnum(num-1) + fibnum(num-2)
}
