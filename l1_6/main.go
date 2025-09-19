package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Способ 1: Выход по условию (просто завершить функцию горутины)
func demoExitByCondition() {
	fmt.Println("Демонстрация: выход по условию")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Горутина %d выполняется\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Горутина выполнилась")
	}()
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 2: Через канал
func demoChannel() {
	fmt.Println("Демонстрация: выход по каналу")
	var wg sync.WaitGroup
	ch := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done() // Исправлено: wg.Done() вместо wg.Wait()
		for i := 0; ; i++ {
			select {
			case <-ch:
				fmt.Println("Горутина получила сигнал остановки через канал")
				return
			default:
				fmt.Printf("Горутина работает %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(ch)
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 3: Через контекст
func demoContext() {
	fmt.Println("Демонстрация: через контекст")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена через контекст:", ctx.Err())
				return
			default:
				fmt.Printf("Горутина работает %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 4: Через runtime.Goexit() (принудительное завершение текущей горутины)
func demoGoexit() {
	fmt.Println("Демонстрация: через runtime.Goexit()")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("Горутина работает %d\n", i)
			time.Sleep(200 * time.Millisecond)
			if i == 2 {
				fmt.Println("Горутина вызывает runtime.Goexit()")
				runtime.Goexit()
			}
		}
	}()
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 5: Через таймер (context.WithTimeout или time.After)
func demoTimer() {
	fmt.Println("Демонстрация: через таймер")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена по таймауту:", ctx.Err())
				return
			default:
				fmt.Printf("Горутина работает %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 6: Через панику (panic/recover, не рекомендуется)
func demoPanic() {
	fmt.Println("Демонстрация: через панику (не рекомендуется)")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Горутина восстановлена после паники:", r)
			}
		}()
		for i := 0; i < 5; i++ {
			fmt.Printf("Горутина работает %d\n", i)
			time.Sleep(200 * time.Millisecond)
			if i == 2 {
				panic("Принудительная остановка через панику")
			}
		}
	}()
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

// Способ 7: Грациозная остановка (graceful shutdown)
func demoGracefulShutdown() {
	fmt.Println("Демонстрация: грациозная остановка")
	done := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			select {
			case <-done:
				fmt.Println("Горутина получила сигнал остановки, завершаем текущую задачу...")
				fmt.Println("Выполняется очистка перед завершением (cleanup)")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Горутина завершила работу грациозно")
				return
			default:
				fmt.Printf("Горутина обрабатывает задачу %d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(done)
	wg.Wait()
	fmt.Println("Демонстрация завершена\n")
}

func main() {
	demoExitByCondition()
	demoChannel()
	demoContext()
	demoGoexit()
	demoTimer()
	demoPanic()
	demoGracefulShutdown()
}
