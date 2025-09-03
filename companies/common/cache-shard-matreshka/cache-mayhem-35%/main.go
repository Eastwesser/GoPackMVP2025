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
	stopChan := make(chan struct{}) // Используем пустую структуру для сигнальных каналов

	// Запускаем Короля-судью
	go func() {
		defer close(verdictChan) // Закрываем verdictChan при завершении

		for {
			select {
			case <-stopChan:
				return
			default:
				// Ждем пока соберутся ВСЕ строки для одной попытки
				collectedLines := make([]string, len(allLines))
				linesReceived := 0

				// Таймаут для одной попытки сборки
				timeout := time.After(1 * time.Second)

			attemptLines:
				for {
					select {
					case <-stopChan:
						return
					case <-timeout:
						select {
						case verdictChan <- false:
						case <-stopChan:
							return
						}
						break attemptLines

					case msg, ok := <-lineChan:
						if !ok {
							// Канал закрыт, выходим
							return
						}
						// Принимаем ЛЮБУЮ строку, но запоминаем только если она ПРАВИЛЬНАЯ для своей позиции
						if msg.Line == allLines[msg.Index] {
							if collectedLines[msg.Index] == "" {
								collectedLines[msg.Index] = msg.Line
								linesReceived++

								// Если собрали все строки - УСПЕХ!
								if linesReceived == len(allLines) {
									select {
									case verdictChan <- true:
									case <-stopChan:
										return
									}
									break attemptLines
								}
							}
						}
						// Неправильные строки просто игнорируем - даем еще шанс
					}
				}
			}
		}
	}()

	// Запускаем горутины-поставщики строк (СЛУЧАЙНЫХ!)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ { // Больше горутин для большего хаоса
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)))

			for {
				select {
				case <-stopChan:
					return
				default:
					// Случайная задержка
					time.Sleep(time.Duration(r.Intn(1000)) * time.Millisecond) // до 1 секунды!

					// Случайная строка из стихотворения!
					randomIndex := r.Intn(len(allLines))

					select {
					case <-stopChan:
						return
					case lineChan <- LineMessage{
						Index: randomIndex,
						Line:  allLines[randomIndex],
					}:
						// Успешно отправили
					case <-time.After(100 * time.Millisecond):
						// Таймаут для отправки, чтобы не блокироваться
					}
				}
			}
		}(i)
	}

	// Главный процесс - собираем статистику
	fmt.Printf("Начало НАСТОЯЩЕЙ поэтической рулетки! Всего строк: %d\n", len(allLines))
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

		// Выводим статистику
		if totalAttempts%10 == 0 {
			successRate := float64(successfulAttempts) / float64(totalAttempts) * 100
			fmt.Printf("📊 Статистика: %d/%d (%.1f%%) успешных попыток\n",
				successfulAttempts, totalAttempts, successRate)
		}

		// Останавливаем после 100 попыток
		if totalAttempts >= 100 {
			close(stopChan) // Сигнал всем горутинам остановиться
			break
		}
	}

	// Ждем завершения всех горутин
	wg.Wait()
	close(lineChan)

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

	if successfulAttempts > 0 {
		fmt.Printf("Шанс собрать стихотворение: 1 к %.0f\n",
			float64(totalAttempts)/float64(successfulAttempts))
		// Теоретический шанс: 1 к 12! (479001600) для 12 строк
		fmt.Printf("Теоретический шанс: 1 к %d (%d!)\n", factorial(len(allLines)), len(allLines))
	} else {
		fmt.Println("Не удалось собрать стихотворение ни разу!")
	}
}

// Вспомогательная функция для расчета факториала (теоретический шанс)
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
