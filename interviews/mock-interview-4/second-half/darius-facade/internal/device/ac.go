package device

import "fmt"

type AC struct{}

func (a *AC) On() {
	fmt.Println("AC is ON")
}

func (a *AC) Off() {
	fmt.Println("AC is OFF")
}
