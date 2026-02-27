package main

import "fmt"

const usdToEuro = 0.85
const usdToRub = 77.00
const rubToEuro = usdToRub / usdToEuro

func userInput() (float64, string, string) {
	var count, from, to = 0.0, "usd", "euro"
	fmt.Print("Введите цифру:")
	fmt.Scan(&count)
	fmt.Print("Введите валюту:")
	fmt.Scan(&from)
	fmt.Print("Введите вторую волюту:")
	fmt.Scan(&to)
	return count, from, to
}

func calculate(count float64, from string, to string) float64 {
	if from == "euro" && to == "rub"{
		return count * rubToEuro
	}
	if from =="usd" && to == "rub"{
		return  count * usdToRub
	}
	if from == "euro" && to == "usd"{
		return count / usdToEuro
	}
	if from == "rub" && to == "usd"{
		return count / usdToRub 
	}
	if from == "rub" && to == "euro"{
		return count / usdToEuro
	}
	if from == "usd" && to == "euro"{
		return count * usdToEuro
	}
	return count
}
	func main() {
	 count,
		from,
		to := userInput()
		fmt.Println(calculate(count, from, to))
	// const usdToEuro = 0.85
	// const usdToRub = 77.00
	// const rubToEuro = usdToRub / usdToEuro
	// fmt.Println("rubToEuro = ", rubToEuro)
}
