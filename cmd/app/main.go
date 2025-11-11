package main

import (
	"bufio"
	"currency_converter/internal/converter"
	"fmt"
	"os"
)

func main() {
	converter.Logo()

	scanner := bufio.NewReader(os.Stdin)

	for {
		choice := converter.Menu()
		switch choice {
		case 1:
			_ = converter.Converter()
		case 0:
			fmt.Println("Нажмите Enter, чтобы выйти...")
			_, _ = scanner.ReadString('\n')
			_, _ = scanner.ReadString('\n')
			return
		}

		fmt.Println("Нажмите Enter, чтобы продолжить...")
		_, _ = scanner.ReadString('\n')
		_, _ = scanner.ReadString('\n')
	}
}
