package internal

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
	Cyan  = "\033[36m"
)

type Row struct {
	Key   string
	Value string
}

func Colorize(color, s string) string {
	return color + s + Reset
}

func PrintRows(rows []Row, typewriter bool) {
	width := labelWidth(rows)

	for _, row := range rows {
		label := Colorize(Bold+Green, fmt.Sprintf("%-*s", width, row.Key))
		line := label + " " + row.Value

		if typewriter {
			TypeString(line)
			fmt.Println() 
		} else {
			fmt.Println(line)
		}
	}
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

func TypeString(s string) {
	for _, char := range s {
		fmt.Print(string(char))

		min := 20
		max := 130
		delay := time.Duration(rand.Intn(max-min)+min) * time.Millisecond
		time.Sleep(delay)
	}
	fmt.Print("\n")
}
