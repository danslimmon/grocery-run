package main

type GroceryListItem string

type GroceryList []GroceryListItem

func (list GroceryList) RemainingItems() []GroceryListItem {
	return []GroceryListItem(list)
}

func groceryList() GroceryList {
	return GroceryList([]GroceryListItem{
		GroceryListItem("vegetable oil"),
		GroceryListItem("eggs"),
		GroceryListItem("english muffins"),
		GroceryListItem("pasta"),
		GroceryListItem("pasta sauce"),
		GroceryListItem("cereal"),
		GroceryListItem("coffee"),
		GroceryListItem("mac and cheese"),
		GroceryListItem("milk"),
		GroceryListItem("frozen pizza"),
		GroceryListItem("frozen meatballs"),
		GroceryListItem("yogurt"),
	})
}
