package converter

import (
	"currency_converter/pkg/utils"
	"fmt"
	"strings"
)

func StartMenu() int {
	fmt.Println(`
______________________________________________________________________________
                              __       ______   _____
                             /  \     |  __  \ |_   _|
                            / /\ \    | |__| |   | |
                           / /__\ \   |  ____/   | |
                          / ______ \  | |       _| |_
                         /_/      \_\ |_|      |_____| 
                       ____    ____    _    _  __      __
                      / ___|  / __ \  |  \ | | \ \    / /
                     / /     | |  | | | \ \| |  \ \  / /
                     \ \___  | |__| | | |\ \ |   \ \/ /
                      \____|  \____/  |_| \ _|    \__/
______________________________________________________________________________`)

	var menuText = `Меню:
- 1. Конвертация
- 2. Меню
- 0. Выход

Введите цифру: `

	fmt.Print(menuText)

	var inputNumber int
	_, err := fmt.Scan(&inputNumber)

	for err != nil || inputNumber > 2 || inputNumber < 1 {
		if err != nil {
			fmt.Println("Введено неверное значение, попробуйте ещё раз.")
		} else if inputNumber > 2 || inputNumber < 0 {
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

	for i := 0; i <= 10; i++ {
		if err != nil {
			rates, err = utils.LoadRates()
		} else {
			break
		}
	}

	// Ввод данных и проверка на ошибки
	var val1, val2 string
	fmt.Print("Введите код валюты из которой перевести (Пример: RUB): ")
	_, err = fmt.Scan(&val1)

	for {
		if err != nil {
			fmt.Println("Неверный ввод.")
			fmt.Print("Введите код валюты из которой перевести (Пример: RUB): ")
			_, err = fmt.Scan(&val1)
		} else if _, ok := rates.Rates[strings.ToUpper(val1)]; !ok {
			fmt.Println("Нет такого кода в системе.")
			fmt.Print("Введите код валюты из которой перевести (Пример: RUB): ")
			_, err = fmt.Scan(&val1)
		} else {
			break
		}
	}

	fmt.Print("Введите код валюты в которую перевести: ")
	fmt.Scan(&val2)

	for {
		if err != nil {
			fmt.Println("Неверный ввод.")
			fmt.Print("Введите код валюты в которую перевести (Пример: RUB): ")
			_, err = fmt.Scan(&val2)
		} else if _, ok := rates.Rates[strings.ToUpper(val1)]; !ok {
			fmt.Println("Нет такого кода в системе.")
			fmt.Print("Введите код валюты в которую перевести (Пример: RUB): ")
			_, err = fmt.Scan(&val2)
		} else {
			break
		}
	}

	val1 = strings.ToUpper(val1)
	val2 = strings.ToUpper(val2)

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

	var rezult float64

	// Рассчет результата
	if val1 == "USD" {
		rezult = rates.Rates[val2] * valCount
	} else if val2 == "USD" {
		rezult = (rates.Rates[val2] / rates.Rates[val1]) * valCount
	} else {
		rezult = (rates.Rates["USD"] / rates.Rates[val1]) * rates.Rates[val2] * valCount
	}

	fmt.Printf("%.2f %s = %.2f %s\n\n", valCount, val1, rezult, val2)

	return nil
}
