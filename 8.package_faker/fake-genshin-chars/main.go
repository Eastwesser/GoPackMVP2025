package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6" // Библиотека для генерации фейковых данных
	"reflect"                         // Для работы с рефлексией (анализ структуры)
	"strings"                         // Для работы со строками
)

// Определение структуры персонажа Genshin Impact
type Character struct {
	ID      int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"` // Уникальный ID персонажа
	Name    string `db_field:"name" db_type:"VARCHAR(100)"`     // Имя персонажа
	Element string `db_field:"element" db_type:"VARCHAR(50)"`   // Элемент персонажа (например, Pyro, Hydro)
	Rarity  int    `db_field:"rarity" db_type:"INT"`            // Редкость персонажа (например, 4 или 5 звезд)
	Region  string `db_field:"region" db_type:"VARCHAR(100)"`   // Регион персонажа (например, Mondstadt, Liyue)
}

// Интерфейс Tabler для получения имени таблицы
type Tabler interface {
	TableName() string
}

// Реализация интерфейса Tabler для структуры Character
func (c Character) TableName() string {
	return "characters" // Имя таблицы в базе данных
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string  // Генерация SQL для создания таблицы
	CreateInsertSQL(model Tabler) string // Генерация SQL для вставки данных
}

// Интерфейс для генерации фейковых данных о персонажах
type FakeDataGenerator interface {
	GenerateFakeCharacter() Character // Генерация фейкового персонажа
}

// Реализация SQLGenerator для SQLite
type SQLiteGenerator struct{}

// Метод для генерации SQL-запроса создания таблицы
func (g SQLiteGenerator) CreateTableSQL(table Tabler) string {
	var columns []string

	// Используем рефлексию для анализа структуры
	v := reflect.ValueOf(table)
	if v.Kind() == reflect.Ptr { // Если передали указатель, разыменовываем
		v = v.Elem()
	}
	t := v.Type()

	// Проходим по всем полям структуры
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		dbField := field.Tag.Get("db_field")                             // Получаем значение тега db_field
		dbType := field.Tag.Get("db_type")                               // Получаем значение тега db_type
		columns = append(columns, fmt.Sprintf("%s %s", dbField, dbType)) // Формируем строку "поле тип"
	}

	// Собираем SQL-запрос
	result := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (\n  %s\n);", // Добавляем переносы строк для читаемости
		table.TableName(),
		strings.Join(columns, ",\n  "), // Разделяем столбцы запятыми и переносами строк
	)

	return result
}

// Метод для генерации SQL-запроса вставки данных
func (g SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	var fields []string // Поля таблицы
	var values []string // Значения для вставки

	// Используем рефлексию для анализа структуры
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr { // Если передали указатель, разыменовываем
		v = v.Elem()
	}
	t := v.Type()

	// Проходим по всем полям структуры
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		dbField := field.Tag.Get("db_field")                                 // Получаем значение тега db_field
		fields = append(fields, dbField)                                     // Добавляем имя поля
		values = append(values, fmt.Sprintf("'%v'", v.Field(i).Interface())) // Добавляем значение поля
	}

	// Собираем SQL-запрос
	result := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		model.TableName(),
		strings.Join(fields, ", "), // Поля через запятую
		strings.Join(values, ", "), // Значения через запятую
	)
	return result
}

// Реализация FakeDataGenerator для генерации фейковых персонажей
type GenshinFakeitGenerator struct{}

// Метод для генерации фейкового персонажа
func (g GenshinFakeitGenerator) GenerateFakeCharacter() Character {
	elements := []string{"Pyro", "Hydro", "Anemo", "Electro", "Geo", "Cryo", "Dendro"}
	regions := []string{"Mondstadt", "Liyue", "Inazuma", "Sumeru", "Fontaine", "Natlan", "Snezhnaya"}

	return Character{
		ID:      gofakeit.Number(1, 1000),                      // Случайный ID
		Name:    gofakeit.FirstName(),                          // Случайное имя
		Element: elements[gofakeit.Number(0, len(elements)-1)], // Случайный элемент
		Rarity:  gofakeit.Number(4, 5),                         // Редкость (4 или 5 звезд)
		Region:  regions[gofakeit.Number(0, len(regions)-1)],   // Случайный регион
	}
}

func main() {
	sqlGenerator := SQLiteGenerator{}             // Создаем генератор SQL-запросов
	fakeDataGenerator := GenshinFakeitGenerator{} // Создаем генератор фейковых данных

	// Генерация SQL-запроса для создания таблицы
	character := Character{}
	sql := sqlGenerator.CreateTableSQL(character)
	fmt.Println("SQL для создания таблицы:")
	fmt.Println(sql)

	// Генерация и вставка фейковых данных
	fmt.Println("\nSQL для вставки данных:")
	for i := 0; i < 5; i++ { // Генерируем 5 персонажей
		fakeCharacter := fakeDataGenerator.GenerateFakeCharacter()
		query := sqlGenerator.CreateInsertSQL(fakeCharacter)
		fmt.Println(query)
	}
}
