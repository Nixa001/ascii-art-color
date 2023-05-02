package utils

import "log"

var Colors = map[string]string{
	"Init":   "\033[0m",
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"Gray":   "\033[37m",
	"Orange": "\033[38;2;255;140;0m",
	"Pink":   "\033[38;2;255;105;108m",
	"Indigo": "\033[38;2;138;43;226m",
	"Brown":  "\033[38;2;165;42;42m",
}

func getColorCode(color string) string {
	switch color {
	case "red":
		return Colors["Red"]
	case "green":
		return Colors["Green"]
	case "yellow":
		return Colors["Yellow"]
	case "blue":
		return Colors["Blue"]
	case "purple":
		return Colors["Purple"]
	case "cyan":
		return Colors["Cyan"]
	case "gray":
		return Colors["Gray"]
	case "orange":
		return Colors["Orange"]
	case "pink":
		return Colors["Pink"]
	case "indigo":
		return Colors["Indigo"]
	case "brown":
		return Colors["Brown"]
	default:
		log.Fatalf("Invalid Color !!")
		return Colors["Init"]
	}
}
