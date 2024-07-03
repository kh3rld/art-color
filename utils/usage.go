package utils

import "fmt"

func PrintUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println("\nExample: go run . --output=<fileName.txt> something standard")
}
