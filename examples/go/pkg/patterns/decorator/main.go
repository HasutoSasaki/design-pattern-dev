package main

import "fmt"

func main() {
	pizza := &VeggieMania{}

	// cheese topping
	pizzaWithCheese := &CheeseTopping{pizza: pizza}
	// tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
