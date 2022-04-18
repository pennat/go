package main

import (
	"fmt"
	"math"
)

func main() {
	var sqr float64

	fmt.Println("Введите площадь круга")
	fmt.Scanln(&sqr)

	d := 2 * math.Sqrt(sqr / math.Pi)

	fmt.Println("Диаметр окружности:", d)

	fmt.Println("Длина окружности:", math.Pi*d)
}
