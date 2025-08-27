// cmd/main.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"go-neuro/internal/models"
	"go-neuro/pkg/neuro"
)

func testNetwork(kurosawa *neuro.Kurosawa, inputs, targets [][]float64) {
	for i, testCase := range inputs {
		result, err := kurosawa.Predict(testCase)
		if err != nil {
			fmt.Printf("Error predicting: %v\n", err)
			continue
		}

		expected := targets[i][0]
		actual := result[0]
		accuracy := 1.0 - math.Abs(expected-actual)

		fmt.Printf("Input: %v, Output: %.4f (expected: %v, accuracy: %.1f%%)\n",
			testCase, actual, expected, accuracy*100)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("🎬 Kurosawa Neural Network Initializing...")

	config := &models.NetworkConfig{
		InputSize:         2,
		OutputSize:        1,
		HiddenLayersCount: 1,
		LearningRate:      0.5,
		Activation:        "sigmoid",
	}

	// Создание нейронной сети Kurosawa
	kurosawa, err := neuro.NewKurosawa(config)
	if err != nil {
		fmt.Printf("Error creating Kurosawa network: %v\n", err)
		return
	}

	// Данные для обучения
	inputs := [][]float64{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	targets := [][]float64{{0}, {1}, {1}, {1}}

	// Обучение сети
	fmt.Println("🎬 Kurosawa is training...")
	if err := kurosawa.Train(inputs, targets, 1000); err != nil {
		fmt.Printf("Error training Kurosawa: %v\n", err)
		return
	}

	// Сохранение модели
	fmt.Println("💾 Saving Kurosawa model...")
	if err := kurosawa.Save("kurosawa_model.json"); err != nil {
		fmt.Printf("Error saving model: %v\n", err)
	} else {
		fmt.Println("✅ Model saved successfully!")
	}

	// Тестирование сети
	fmt.Println("\n🎬 Kurosawa is testing...")
	testNetwork(kurosawa, inputs, targets)

	// Загрузка модели (демонстрация)
	fmt.Println("\n🔄 Loading Kurosawa model...")
	kurosawa2, err := neuro.NewKurosawa(config)
	if err != nil {
		fmt.Printf("Error creating new Kurosawa: %v\n", err)
		return
	}

	if err := kurosawa2.Load("kurosawa_model.json"); err != nil {
		fmt.Printf("Error loading model: %v\n", err)
		return
	}
	fmt.Println("✅ Model loaded successfully!")

	// Тестирование загруженной модели
	fmt.Println("\n🎬 Loaded Kurosawa is testing...")
	testNetwork(kurosawa2, inputs, targets)

	// Очистка
	if err := os.Remove("kurosawa_model.json"); err != nil {
		fmt.Printf("Note: Could not clean up test file: %v\n", err)
	}

	fmt.Println("\n🎬 Kurosawa training complete! Ready for action!")
}
