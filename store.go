package main

import (
	"fmt"
)

type Aisle string

type AisleRegion string

type AisleSide string

func (s AisleSide) Flip() AisleSide {
	switch string(s) {
	case "left":
		return AisleSide("right")
	case "right":
		return AisleSide("left")
	default:
		// Some locations (e.g. the produce section) don't have "sides"
		return AisleSide("")
	}
}

type Row string

func (r Row) Flip() Row {
	if r == Row("front") {
		return Row("back")
	}
	return Row("front")
}

type StoreLayout struct {
	Aisles []Aisle
}

type Location struct {
	Aisle  Aisle
	Region AisleRegion
	// The side of the aisle that the item is on. "left" means the left side when you enter the
	// store.
	Side AisleSide
}

// Sorts locations from front to back of store.
type LocationSorter []Location

func (s LocationSorter) ordering() []AisleRegion {
	return []AisleRegion{
		AisleRegion("front"),
		AisleRegion("middle"),
		AisleRegion("back"),
		AisleRegion("behind"),
	}
}

func (s LocationSorter) Len() int {
	return len(s)
}

func (s LocationSorter) Less(i, j int) bool {
	fmt.Printf("s: %v; i: %d; j: %d\n", s, i, j)
	if s[i].Region == s[j].Region {
		return false
	}
	o := s.ordering()
	for _, region := range o {
		if region == s[i].Region {
			return true
		}
		if region == s[j].Region {
			return false
		}
	}
	panic("unkown aisle regions " + s[i].Region + " and " + s[j].Region)
}

func (s LocationSorter) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func stopAndShopNorthHavenLayout() StoreLayout {
	return StoreLayout{
		Aisles: []Aisle{
			Aisle("Aisle 24"),
			Aisle("Aisle 23"),
			Aisle("Aisle 22"),
			Aisle("Aisle 21"),
			Aisle("Aisle 20"),
			Aisle("Aisle 19"),
			Aisle("Aisle 18"),
			Aisle("Aisle 17"),
			Aisle("Aisle 16"),
			Aisle("Aisle 15"),
			Aisle("Aisle 14"),
			Aisle("Aisle 13"),
			Aisle("Aisle 12"),
			Aisle("Aisle 11"),
			Aisle("Aisle 10"),
			Aisle("Aisle 9"),
			Aisle("Aisle 8"),
			Aisle("Aisle 7"),
			Aisle("Aisle 6"),
			Aisle("Aisle 5"),
			Aisle("Aisle 4"),
			Aisle("Aisle 3"),
			Aisle("Aisle 2"),
			Aisle("Aisle 1"),
			Aisle("Produce"),
			Aisle("Checkout"),
		},
	}
}

// maps items to their locations within a given store arrangement
type StoreArrangement struct {
	layout     StoreLayout
	placements map[GroceryListItem]Location
}

type itemLocation struct {
	Item     GroceryListItem
	Location Location
}

// Sorts itemLocations from front to back of store.
type itemLocationSorter []itemLocation

func (s itemLocationSorter) Len() int {
	return len(s)
}

func (s itemLocationSorter) Less(i, j int) bool {
	b := LocationSorter([]Location{s[i].Location, s[j].Location}).Less(0, 1)
	fmt.Printf("b: %v\n", b)
	return b
}

func (s itemLocationSorter) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

// FindItems returns the locations of items on the list that are in the aisle.
//
// The returned items are ordered by aisle (in the same order specified in the layout), but within
// aisles they are not guaranteed to be ordered in any particular way.
func (arr StoreArrangement) FindItems(aisle Aisle, list GroceryList) []itemLocation {
	itemLocations := []itemLocation{}
	for item, location := range arr.placements {
		if location.Aisle == aisle {
			itemLocations = append(itemLocations, itemLocation{Item: item, Location: location})
		}
	}

	rslt := []itemLocation{}
	for _, itemFromList := range list.RemainingItems() {
		for _, itemLocation := range itemLocations {
			if itemFromList == itemLocation.Item {
				rslt = append(rslt, itemLocation)
			}
		}
	}

	return rslt
}

func stopAndShopNorthHavenArrangement() StoreArrangement {
	return StoreArrangement{
		layout: stopAndShopNorthHavenLayout(),
		placements: map[GroceryListItem]Location{
			GroceryListItem("milk"):            Location{Aisle: Aisle("Aisle 24"), Region: AisleRegion("behind")},
			GroceryListItem("coffee"):          Location{Aisle: Aisle("Aisle 8"), Region: AisleRegion("back"), Side: AisleSide("left")},
			GroceryListItem("bars"):            Location{Aisle: Aisle("Aisle 14"), Region: AisleRegion("front"), Side: AisleSide("right")},
			GroceryListItem("carrots"):         Location{Aisle: Aisle("Produce")},
			GroceryListItem("dip"):             Location{Aisle: Aisle("Produce"), Region: AisleRegion("middle")},
			GroceryListItem("red onion"):       Location{Aisle: Aisle("Produce")},
			GroceryListItem("english muffins"): Location{Aisle: Aisle("Aisle 10"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("frozen pizza"):    Location{Aisle: Aisle("Aisle 12"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("pasta"):           Location{Aisle: Aisle("Aisle 4"), Region: AisleRegion("front"), Side: AisleSide("left")},
			GroceryListItem("pasta sauce"):     Location{Aisle: Aisle("Aisle 4"), Region: AisleRegion("front"), Side: AisleSide("right")},
			GroceryListItem("crackers"):        Location{Aisle: Aisle("Aisle 8"), Region: AisleRegion("front"), Side: AisleSide("right")},
			GroceryListItem("yogurt"):          Location{Aisle: Aisle("Aisle 14"), Region: AisleRegion("behind")},
		},
	}
}
