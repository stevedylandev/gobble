package main

import (
	internal "github.com/stevedylandev/gobble/internal"
)

// Gobble: fetch in go for fun

func main() {
	rows := []internal.Row {
		internal.GetHostname(),
		{"Machine", internal.GetMachine()},
		{"Kernel", internal.GetKernel()},
		{"Distro", internal.GetDistro()},
		{"Packages", internal.GetPackages()},
		{"Shell", internal.GetShell()},
		{"CPU", internal.GetCpu()},
		{"Memory", internal.GetMemory()},
		{"Uptime", internal.GetUptime()},
	}
	internal.PrintRows(rows, false)
}

