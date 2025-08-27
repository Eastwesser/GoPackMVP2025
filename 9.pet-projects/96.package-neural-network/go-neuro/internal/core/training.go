// internal/core/training.go
package core

import (
	"fmt"
	"math"

	"go-neuro/internal/models"
)

// FeedForward выполняет прямое распространение сигнала
func (s *NeuralNetworkService) FeedForward(inputs []float64) ([]float64, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(inputs) != s.network.Config.InputSize {
		return nil, &models.NetworkError{
			Message:   "input size mismatch",
			Operation: "FeedForward",
		}
	}

	// Устанавливаем значения входного слоя
	inputLayer := s.network.Layers[0]
	for i, neuron := range inputLayer.Neurons {
		neuron.Value = inputs[i]
	}

	// Распространение по скрытым и выходному слоям
	for i := 1; i < len(s.network.Layers); i++ {
		layer := s.network.Layers[i]
		for _, neuron := range layer.Neurons {
			if err := s.calculateNeuronValue(neuron); err != nil {
				return nil, err
			}
		}
	}

	// Собираем выходные значения
	outputLayer := s.network.Layers[len(s.network.Layers)-1]
	outputs := make([]float64, len(outputLayer.Neurons))
	for i, neuron := range outputLayer.Neurons {
		outputs[i] = neuron.Value
	}

	return outputs, nil
}

// calculateNeuronValue вычисляет значение нейрона
func (s *NeuralNetworkService) calculateNeuronValue(neuron *models.Neuron) error {
	// Используем прямое обращение к полям, так как mu не экспортируется
	sum := neuron.Bias
	for _, input := range neuron.Inputs {
		if input.Source == nil {
			return &models.NetworkError{
				Message:   "nil source neuron in synapse",
				Operation: "calculateValue",
			}
		}
		sum += input.Source.Value * input.Weight
	}

	// Функция активации
	neuron.Value = s.activationFunction(sum)
	return nil
}

// activationFunction возвращает функцию активации
func (s *NeuralNetworkService) activationFunction(x float64) float64 {
	switch s.network.Config.Activation {
	case "relu":
		return math.Max(0, x)
	case "tanh":
		return math.Tanh(x)
	default: // sigmoid
		return 1.0 / (1.0 + math.Exp(-x))
	}
}

// activationDerivative возвращает производную функции активации
func (s *NeuralNetworkService) activationDerivative(x float64) float64 {
	switch s.network.Config.Activation {
	case "relu":
		if x > 0 {
			return 1
		}
		return 0
	case "tanh":
		return 1 - math.Pow(math.Tanh(x), 2)
	default: // sigmoid
		return x * (1.0 - x)
	}
}

// Backpropagate выполняет обратное распространение ошибки
func (s *NeuralNetworkService) Backpropagate(targets []float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(targets) != s.network.Config.OutputSize {
		return &models.NetworkError{
			Message:   "target size mismatch",
			Operation: "Backpropagate",
		}
	}

	outputLayer := s.network.Layers[len(s.network.Layers)-1]

	// Вычисляем ошибку для выходного слоя
	for i, neuron := range outputLayer.Neurons {
		neuron.Delta = (targets[i] - neuron.Value) * s.activationDerivative(neuron.Value)
	}

	// Распространяем ошибку назад по скрытым слоям
	for i := len(s.network.Layers) - 2; i >= 0; i-- {
		layer := s.network.Layers[i]
		for _, neuron := range layer.Neurons {
			var errorSum float64
			for _, output := range neuron.Outputs {
				errorSum += output.Weight * output.Target.Delta
			}
			neuron.Delta = errorSum * s.activationDerivative(neuron.Value)
		}
	}

	// Обновляем веса и смещения
	for i := 1; i < len(s.network.Layers); i++ {
		layer := s.network.Layers[i]
		for _, neuron := range layer.Neurons {
			neuron.Bias += neuron.Delta * s.network.Config.LearningRate
			for _, input := range neuron.Inputs {
				input.Weight += input.Source.Value * neuron.Delta * s.network.Config.LearningRate
			}
		}
	}

	return nil
}

// Train обучает нейронную сеть на наборе данных
func (s *NeuralNetworkService) Train(inputs, targets [][]float64, epochs int) error {
	if len(inputs) != len(targets) {
		return &models.NetworkError{
			Message:   "inputs and targets must have same length",
			Operation: "Train",
		}
	}

	for epoch := 0; epoch < epochs; epoch++ {
		totalError := 0.0

		for i := range inputs {
			// Прямое распространение
			outputs, err := s.FeedForward(inputs[i])
			if err != nil {
				return err
			}

			// Вычисление ошибки
			for j := range outputs {
				totalError += math.Pow(targets[i][j]-outputs[j], 2)
			}

			// Обратное распространение
			if err := s.Backpropagate(targets[i]); err != nil {
				return err
			}
		}

		if epoch%100 == 0 {
			fmt.Printf("Epoch %d, Error: %f\n", epoch, totalError/float64(len(inputs)))
		}
	}

	return nil
}

// Predict выполняет предсказание на новых данных
func (s *NeuralNetworkService) Predict(inputs []float64) ([]float64, error) {
	return s.FeedForward(inputs)
}
