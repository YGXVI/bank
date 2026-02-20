package main

import "fmt"

func Calculate(amount float64, currency string) float64 {
	if amount <= 0 {
		return 0
	}

	annualRate := 0.05
	return amount * annualRate / 12
}

func main() {
	fmt.Println(
		Calculate(0, "TJS"),
		Calculate(0, "USD"),
		Calculate(500000, "TJS"),
		Calculate(500000, "USD"),
		Calculate(1000000, "TJS"),
		Calculate(1000000, "USD"),
	)
}
