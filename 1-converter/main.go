package main

import "fmt"

	func main () {
		const usdToEuro  = 0.85
		const usdToRub = 77.00
		const rubToEuro = usdToRub / usdToEuro
		fmt.Println("rubToEuro = ", rubToEuro)
	}