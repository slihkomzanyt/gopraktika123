package main

import (
	"errors"
	"fmt"
)

// создаем слайз в функции
func findSecondMax(slice []int) (int, error) {
	if len(slice) < 2 {
		return 0, errors.New("что то не так")
	}
	//задаем две переменные макс1 и макс2 принимающие значения слайса
	// если макс1 больше макс2, то сравниваем слайс2 со слайсом 3
	max1, max2 := slice[0], slice[1]
	if max2 > max1 {
		max1, max2 = max2, max1
	}

	// Проходим по остальным элементам
	//для этого задаем переменную i которая каждый раз увеличивается на 1
	//создаем переменную курент для простоты сравнения и передаем в нее переменную i
	for i := 2; i < len(slice); i++ {
		current := slice[i]

		if current > max1 {
			max2 = max1
			max1 = current
		} else if current > max2 && current != max1 {
			max2 = current
		}
	}

	return max2, nil
}

func main() {
	slice := []int{3, 5, 7, 1, 9}
	result, err := findSecondMax(slice)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Second max:", result)
	}
}
