package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := 5 * time.Second
	go func() {
		for i := 1; ; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт!")
				return
			}
			fmt.Printf("Значение получено: %d\n", val)
		case <-timer.C:
			fmt.Println("Программа завершена по таймауту")
			return
		}
	}
}
