package main

import (
	"fmt"
)

func main() {
	arrangement := stopAndShopArrangement()
	list := groceryList()
	route := calculateRoute(arrangement, list)
	for _, step := range route.Steps() {
		fmt.Printf(
			"%-10s %-10s %s\n",
			step.Location.Aisle,
			step.Location.Side,
			step.Item,
		)
	}
}
