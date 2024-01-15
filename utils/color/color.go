package color

import (
	"fmt"
	"log"
	"runtime"
)

var reset = "\033[0m"

var red = "\033[31m"

var green = "\033[32m"

var yellow = "\033[33m"

var blue = "\033[34m"

var purple = "\033[35m"

var cyan = "\033[36m"

var gray = "\033[37m"

var white = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}

// Color type constants.
type Color string

const (
	// RED represents red color.
	RED Color = "Red"
	// GREEN represents green color.
	GREEN Color = "Green"
	// YELLOW represents yellow color.
	YELLOW Color = "Yellow"
	// BLUE represents blue color.
	BLUE Color = "Blue"
	// PURPLE represents purple color.
	PURPLE Color = "Purple"
	// CYAN represents cyan color.
	CYAN Color = "Cyan"
	// GRAY represents gray color.
	GRAY Color = "Gray"
	// WHITE represents white color.
	WHITE Color = "White"
)

// This can colorize your str
// Available color are:  Red, Green, Yellow, Blue, Purple, Cyan, Gray, White.
func This(color Color, str ...any) string {
	switch color {
	case RED:
		return red + fmt.Sprint(str...) + reset
	case GREEN:
		return green + fmt.Sprint(str...) + reset
	case YELLOW:
		return yellow + fmt.Sprint(str...) + reset
	case BLUE:
		return blue + fmt.Sprint(str...) + reset
	case PURPLE:
		return purple + fmt.Sprint(str...) + reset
	case CYAN:
		return cyan + fmt.Sprint(str...) + reset
	case GRAY:
		return gray + fmt.Sprint(str...) + reset
	case WHITE:
		return white + fmt.Sprint(str...) + reset
	}
	log.Println("[stdout] invalid or unavailable color")
	return fmt.Sprint(str...)
}

// ThisF can colorize your str
// Available color are:  Red, Green, Yellow, Blue, Purple, Cyan, Gray, White.
func ThisF(color Color, str string, args ...any) string {
	switch color {
	case RED:
		return fmt.Sprintf("%s%s%s", red, fmt.Sprintf(str, args...), reset)
	case GREEN:
		return fmt.Sprintf("%s%s%s", green, fmt.Sprintf(str, args...), reset)
	case YELLOW:
		return fmt.Sprintf("%s%s%s", yellow, fmt.Sprintf(str, args...), reset)
	case BLUE:
		return fmt.Sprintf("%s%s%s", blue, fmt.Sprintf(str, args...), reset)
	case PURPLE:
		return fmt.Sprintf("%s%s%s", purple, fmt.Sprintf(str, args...), reset)
	case CYAN:
		return fmt.Sprintf("%s%s%s", cyan, fmt.Sprintf(str, args...), reset)
	case GRAY:
		return fmt.Sprintf("%s%s%s", gray, fmt.Sprintf(str, args...), reset)
	case WHITE:
		return fmt.Sprintf("%s%s%s", white, fmt.Sprintf(str, args...), reset)
	}
	log.Println("[stdout] invalid or unavailable color")
	return str
}
