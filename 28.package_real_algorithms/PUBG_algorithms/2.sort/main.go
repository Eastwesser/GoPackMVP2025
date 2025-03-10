package main

import (
	"fmt"
	"sort"
)

/*
2. Сортировки
Пример: Сортировка лута по качеству.

Объяснение: Игрок сортирует лут, чтобы выбрать лучшее.
*/

func main() {
	loot := []string{"Helmet Lv3", "Vest Lv1", "Backpack Lv2"}
	sort.Strings(loot)
	fmt.Println("Sorted Loot:", loot)
}
