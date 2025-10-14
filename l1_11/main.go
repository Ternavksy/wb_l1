package main

import "fmt"

func intersectionMap(a, b []int) []int {
	// Создаем мапу для быстрого поиска
	setA := make(map[int]bool)

	// Заполняем мапу элементами первого множества
	for _, elem := range a {
		setA[elem] = true
	}

	// Ищем общие элементы
	var result []int
	for _, elem := range b {
		if setA[elem] {
			result = append(result, elem)
			delete(setA, elem) // исключаем дубликаты
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	result := intersectionMap(A, B)
	fmt.Printf("Пересечение: %v\n", result)
}
