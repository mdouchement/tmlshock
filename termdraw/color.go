package termdraw

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var FlagColorUndefined lipgloss.TerminalColor = lipgloss.NoColor{}

func FlagColor(color string) lipgloss.TerminalColor {
	switch color {
	case "none":
		return FlagColorUndefined
	case "black":
		return lipgloss.Color("0")
	case "blue":
		return lipgloss.Color("4")
	case "cyan":
		return lipgloss.Color("6")
	case "dark-gray":
		return lipgloss.Color("8")
	case "green":
		return lipgloss.Color("2")
	case "light-green":
		return lipgloss.Color("10")
	case "light-blue":
		return lipgloss.Color("12")
	case "light-cyan":
		return lipgloss.Color("14")
	case "light-gray":
		return lipgloss.Color("15")
	case "light-magenta":
		return lipgloss.Color("13")
	case "light-red":
		return lipgloss.Color("9")
	case "light-yellow":
		return lipgloss.Color("11")
	case "magenta":
		return lipgloss.Color("5")
	case "red":
		return lipgloss.Color("1")
	case "white":
		return lipgloss.Color("7")
	case "yellow":
		return lipgloss.Color("3")
	default:
		if color == "" {
			return lipgloss.Color("5") // Default magenta
		}

		num, err := strconv.Atoi(color)
		if err == nil {
			return lipgloss.Color(strconv.Itoa(num)) // ANSI color
		}

		if !strings.HasPrefix(color, "#") {
			color = "#" + color // Try to convert to hexa color like ff00ff -> #ff00ff
		}

		return lipgloss.Color(color)
	}
}
