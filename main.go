package main

import (
	"coffeeshop/coffee"
	"fmt"
)

func main() {
	coffees, err := coffee.GetCoffees()
	if err != nil {
		fmt.Println(err)
	}
	for _, element := range coffees.List {
		result := fmt.Sprintf("%s for $%v", element.Name, element.Price)
		fmt.Println(result)
	}
	fmt.Println("Is Milk available?", func() string {if coffee.IsCoffeeAvailable("Milk") {return "Yes"} else {return "No"}}())
}