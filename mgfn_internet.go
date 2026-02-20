package main

import "fmt"

func calculate(used, limit, base float64) float64 {
	if used <= limit {
		return base
	}
	return base + (used-limit)*0.06
}

func Calculate1000(mb float64) float64  { return calculate(mb, 1000, 35) }
func Calculate2000(mb float64) float64  { return calculate(mb, 2000, 55) }
func Calculate3000(mb float64) float64  { return calculate(mb, 3000, 70) }
func Calculate5000(mb float64) float64  { return calculate(mb, 5000, 95) }
func Calculate10000(mb float64) float64 { return calculate(mb, 10000, 170) }

func main() {
	fmt.Println(Calculate1000(9999))
}
