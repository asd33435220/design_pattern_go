package main

import "fmt"

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Patient " + p.name + " Payment done")
	}
	p.paymentDone = true
	fmt.Println("Patient " + p.name + " Cashier getting money from patient")
}
func (c *cashier) setNext(next department) {
	c.next = next
}
