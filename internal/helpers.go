package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func GetHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	user := os.Getenv("USER")
	if user == "" {
		return "", err
	}
	return fmt.Sprintf("%s@%s\n", user, hostname), nil
}

func GetDistro() (string, error) {
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

func GetMachine() (string, error) {
	brandFile, err := os.ReadFile("/sys/class/dmi/id/sys_vendor")
	if err != nil {
		return "", err
	}
	productFile, err := os.ReadFile("/sys/class/dmi/id/product_name")
	if err != nil {
		return "", err
	}

	brand := strings.TrimSpace(string(brandFile))
	product := strings.TrimSpace(string(productFile))

	return fmt.Sprintf("%s %s", brand, product), nil
}

func GetKernel() (string, error) {
	kernelFile, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "", err
	}
	kernel := strings.TrimSpace(string(kernelFile))
	return kernel, nil
}

func GetPackages() (int, error) {
	// Packages - /lib/apk/db/installed count P:
	var count int
	file, err := os.Open("/lib/apk/db/installed")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "P:") {
			count++
		}
	}

	return count, nil
}

func GetMemory() (string, error) {
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

func GetUptime() (string, error) {
	var si syscall.Sysinfo_t
	syscall.Sysinfo(&si)

	uptime := time.Duration(float64(si.Uptime) * float64(time.Second))

	hours := int(uptime.Hours())
	minutes := int(uptime.Minutes()) % 60

	return fmt.Sprintf("%d hours, %d mins", hours, minutes), nil
}
