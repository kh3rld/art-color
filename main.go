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

var bannerMap = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

func main() {
	color := flag.String("color", "", "Color to apply to the substring.")
	flag.Parse()

	// Check if the number of arguments is valid
	if len(flag.Args()) < 1 || len(flag.Args()) > 3 {
		utils.PrintUsage()
		return
	}
	var stringInput string
	var substring string

	bannerFile := "standard.txt"
	isBanner := false

	if bannerMap[flag.Args()[len(flag.Args())-1]] && len(flag.Args()) > 1 {
		bannerFile = flag.Args()[len(flag.Args())-1] + ".txt"
		isBanner = true
	}

	if len(flag.Args()) == 3 && !bannerMap[flag.Args()[2]] {
		fmt.Println(flag.Args()[2], "is an Invalid Banner")
		return
	}

	substring = flag.Args()[0]

	if len(flag.Args()) == 1 {
		stringInput = flag.Args()[0]
	} else if len(flag.Args()) == 2 && !isBanner {
		stringInput = flag.Args()[1]
	} else if len(flag.Args()) == 3 {
		stringInput = flag.Args()[1]
	} else if len(flag.Args()) == 2 && isBanner {
		stringInput = flag.Args()[0]
	}

	// Variable to track if the flag was set
	var nameSet bool

	// Enforce the flag format to be used to be --color=<color>
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "color" {
			result := strings.Replace(os.Args[1], *color, "", 1)
			if !(result == "--color=") {
				nameSet = true
			}
		}
	})

	// defining usage and error handling
	if nameSet {
		utils.PrintUsage()
		return
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

	var asciioutput string
	if *color != "" {
		asciioutput = functions.AsciiArt(stringInput, substring, fileLines)
	} else {
		asciioutput = functions.AsciiArt(stringInput, substring, fileLines)
	}
	fmt.Print(asciioutput)
	// Print to console if output file not specified
}
