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
// Kernal - /proc/sys/kernal/osrelease
// Distro - Done
// Packages - /lib/apk/db/installed count P:
// Shell - env
// Terminal?
// Uptime - Done
// CPU - /proc/cpuinfo
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
	distro, err := internal.GetDistro()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Disto: ", distro)
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

