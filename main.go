package main

import "os"

// Gobble: fetch in go for fun

// Add or update a row in system.go
func main() {

	if len(os.Args) > 1 && os.Args[1] == "debug" {
		debug = true
	}

	rows := []Row{
		GetHostname(),
		GetMachine(),
		GetKernel(),
		GetDistro(),
		GetDe(),
		GetPackages(),
		GetShell(),
		GetCpu(),
		GetMemory(),
		GetUptime(),
	}
	PrintRows(rows)
	PrintPalette()
}
