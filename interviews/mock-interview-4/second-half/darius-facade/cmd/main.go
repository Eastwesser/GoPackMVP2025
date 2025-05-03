package main

import "darius-facade/internal/facade"

func main() {
	home := facade.NewSmartHomeFacade()

	// CAME BACK HOME
	home.TurnOnEverything()

	println()

	// LEFT HOME
	home.TurnOffEverything()
}
