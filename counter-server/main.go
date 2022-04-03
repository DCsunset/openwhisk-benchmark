package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	port := 8000
	fmt.Printf("Listen at port %d\n", port)

	counter := 0
	var maxThroughput float64 = 0

	start := false
	var startTime time.Time

	f, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write result
	writer := csv.NewWriter(f)
	writer.Write([]string{"time", "throughput"})
	samples := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter += 1
		if !start {
			fmt.Println("Start!")
			startTime = time.Now()
			start = true
		} else {
			elapsed := time.Since(startTime).Seconds()
			// Sample every 5 seconds
			if elapsed > 5 {
				throughput := float64(counter) / elapsed
				if throughput > maxThroughput {
					maxThroughput = throughput
				}
				samples += 5
				writer.Write([]string{strconv.Itoa(samples), fmt.Sprintf("%f", throughput)})
				writer.Flush()

				// Reset
				counter = 0
				startTime = time.Now()

				fmt.Printf("Throughput: %f, Max: %f\n", throughput, maxThroughput)
			}
		}
		fmt.Fprintln(w, "OK!")
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}
