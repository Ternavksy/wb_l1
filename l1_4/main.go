package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	numWorkers := 5
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Получен сигнал прерывания")
	cancel()
	wg.Wait()
	fmt.Println("Все воркеры завершены")
}

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершен\n", id)
			return
		case <-ticker.C:
			req, err := http.NewRequestWithContext(ctx, "GET", "https://catfact.ninja/fact", nil)
			if err != nil {
				fmt.Printf("Воркер %d: ошибка создания запроса: %v\n", id, err)
				continue
			}
			req.Header.Set("Authorization", "Bearer YOUR_API_TOKEN")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Воркер %d: ошибка запроса: %v\n", id, err)
				continue
			}
			resp.Body.Close()
			fmt.Printf("Воркер %d: статус %d\n", id, resp.StatusCode)
		}
	}
}
