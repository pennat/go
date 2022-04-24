package main

import (
	"fmt"
)

func main() {
	var len int

	fmt.Println("Введите длину массива: ")
	fmt.Scanln(&len)

	fmt.Println("Введите элементы массива: ")

	var arr = make([]int, len)

	for i := 0; i < len; i++ {
		fmt.Scanln(&arr[i])
	}

	insertSort(arr, len)

	fmt.Println("Отсортированный массив: ")
	for i := 0; i < len; i++ {
		fmt.Println(arr[i])
	}
}

func insertSort(arr []int, len int) {
	for i := 1; i < len; i++ {
		x := i
		y := arr[i]
		for ; x > 0 && arr[x-1] > y; x-- {
			arr[x] = arr[x-1]
		}
		arr[x] = y
	}
}
