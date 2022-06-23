package main

import (
	"fmt"
)

func main() {
	arrangement := stopAndShopNorthHavenArrangement()
	list := groceryList()
	route := calculateRoute(arrangement, list)
	for _, step := range route.Steps() {
		fmt.Printf(
			"%-10s %-10s %-10s %s\n",
			step.Location.Aisle,
			step.Location.Side,
			step.Location.Region,
			step.Item,
		)
	}

	fmt.Printf("\nUnfound:\n")
	for _, unfoundItem := range route.Unfound() {
		fmt.Printf("%s\n", unfoundItem)
	}
}
