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

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("У меня %d %s!\n", i, getFriendWord(i))
	}
}
