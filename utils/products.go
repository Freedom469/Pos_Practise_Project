package pos

func ProductList() map[int]map[string]int {
	List := map[int]map[string]int{
		1: {"banana": 80},
		2: {"Watermelon": 300},
		3: {"Mango": 100},
		4: {"Apple": 180},
		5: {"Orange": 50},
		6: {"PineApple": 200},
		// Add more products here following the same pattern
	}

	return List
}