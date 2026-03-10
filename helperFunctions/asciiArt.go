package helperFunctions

import (
	"fmt"
)

func PrintAsciiArt(text string, banner []string) {
	for row := 0; row < 8; row++ {
		for _, char := range text {
			block := GetCharacter(banner, char)
			fmt.Print(block[row])
		}
		fmt.Println()
	}
}
