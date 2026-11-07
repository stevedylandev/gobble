package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"syscall"
	"runtime"
)

// Gobble: fetch in go for fun
// Shell env
//
// Order:
// Host - Done
// Machine /sys/class/dmi/id/sys_vendor + /sys/class/dmi/id/product_name
// Kernal - /proc/sys/kernal/osrelease
// Distro - Done
// Packages - /lib/apk/db/installed count P:
// Shell - env
// Terminal?
// Uptime syscall.Sysinfo()
// CPU - /proc/cpuinfo
// GPU - 
// Memory - Done 
//
// Need to figure out colors and printing design

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error fetching hostname: %s", err)
		os.Exit(1)
	}
	user := os.Getenv("USER")
	if user == "" {
		fmt.Printf("Error fetching User: %s", err)
		os.Exit(1)
	}
	fmt.Printf("%s@%s\n", user ,hostname)
	fmt.Println("-----------------------------")
	distro, err := getDistro()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Disto: ", distro)
	memory, err := getMemory()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Memory: ", memory)

	uptime, err := getUptime()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Uptime: ", uptime)

	fmt.Println("------------TEST--------")
	var si syscall.Sysinfo_t
	syscall.Sysinfo(&si)
	fmt.Println(si.Uptime)
	fmt.Println(runtime.GOARCH)


	os.Exit(0)
}

func getDistro() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(value, `"`), nil
		}
	}
	return "", scanner.Err()
}

func getMemory() (string, error) {
	var memTotal string
	var memAvailable string
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			value := strings.TrimPrefix(line, "MemTotal:")

			memTotal = strings.Trim(value, ` kB`)
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			value := strings.TrimPrefix(line, "MemAvailable:")

			memAvailable = strings.Trim(value, ` kB`)
		}
	}
	memTotalInt, err := strconv.Atoi(memTotal)
	if err != nil {
		fmt.Println("Problem convering mem total to int", err)
		os.Exit(1)
	}
	memAvailableInt, err := strconv.Atoi(memAvailable)
	if err != nil {
		fmt.Println("Problem convering mem available to int", err)
		os.Exit(1)
	}

	memTotalGb := float64(memTotalInt) / float64(1024) / float64(1024)
	memAvailableGb := float64(memAvailableInt) / float64(1024) / float64(1024)
	memUsed := memTotalGb - memAvailableGb

	result := fmt.Sprintf("%.2f Gb / %.2f Gb", memUsed, memTotalGb)
	return result, scanner.Err()
}

func getUptime() (string, error) {
	content, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "", err
	}

	fields := strings.Fields(string(content))
	uptimeSeconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return "", err
	}

	uptime := time.Duration(uptimeSeconds * float64(time.Second))

	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60

	return fmt.Sprintf("%d hours, %d mins", hours, minutes), nil
}
