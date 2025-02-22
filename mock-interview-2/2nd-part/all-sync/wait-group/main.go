package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	// Путь к директории с данными о персонажах
	characterDataDir := "./character_data" // Замените на реальный путь к директории

	// Чтение списка файлов в директории
	files, err := ioutil.ReadDir(characterDataDir)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(files)) // Устанавливаем количество задач

	// Обработка каждого файла в отдельной горутине
	for _, file := range files {
		go func(f os.FileInfo) {
			defer wg.Done() // Уменьшаем счетчик при завершении горутины

			// Выводим имя файла
			fmt.Printf("Найден файл с данными персонажа: %s\n", f.Name())

			// Читаем содержимое файла
			filePath := fmt.Sprintf("%s/%s", characterDataDir, f.Name())
			data, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Ошибка при чтении файла %s: %v\n", f.Name(), err)
				return
			}
			fmt.Printf("Данные персонажа %s: %s\n", f.Name(), string(data))
		}(file)
	}

	wg.Wait() // Блокируем выполнение до завершения всех горутин
	fmt.Println("Все файлы обработаны.")
}
