// internal/models/activation.go
package core

import "math"

type ActivationFunction func(float64) float64
type ActivationDerivative func(float64) float64

var ActivationFunctions = map[string]struct {
	Function   ActivationFunction
	Derivative ActivationDerivative
}{
	"sigmoid": {
		Function:   sigmoid,
		Derivative: sigmoidDerivative,
	},
	"relu": {
		Function:   relu,
		Derivative: reluDerivative,
	},
	"tanh": {
		Function:   tanh,
		Derivative: tanhDerivative,
	},
}

func sigmoid(x float64) float64           { return 1.0 / (1.0 + math.Exp(-x)) }
func sigmoidDerivative(x float64) float64 { return x * (1.0 - x) }

func relu(x float64) float64 { return math.Max(0, x) }
func reluDerivative(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}

func tanh(x float64) float64           { return math.Tanh(x) }
func tanhDerivative(x float64) float64 { return 1 - math.Pow(math.Tanh(x), 2) }
