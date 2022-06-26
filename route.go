package main

import (
	"sort"
)

type RouteStep struct {
	Location Location
	Item     GroceryListItem
}

type Route struct {
	steps   []RouteStep
	unfound []GroceryListItem
}

func (rt *Route) AddStep(step RouteStep) {
	rt.steps = append(rt.steps, step)
}

func (rt *Route) Steps() []RouteStep {
	return rt.steps
}

func (rt *Route) Unfound() []GroceryListItem {
	return rt.unfound
}

func (rt *Route) getUnfound(list GroceryList, found []GroceryListItem) []GroceryListItem {
	rslt := make([]GroceryListItem, 0)
OUTER:
	for _, listItem := range list {
		for _, foundItem := range found {
			if listItem == foundItem {
				continue OUTER
			}
		}
		rslt = append(rslt, listItem)
	}
	return rslt
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

	found := make([]GroceryListItem, 0)
	for _, aisle := range arr.layout.Aisles {
		itemLocs := arr.FindItems(aisle, list)
		if len(itemLocs) == 0 {
			// None of the items on the list are in this aisle; move on
			continue
		}

		sort.Sort(itemLocationSorter(itemLocs))

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
			found = append(found, itemLoc.Item)
		}

		// Flip y if the shopper has passed all the way through an aisle to the opposite row of the
		// store.
		//
		// Don't flip y if last two items have both been behind the aisle – e.g. if you get milk
		// behind aisle 24 then yogurt behind aisle 14, that doesn't necessitate walking all the way
		// through aisle 14 to the front of the store.
		steps := route.Steps()
		if len(route.Steps()) == 1 {
			continue
		}
		if steps[len(steps)-1].Location.Region == AisleRegion("behind") && s.y == Row("back") {
			continue
		}
		s.y = s.y.Flip()
	}

	route.unfound = route.getUnfound(list, found)
	return route
}
