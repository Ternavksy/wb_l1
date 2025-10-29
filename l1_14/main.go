package main

import "fmt"

func detectType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("Тип переменной int")
	case string:
		fmt.Println("Тип переменной string")
	case bool:
		fmt.Println("Тип переменной bool")
	case chan int:
		fmt.Println("Тип переменной chan int")
	default:
		fmt.Printf("Немзвестный тип: %T", t)
	}
}

func main() {
	a := 42
	b := "WB"
	c := true
	d := make(chan int)
	detectType(a)
	detectType(b)
	detectType(c)
	detectType(d)
}
