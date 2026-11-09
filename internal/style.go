package internal

import (
	"fmt"
)

const (
	Reset   = "\033[0m"
	Bold    = "\033[1m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"
)

type Color struct {
	Name string
	Code string
}

var Palette = []Color{
	{"Red", Red},
	{"Green", Green},
	{"Yellow", Yellow},
	{"Blue", Blue},
	{"Magenta", Magenta},
	{"Cyan", Cyan},
	{"White", White},
	{"Gray", Gray},
}

func Colorize(color, s string) string {
	return color + s + Reset
}

func PrintRows(rows []Row) {
	width := labelWidth(rows)

	fmt.Println()
	for _, row := range rows {
		label := Colorize(Bold+Yellow, fmt.Sprintf("%-*s", width, row.Key))
		line := "  " + label + "    " + row.Value

		fmt.Println(line)
	}
	fmt.Println()
}

func labelWidth(rows []Row) int {
	width := 0
	for _, row := range rows {
		if w := len(row.Key) + 1; w > width {
			width = w
		}
	}
	return width
}

func PrintPalette() {
	fmt.Print("  ")
	for _, c := range Palette {
		swatch := Colorize(c.Code, "●  ")
		fmt.Print(swatch)
	}
	fmt.Println()
}
