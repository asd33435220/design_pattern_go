package main

import "fmt"

type mac struct {
	printer Printer
}

func (m *mac) print() {
	fmt.Println("Print request from mac")
	m.printer.printFile()
}
func (m *mac) setPrint(p Printer) {
	m.printer = p
}
