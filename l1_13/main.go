package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Printf("До: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("После: a = %d, b = %d\n", a, b)
}
