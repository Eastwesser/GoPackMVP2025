package main

import (
	"errors"
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

func findFirstPNG(dir string) (string, error) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".png" {
			return filepath.Join(dir, f.Name()), nil
		}
	}

	return "", errors.New("no PNG files found")
}

func convertImageToASCII(path string, width uint) error {

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

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

	imagePath, err := findFirstPNG("pics")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = convertImageToASCII(imagePath, 100)
	if err != nil {
		fmt.Println("Ошибка при конвертации:", err)
	}
}
