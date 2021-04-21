package main

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		}}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		Shirt{
			logo: "adidas",
			size: 14,
		}}
}
