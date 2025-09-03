package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}

type Poem struct {
	XMLName xml.Name `xml:"poem"`
	Title   string   `xml:"title"`
	Stanzas []Stanza `xml:"stanza"`
}

type Stanza struct {
	XMLName xml.Name `xml:"stanza"`
	Lines   []string `xml:"line"`
}

type InMemoryCache struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]string),
	}
}

func parsePoemFromXML(filename string) (string, []string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", nil, fmt.Errorf("failed to read poem file %s: %w", filename, err)
	}
	var poem Poem
	err = xml.Unmarshal(data, &poem)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse poem file %s: %w", filename, err)
	}

	var allLines []string
	for _, stanza := range poem.Stanzas {
		allLines = append(allLines, stanza.Lines...)
	}

	completePoem := strings.Join(allLines, "\n")

	return completePoem, allLines, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *InMemoryCache) Set(k string, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[k] = v
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[k]

	return val, ok
}

//func (c *InMemoryCache) GetPoem(keys []string) string {
//	c.mu.RLock()
//	defer c.mu.RUnlock()
//
//	var b strings.Builder
//
//	//for _, key := range keys { // почему не так?
//	//	if v, ok := c.data[key]; ok {
//	//		b.WriteString(v)
//	//		b.WriteRune('\n')
//	//	}
//	//}
//
//	for i, key := range keys {
//		if v, ok := c.data[key]; ok {
//			b.WriteString(v)
//			if i < len(keys)-1 {
//				b.WriteRune('\n')
//			}
//		}
//	}
//
//	return b.String()
//}

func (c *InMemoryCache) GetPoem(keys []string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	lines := make([]string, len(keys))
	for i, key := range keys {
		if v, ok := c.data[key]; ok {
			lines[i] = v
		}
	}
	return strings.Join(lines, "\n")
}

func main() {
	var resultPoem string
	tries := 0
	timer := time.Now()

	originalPoem, lines, err := parsePoemFromXML("poem.xml")
	if err != nil {
		fmt.Printf("An error occured in main: %v\n", err)
		return
	}

	cache := NewInMemoryCache()

	keys := make([]string, len(lines))
	for i := range keys {
		keys[i] = fmt.Sprintf("line%d", i)
	}

	fmt.Printf("Ожидаемая длина: %d символов\n", len(originalPoem))
	fmt.Printf("Ожидаемый текст в hex:\n")
	for i, r := range originalPoem {
		if i < 50 { // Покажем первые 50 символов в hex
			fmt.Printf("%02x ", r)
		}
	}
	fmt.Println()

	success := false
	for {
		tries++

		if tries%100 == 0 {
			fmt.Printf("Сборка №%d...\n", tries)
		}

		var wg sync.WaitGroup
		wg.Add(len(lines))

		for i, line := range lines {
			go func(k, v string) {
				defer wg.Done()
				cache.Set(k, v)
				time.Sleep(time.Millisecond * time.Duration(10))
			}(keys[i], line)
		}
		wg.Wait()

		resultPoem = cache.GetPoem(keys)

		if resultPoem != originalPoem && tries%100 == 0 {
			fmt.Printf("Длина результата: %d символов\n", len(resultPoem))

			// Сравним побайтово
			fmt.Printf("Результат в hex (первые 50):\n")
			for i := 0; i < min(50, len(resultPoem)); i++ {
				fmt.Printf("%02x ", resultPoem[i])
			}
			fmt.Println()

			fmt.Printf("Ожидаемый в hex (первые 50):\n")
			for i := 0; i < min(50, len(originalPoem)); i++ {
				fmt.Printf("%02x ", originalPoem[i])
			}
			fmt.Println()

			// Найдем первую отличающуюся позицию
			for i := 0; i < min(len(resultPoem), len(originalPoem)); i++ {
				if resultPoem[i] != originalPoem[i] {
					fmt.Printf("Первое отличие на позиции %d: %q vs %q\n",
						i, resultPoem[i], originalPoem[i])
					fmt.Printf("Hex: %02x vs %02x\n", resultPoem[i], originalPoem[i])
					break
				}
			}
		}

		if resultPoem == originalPoem {
			success = true
			break
		}

		for _, key := range keys {
			cache.Set(key, "")
		}

		// Защита от бесконечного цикла
		if tries >= 1000 {
			fmt.Println("Превышено максимальное количество попыток (1000)")
			break
		}
	}

	elapsed := time.Since(timer)

	if success {
		fmt.Printf("\n=== ПОЭЗИЯ СОБРАНА! ===\n\n")
		fmt.Println(resultPoem)
	} else {
		fmt.Printf("\n=== ПОЭЗИЯ НЕ СОБРАНА! ===\n\n")
		fmt.Println("Последняя попытка:")
		fmt.Println(resultPoem)
		fmt.Println("\nОжидалось:")
		fmt.Println(originalPoem)
	}

	fmt.Printf("\nПопыток: %d\n", tries)
	fmt.Printf("Общее время: %v\n", elapsed)
	if tries > 0 {
		fmt.Printf("Время на попытку: %v\n", time.Duration(int64(elapsed)/int64(tries)))
	}
}
