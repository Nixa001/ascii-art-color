package asciiart

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var HeightChar = 9
var tabChar = map[int][]string{}

func AsciiArt(textInput, bannerName string) string {
	var tabChars []string
	scanner, err := ioutil.ReadFile("./banner/" + bannerName + ".txt")
	if err != nil {
		fmt.Println ("Invalid Banner!!!")
		os.Exit(0)
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
	return generateAsciiArt(textInput)

}

func generateAsciiArt(input string) string {
	var result string
	inputLines := strings.Split(input, "\r\n")
	for _, inputLine := range inputLines {
		for i := 1; i < HeightChar; i++ {
			for _, char := range inputLine {
				if int(char) >= 32 && int(char) <= 126 {
					chars := int(char) - 32
					line := tabChar[chars][i]
					result += string(line)
				} else {
					return ("\"" + string(char) + "\"" + " is not printable in Ascii-Art")
				}
			}
			result += "\n"
		}
	}
	return result
}
