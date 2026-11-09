package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"path/filepath"
)

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

func GetDistro() string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			value := strings.TrimPrefix(line, "PRETTY_NAME=")
			return strings.Trim(value, `"`)
		}
	}
	return ""
}

func GetMachine() string {
	brandFile, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		return ""
	}
	productFile, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		return ""
	}

	brand := strings.TrimSpace(string(brandFile))
	product := strings.TrimSpace(string(productFile))

	return fmt.Sprintf("%s %s", brand, product)
}

func GetKernel() string {
	kernelFile, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return ""
	}
	kernel := strings.TrimSpace(string(kernelFile))
	return kernel
}

func GetPackages() string {
	var count int
	file, err := os.Open("/lib/apk/db/installed")
	if err != nil {
		return ""
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

	return stringCount + " (apk)"
}

func GetShell() string {
	shellEnv := os.Getenv("SHELL")
	if shellEnv == "" {
		return ""
	}

	shell := filepath.Base(shellEnv)

	return shell
}

func GetCpu() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "model name") {
			value := strings.TrimPrefix(line, "model name")
			cleanSpace := strings.TrimSpace(value)
			valueClean := strings.TrimPrefix(cleanSpace, ": ")
			return valueClean
		}
	}
	return ""
}

func GetMemory() string {
	var memTotal string
	var memAvailable string
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return ""
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

	result := fmt.Sprintf("%.2f GiB / %.2f GiB", memUsed, memTotalGb)
	return result
}

func GetUptime() string {
	var si syscall.Sysinfo_t
	syscall.Sysinfo(&si)

	uptime := time.Duration(float64(si.Uptime) * float64(time.Second))

	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60

	return fmt.Sprintf("%d hours, %d mins", hours, minutes)
}
