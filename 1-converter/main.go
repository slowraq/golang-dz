package main

import (
	"fmt"
)

// const usdToEuro = 0.85
// const usdToRub = 77.00
var m = map[string]float64{"usdToEuro": 0.85, "usdToRub": 77.00}

// 1. Функция ввода и проверки валюты
// Параметр forbidden нужен, чтобы целевая валюта не совпадала с исходной
func inputCurrency(label string, forbidden string) string {
	var cur string
	for {
		fmt.Printf("Введите %s валюту (usd, rub, euro): ", label)
		fmt.Scan(&cur)
		for key := range m {
			if cur != key {
				fmt.Println("Ошибка: некорректная валюта. Попробуйте еще раз.")
				continue
			}
		}
		// Проверка: входит ли валюта в список разрешенных

		// Проверка: не совпадает ли валюта с уже выбранной (для второго шага)
		if cur == forbidden {
			fmt.Println("Ошибка: целевая валюта не может быть такой же, как исходная.")
			continue
		}

		return cur
	}
}

// 2. Функция ввода и проверки числа
func inputAmount() float64 {
	var count float64
	for {
		fmt.Print("Введите количество: ")
		fmt.Scan(&count)

		if count > 0 {
			return count
		}
		fmt.Println("Ошибка: количество должно быть больше нуля.")
	}
}

// 3. Функция расчета (без использования map)
func calculate(count float64, from string, to string) float64 {
	// Конвертация, если исходная - USD
	if from == "usd" {
		if to == "euro" {
			return count * m["usdToEuro"]
		}
		if to == "rub" {
			return count * m["usdToRub"]
		}
	}
	// Конвертация, если исходная - EURO
	if from == "euro" {
		if to == "usd" {
			return count / m["usdToEuro"]
		}
		if to == "rub" {
			return (count / m["usdToEuro"]) * m["usdToRub"]
		}
	}
	// Конвертация, если исходная - RUB
	if from == "rub" {
		if to == "usd" {
			return count / m["usdToRub"]
		}
		if to == "euro" {
			return (count / m["usdToRub"]) * m["usdToEuro"]
		}
	}
	return count
}

func main() {
	fmt.Println("--- Конвертер валют ---")

	// Шаг 1: Исходная валюта (запретов нет, передаем пустую строку)
	from := inputCurrency("исходную", "")

	// Шаг 2: Количество
	count := inputAmount()

	// Шаг 3: Целевая валюта (запрещаем вводить ту же, что в переменной from)
	to := inputCurrency("целевую", from)

	// Расчет и вывод
	result := calculate(count, from, to)
	fmt.Printf("Результат: %.2f %s\n", result, to)
}
