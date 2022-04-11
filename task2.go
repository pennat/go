package main

import (
	"fmt"
	"math"
)

func main() {
	var sqr float64
	pi := math.Pi

	fmt.Println("Введите площадь круга")
	fmt.Scanln(&sqr)

	d := math.Sqrt(sqr / pi)

	fmt.Println("Диаметр окружности:", d)

	fmt.Println("Радиус окружности:", d/2)
}
