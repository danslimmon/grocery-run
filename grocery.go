package main

type GroceryListItem string

type GroceryList []GroceryListItem

func (list GroceryList) RemainingItems() []GroceryListItem {
	return []GroceryListItem(list)
}

func groceryList() GroceryList {
	return GroceryList([]GroceryListItem{
		GroceryListItem("bell pepper"),
		GroceryListItem("broccoli"),
		GroceryListItem("burger buns"),
		GroceryListItem("butter"),
		GroceryListItem("carrots"),
		//GroceryListItem("cereal"),
		GroceryListItem("coffee"),
		GroceryListItem("crackers"),
		GroceryListItem("croutons"),
		GroceryListItem("cucumber"),
		//GroceryListItem("eggs"),
		//GroceryListItem("english muffins"),
		GroceryListItem("frozen french fries"),
		//GroceryListItem("frozen meatballs"),
		GroceryListItem("frozen pizza"),
		GroceryListItem("frozen meals"),
		GroceryListItem("fruit"),
		GroceryListItem("hummus"),
		//GroceryListItem("mac and cheese"),
		GroceryListItem("milk"),
		GroceryListItem("pasta"),
		GroceryListItem("pasta sauce"),
		GroceryListItem("peanut butter"),
		GroceryListItem("salad mix"),
		GroceryListItem("string cheese"),
		GroceryListItem("veggie burgers"),
		//GroceryListItem("vegetable oil"),
		GroceryListItem("yogurt"),
	})
}
