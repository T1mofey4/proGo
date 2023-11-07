package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// kayak := store.NewProduct("Kayak", "Watersports", 275)
	// lifejacket := &store.Product{Name: "Lifejacket", Category: "Watersports"}

	// for _, p := range []*store.Product{kayak, lifejacket} {
	// 	fmt.Println("Name: ", p.Name, "Category: ", p.Category, "Price: ", p.Price(0.2))
	// }

	// boats := []*store.Boat{
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Kayak", 650.25, 2, true),
	// }

	// for _, b := range boats {
	// 	fmt.Println("Conventional: ", b.Product.Name, "Direct: ", b.Name)
	// }

	// rentals := []*store.RentalBoat{
	// 	store.NewRentalBoat("RubberRing", 10, 1, false, false),
	// 	store.NewRentalBoat("Yacht", 50000, 5, true, true),
	// 	store.NewRentalBoat("Super Yacht", 100000, 15, true, true),
	// }

	// for _, r := range rentals {
	// 	fmt.Println("Rental Boat: ", r.Name, "Rental Price: ", r.Price(0.2))
	// }

	// product := store.NewProduct("Kayak", "Watersports", 279)

	// deal := store.NewSpecialDeal("Weekend Special", product, 50)

	// Name, price, Price := deal.GetDetails()

	// fmt.Println("Name: ", Name)
	// fmt.Println("Price field ", price)
	// fmt.Println("Price method ", Price)

	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball":  store.NewProduct("Soccer ball", "Soccer", 19.50),
	}

	for key, p := range products {
		fmt.Println("Key:", key, "Price:", p.Price(0.2))
	}

}
