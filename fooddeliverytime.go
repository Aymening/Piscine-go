package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var ft food

	if order == "burger" {
		ft.preptime = 15
	} else if order == "chips" {
		ft.preptime = 10
	} else if order == "nuggets" {
		ft.preptime = 12
	} else {
		ft.preptime = 404
	}
	return ft.preptime
}
