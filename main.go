package main

import(
	"fmt"
	"pos/utils"
)

func main() {
	render := true
	
	for render {
		var input int
		fmt.Println(`
			WELCOME TO THE PoS`)
			fmt.Println("..............................................................................")
		fmt.Println(`
			0 : Place Order
			1 : View Menu
			2 : Exit
		`)
		_, err := fmt.Scan(&input)

		pos.ErrHandler(err, "Ivalid input Use The values next to the option")
		
		if input < 0 && input > 2 {
			fmt.Println("Invalid choice. Please select a valid option.")
		}

		switch input {
		case 0 : {
			pos.TakeOrder()
		}
		case 2: {
			fmt.Println("Exiting ....")
			render = false
		}
		case 1: {
			fmt.Println(`
				Viewing menu ...`+"\n")
			pos.DisplayMenu()
		}
		default: {
			fmt.Println("Invalid choice. Please select a valid option.")
		}
			
		}
	}
}