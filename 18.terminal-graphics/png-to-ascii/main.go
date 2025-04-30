package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

var asciiChars = "@%#*+=-:. "

func pixelToASCII(c color.Color) byte {
	r, g, b, _ := c.RGBA()
	gray := uint8((r + g + b) / 3 >> 8)
	scale := float64(gray) / 255.0
	index := int(scale * float64(len(asciiChars)-1))
	return asciiChars[index]
}

func convertImageToASCII(path string, width uint) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Преобразуем до квадратной ширины (например, 100 символов)
	resized := resize.Resize(width, width/2, img, resize.Lanczos3)
	bounds := resized.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			fmt.Print(string(pixelToASCII(resized.At(x, y))))
		}
		fmt.Println()
	}

	return nil
}

func main() {
	// Меняй имя файла здесь
	imagePath := filepath.Join("pics", "ascii.png")

	err := convertImageToASCII(imagePath, 100) // ширина в символах
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
