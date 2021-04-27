package main

import "fmt"

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	windowsComputer := &windows{}

	macComputer.setPrint(hpPrinter)
	macComputer.print()

	fmt.Println()

	windowsComputer.setPrint(epsonPrinter)
	windowsComputer.print()

	fmt.Println()

	macComputer.setPrint(epsonPrinter)
	macComputer.print()

	fmt.Println()

	windowsComputer.setPrint(hpPrinter)
	windowsComputer.print()
}
