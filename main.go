package main

// Gobble: fetch in go for fun

// Add or update a row in system.go
func main() {
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
