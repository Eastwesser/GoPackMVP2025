package main

import "fmt"

// Подсистема освещения
type Lighting struct{}

func (l *Lighting) TurnOn() {
	fmt.Println("Свет включён")
}

func (l *Lighting) TurnOff() {
	fmt.Println("Свет выключен")
}

// Подсистема климат-контроля
type ClimateControl struct{}

func (c *ClimateControl) SetTemperature(temp int) {
	fmt.Printf("Температура установлена на %d°C\n", temp)
}

// Подсистема безопасности
type Security struct{}

func (s *Security) Arm() {
	fmt.Println("Сигнализация включена")
}

func (s *Security) Disarm() {
	fmt.Println("Сигнализация выключена")
}

// Фасад для управления умным домом
type SmartHomeFacade struct {
	lighting       *Lighting
	climateControl *ClimateControl
	security       *Security
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lighting:       &Lighting{},
		climateControl: &ClimateControl{},
		security:       &Security{},
	}
}

// Метод "Уйти из дома"
func (s *SmartHomeFacade) LeaveHome() {
	fmt.Println("Уходим из дома...")
	s.lighting.TurnOff()
	s.climateControl.SetTemperature(18) // Устанавливаем комфортную температуру
	s.security.Arm()
}

// Метод "Вернуться домой"
func (s *SmartHomeFacade) ReturnHome() {
	fmt.Println("Возвращаемся домой...")
	s.security.Disarm()
	s.lighting.TurnOn()
	s.climateControl.SetTemperature(22) // Устанавливаем комфортную температуру
}

func main() {
	// Создаем фасад
	smartHome := NewSmartHomeFacade()

	// Уходим из дома
	smartHome.LeaveHome()

	// Возвращаемся домой
	smartHome.ReturnHome()
}
