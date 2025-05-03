package device

import "fmt"

type MusicPlayer struct{}

func (m *MusicPlayer) Play() {
	fmt.Println("Music Player is PLAYING")
}

func (m *MusicPlayer) Stop() {
	fmt.Println("Music Player is STOPPED")
}
