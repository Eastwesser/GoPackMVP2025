package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/ascii"
	"image"
	_ "image/png" // Для поддержки PNG
	"os"
)

func main() {
	// Указываем путь к изображению
	imagePath := "GoPackMVP2025/18.terminal-graphics/anime-tyan/anime.png"

	// Открываем изображение
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Декодируем изображение
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка при декодировании изображения:", err)
		return
	}

	// Создаем конвертер
	converter := ascii.NewPixelConverter()

	// Настройки для конвертации
	options := ascii.NewOptions()
	options.Pixels = []byte(" .,:;i1tfLCG08@")
	options.Reversed = false
	options.Colored = true

	// Проходим по каждому пикселю изображения
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y)
			asciiChar := converter.ConvertPixelToASCII(pixel, &options)
			fmt.Print(asciiChar)
		}
		fmt.Println()
	}
}
