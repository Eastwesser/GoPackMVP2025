package main

import (
	"context"
	"errors"
	"fmt"
)

// GPU представляет данные о видеокарте
type GPU struct {
	Model       string // Модель видеокарты
	ClockSpeed  string // Тактовая частота
	MemorySpeed string // Скорость памяти
}

// GPURepositoryOld — это старый интерфейс, который нужно адаптировать.
type GPURepositoryOld interface {
	FindByModel(ctx context.Context, model string) (*GPU, error)
}

// GPUService — это новый интерфейс, который ожидает клиент.
type GPUService interface {
	GetGPU(ctx context.Context, model string) (*GPU, error)
}

// GPURepositoryAdapter — это адаптер, который делает GPURepositoryOld совместимым с GPUService.
type GPURepositoryAdapter struct {
	repo GPURepositoryOld // Адаптируемый объект (GPURepositoryOld)
}

// GetGPU — метод, который реализует интерфейс GPUService.
func (gra *GPURepositoryAdapter) GetGPU(ctx context.Context, model string) (*GPU, error) {
	gpu, err := gra.repo.FindByModel(ctx, model)
	if err != nil {
		return nil, err
	}
	return gpu, nil
}

// =====================================================================================================================

// MockGPURepository — моковая реализация GPURepositoryOld для тестирования.
type MockGPURepository struct{}

func (mgr *MockGPURepository) FindByModel(ctx context.Context, model string) (*GPU, error) {
	if model == "RTX 3080" {
		return &GPU{
			Model:       "RTX 3080",
			ClockSpeed:  "1800 MHz",
			MemorySpeed: "14 Gbps",
		}, nil
	}
	return nil, errors.New("GPU not found")
}

// =====================================================================================================================

// Пример использования
func main() {
	// Создаем моковый GPURepository
	mockRepo := &MockGPURepository{}

	// Создаем адаптер
	adapter := &GPURepositoryAdapter{
		repo: mockRepo,
	}

	// Используем GPUService через адаптер
	gpu, err := adapter.GetGPU(context.Background(), "RTX 3080")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("GPU Details:")
	fmt.Printf("Model: %s\n", gpu.Model)
	fmt.Printf("Clock Speed: %s\n", gpu.ClockSpeed)
	fmt.Printf("Memory Speed: %s\n", gpu.MemorySpeed)
}
