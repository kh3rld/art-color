package functions

import (
	"fmt"
	"strings"
)

func AsciiArt(stringInput, substring string, fileLine []string) string {
	result := ""
	index := []int{}

	// replacing every instance of new line with the newline character (\n)
	stringInput = strings.Replace(stringInput, "\n", "\\n", -1)

	if !ValidSentence(stringInput) {
		return ""
	}

	// slicing the input based on the presence of the string "\n"
	words := strings.Split(stringInput, "\\n")

	empty := EmptyArray(words)
	if empty != "false" {
		return empty
	}

	sub := false
	subEnd := 0

	for _, word := range words {
		if substring != "" {
			index = subStringIndex(word, substring)
		}
		if word == "" {
			result += "\n"
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(word); j++ {
					start := (int(word[j]-' ') * 9) + 1 // calculating the begining of a character based on data from standard.txt

					if subIndex(j, index) {
						sub = true
						subEnd = j + len(substring)
						result += "\033[38;2;0;0;255m"
					}

					result += fileLine[start+i]

					if sub && subEnd == j+1 {
						sub = false
						result += "\033[0m"
					}

				}
				result += "\n"
			}
		}
	}
	return result
}

func ValidSentence(word string) bool {
	for _, letter := range word {
		if !(letter >= ' ' && letter <= '~') {
			fmt.Println("Error, character", string(letter), "is an invalid character!!!!")
			return false
		}
	}
	return true
}

func EmptyArray(words []string) string {
	result := ""

	for i, word := range words {
		if word != "" {
			result = "false"
			return result
		}
		if i != 0 {
			result += "\n"
		}
	}
	return result
}

func subStringIndex(s, subStr string) []int {
	index := []int{}
	position := 0

	for {
		idx := strings.Index(s, subStr)
		if idx == -1 {
			break
		}
		index = append(index, idx+position)
		s = s[len(subStr)+idx:]
		position += len(subStr) + idx
	}
	fmt.Println(index)

	return index
}

func subIndex(num int, index []int) bool {
	for _, idx := range index {
		if idx == num {
			return true
		}
	}
	return false
}
