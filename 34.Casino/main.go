package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Символы для игры
const (
	SymbolCherry  = "🍒"
	SymbolBar     = "🍫"
	SymbolDiamond = "💎"
	SymbolSeven   = "7"
	SymbolWild    = "🌟"
)

// Коэффициенты выигрыша для каждого символа
var symbolPayouts = map[string]int{
	SymbolCherry:  5,
	SymbolBar:     10,
	SymbolDiamond: 15,
	SymbolSeven:   50,
	SymbolWild:    0, // Wild не имеет собственного коэффициента
}

// Игровое поле: 5 колонок, в каждой 3 строки
type SlotMachine [5][3]string

// Генерация случайного символа
func getRandomSymbol() string {
	symbols := []string{SymbolCherry, SymbolBar, SymbolDiamond, SymbolSeven, SymbolWild}
	return symbols[rand.Intn(len(symbols))]
}

// Создание игрового поля
func generateSlotMachine() SlotMachine {
	var slot SlotMachine
	rand.Seed(time.Now().UnixNano())
	for col := 0; col < 5; col++ {
		for row := 0; row < 3; row++ {
			slot[col][row] = getRandomSymbol()
		}
	}
	return slot
}

// Вывод игрового поля
func (s SlotMachine) Display() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("%s ", s[col][row])
		}
		fmt.Println()
	}
}

// Проверка выигрышной комбинации
func checkWinningLine(slot SlotMachine) int {
	totalWin := 0

	// Проверяем среднюю горизонтальную линию (можно добавить другие линии)
	line := [5]string{slot[0][1], slot[1][1], slot[2][1], slot[3][1], slot[4][1]}
	fmt.Println("Проверяемая линия:", line)

	// Ищем комбинации
	symbol := ""
	count := 0
	for _, s := range line {
		if s == SymbolWild {
			count++ // Wild считается как часть комбинации
		} else if symbol == "" && s != SymbolWild {
			symbol = s // Первый не-Wild символ
			count++
		} else if s == symbol || s == SymbolWild {
			count++
		} else {
			break // Комбинация прервана
		}
	}

	// Если найдена комбинация из 3 или более символов
	if count >= 3 {
		payout := symbolPayouts[symbol]
		if payout > 0 {
			totalWin += payout * count
			fmt.Printf("Выигрышная комбинация: %d x %s! Выигрыш: %d\n", count, symbol, payout*count)
		}
	}

	return totalWin
}

func main() {
	// Генерация игрового поля
	slot := generateSlotMachine()
	fmt.Println("Игровое поле:")
	slot.Display()

	// Проверка выигрыша
	totalWin := checkWinningLine(slot)
	if totalWin > 0 {
		fmt.Printf("Общий выигрыш: %d\n", totalWin)
	} else {
		fmt.Println("Повезет в следующий раз, приятель!")
	}
}
