package main

import (
	"fmt"
	"runtime"
)

func showMemStats() {
	var memStats runtime.MemStats

	runtime.ReadMemStats(&memStats)

	fmt.Printf(
		"Alloc = %v MiB",
		memStats.Alloc/1024/1024,
	)

	fmt.Printf(
		"\tTotalAlloc = %v MiB",
		memStats.TotalAlloc/1024/1024,
	)

	fmt.Printf(
		"\tSys = %v MiB",
		memStats.Sys/1024/1024,
	)

	fmt.Printf(
		"\tNumGC = %v\n",
		memStats.NumGC,
	)
}

func main() {
	fmt.Println("Memory before:")
	showMemStats()

	s := make([]int, 10_000_000)

	for i := range s {
		s[i] = i
	}

	fmt.Println("Memory after:")
	showMemStats()
}
