// pkg/neuro/neuro.go
package neuro

import (
	"fmt"

	"go-neuro/internal/core"
	"go-neuro/internal/models"
	"go-neuro/internal/storage"
)

// Kurosawa представляет нейронную сеть с японским именем
type Kurosawa struct {
	service *core.NeuralNetworkService
	storage *storage.NeuralNetworkStorage
}

// NewKurosawa создает новую нейронную сеть Kurosawa
func NewKurosawa(config *models.NetworkConfig) (*Kurosawa, error) {
	service, err := core.NewNeuralNetworkService(config)
	if err != nil {
		return nil, err
	}

	return &Kurosawa{
		service: service,
		storage: storage.NewNeuralNetworkStorage(),
	}, nil
}

// Train обучает нейронную сеть
func (k *Kurosawa) Train(inputs, targets [][]float64, epochs int) error {
	return k.service.Train(inputs, targets, epochs)
}

// Predict выполняет предсказание
func (k *Kurosawa) Predict(inputs []float64) ([]float64, error) {
	return k.service.Predict(inputs)
}

// Save сохраняет модель в файл
func (k *Kurosawa) Save(filename string) error {
	network := k.service.GetNetwork()
	if err := k.storage.SaveWithWeights(network, filename); err != nil {
		return fmt.Errorf("failed to save Kurosawa model: %w", err)
	}
	return nil
}

// Load загружает модель из файла
func (k *Kurosawa) Load(filename string) error {
	network, err := k.storage.LoadWithWeights(filename)
	if err != nil {
		return fmt.Errorf("failed to load Kurosawa model: %w", err)
	}

	k.service.SetNetwork(network)
	return nil
}

// GetConfig возвращает конфигурацию сети
func (k *Kurosawa) GetConfig() *models.NetworkConfig {
	return k.service.GetNetwork().Config
}

// GetNetwork возвращает внутреннее представление сети (для продвинутого использования)
func (k *Kurosawa) GetNetwork() *models.NeuralNetwork {
	return k.service.GetNetwork()
}
