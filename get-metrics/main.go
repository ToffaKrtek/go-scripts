package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	//"github.com/shirou/gopsutil/cpu"
)

const (
	csvFile  = "metrics.csv"
	interval = 15 * time.Second
)

func getMetrics() {
	for {
		loadAvg, err := load.Avg()
		if err != nil {
			fmt.Println("Error getting load average:", err)
			continue
		}
		memStat, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error getting virtual memory:", err)
			continue
		}
		diskStats, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Error getting disk usage:", err)
			continue
		}
		// Запись метрик в CSV
		err = writeMetricsToCSV(loadAvg, memStat, diskStats)
		if err != nil {
			fmt.Println("Error writing metrics to CSV:", err)
		}

		// Ждемs
		time.Sleep(interval)
	}
}

func clearCSV() {
	file, err := os.OpenFile(csvFile, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error clearing CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Timestamp", "Load1", "Load5", "Load15", "MemoryUsed", "MemoryTotal", "DiskUsed", "DiskTotal"})
	writer.Flush()
}

func writeMetricsToCSV(loadAvg *load.AvgStat, memStats *mem.VirtualMemoryStat, diskStats *disk.UsageStat) error {
	file, err := os.OpenFile(csvFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	timestamp := time.Now().Format(time.RFC3339)
	memoryUsedGB := roundToOneDecimal(float64(memStats.Used) / (1024 * 1024 * 1024))
	memoryTotalGB := roundToOneDecimal(float64(memStats.Total) / (1024 * 1024 * 1024))
	diskUsedGB := roundToOneDecimal(float64(diskStats.Used) / (1024 * 1024 * 1024))
	diskTotalGB := roundToOneDecimal(float64(diskStats.Total) / (1024 * 1024 * 1024))
	record := []string{
		timestamp,
		fmt.Sprintf("%f", loadAvg.Load1),
		fmt.Sprintf("%f", loadAvg.Load5),
		fmt.Sprintf("%f", loadAvg.Load15),
		fmt.Sprintf("%.1f", memoryUsedGB),
		fmt.Sprintf("%.1f", memoryTotalGB),
		fmt.Sprintf("%.1f", diskUsedGB),
		fmt.Sprintf("%.1f", diskTotalGB),
	}
	writer.Write(record)
	writer.Flush()

	return nil
}

func roundToOneDecimal(value float64) float64 {
	return float64(int(value*10+0.5)) / 10
}

func main() {
	clearCSV()
	getMetrics()
}
