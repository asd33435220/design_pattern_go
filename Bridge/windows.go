package main

import "fmt"

type windows struct {
	printer Printer
}

func (w *windows) print() {
	fmt.Println("Print request from windows")
	w.printer.printFile()
}
func (w *windows) setPrint(p Printer) {
	w.printer = p
}
