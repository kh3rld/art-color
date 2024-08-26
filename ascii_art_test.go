package main

import (
	"os"
	"strings"
	"testing"

	"ascii-art-output/functions"
)

type inputs struct {
	str      string
	subStr   string
	path     string
	color    string
}

func TestAsciiArtStandard(t *testing.T) {
	inputFile := "standard.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Error reading file '%s': %v", inputFile, err)
	}

	fileLines := strings.Split(string(file), "\n")

	testCases := []inputs{
		{"hello", "ll", "test_cases/expectedOutput1.txt", "red"},
		{"you & me", "&", "test_cases/expectedOutput2.txt", "green"},
		{"Hello\\nThere", "e", "test_cases/expectedOutput3.txt", "blue"},
		{"{Hello There}", "o T", "test_cases/expectedOutput4.txt", "yellow"},
		{"1Hello 2There", "el", "test_cases/expectedOutput5.txt", "orange"},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			expectedContent, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", tc.path, tc.str, err)
			}

			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(tc.str, tc.subStr, fileLines, tc.color)
			
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", tc.str, expectedContentStr, result)
			}
		})
	}
}

func TestAsciiArtShadow(t *testing.T) {
	inputFile := "shadow.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Error reading file '%s': %v", inputFile, err)
	}

	fileLines := strings.Split(string(file), "\n")

	testCases := []inputs{
		{"hello world", "rl", "test_cases/expectedOutput6.txt", "green"},
		{"123", "2", "test_cases/expectedOutput7.txt", "yellow"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "LMN", "test_cases/expectedOutput8.txt", "red"},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			expectedContent, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", tc.path, tc.str, err)
			}

			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(tc.str, tc.subStr, fileLines, tc.color)			
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", tc.str, expectedContentStr, result)
			}
		})
	}
}

func TestAsciiArtThinkertoy(t *testing.T) {
	inputFile := "thinkertoy.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Error reading file '%s': %v", inputFile, err)
	}

	fileLines := strings.Split(string(file), "\n")

	testCases := []inputs{
		{"nice 2 meet you", "meet", "test_cases/expectedOutput9.txt", "red"},
		{"/(\\\")", "(", "test_cases/expectedOutput10.txt", "yellow"},
		{"\"#$%&/()*+,-./", "./", "test_cases/expectedOutput11.txt", "green"},
		{"It's Working", "t's Wor", "test_cases/expectedOutput12.txt", "orange"},
	}

	for _, tc := range testCases {
		t.Run(tc.str, func(t *testing.T) {
			expectedContent, err := os.ReadFile(tc.path)
			if err != nil {
				t.Fatalf("Failed to read expected output file '%s' for input '%s': %v", tc.path, tc.str, err)
			}

			expectedContentStr := string(expectedContent)

			result := functions.AsciiArt(tc.str, tc.subStr, fileLines, tc.color)
			
			if result != expectedContentStr {
				t.Errorf("For input:\n'%s'\nExpected:\n%s\n but got:\n%s", tc.str, expectedContentStr, result)
			}
		})
	}
}
