package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Создание таблицы
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS characters (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		constellations TEXT NOT NULL,
		weapon_refinement TEXT NOT NULL,
		artifact_set TEXT NOT NULL,
		crit_rate REAL NOT NULL,
		crit_damage REAL NOT NULL,
		cv REAL NOT NULL,
		mastery_bonus REAL DEFAULT 0,
		attack INTEGER NOT NULL,
		attack_bonus REAL NOT NULL,
		total_damage REAL NOT NULL
	);`
	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}
	fmt.Println("Таблица успешно создана!")
}

// Вставка данных персонажа
func insertCharacter(db *sql.DB, name, constellations, weaponRefinement, artifactSet string, critRate, critDamage, cv, masteryBonus float64, attack int, attackBonus, totalDamage float64) error {
	query := `
	INSERT INTO characters (name, constellations, weapon_refinement, artifact_set, crit_rate, crit_damage, cv, mastery_bonus, attack, attack_bonus, total_damage)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	_, err := db.Exec(query, name, constellations, weaponRefinement, artifactSet, critRate, critDamage, cv, masteryBonus, attack, attackBonus, totalDamage)
	return err
}

// Получение всех персонажей
func getAllCharacters(db *sql.DB) {
	query := `SELECT id, name, crit_rate, crit_damage, cv FROM characters ORDER BY cv DESC;`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Ошибка при получении данных: %v", err)
	}
	defer rows.Close()

	fmt.Println("Персонажи:")
	for rows.Next() {
		var id int
		var name string
		var critRate, critDamage, cv float64
		if err := rows.Scan(&id, &name, &critRate, &critDamage, &cv); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s - CR: %.2f, CD: %.2f, CV: %.2f\n", id, name, critRate, critDamage, cv)
	}
}

//// Обновление CV для персонажа
//func updateCharacterCV(db *sql.DB, id int, newCV float64) error {
//	query := `UPDATE characters SET cv = ? WHERE id = ?;`
//	_, err := db.Exec(query, newCV, id)
//	return err
//}

// Вычисление урона
func calculateDamage(critRate, critDamage, attack, masteryBonus, attackBonus float64) float64 {
	baseDamage := attack * (1 + attackBonus/100)
	critMultiplier := 1 + (critRate / 100 * critDamage / 100)
	return baseDamage * critMultiplier * (1 + masteryBonus/100)
}

func main() {
	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "./genshin_characters.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создаем таблицу
	createTable(db)

	// Пример данных персонажей
	characters := []struct {
		name, constellations, weaponRefinement, artifactSet string
		critRate, critDamage, cv, masteryBonus              float64
		attack                                              int
		attackBonus, totalDamage                            float64
	}{
		{"Нахида", "C0", "R5", "4", 66.4, 227.1, 244.8, 15.0, 16914, 105.8, 0},
		{"Эола", "C2", "R1", "4", 67.6, 193.3, 230.1, 104.0, 19507, 105.8, 0},
		{"Мона", "C4", "R1", "2", 79.7, 196.1, 229.3, 109.3, 1818, 163.7, 0},
		{"Аяка", "C0", "R1", "4", 78.5, 236.1, 228.6, 27.0, 19450, 116.8, 0},
		{"Гань Юй", "C1", "R1", "4", 51.3, 224.4, 228.5, 46.6, 2366, 105.2, 0},
		{"Мавуика", "C0", "R1", "4", 62.2, 200.3, 226.3, 46.6, 2017, 106.5, 0},
		{"Арлекино", "C0", "R5", "4", 82.2, 231.4, 223.9, 46.6, 1775, 111.7, 0},
		{"Навия", "C1", "R5", "4", 76.9, 222.1, 222.3, 46.6, 20519, 104.5, 0},
		{"Аль-Хайтам", "C1", "R1", "4", 59.2, 259.5, 221.6, 75.4, 1355, 118.8, 0},
		{"Райдэн", "C0", "R1", "4", 60.6, 158.0, 219.2, 65.9, 2261, 264.8, 0},
		{"Джинн", "C3", "R1", "2", 80.0, 172.0, 216.9, 61.6, 1653, 120.1, 0},
		{"Фишль", "C6", "R1", "4", 75.0, 160.6, 213.8, 46.6, 1646, 124.0, 0},
		{"Ёимия", "C2", "R1", "4", 82.1, 213.3, 213.0, 46.6, 2314, 110.4, 0},
		{"Дилюк", "C4", "R1", "2", 70.1, 170.5, 212.2, 0, 18768, 135.0, 0},
		{"Яэ Мико", "C0", "R5", "4", 63.9, 237.3, 211.5, 46.6, 1744, 110.4, 0},
		{"Кэ Цин", "C4", "R5", "2", 55.2, 195.7, 207.6, 61.6, 1836, 116.2, 0},
		{"Е Лань", "C1", "R1", "4", 89.1, 215.9, 207.6, 46.6, 30254, 131.0, 0},
		{"Тигнари", "C0", "R1", "4", 57.4, 211.5, 206.0, 75.4, 1337, 105.2, 0},
		{"Сян Лин", "C6", "R5", "4", 55.9, 153.4, 205.3, 46.6, 1240, 232.6, 0},
		{"Ху Тао", "C0", "R5", "4", 55.9, 187.1, 200.5, 61.6, 28775, 104.5, 0},
	}

	// Вставка данных персонажей
	for _, char := range characters {
		char.totalDamage = calculateDamage(char.critRate, char.critDamage, float64(char.attack), char.masteryBonus, char.attackBonus)
		err := insertCharacter(db, char.name, char.constellations, char.weaponRefinement, char.artifactSet, char.critRate, char.critDamage, char.cv, char.masteryBonus, char.attack, char.attackBonus, char.totalDamage)
		if err != nil {
			log.Printf("Ошибка вставки персонажа %s: %v", char.name, err)
		}
	}

	// Чтение данных из таблицы
	getAllCharacters(db)
}
