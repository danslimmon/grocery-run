package main

type GroceryListItem string

type GroceryList []GroceryListItem

func (list GroceryList) RemainingItems() []GroceryListItem {
	return []GroceryListItem(list)
}

func groceryList() GroceryList {
	return GroceryList([]GroceryListItem{
		GroceryListItem("milk"),
		GroceryListItem("coffee"),
		GroceryListItem("carrots"),
		GroceryListItem("cheddar"),
		GroceryListItem("english muffins"),
		GroceryListItem("red onion"),
		GroceryListItem("bars"),
		GroceryListItem("frozen pizza"),
		GroceryListItem("pasta sauce"),
		GroceryListItem("pasta"),
	})
}
