package main

import "fmt"

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Patient " + p.name + " Doctor checkup already done")
		d.next.execute(p)
		return
	}
	p.doctorCheckUpDone = true
	fmt.Println("Patient " + p.name + " Doctor checking patient")
	d.next.execute(p)
}
func (d *doctor) setNext(next department) {
	d.next = next
}
