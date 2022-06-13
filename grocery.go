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
		GroceryListItem("bars"),
		GroceryListItem("pasta sauce"),
		GroceryListItem("pasta"),
		GroceryListItem("dip"),
		GroceryListItem("crackers"),
		GroceryListItem("red onion"),
		GroceryListItem("carrots"),
		GroceryListItem("yogurt"),
	})
}
