package main

import "fmt"

func main() {
	Pizza := &veggieMania{}

	pizzaWithCheese := &cheeseTopping{myPizza: Pizza}

	pizzaWithCheeseAndTomato := &tomatoTopping{myPizza: pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())

}
