package main

import "fmt"

type iSportFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportFactory(brand string) (iSportFactory, error) {
	if brand == "nike" {
		return &nike{}, nil
	}
	if brand == "adidas" {
		return &adidas{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
