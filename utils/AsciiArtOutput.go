package utils

import (
	"ascii-art-color/asciiart"
	"fmt"
	"os"
	"strings"
)

func AsciiArtOutput() {
	args := os.Args
	banner := "standard"
	var text = ""
	if len(args)-1 > 2 {
		banner = args[3]
	}
	if len(args)-1 > 1 {
		text = args[2]
	}
	if len(args)-1 > 3 {

		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	output := asciiart.AsciiArt(text, banner)
	var nameFile string
	if len(os.Args[1]) > 9 && strings.HasPrefix(os.Args[1], "--output=") {
		nameFile = os.Args[1][9:]
	}
	file, err := os.Create(nameFile)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	defer file.Close()
	_, err = file.WriteString(output)

	if err != nil {
		fmt.Println("Error")
	}
}
