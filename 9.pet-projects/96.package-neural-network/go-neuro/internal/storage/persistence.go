// internal/storage/persistence.go
package storage

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"

	"go-neuro/internal/core"
	"go-neuro/internal/models"
)

type NeuralNetworkStorage struct {
	mu sync.Mutex
}

func NewNeuralNetworkStorage() *NeuralNetworkStorage {
	return &NeuralNetworkStorage{}
}

// Save сохраняет нейронную сеть в файл
func (s *NeuralNetworkStorage) Save(nn *models.NeuralNetwork, filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Создаем безопасную копию для сериализации
	networkCopy := &models.NeuralNetwork{
		Config: &models.NetworkConfig{
			InputSize:         nn.Config.InputSize,
			OutputSize:        nn.Config.OutputSize,
			HiddenLayersCount: nn.Config.HiddenLayersCount,
			HiddenLayerSize:   nn.Config.HiddenLayerSize,
			LearningRate:      nn.Config.LearningRate,
			Activation:        nn.Config.Activation,
		},
		Layers: make([]*models.Layer, len(nn.Layers)),
	}

	// Копируем слои и нейроны (без мьютексов и циклических ссылок)
	for i, layer := range nn.Layers {
		networkCopy.Layers[i] = &models.Layer{
			Type:    layer.Type,
			Neurons: make([]*models.Neuron, len(layer.Neurons)),
		}

		for j, neuron := range layer.Neurons {
			neuronCopy := &models.Neuron{
				Bias:  neuron.Bias,
				Value: neuron.Value,
				Delta: neuron.Delta,
				// Inputs и Outputs не сериализуем для избежания циклических ссылок
				Inputs:  nil,
				Outputs: nil,
			}
			networkCopy.Layers[i].Neurons[j] = neuronCopy
		}
	}

	data, err := json.MarshalIndent(networkCopy, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal network: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Load загружает нейронную сеть из файла
func (s *NeuralNetworkStorage) Load(filename string) (*models.NeuralNetwork, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var nn models.NeuralNetwork
	if err := json.Unmarshal(data, &nn); err != nil {
		return nil, fmt.Errorf("failed to unmarshal network: %w", err)
	}

	// После загрузки нужно восстановить связи между нейронами
	if err := s.reconstructNetworkConnections(&nn); err != nil {
		return nil, fmt.Errorf("failed to reconstruct network connections: %w", err)
	}

	return &nn, nil
}

// reconstructNetworkConnections восстанавливает связи после загрузки
func (s *NeuralNetworkStorage) reconstructNetworkConnections(nn *models.NeuralNetwork) error {
	// Восстанавливаем связи между слоями
	for i := 1; i < len(nn.Layers); i++ {
		currentLayer := nn.Layers[i]
		previousLayer := nn.Layers[i-1]

		for _, currNeuron := range currentLayer.Neurons {
			currNeuron.Inputs = make([]*models.Synapse, 0)
			currNeuron.Outputs = make([]*models.Synapse, 0)

			for _, prevNeuron := range previousLayer.Neurons {
				// Создаем синапс с случайным весом (при загрузке веса будут обновляться при обучении)
				synapse := &models.Synapse{
					Weight: rand.Float64()*2 - 1,
					Source: prevNeuron,
					Target: currNeuron,
				}

				prevNeuron.Outputs = append(prevNeuron.Outputs, synapse)
				currNeuron.Inputs = append(currNeuron.Inputs, synapse)
			}
		}
	}

	return nil
}

// SaveWithWeights сохраняет сеть с весами (более сложная реализация)
func (s *NeuralNetworkStorage) SaveWithWeights(nn *models.NeuralNetwork, filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	type SerializableNetwork struct {
		Config  *models.NetworkConfig `json:"config"`
		Weights [][][]float64         `json:"weights"` // [layer][neuron][weight]
		Biases  [][]float64           `json:"biases"`  // [layer][neuron]
	}

	serializable := &SerializableNetwork{
		Config:  nn.Config,
		Weights: make([][][]float64, len(nn.Layers)),
		Biases:  make([][]float64, len(nn.Layers)),
	}

	// Сохраняем только веса и смещения
	for i, layer := range nn.Layers {
		serializable.Biases[i] = make([]float64, len(layer.Neurons))
		serializable.Weights[i] = make([][]float64, len(layer.Neurons))

		for j, neuron := range layer.Neurons {
			serializable.Biases[i][j] = neuron.Bias

			if i > 0 { // У входного слоя нет входных весов
				serializable.Weights[i][j] = make([]float64, len(neuron.Inputs))
				for k, input := range neuron.Inputs {
					serializable.Weights[i][j][k] = input.Weight
				}
			}
		}
	}

	data, err := json.MarshalIndent(serializable, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal network: %w", err)
	}

	return os.WriteFile(filename, data, 0644)
}

// LoadWithWeights загружает сеть с весами
func (s *NeuralNetworkStorage) LoadWithWeights(filename string) (*models.NeuralNetwork, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var serializable struct {
		Config  *models.NetworkConfig `json:"config"`
		Weights [][][]float64         `json:"weights"`
		Biases  [][]float64           `json:"biases"`
	}

	if err := json.Unmarshal(data, &serializable); err != nil {
		return nil, fmt.Errorf("failed to unmarshal network: %w", err)
	}

	// Создаем новую сеть с теми же параметрами
	service, err := core.NewNeuralNetworkService(serializable.Config)
	if err != nil {
		return nil, err
	}

	nn := service.GetNetwork()

	// Восстанавливаем веса и смещения
	for i, layer := range nn.Layers {
		for j, neuron := range layer.Neurons {
			if i < len(serializable.Biases) && j < len(serializable.Biases[i]) {
				neuron.Bias = serializable.Biases[i][j]
			}

			if i > 0 && i < len(serializable.Weights) &&
				j < len(serializable.Weights[i]) &&
				len(neuron.Inputs) == len(serializable.Weights[i][j]) {
				for k := range neuron.Inputs {
					neuron.Inputs[k].Weight = serializable.Weights[i][j][k]
				}
			}
		}
	}

	return nn, nil
}
