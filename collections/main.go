package main

import "fmt"

func main() {
	products := map[string]float64{
		"kayak":      279,
		"lifejacket": 48.95,
		"hat":        0,
	}

	value, ok := products["hat"]
	if ok {
		fmt.Println("Ok", value)
	} else {
		fmt.Println("Not ok")
	}
}
