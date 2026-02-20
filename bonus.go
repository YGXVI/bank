
package main

import "fmt"

func Total(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	return amount + amount*0.005
}

func main() {
	fmt.Println(
		Total(0),
		Total(5000),
		Total(10000),
	)
}
