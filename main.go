package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ascii-art-output/functions"
	utils "ascii-art-output/utils"
)

func main() {
	// Define and parse flags
	output := flag.String("output", "output.txt", "File that stores the output.")
	flag.Parse()

	// Check if the number of arguments is valid
	if len(flag.Args()) < 1 || len(flag.Args()) > 2 {
		utils.PrintUsage()
		return
	}

	stringInput := flag.Args()[0]

	// Variable to track if the flag was set
	var nameSet bool
	var flagSet bool = false

	// Enforce the flag format to be used to be --output=<filename.txt>
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "output" {
			flagSet = true
			result := strings.Replace(os.Args[1], *output, "", 1)
			if !(result == "--output=") {
				nameSet = true
			}
		}
	})
	// defining usage and error handling
	if nameSet {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Validate output file extension
	if !strings.HasSuffix(*output, ".txt") {
		utils.PrintUsage()
		return
	}

	bannerFile := "standard.txt"
	if len(flag.Args()) == 2 {
		bannerFile = flag.Args()[1] + ".txt"
	}

	err := utils.ValidateFileChecksum(bannerFile)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return
	}

	file, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error opening", bannerFile, ":", err)
		return
	}
	fileLines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	asciiOutput := functions.AsciiArt(stringInput, "he", fileLines)

	// Write to output file
	err = os.WriteFile(*output, []byte(asciiOutput), 0o644)
	if err != nil {
		fmt.Println("Error writing to", *output, ":", err)
		return
	}

	// Print to console if output file not specified
	if !flagSet {
		fmt.Print(asciiOutput)
	}
}
