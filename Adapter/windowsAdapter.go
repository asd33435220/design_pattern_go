package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightingPort() {
	fmt.Println("Adapter converts Lighting signal to USB")
	w.windowsMachine.insertIntoUSBPort()
}
