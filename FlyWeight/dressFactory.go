package main

import "fmt"

const (
	TerroristDressType = "tDress"

	CounterTerroristType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		dress := newTerroristDress()
		d.dressMap[dressType] = dress
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerroristType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
