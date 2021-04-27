package main

import "fmt"

func main() {
	game := newGame()
	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)
	game.addTerrorists(TerroristDressType)

	game.addCounterTerrorists(CounterTerroristType)
	game.addCounterTerrorists(CounterTerroristType)
	game.addCounterTerrorists(CounterTerroristType)
	game.addCounterTerrorists(CounterTerroristType)
	game.addCounterTerrorists(CounterTerroristType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType : %s \nDressColor : %s \n", dressType, dress.getColor())
	}

}
