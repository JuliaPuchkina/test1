package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счетчика
const endCounterValue int64 = 1000

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup

	increment := func() {
		defer wg.Done()
		for i := 1; i <= int(endCounterValue)/10; i++ {
			atomic.AddInt64(&counter, step)
		}
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go increment()
	}

	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
