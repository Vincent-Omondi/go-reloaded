package transformWords

import (
	"fmt"
	"strconv"
)

func TransformWord(words []string, i int, transform func(string) string) []string {
	if i == 0 {
		return words
	}

	derivative := words[i] == "(low)" || words[i] == "(up)" || words[i] == "(cap)" || words[i] == "(bin)" || words[i] == "(hex)"

	derivativeWithN := words[i] == "(low," || words[i] == "(up," || words[i] == "(cap,"

	nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if derivative {
		// check if the last word is a derivative
		if i == len(words)-1 {
			words[i-1] = transform(words[i-1])
			return words[:len(words)-1]
		}
		// if the current word is a derivative
		words[i-1] = transform(words[i-1])
		return append(words[:i], words[i+1:]...)
	}

	// handle derivatives with numbers
	if derivativeWithN {
		val := ""
		word := words[i+1]
		for i, value := range word {
			for _, m := range nums {
				if string(value) == m {
					val += string(word[i])
				}
			}
		}
		number, _ := strconv.Atoi(val)
		for j := 1; j <= number; j++ {
			words[i-j] = transform(words[i-j])
		}
		return append(words[:i], words[i+2:]...)

	}

	return words
}

// convert hexadecimal string to decimal string
func ConvertHexToInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

// convert binary string to decimal string
func ConvertBinToInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}
