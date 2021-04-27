package main

type counterTerroristDress struct {
	color string
}

func (t *counterTerroristDress) getColor() string {
	return t.color
}
func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}
