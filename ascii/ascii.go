package ascii

import (
	"strings"
)

func Ascii(word string, ban [][]string) string {
	// Разделение по знакам \n
	symbols := strings.Split(word, "\\n")

	// Разделение по фактическим символам новой строки
	symbols = strings.Split(strings.Join(symbols, "\n"), "\n")

	// Если дан символ новой строки и ничего больше
	hasSymbol := false
	for _, v := range symbols {
		if v != "" {
			hasSymbol = true
		}
	}

	totalStr := ""

	if !hasSymbol {
		for i := 0; i < len(symbols)-1; i++ {
			totalStr += "\n"
		}
		return totalStr
	}

	// Вывод ASCII символов
	for _, word := range symbols {
		if word != "" {
			for j := 0; j < 8; j++ {
				for _, char := range word {
					if char == 32 {
						totalStr += "        "
					} else {
						totalStr += ban[char-32][j]
					}
				}
				totalStr += "\n"
			}
		} else {
			totalStr += "\n"
		}
	}
	return totalStr
}
