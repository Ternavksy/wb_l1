package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, x := range nums {
			ch1 <- x
		}
		close(ch1)
	}()

	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2)
	}()

	for res := range ch2 {
		fmt.Println(res)
	}
}
