package utils

import "fmt"

func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("\nEX: go run . --color=<color> <substring to be colored> 'something'")
}
