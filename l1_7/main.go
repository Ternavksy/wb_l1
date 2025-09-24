package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap - конкурентно-безопасная структура с sync.Mutex
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, ok := sm.data[key]
	return value, ok
}

// SafeMapRW - конкурентно-безопасная структура с sync.RWMutex
type SafeMapRW struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMapRW() *SafeMapRW {
	return &SafeMapRW{
		data: make(map[string]int),
	}
}

func (sm *SafeMapRW) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMapRW) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, ok := sm.data[key]
	return value, ok
}

// stressTest - общий стресс-тест для проверки конкурентности
func stressTest(name string, setFunc func(string, int), getFunc func(string) (int, bool)) {
	var wg sync.WaitGroup
	start := time.Now()

	// Конкурентная запись
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%10)
			setFunc(key, i)
		}(i)
	}
	wg.Wait()

	// Конкурентное чтение
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%10)
			if value, ok := getFunc(key); ok {
				if i < 5 {
					fmt.Printf("%s: %s = %d\n", name, key, value)
				}
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("%s: Stress test completed in %v\n", name, time.Since(start))
}

func main() {
	// Тест для SafeMap
	fmt.Println("Testing SafeMap with sync.Mutex...")
	safeMap := NewSafeMap()
	stressTest("SafeMap", safeMap.Set, safeMap.Get)

	// Тест для SafeMapRW
	fmt.Println("\nTesting SafeMapRW with sync.RWMutex...")
	safeMapRW := NewSafeMapRW()
	stressTest("SafeMapRW", safeMapRW.Set, safeMapRW.Get)

	// Тест для sync.Map
	fmt.Println("\nTesting sync.Map...")
	var syncMap sync.Map
	stressTest("sync.Map",
		func(key string, value int) { syncMap.Store(key, value) },
		func(key string) (int, bool) {
			if value, ok := syncMap.Load(key); ok {
				return value.(int), true
			}
			return 0, false
		})
}
