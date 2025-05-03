package main

import "fmt"

// Light Subsystem
type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is ON")
}

func (l *Light) Off() {
	fmt.Println("Light is OFF")
}

// AirConditioner Subsystem
type AirConditioner struct{}

func (a *AirConditioner) On() {
	fmt.Println("AC is ON")
}

func (a *AirConditioner) Off() {
	fmt.Println("AC is OFF")
}

// MusicPlayer Subsystem
type MusicPlayer struct{}

func (m *MusicPlayer) Play() {
	fmt.Println("Music is playing")
}

func (m *MusicPlayer) Stop() {
	fmt.Println("Music stopped")
}

// SmartHomeFacade our facade pattern
type SmartHomeFacade struct {
	light *Light
	ac    *AirConditioner
	music *MusicPlayer
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		light: &Light{},
		ac:    &AirConditioner{},
		music: &MusicPlayer{},
	}
}

func (s *SmartHomeFacade) TurnOnEverything() {
	fmt.Println("\nI'm home! Turning everything ON...")
	s.light.On()
	s.ac.On()
	s.music.Play()
}

func (s *SmartHomeFacade) TurnOffEverything() {
	fmt.Println("\nI'm leaving! Turning everything OFF...")
	s.light.Off()
	s.ac.Off()
	s.music.Stop()
}

func main() {
	home := NewSmartHomeFacade()

	// arrive-leave scenario
	home.TurnOnEverything()
	home.TurnOffEverything()
}
