// internal/core/neuralnet.go
package core

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"go-neuro/internal/models"
)

type NeuralNetworkService struct {
	network *models.NeuralNetwork
	mu      sync.RWMutex
}

// calculateHiddenLayerSize вычисляет размер скрытого слоя
func calculateHiddenLayerSize(inputSize, outputSize int) int {
	size1 := inputSize*2 - 1
	size2 := int(math.Ceil(float64(inputSize)*2/3 + float64(outputSize)))
	return int(math.Min(float64(size1), float64(size2)))
}

// createLayer создает слой нейронов
func createLayer(layerType models.LayerType, size int) (*models.Layer, error) {
	if size <= 0 {
		return nil, &models.NetworkError{
			Message:   "layer size must be positive",
			Operation: "createLayer",
		}
	}

	layer := &models.Layer{
		Neurons: make([]*models.Neuron, size),
		Type:    layerType,
	}

	for i := 0; i < size; i++ {
		layer.Neurons[i] = &models.Neuron{
			Bias:    rand.Float64()*2 - 1,
			Value:   0,
			Delta:   0,
			Inputs:  make([]*models.Synapse, 0),
			Outputs: make([]*models.Synapse, 0),
		}
	}

	return layer, nil
}

// createLayers создает все слои сети
func createLayers(config *models.NetworkConfig) ([]*models.Layer, error) {
	layers := make([]*models.Layer, config.HiddenLayersCount+2)

	// Создаем входной слой
	inputLayer, err := createLayer(models.InputLayer, config.InputSize)
	if err != nil {
		return nil, err
	}
	layers[0] = inputLayer

	// Создаем скрытые слои
	for i := 0; i < config.HiddenLayersCount; i++ {
		hiddenLayer, err := createLayer(models.HiddenLayer, config.HiddenLayerSize)
		if err != nil {
			return nil, err
		}
		layers[i+1] = hiddenLayer
	}

	// Создаем выходной слой
	outputLayer, err := createLayer(models.OutputLayer, config.OutputSize)
	if err != nil {
		return nil, err
	}
	layers[config.HiddenLayersCount+1] = outputLayer

	return layers, nil
}

// connectTwoLayers связывает два соседних слоя
func connectTwoLayers(prevLayer, currLayer *models.Layer) error {
	for _, currNeuron := range currLayer.Neurons {
		for _, prevNeuron := range prevLayer.Neurons {
			synapse := &models.Synapse{
				Weight: rand.Float64()*2 - 1,
				Source: prevNeuron,
				Target: currNeuron,
			}

			prevNeuron.Outputs = append(prevNeuron.Outputs, synapse)
			currNeuron.Inputs = append(currNeuron.Inputs, synapse)
		}
	}
	return nil
}

// connectLayers связывает слои сети между собой
func connectLayers(nn *models.NeuralNetwork) error {
	for i := 1; i < len(nn.Layers); i++ {
		currentLayer := nn.Layers[i]
		previousLayer := nn.Layers[i-1]

		if err := connectTwoLayers(previousLayer, currentLayer); err != nil {
			return err
		}
	}
	return nil
}

func NewNeuralNetworkService(config *models.NetworkConfig) (*NeuralNetworkService, error) {
	rand.Seed(time.Now().UnixNano())

	if config.InputSize <= 0 || config.OutputSize <= 0 {
		return nil, &models.NetworkError{
			Message:   "input and output sizes must be positive",
			Operation: "NewNeuralNetwork",
		}
	}

	if config.HiddenLayerSize <= 0 {
		config.HiddenLayerSize = calculateHiddenLayerSize(config.InputSize, config.OutputSize)
	}

	nn := &models.NeuralNetwork{
		Config: config,
		Layers: make([]*models.Layer, 0),
	}

	// Создаем слои
	layers, err := createLayers(config)
	if err != nil {
		return nil, err
	}
	nn.Layers = layers

	// Связываем слои
	if err := connectLayers(nn); err != nil {
		return nil, err
	}

	return &NeuralNetworkService{network: nn}, nil
}

// GetNetwork возвращает внутреннюю модель сети (для persistence)
func (s *NeuralNetworkService) GetNetwork() *models.NeuralNetwork {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.network
}

// SetNetwork устанавливает внутреннюю модель сети (для persistence)
func (s *NeuralNetworkService) SetNetwork(nn *models.NeuralNetwork) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.network = nn
}
