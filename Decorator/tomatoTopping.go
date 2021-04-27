package main

type tomatoTopping struct {
	myPizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.myPizza.getPrice()
	return pizzaPrice
}
