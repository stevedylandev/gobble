package internal

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type Row struct {
	Key   string
	Value string
}

// Create a new function that returns a row 
//
// func MyInfo() Row {
// 		var myValue string
//
// 		// get the value here
//
// 		return Row{Key: "Foo", Value: myValue}
// }

func GetHostname() Row {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}
	user := os.Getenv("USER")
	if user == "" {
		user = "Unknown"
	}
	host := fmt.Sprintf("%s@%s", user, hostname)
	return Row{Key: "Host", Value: host}
}

func GetDistro() Row {
	var distro string
	file, err := os.Open("/etc/os-release")
	if err != nil {
		distro = "Unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			distro = strings.Trim(value, `"`)
		}
	}
	return Row{Key: "Distro", Value: distro}
}

func GetMachine() Row {
	var machine string
	brandFile, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		machine = "Unknown"
	}
	productFile, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		machine = "Unknown"
	}

	brand := strings.TrimSpace(string(brandFile))
	product := strings.TrimSpace(string(productFile))

	machine = fmt.Sprintf("%s %s", brand, product)
	return Row{Key: "Machine", Value: machine}
}

func GetKernel() Row {
	var kernel string
	kernelFile, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		kernel = "Unknown"
	}
	kernel = strings.TrimSpace(string(kernelFile))
	return Row{Key: "Kernel", Value: kernel}
}

func GetPackages() Row {
	var packages string
	var count int
	file, err := os.Open("/lib/apk/db/installed")
	if err != nil {
		packages = "Unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "P:") {
			count++
		}
	}

	stringCount := strconv.Itoa(count)
	packages = stringCount + " (apk)"
	return Row{Key: "Packages", Value: packages}
}

func GetDe() Row {
	var de string
	de = os.Getenv("DESKTOP_SESSION")
	if de == "" {
		de = "Unknown"
	}
	return Row{Key: "DE", Value: de}
}

func GetShell() Row {
	var shell string
	shellEnv := os.Getenv("SHELL")
	if shellEnv == "" {
		shell = "Unknown"
	}

	shell = filepath.Base(shellEnv)

	return Row{Key: "Shell", Value: shell}
}

func GetCpu() Row {
	var cpu string
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		cpu = "Unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			value := strings.TrimPrefix(line, "model name")
			cleanSpace := strings.TrimSpace(value)
			cpu = strings.TrimPrefix(cleanSpace, ": ")
		}
	}
	return Row{Key: "CPU", Value: cpu}
}

func GetMemory() Row {
	var memory string
	var memTotal string
	var memAvailable string
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		memory = "Unknown"
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

	memory = fmt.Sprintf("%.2f GiB / %.2f GiB", memUsed, memTotalGb)
	return Row{Key: "Memory", Value: memory}
}

func GetUptime() Row {
	var uptime string
	var si syscall.Sysinfo_t
	syscall.Sysinfo(&si)

	uptimeValues := time.Duration(float64(si.Uptime) * float64(time.Second))

	days := int(uptimeValues.Hours() / 24)
	hours := int(uptimeValues.Hours()) % 24
	minutes := int(uptimeValues.Minutes()) % 60

	if days > 0 {
		uptime = fmt.Sprintf("%d days, %d hours, %d mins", days, hours, minutes)
	} else {
		uptime = fmt.Sprintf("%d hours, %d mins", hours, minutes)
	}
	return Row{Key: "Uptime", Value: uptime}
}
