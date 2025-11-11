package converter

import (
	"currency_converter/pkg/utils"
	"fmt"
	"strings"
)

func Logo() {
	fmt.Println(
		`___________________________________________________________________________________
                               __       ______   _____
                              /  \     |  __  \ |_   _|
                             / /\ \    | |__| |   | |
                            / /__\ \   |  ____/   | |
                           / ______ \  | |       _| |_
                          /_/      \_\ |_|      |_____|
   ____    ____    _    _  __      __  ______   _____   _______   ______   _____
  / ___|  / __ \  |  \ | | \ \    / / | _____| |  _  | |__   __| | _____| |  _  | 
 / /     | |  | | | \ \| |  \ \  / /  | |___   |  ___/    | |    | |___   |  ___/
 \ \___  | |__| | | |\ \ |   \ \/ /   | |____  |  _ \     | |    | |____  |  _ \  
  \____|  \____/  |_| \ _|    \__/    |______| |_| \_\    |_|    |______| |_| \_\
___________________________________________________________________________________`)

	fmt.Println()
}

func Menu() int {
	var menuText = `Меню:
- 1. Конвертация
- 0. Выход

Введите цифру: `

	fmt.Print(menuText)

	// Ввод данных и обработка ошибок
	var inputNumber int
	_, err := fmt.Scan(&inputNumber)

	for err != nil || inputNumber > 1 || inputNumber < 0 {
		if err != nil {
			fmt.Println("Введено неверное значение, попробуйте ещё раз.")
		} else if inputNumber > 1 || inputNumber < 0 {
			fmt.Println("Вы ввели число, не входящее в диапазон от 0 до 2.")
		}
		fmt.Print(menuText)
		_, err = fmt.Scan(&inputNumber)
	}

	return inputNumber
}

func Converter() error {
	// Загрузка данных
	rates, err := utils.LoadRates()

	if err != nil {
		fmt.Println("Локальный файл с курсами не найден. Загружаю с API...")
		rates = Request()
		utils.SaveFile(rates)
	}

	// Ввод данных и обработка ошибок
	val1 := getValidCurrency("Введите код валюты из которой перевести (Пример: RUB): ", rates.Rates)
	val2 := getValidCurrency("Введите код валюты в которую перевести (Пример: RUB): ", rates.Rates)

	var valCount float64
	fmt.Printf("Введите количество %s: ", val1)
	_, err = fmt.Scan(&valCount)

	for {
		if err != nil {
			fmt.Println("Неверный ввод.")
			fmt.Printf("Введите количество %s: ", val1)
			_, err = fmt.Scan(&valCount)
		} else {
			break
		}
	}

	// Рассчет результата
	if val1 == val2 {
		fmt.Printf("%.2f %s = %.2f %s (одинаковые валюты)\n", valCount, val1, valCount, val2)
		return nil
	}

	result := (valCount / rates.Rates[val1]) * rates.Rates[val2]
	fmt.Printf("%.2f %s = %.2f %s\n", valCount, val1, result, val2)

	return nil
}

func getValidCurrency(prompt string, validCurrencies map[string]float64) string {
	var code string
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&code)
		if err != nil {
			fmt.Println("Неверный ввод. Попробуйте снова.")
			continue
		}
		code = strings.ToUpper(code)
		if _, ok := validCurrencies[code]; ok {
			return code
		}
		fmt.Println("Нет такого кода валюты.")
	}
}
