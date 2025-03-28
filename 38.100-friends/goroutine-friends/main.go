package main

import (
	"fmt"
	"runtime"
	"sync"
)

// getFriendWord возвращает правильную форму слова "друг" для числа n.
// Примеры:
//
//	1 → "друг",  2 → "друга",  5 → "друзей",  21 → "друг",  22 → "друга".
func getFriendWord(n int) string {
	if n%100 >= 11 && n%100 <= 14 {
		return "друзей"
	}
	switch n % 10 {
	case 1:
		return "друг"
	case 2, 3, 4:
		return "друга"
	default:
		return "друзей"
	}
}

func processFriendsConcurrent(max int) []string {
	result := make([]string, max)
	ch := make(chan struct {
		index int
		word  string
	}, max) // Буферизованный канал

	sem := make(chan struct{}, runtime.NumCPU()*2)
	var wg sync.WaitGroup

	for i := 0; i < max; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(n int) {
			defer func() {
				<-sem
				wg.Done()
			}()
			word := getFriendWord(n + 1)
			ch <- struct {
				index int
				word  string
			}{n, word}
		}(i)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		result[r.index] = r.word
	}
	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("У меня %d %s!\n", i, getFriendWord(i))
	}
}
