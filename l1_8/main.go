package main

import (
	"fmt"
)

// функция setBit устанавливает i-й бит числа num в значение value (0 или 1)
func setBit(num int64, i uint, value int) int64 {
	if value == 1 {
		// Установка бита в 1 с помощью побитового ИЛИ
		return num | (1 << i)
	}
	// Установка бита в 0 с помощью побитового И НЕ
	return num &^ (1 << i)
}

func main() {
	var num int64 = 5 // 0101 в двоичной системе
	var i uint = 1    // позиция бита
	var value int = 0 // значение для установки (0 или 1)

	result := setBit(num, i, value)
	fmt.Printf("Число %d (%b) после установки %d-го бита в %d = %d (%b)\n",
		num, num, i, value, result, result)
}
