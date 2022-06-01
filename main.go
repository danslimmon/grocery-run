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
			"%-20v %s\n",
			step.Location,
			string(step.Item),
		)
	}
}
