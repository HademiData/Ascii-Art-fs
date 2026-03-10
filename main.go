package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// ------------------- THINKERTOY PARSER -------------------
func parseThinkertoyBanner(banner []string) map[rune][]string {
	charMap := make(map[rune][]string)
	var block []string
	code := 32 // ASCII space

	for _, line := range banner {
		line = strings.TrimRight(line, "\r")
		if line == "" {
			if len(block) > 0 {
				charMap[rune(code)] = block
				block = []string{}
				code++
			}
		} else {
			block = append(block, line)
		}
	}
	if len(block) > 0 {
		charMap[rune(code)] = block
	}
	return charMap
}

// ------------------- PRINT THINKERTOY -------------------
func printThinkertoyArt(text string, charMap map[rune][]string) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		// find max height for this line
		maxHeight := 0
		for _, c := range line {
			if block, ok := charMap[c]; ok {
				if len(block) > maxHeight {
					maxHeight = len(block)
				}
			}
		}
		// print row by row
		for row := 0; row < maxHeight; row++ {
			for _, c := range line {
				block, ok := charMap[c]
				if ok {
					if row < len(block) {
						fmt.Print(block[row])
					} else {
						fmt.Print(strings.Repeat(" ", len(block[0])))
					}
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}

// ------------------- STANDARD / SHADOW -------------------
func getCharacter(lines []string, c rune) []string {
	index := int(c) - 32
	start := (index * 9) + 1
	return lines[start : start+8]
}

func printAsciiArt(text string, banner []string) {
	for row := 0; row < 8; row++ {
		for _, char := range text {
			block := getCharacter(banner, char)
			fmt.Print(block[row])
		}
		fmt.Println()
	}
}

// ------------------- MAIN -------------------
func main() {
	if len(os.Args) != 3 {
		log.Fatal("\nUsage: go run . [STRING] [BANNER]\n\nEX: go run . hello standard")
	}

	inputText := os.Args[1]
	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	bannerFile := os.Args[2] + ".txt"

	buffer, err := os.ReadFile(bannerFile)
	if err != nil {
		log.Fatal("Error reading banner file: ", err)
	}

	banner := strings.Split(string(buffer), "\n")

	switch strings.ToLower(os.Args[2]) {
	case "thinkertoy":
		charMap := parseThinkertoyBanner(banner)
		printThinkertoyArt(inputText, charMap)
	default: // standard, shadow
		lines := strings.Split(inputText, "\n")
		for _, line := range lines {
			if line == "" {
				fmt.Println()
				continue
			}
			printAsciiArt(line, banner)
		}
	}
}