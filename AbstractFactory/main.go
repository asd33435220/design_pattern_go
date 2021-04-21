package main

import "fmt"

func main() {
	adidasFactory, _ := getSportFactory("adidas")
	nikeFactory, _ := getSportFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	nikeShoe := nikeFactory.makeShoe()

	adidasShirt := adidasFactory.makeShirt()
	nikeShirt := nikeFactory.makeShirt()

	printShoeDetail(adidasShoe)
	printShoeDetail(nikeShoe)

	printShirtDetail(adidasShirt)
	printShirtDetail(nikeShirt)

}

func printShoeDetail(shoe iShoe) {
	fmt.Println("Shoe Logo:", shoe.getLogo())
	fmt.Println(" Shoe Size:", shoe.getSize())
}

func printShirtDetail(shirt iShirt) {
	fmt.Println("Shirt Logo:", shirt.getLogo())
	fmt.Println("Shirt Size:", shirt.getSize())
}
