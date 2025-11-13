package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 5, 9, 8, 25, 11, 10}
	sort.Ints(arr)
	fmt.Println("Отсортированный массив:", arr)
	target := 11
	fmt.Println("Индекс:", binarySearch(arr, target))
}
