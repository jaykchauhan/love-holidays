package main

func getCostMatrix() [][]int {

	return [][]int{
		{0, 15, 80, 90},
		{0, 0, 40, 50},
		{0, 0, 0, 70},
		{0, 0, 0, 0},
	}
}

func getAllRoutes() [4]string {
	return [4]string{"Castle Black", "Winterfell", "Riverrun", "King's Landing"}
}
