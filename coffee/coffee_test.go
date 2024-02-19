package coffee

import (
	"testing"
)

func init() {
	Coffees = CoffeeList{
		
		List: []CoffeeDetails{
			{
				Name: "Cappuccino",
				Price: 3.50,
			},
			{
				Name: "Latte",
				Price: 3.50,
			},
			{
				Name: "Espresso",
				Price: 2.50,
			},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	got := IsCoffeeAvailable("Cappuccino")
	if got != true {
		t.Errorf("IsCoffeeAvailable() = %t; want %t", got, true)
	}

}