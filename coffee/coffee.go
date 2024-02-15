package coffee

import (
	"fmt"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name string
	Price float64
}

type CoffeeList struct {
	List []CoffeeDetails
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: %w", err)
		return nil, err
	}
	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(name string) bool {
	for i:= 0; i< len(Coffees.List); i++ {
		if Coffees.List[i].Name == name {
			return true
		}
	}
	return false
}