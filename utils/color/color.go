package color

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// reset color
var reset  = "\033[0m"

// Red color
var red    = "\033[31m"

// Green color
var green  = "\033[32m"

// Yellow color
var yellow = "\033[33m"

// Blue color
var blue   = "\033[34m"

// Purple color
var purple = "\033[35m"

// Cyan color
var cyan   = "\033[36m"

// Gray color
var gray   = "\033[37m"

// White color
var white  = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		reset  = ""
		red    = ""
		green  = ""
		yellow = ""
		blue   = ""
		purple = ""
		cyan   = ""
		gray   = ""
		white  = ""
	}
}

// This can colorize your str
// Available color are:  Red, Green, Yellow, Blue, Purple, Cyan, Gray, White
func This(color, str string) string {
	
	switch cases.Title(language.English).String(strings.ToLower(color)) {
	case "Red":
		return fmt.Sprintf("%s%s%s", red, str, reset)
	case "Green":
		return fmt.Sprintf("%s%s%s", green, str, reset)
	case "Yellow":
		return fmt.Sprintf("%s%s%s", yellow, str, reset)
	case "Blue":
		return fmt.Sprintf("%s%s%s", blue, str, reset)
	case "Purple":
		return fmt.Sprintf("%s%s%s", purple, str, reset)
	case "Cyan":
		return fmt.Sprintf("%s%s%s", cyan, str, reset)
	case "Gray":
		return fmt.Sprintf("%s%s%s", gray, str, reset)
	case "White":
		return fmt.Sprintf("%s%s%s", white, str, reset)
	}
	log.Println("[stdout] invalid or unavailable color")
	return str
}

// ThisF can colorize your str
// Available color are:  Red, Green, Yellow, Blue, Purple, Cyan, Gray, White
func ThisF(color, str string, args ...any) string {
	switch cases.Title(language.English).String(strings.ToLower(color)) {
	case "Red":
		return fmt.Sprintf("%s%s%s", red, fmt.Sprintf(str, args...), reset)
	case "Green":
		return fmt.Sprintf("%s%s%s", green, fmt.Sprintf(str, args...), reset)
	case "Yellow":
		return fmt.Sprintf("%s%s%s", yellow, fmt.Sprintf(str, args...), reset)
	case "Blue":
		return fmt.Sprintf("%s%s%s", blue, fmt.Sprintf(str, args...), reset)
	case "Purple":
		return fmt.Sprintf("%s%s%s", purple, fmt.Sprintf(str, args...), reset)
	case "Cyan":
		return fmt.Sprintf("%s%s%s", cyan, fmt.Sprintf(str, args...), reset)
	case "Gray":
		return fmt.Sprintf("%s%s%s", gray, fmt.Sprintf(str, args...), reset)
	case "White":
		return fmt.Sprintf("%s%s%s", white, fmt.Sprintf(str, args...), reset)
	}
	log.Println("[stdout] invalid or unavailable color")
	return str
}