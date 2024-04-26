package pos

import "fmt"

func DisplayMenu() {
	Menu := ProductList()
	for _, menu := range Menu {
		for item, price := range menu {
		fmt.Println("                                "+item + ": ", price, )
	}
	}

}