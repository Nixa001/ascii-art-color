package main

import (
	// "ascii-art-color/asciiart"
	"ascii-art-color/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	// args := os.Args

	// if len(args)-1 == 1 {
	// 	fmt.Print(asciiart.AsciiArt(args[1], "standard"))
	// }
	var check = true
	var option string
	if len(os.Args)-1 > 0 {
		option = os.Args[1]
	}
	switch check {
	case strings.HasPrefix(option, "--output="):
		utils.AsciiArtOutput()

	case strings.HasPrefix(option, "--color="):
		utils.AsciiArtColor()

	case len(os.Args)-1 < 3:
		fmt.Print(utils.AsciiArtFs())
	default:
		fmt.Println("INVALIDE INPUT")
		fmt.Println("Try one of these formats")
		fmt.Println("\nUsage: go run . [OPTION] [STRING]")
		fmt.Println("\nEX: go run . --color=COLOR LetterToColor Text")
		fmt.Println("\nEX: go run . --output=<fileName.txt> something standard")
		fmt.Println("\nEX: go run . TEXT BANNER)")
	}

}
