package main

import (
	"errors"
	"fmt"
)

const usdToEuro = 0.85
const usdToRub = 77.00
const rubToEuro = usdToRub / usdToEuro

func userInput() (float64, string, string) {
	var count float64
	var from string
	var to string
	for {
		fmt.Print("Введите количество: ")
		fmt.Scan(&count)
		fmt.Print("Введите исходную валюту (usd, rub, euro): ")
		fmt.Scan(&from)
		fmt.Print("Введите целевую валюту (usd, rub, euro): ")
		fmt.Scan(&to)
		err := checkoutValue(count, from, to)
		if err != nil {
			fmt.Println("Неправильное значение, попробуйте ещё раз")
			continue
		}
		return count, from, to
	}

}

func calculate(count float64, from string, to string) float64 {

	if from == "euro" && to == "rub" {
		return count * rubToEuro
	}
	if from == "usd" && to == "rub" {
		return count * usdToRub
	}
	if from == "euro" && to == "usd" {
		return count / usdToEuro
	}
	if from == "rub" && to == "usd" {
		return count / usdToRub
	}
	if from == "rub" && to == "euro" {
		return count / usdToEuro
	}
	if from == "usd" && to == "euro" {
		return count * usdToEuro
	}
	return count

}

func checkoutValue(count float64, from string, to string) error {
	if count <= 0 {
		return errors.New("ERROR")
	}
	if from != "euro" && from != "usd" && from != "rub" {
		return errors.New("ERROR")
	}

	if to != "euro" && to != "usd" && to != "rub" {
		return errors.New("ERROR")
	}

	if from == to {
		return errors.New("ERROR")
	}

	return nil
}

func main() {
	count,
		from,
		to := userInput()
	fmt.Println(calculate(count, from, to))

}
