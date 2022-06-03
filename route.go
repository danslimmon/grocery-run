package main

type RouteStep struct {
	Location Location
	Item     GroceryListItem
}

type Route struct {
	steps []RouteStep
}

func (rt *Route) AddStep(step RouteStep) {
	rt.steps = append(rt.steps, step)
}

func (rt *Route) Steps() []RouteStep {
	return rt.steps
}

func calculateRoute(arr StoreArrangement, list GroceryList) *Route {
	route := new(Route)
	type shopper struct {
		x Aisle
		y Row
	}
	s := shopper{
		// we're at the far right of the store
		x: Aisle("Bakery"),
		// we're at the front of the store
		y: Row("front"),
	}

	for _, aisle := range arr.layout.Aisles {
		itemLocs := arr.FindItems(aisle, list)
		if len(itemLocs) == 0 {
			// None of the items on the list are in this aisle; move on
			continue
		}

		// Go up the aisle and grab items as they occur
		for i := range itemLocs {
			if s.y == Row("back") {
				i = len(itemLocs) - i - 1
			}
			itemLoc := itemLocs[i]
			if s.y == Row("back") {
				itemLoc.Location.Side = itemLoc.Location.Side.Flip()
			}
			route.AddStep(RouteStep{Location: itemLoc.Location, Item: itemLoc.Item})
		}
		s.y = s.y.Flip()
	}

	return route
}
