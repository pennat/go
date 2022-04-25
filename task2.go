package main

import "fmt"

func main() {
	var num int

	fmt.Println("Введите порядковый номер: ")
	fmt.Scanln(&num)

	cache := map[int]int{
		1: 0,
		2: 1,
	}
	fmt.Println("Число Фибоначчи: ", fibcache(num, cache))
}

func fibcache(num int, cache map[int]int) int {
	if num == 1 {
		return 0
	} else if num == 2 {
		return 1
	}
	val, ok := cache[num]
	if ok {
		return val
	}

	cache[num] = fibcache(num-1, cache) + fibcache(num-2, cache)
	return cache[num]
}
