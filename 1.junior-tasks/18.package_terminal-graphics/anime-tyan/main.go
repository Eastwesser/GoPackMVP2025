package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"image"
	_ "image/png"
	"os"
)

func main() {
	imagePath := "anime.png"

	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Используем готовый конвертер с настройками
	converter := convert.NewImageConverter()

	options := convert.DefaultOptions
	options.FixedWidth = 70
	options.FixedHeight = 35
	options.Colored = true
	options.Reversed = false

	// Конвертируем и выводим
	asciiArt := converter.Image2ASCIIString(img, &options)
	fmt.Print(asciiArt)
}
