package main

import "fmt"

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
	}, max)

	for i := 0; i < max; i++ {
		go func(n int) {
			word := getFriendWord(n + 1)
			ch <- struct {
				index int
				word  string
			}{n, word}
		}(i)
	}

	for i := 0; i < max; i++ {
		r := <-ch
		result[r.index] = r.word
	}
	return result
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("У меня %d %s!\n", i, getFriendWord(i))
	}
}
