package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

var start = time.Now()

func EastwesserBadge(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "🚀 Eastwesser | Go version: %s | Uptime: %s",
		runtime.Version(),
		time.Since(start).Round(time.Second))
}

func main() {
	http.HandleFunc("/", EastwesserBadge)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
