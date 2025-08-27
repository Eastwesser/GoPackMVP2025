// internal/models/models.go
package models

import (
	"fmt"
	"sync"
)

type NeuralNetwork struct {
	Layers []*Layer
	Config *NetworkConfig
}

type NetworkConfig struct {
	InputSize         int
	OutputSize        int
	HiddenLayersCount int
	HiddenLayerSize   int
	LearningRate      float64
	Activation        string
}

type LayerType int

const (
	InputLayer LayerType = iota
	HiddenLayer
	OutputLayer
)

type Layer struct {
	Neurons []*Neuron
	Type    LayerType
}

type Neuron struct {
	Inputs  []*Synapse
	Outputs []*Synapse
	Bias    float64
	Value   float64
	Delta   float64
	Mu      sync.Mutex // Изменено на заглавную букву
}

type Synapse struct {
	Weight    float64
	Source    *Neuron
	Target    *Neuron
	LastDelta float64
}

type NetworkError struct {
	Message   string
	Operation string
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("Network error in %s: %s", e.Operation, e.Message)
}
