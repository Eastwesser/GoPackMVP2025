package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var ready bool
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	wg.Add(2)

	go func() {
		defer wg.Done()
		mutex.Lock()
		for !ready {
			cond.Wait()
		}
		mutex.Unlock()
		fmt.Println("Горутина 1: Работа завершена")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		mutex.Lock()
		ready = true
		cond.Signal()
		mutex.Unlock()
		fmt.Println("Горутина 2: Сигнал отправлен")
	}()

	wg.Wait()
}
