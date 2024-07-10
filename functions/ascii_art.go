package functions

import (
	"fmt"
	"strings"
)

// AsciiArt generates ASCII art with optional substring highlighting and color.
func AsciiArt(stringInput, substring string, fileLines []string, color string) string {
	result := ""
	index := []int{}

	// Replace newline characters with "\\n"
	stringInput = strings.Replace(stringInput, "\n", "\\n", -1)

	if !ValidSentence(stringInput) {
		return ""
	}

	// Split stringInput into words based on "\\n"
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
					start := (int(word[j]-' ') * 9) + 1 // Calculate start based on ' ' to ASCII offset

					if subIndex(j, index) {
						sub = true
						subEnd = j + len(substring)
						result += GetColorEscapeCode(color)
					}

					result += fileLines[start+i]

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

// ANSI escape code for the specified color
func GetColorEscapeCode(color string) string {
	switch color {
	case "red":
		return "\033[38;2;255;0;0m"
	case "green":
		return "\033[38;2;0;255;0m"
	case "blue":
		return "\033[38;2;0;0;255m"
	case "yellow":
		return "\033[38;2;255;255;0m"
	case "cyan":
		return "\033[38;2;0;255;255m"
	case "magenta":
		return "\033[38;2;255;0;255m"
	case "white":
		return "\033[38;2;255;255;255m"
	case "black":
		return "\033[38;2;0;0;0m"
	case "orange":
		return "\033[38;2;255;165;0m"
	case "purple":
		return "\033[38;2;128;0;128m"
	case "brown":
		return "\033[38;2;165;42;42m"
	case "pink":
		return "\033[38;2;255;192;203m"
	case "gray":
		return "\033[38;2;128;128;128m"
	case "navy":
		return "\033[38;2;0;0;128m"
	case "teal":
		return "\033[38;2;0;128;128m"
	case "olive":
		return "\033[38;2;128;128;0m"
	default:
		return "\033[0m"
	}
}

// ValidSentence checks if a sentence contains valid characters
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
