package main

import (
	"os"
	"testing"
)

// Test_calculateRevenue: Проверяет корректность расчёта выручки для различных случаев.
func Test_calculateRevenue(t *testing.T) {
	type args struct {
		visitors int
		price    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Zero visitors", args: args{visitors: 0, price: 20.5}, want: 0},
		{name: "Standard case", args: args{visitors: 100, price: 10}, want: 1000},
		{name: "Decimal price", args: args{visitors: 50, price: 15.75}, want: 787.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateRevenue(tt.args.visitors, tt.args.price); got != tt.want {
				t.Errorf("calculateRevenue() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Test_checkCapacity: Убеждается, что функция корректно обрабатывает вместимость парка.
func Test_checkCapacity(t *testing.T) {
	type args struct {
		currentVisitors int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Below capacity", args: args{currentVisitors: 800}, wantErr: false},
		{name: "At capacity", args: args{currentVisitors: 1000}, wantErr: false},
		{name: "Exceeds capacity", args: args{currentVisitors: 1200}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkCapacity(tt.args.currentVisitors); (err != nil) != tt.wantErr {
				t.Errorf("checkCapacity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test_logToFile: Создаёт файл, записывает сообщение, проверяет содержимое, затем удаляет файл.
func Test_logToFile(t *testing.T) {
	type args struct {
		filename string
		message  string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Log message", args: args{filename: "test_log.txt", message: "Test log entry"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logToFile(tt.args.filename, tt.args.message)

			// Проверка файла
			content, err := os.ReadFile(tt.args.filename)
			if err != nil {
				t.Errorf("Failed to read log file: %v", err)
			}
			if string(content) != tt.args.message {
				t.Errorf("logToFile() content = %v, want %v", string(content), tt.args.message)
			}

			// Удаление тестового файла
			_ = os.Remove(tt.args.filename)
		})
	}
}

// Test_printAttractions: Проверяет вывод для разных списков аттракционов.
func Test_printAttractions(t *testing.T) {
	type args struct {
		attractions []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "No attractions", args: args{attractions: []string{}}},
		{name: "One attraction", args: args{attractions: []string{"Roller Coaster"}}},
		{name: "Multiple attractions", args: args{attractions: []string{"Ferris Wheel", "Haunted House", "Bumper Cars"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printAttractions(tt.args.attractions)
		})
	}
}

// Test_printVisitorStats: Проверяет вывод статистики посетителей.
func Test_printVisitorStats(t *testing.T) {
	type args struct {
		stats map[string]int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Empty stats", args: args{stats: map[string]int{}}},
		{name: "Single group", args: args{stats: map[string]int{"Adults": 100}}},
		{name: "Multiple groups", args: args{stats: map[string]int{"Children": 50, "Adults": 100, "Seniors": 25}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printVisitorStats(tt.args.stats)
		})
	}
}

// Test_promoteAttraction: Тестирует изменения названия аттракциона через указатель.
func Test_promoteAttraction(t *testing.T) {
	type args struct {
		attraction *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Promote simple name", args: args{attraction: ptr("Roller Coaster")}, want: "Roller Coaster - САМЫЙ ПОПУЛЯРНЫЙ!"},
		{name: "Promote empty name", args: args{attraction: ptr("")}, want: " - САМЫЙ ПОПУЛЯРНЫЙ!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			promoteAttraction(tt.args.attraction)
			if *tt.args.attraction != tt.want {
				t.Errorf("promoteAttraction() = %v, want %v", *tt.args.attraction, tt.want)
			}
		})
	}
}

// Вспомогательная функция для получения указателя на строку
func ptr(s string) *string {
	return &s
}
