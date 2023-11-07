package main

import "fmt"

type calcFunc func(float64) float64

// func calcWithTax(price float64) float64 {
// 	return price + (price * 0.2)
// }

// func calcWithoutTax(price float64) float64 {
// 	return price
// }

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

// func selectCalculator(price float64) calcFunc {
// 	if price > 100 {
// 		var withTax calcFunc = func(price float64) float64 {
// 			return price + (price * 0.2)
// 		}
// 		return withTax
// 	}
// 	withoutTax := func(price float64) float64 {
// 		return price
// 	}
// 	return withoutTax
// }

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if price > threshold {
			return price + (price * rate)
		}
		return price
	}
}

func main() {
	watersportsProducts := map[string]float64{
		"Kayak":      275,
		"Lifejacket": 48.95,
	}

	soccerProducts := map[string]float64{
		"Soccer Ball": 19.50,
		"Stadium":     79500,
	}

	waterCalc := priceCalcFactory(100, 0.2)
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}

	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}

	// calc := func(price float64) float64 {
	// 	if price > 100 {
	// 		return price + (price * 0.2)
	// 	}
	// 	return price
	// }

	// for product, price := range watersportsProducts {
	// 	printPrice(product, price, calc)
	// }

	// calc = func(price float64) float64 {
	// 	if price > 50 {
	// 		return price + (price * 0.1)
	// 	}
	// 	return price
	// }

	// for product, price := range soccerProducts {
	// 	printPrice(product, price, calc)
	// }

}
