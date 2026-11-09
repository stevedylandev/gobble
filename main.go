package main

import (
	system "github.com/stevedylandev/gobble/internal"
)

// Gobble: fetch in go for fun

func main() {
	rows := []system.Row {
		system.GetHostname(),
		system.GetMachine(),
		system.GetKernel(),
		system.GetDistro(),
		system.GetDe(),
		system.GetPackages(),
		system.GetShell(),
		system.GetCpu(),
		system.GetMemory(),
		system.GetUptime(),
	}
	system.PrintRows(rows)
	system.PrintPalette()
}

