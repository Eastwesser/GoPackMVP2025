package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	// URL страницы, с которой будем скачивать видео
	url := "https://knower-sculpt.ru/intro#tlection=885140707_24" // Замени на нужный URL

	// Создаем коллектор
	c := colly.NewCollector()

	// Слайс для хранения ссылок на видео
	var videoLinks []string

	// Ищем все теги <video> и <a> с ссылками на видео
	c.OnHTML("video source[src], a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if link == "" {
			link = e.Attr("href")
		}
		// Проверяем, что ссылка ведет на видео
		if strings.HasSuffix(link, ".mp4") || strings.HasSuffix(link, ".webm") || strings.HasSuffix(link, ".ogg") {
			videoLinks = append(videoLinks, link)
		}
	})

	// Обрабатываем ошибки
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Ошибка:", err)
	})

	// Запускаем парсинг
	c.Visit(url)

	// Скачиваем все найденные видео
	for i, link := range videoLinks {
		fmt.Printf("Скачиваем видео %d/%d: %s\n", i+1, len(videoLinks), link)
		err := downloadFile(link, fmt.Sprintf("video_%d%s", i+1, filepath.Ext(link)))
		if err != nil {
			fmt.Printf("Ошибка при скачивании %s: %v\n", link, err)
		}
	}

	fmt.Println("Скачивание завершено!")
}

// Функция для скачивания файла
func downloadFile(url, filename string) error {
	// Создаем файл для сохранения
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Выполняем HTTP-запрос
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Копируем данные из ответа в файл
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
