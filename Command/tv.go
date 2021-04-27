package main

import "fmt"

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	fmt.Println("Turning TV ON")
	t.isRunning = true
}
func (t *tv) off() {
	fmt.Println("Turning TV OFF")
	t.isRunning = false
}
