package facade

import (
	"darius-facade/internal/device"
)

type SmartHomeFacade struct {
	light       *device.Light
	ac          *device.AC
	musicPlayer *device.MusicPlayer
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		light:       &device.Light{},
		ac:          &device.AC{},
		musicPlayer: &device.MusicPlayer{},
	}
}

func (s *SmartHomeFacade) TurnOnEverything() {
	s.light.On()
	s.ac.On()
	s.musicPlayer.Play()
}

func (s *SmartHomeFacade) TurnOffEverything() {
	s.light.Off()
	s.ac.Off()
	s.musicPlayer.Stop()
}
