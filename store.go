package main

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

type LocationSorter []Location

func (sorter LocationSorter) ordering() []AisleRegion {
	return []AisleRegion{
		AisleRegion("front"),
		AisleRegion("middle"),
		AisleRegion("back"),
		AisleRegion("behind"),
	}
}

func (sorter LocationSorter) Len() int {
	return len(sorter)
}

func (sorter LocationSorter) Less(i, j int) bool {
	o := sorter.ordering()
	if o[i] == o[j] {
		return false
	}
	for _, region := range o {
		if region == o[i] {
			return true
		}
		if region == o[j] {
			return false
		}
	}
	panic("unkown aisle regions " + o[i] + " and " + o[j])
}

func (sorter LocationSorter) Swap(i, j int) {
	tmp := sorter[i]
	sorter[i] = sorter[j]
	sorter[j] = tmp
}

func stopAndShopLayout() StoreLayout {
	return StoreLayout{
		Aisles: []Aisle{
			Aisle("Bakery"),
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

// FindItems returns the locations of items on the list that are in the aisle.
//
// The return value is not ordered.
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

func stopAndShopArrangement() StoreArrangement {
	return StoreArrangement{
		layout: stopAndShopLayout(),
		placements: map[GroceryListItem]Location{
			GroceryListItem("milk"):            Location{Aisle: Aisle("Aisle 13"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("cheddar"):         Location{Aisle: Aisle("Aisle 13"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("coffee"):          Location{Aisle: Aisle("Aisle 6"), Region: AisleRegion("front"), Side: AisleSide("left")},
			GroceryListItem("bars"):            Location{Aisle: Aisle("Aisle 6"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("carrots"):         Location{Aisle: Aisle("Produce")},
			GroceryListItem("red onion"):       Location{Aisle: Aisle("Produce")},
			GroceryListItem("english muffins"): Location{Aisle: Aisle("Aisle 10"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("frozen pizza"):    Location{Aisle: Aisle("Aisle 12"), Region: AisleRegion("back"), Side: AisleSide("right")},
			GroceryListItem("pasta sauce"):     Location{Aisle: Aisle("Aisle 7"), Region: AisleRegion("middle"), Side: AisleSide("right")},
			GroceryListItem("pasta"):           Location{Aisle: Aisle("Aisle 7"), Region: AisleRegion("middle"), Side: AisleSide("left")},
		},
	}
}
