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
		for x := 0; x < i; x++ {
			if arr[x] > arr[i] {
				arr[x], arr[i] = arr[i], arr[x]
			}
		}
	}
}
