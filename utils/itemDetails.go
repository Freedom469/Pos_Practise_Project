package pos

func GetItemsNameAndPrice(code int) (string, int) {
	List := ProductList()
	name := List[code]
	itemName := ""
	for key := range name {
		itemName = key
	}
	price := List[code][itemName]

	return itemName, price
}