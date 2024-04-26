package pos

import "fmt"
import "strings"
import "os"

func TakeOrder() {

	//Output Messages
	BuyMessage := `How Many %s(s) Would You Like To Buy:  `
	ProceesMessage := "Place Another Order? (y/n): "
	ProceesErrorMessage := "Please reply with y/n"
	QuantityErrorMessage := "Invalid input. Please enter a valid quantity."
	OrderChoiceErrorMessage := "Ivalid input Use The values next to the option"
	OrderRangeErrorMessage := "Invalid Input, TRY AGAIN"

	Istrue := true
	total := 0
	orderdeItems := make(map[int]int)
	List := ProductList()

	for Istrue {
			fmt.Println(`
		       Placing Your order ...`)
	var order int
	var quantity int

	//Print The available Products 
	for i, product := range List {
		for item, price := range product {
			fmt.Printf(`
			%d : %s: %d
		`, i, item, price)
		}
	}

	//Taking Input / users order
	fmt.Println("\nChoose Your Order: (ex: 1 for Banana):")
		_, err := fmt.Scan(&order)

		//checking if users order is within the scope of available items
		if order < 1 || order > len(List) {
			fmt.Println(OrderRangeErrorMessage)
		} 
		ErrHandler(err, OrderChoiceErrorMessage)

		var res string
		var sub int

	switch order {
	case 1, 2, 3, 4, 5, 6:
		item, price := GetItemsNameAndPrice(order)
		fmt.Printf(BuyMessage, item)
		_, err := fmt.Scan(&quantity)
		if err != nil || quantity <= 0 {
			fmt.Println(QuantityErrorMessage)
			continue
		}
		// Retrieve item price and calculate subtotal
		sub = quantity * price
		total += sub
		// Process message and user input for placing another order
		fmt.Println(ProceesMessage)
		_, error := fmt.Scan(&res)
		ErrHandler(error, ProceesErrorMessage)
		if strings.ToLower(res) != "y" {
			Istrue = false
		}
	default:
		fmt.Println("Invalid choice. Please select a valid option.")
}

		//string orderd items
		orderdeItems[order] += quantity
	}

	cash := 0

	for cash < total {
	fmt.Printf("Total: %d\n",total)
	fmt.Println("Enter The Cash Amount Paid: ")
	fmt.Scan(&cash)
		
	}

	

	fmt.Println(`
			Your Purchase`)
	fmt.Println("...........................................................................")
	fmt.Println(`Item:                     Price(Ksh):                          Quantity:`)
	fmt.Println("...........................................................................")
	for order, number := range orderdeItems {
		if product, ok := List[order]; ok {
			for key, value := range product {
				fmt.Println(key, "                     ", value, "                           ", number)
			}
		}
	}
	fmt.Println("...........................................................................")

	fmt.Printf("Your Orders: %v             Cash Paid: %d                       Bal: %d", len(orderdeItems), cash, cash-total)

		fmt.Println(
			`      
			   Total: Ksh`, total)

		 fmt.Println(`...........................................................................`)
		 fmt.Println(`
                     _______   _
                    |__   __| | |                      _
                       | |    | |__     __ _   _ __   | | _         _   _    ___    _   _
                       | |    | _` + "`" + `  | |  \  | |/ /       | | | |  / _ \  | | | |
                       | |    | | | | | (_| | | | | | |   <        | |_| | | (_) | | |_| |
                       |_|    |_| |_|  \__,_| |_| |_| |_|\_\        \__, |  \___/   \__,_|
                                                                      __/ /
                                                                     |___/
`)
os.Exit(0)

}