package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var tabChar = map[int][]string{}
var HeightChar = 9
var color string

func AsciiArtColor() {
	args := os.Args
	if len(args)-1 < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	ReadBanner()
	color = getColorCode(strings.ToLower(os.Args[1][8:]))
	text := os.Args[2]
	letterToColor := text
	if len(os.Args)-1 > 2 {
		text = os.Args[3]
		letterToColor = os.Args[2]
	}
	GenerateAscii(text, letterToColor)
}
func ReadBanner() {
	banner := "banner/standard.txt"
	var tabChars []string
	scanner, err := ioutil.ReadFile(banner)
	if err != nil {
		fmt.Println("Invalid banner")
		return
	}
	data := bufio.NewScanner(strings.NewReader(string(scanner)))
	for data.Scan() {
		lines := data.Text()
		tabChars = append(tabChars, lines)
	}
	characters := len(tabChars) / HeightChar
	for i := 0; i < characters; i++ {
		charLines := tabChars[i*HeightChar : (i+1)*HeightChar]
		tabChar[int(i)] = charLines
	}
}
func GenerateAscii(input, letterToColor string) {
	var output string
	divInput := strings.Split(input, "\\n")
	for _, InputPart := range divInput {
		for i := 1; i < HeightChar; i++ {
			for _, runeValue := range InputPart {
				if int(runeValue) >= 32 && int(runeValue) <= 126 {
					NumCharsInTab := int(runeValue) - 32
					outputLine := ""
					if strings.ContainsRune(letterToColor, runeValue) {
						outputLine = (color + tabChar[NumCharsInTab][i] + Colors["Init"])
					} else {
						outputLine = tabChar[NumCharsInTab][i]

					}
					if outputLine != "" {
						output += outputLine
					}
				} else {
					fmt.Println("Invalid Input")
					return
				}
			}
			output += "\n"
		}
		if InputPart == "" {
			fmt.Println()
		} else {
			fmt.Print(output)
		}
		output = ""
	}
}
