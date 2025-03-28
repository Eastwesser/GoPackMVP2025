package main

import (
	"fmt"
	"runtime"
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
	})

	// Ограничиваем количество одновременно работающих горутин
	sem := make(chan struct{}, runtime.NumCPU()*2)

	for i := 0; i < max; i++ {
		sem <- struct{}{}
		go func(n int) {
			defer func() { <-sem }()
			word := getFriendWord(n + 1)
			ch <- struct {
				index int
				word  string
			}{n, word}
		}(i)
	}

	// Дополнительные горутины для закрытия канала
	go func() {
		for i := 0; i < max; i++ {
			<-sem
		}
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
