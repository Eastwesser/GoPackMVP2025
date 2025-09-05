// ALWAYS EMPTY FILE FOR TEST RUNS
package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Найти максимальный элемент в слайсе

// С числами
func maxIntSliceElement(inter []int) int {
	var result int // это то самое большое число из слайса, которое мы насчитали - 12314

	// нам надо в цикле сравнивать элемент с самым высоким.
	// И когда никого выше не найдется, а цикл закончится, это число и будет победителем
	for i := 0; i < len(inter); i++ {
		if inter[i] < inter[i] {
			return inter[i]
		}
	}

	return result
}

// Со строками
func maxStringSliceElement(stringer []string) string {
	var victoriousLine strings.Builder

	// нам надо в цикле сравнивать элемент с самым высоким количеством букв в слове.
	// И когда никого выше не найдется, а цикл закончится, это слово и будет победителем
	for i, r := range stringer {
		if len(stringer[i]) < len(stringer[i]) {
			if unicode.IsLetter(rune(r[0])) {
				victoriousLine.WriteRune(rune(r[0]))
			}
		}
	}

	return victoriousLine.String()
}

func main() {
	integers := maxIntSliceElement([]int{1, 2, 33, 41, 12314, 42})
	words := maxStringSliceElement([]string{"aaaaddddddbaba", "ababssas", "ababaaa", "ababaa", "abasba"})

	fmt.Println(integers)
	fmt.Println(words)
}
