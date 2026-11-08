package main

import (
	"fmt"
	"os"
	internal "github.com/stevedylandev/gobble/internal"
)

// Gobble: fetch in go for fun
// Shell env
//
// Order:
// Host - Done
// Machine - Done
// Kernel - Done
// Distro - Done
// Packages - Done
// Shell - Done
// Terminal?
// Uptime - Done
// CPU - Done
// GPU -
// Memory - Done
//
// Need to figure out colors and printing design

func main() {
	hostname, err := internal.GetHostname()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Printf(hostname)

	fmt.Println("-----------------------------")

	machine, err := internal.GetMachine()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Machine: ", machine)

	kernel, err := internal.GetKernel()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Kernel: ", kernel)

	distro, err := internal.GetDistro()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Disto: ", distro)

	packages, err := internal.GetPackages()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Packages: ", packages, "(apk)")

	shell, err := internal.GetShell()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Shell: ", shell)

	cpu, err := internal.GetCpu()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("CPU: ", cpu)

	memory, err := internal.GetMemory()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Memory: ", memory)

	uptime, err := internal.GetUptime()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Uptime: ", uptime)

	os.Exit(0)
}

