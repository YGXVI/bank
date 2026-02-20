package main

import "fmt"

func Calculate(sum float64) float64 {
	if sum <= 5000 {
		return 0
	}

	commission := sum * 0.005
	if commission < 0.25 {
		return 0.25
	}

	return commission
}

func main() {
	fmt.Println(Calculate(9999))
}
