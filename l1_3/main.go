package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Воркер %d: %s\n", id, data)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Ошибка: укажите количество воркеров")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Ошибка: количество воркеров должно быть положительным числом")
		os.Exit(1)
	}

	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func() {
		for i := 1; ; i++ {
			ch <- fmt.Sprintf("Сообщение %d", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
