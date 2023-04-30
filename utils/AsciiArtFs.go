package utils

import (
	"ascii-art-color/asciiart"
	"os"
)

func AsciiArtFs() string {
	args := os.Args
	if len(args)-1 > 0 {
		value := "standard"
		if len(args)-1 == 2 {
			value = os.Args[2]
		}
		return asciiart.AsciiArt(os.Args[1], value)

	}
	return ""
}
