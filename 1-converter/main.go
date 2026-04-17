package main

import (
	"fmt"
)

// const usdToEuro = 0.85
// const usdToRub = 77.00
var rates = map[string]float64{
	"usd_euro": 0.85,
	"usd_rub":  77.00,
	"euro_usd": 1.18,  // 1 / 0.85
	"euro_rub": 90.59, // (1 / 0.85) * 77
	"rub_usd":  0.013, // 1 / 77
	"rub_euro": 0.011, // (1 / 77) * 0.85
}

// 1. Функция ввода и проверки валюты
// Параметр forbidden нужен, чтобы целевая валюта не совпадала с исходной
func inputCurrency(label string, forbidden string) string {
	var cur string
	for {
		fmt.Printf("Введите %s валюту (usd, rub, euro): ", label)
		fmt.Scan(&cur)
		if cur != "usd" && cur != "rub" && cur != "euro" {
			fmt.Println("Ошибка: некорректная валюта. Попробуйте еще раз.")
			continue
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

func calculate(count float64, from string, to string, rates *map[string]float64) float64 {
	key := from + "_" + to
	rate, exists := (*rates)[key]

	if !exists {
		return count // Если пара не найдена, возвращаем исходную сумму
	}

	return count * rate
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
	result := calculate(count, from, to, &rates)
	fmt.Printf("Результат: %.2f %s\n", result, to)
}
