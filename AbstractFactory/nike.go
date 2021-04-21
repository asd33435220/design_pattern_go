package main

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		Shoe{
			logo: "nike",
			size: 14,
		}}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		Shirt{
			logo: "nike",
			size: 14,
		}}
}
