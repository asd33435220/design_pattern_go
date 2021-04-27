package main

type cheeseTopping struct {
	myPizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.myPizza.getPrice()
	return pizzaPrice
}
