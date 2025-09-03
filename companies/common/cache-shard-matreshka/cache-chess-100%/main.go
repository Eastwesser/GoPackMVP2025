package main

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Poem struct {
	XMLName xml.Name `xml:"poem"`
	Title   string   `xml:"title"`
	Stanzas []Stanza `xml:"stanza"`
}

type Stanza struct {
	XMLName xml.Name `xml:"stanza"`
	Lines   []string `xml:"line"`
}

type LineMessage struct {
	Index int
	Line  string
}

func parsePoemFromXML(filename string) (string, []string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read poem file: %w", err)
	}

	var poem Poem
	err = xml.Unmarshal(data, &poem)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse XML: %w", err)
	}

	var allLines []string
	for _, stanza := range poem.Stanzas {
		allLines = append(allLines, stanza.Lines...)
	}

	return strings.Join(allLines, "\n"), allLines, nil
}

func main() {
	// Парсим стихотворение
	originalPoem, allLines, err := parsePoemFromXML("poem.xml")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	// Статистика
	totalAttempts := 0
	successfulAttempts := 0
	startTime := time.Now()

	// Каналы для communication
	lineChan := make(chan LineMessage, len(allLines))
	verdictChan := make(chan bool)
	stopChan := make(chan bool)

	// Запускаем Короля-судью
	go func() {
		collectedLines := make([]string, len(allLines))

		for {
			select {
			case <-stopChan:
				return
			case msg := <-lineChan:
				collectedLines[msg.Index] = msg.Line

				// Проверяем, все ли строки собраны
				complete := true
				for _, line := range collectedLines {
					if line == "" {
						complete = false
						break
					}
				}

				if complete {
					assembledPoem := strings.Join(collectedLines, "\n")
					verdictChan <- (assembledPoem == originalPoem)

					// Сбрасываем для следующей попытки
					collectedLines = make([]string, len(allLines))
				}
			}
		}
	}()

	// Запускаем горутины-поставщики строк
	var wg sync.WaitGroup
	for i := 0; i < len(allLines); i++ {
		wg.Add(1)
		go func(lineIndex int) {
			defer wg.Done()

			rand.Seed(time.Now().UnixNano())

			for {
				select {
				case <-stopChan:
					return
				default:
					// Имитируем "поиск" строки (рандомная задержка)
					searchTime := time.Duration(rand.Intn(100)) * time.Millisecond
					time.Sleep(searchTime)

					// "Находим" строку и отправляем Королю
					lineChan <- LineMessage{
						Index: lineIndex,
						Line:  allLines[lineIndex],
					}
				}
			}
		}(i)
	}

	// Главный процесс - собираем статистику
	fmt.Printf("Начало поэтической рулетки! Всего строк: %d\n", len(allLines))
	fmt.Printf("Оригинальное стихотворение (%d символов):\n", len(originalPoem))
	fmt.Println("========================================")

	for verdict := range verdictChan {
		totalAttempts++

		if verdict {
			successfulAttempts++
			fmt.Printf("🎉 УСПЕХ! Попытка #%d\n", totalAttempts)
		} else {
			fmt.Printf("❌ ПРОВАЛ! Попытка #%d\n", totalAttempts)
		}

		// Выводим статистику каждые 10 попыток
		if totalAttempts%10 == 0 {
			successRate := float64(successfulAttempts) / float64(totalAttempts) * 100
			fmt.Printf("📊 Статистика: %d/%d (%.1f%%) успешных попыток\n",
				successfulAttempts, totalAttempts, successRate)
		}

		// Останавливаем после 100 попыток или по Ctrl+C
		if totalAttempts >= 100 {
			close(stopChan)
			break
		}
	}

	wg.Wait()
	close(lineChan)
	close(verdictChan)

	// Финальная статистика
	duration := time.Since(startTime)
	successRate := float64(successfulAttempts) / float64(totalAttempts) * 100

	fmt.Println("\n========================================")
	fmt.Println("🏁 ПОЭТИЧЕСКАЯ РУЛЕТКА ЗАВЕРШЕНА!")
	fmt.Println("========================================")
	fmt.Printf("Общее время: %v\n", duration)
	fmt.Printf("Всего попыток сборки: %d\n", totalAttempts)
	fmt.Printf("Успешных попыток: %d\n", successfulAttempts)
	fmt.Printf("Процент успеха: %.1f%%\n", successRate)
	fmt.Printf("Среднее время на попытку: %v\n", duration/time.Duration(totalAttempts))

	if successfulAttempts > 0 {
		fmt.Printf("Шанс собрать стихотворение: 1 к %.0f\n",
			float64(totalAttempts)/float64(successfulAttempts))
	}
}
