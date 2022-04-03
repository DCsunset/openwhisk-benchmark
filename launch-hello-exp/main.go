package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/apache/openwhisk-client-go/whisk"
)

func main() {
	// Share transport
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	wskClient, err := whisk.NewClient(client, nil)
	if err != nil {
		fmt.Println("whisk newclient():", err)
		os.Exit(-1)
	}

	warmUp := 15
	expLength := 90
	coolDown := 15

	counter := 0
	// errCount := make(map[string]int)
	errCount := 0
	var latencies []int64

	start := time.Now()

	for {
		if time.Since(start) >= time.Duration(expLength)*time.Second {
			break
		}

		m := make(map[string]interface{})

		s := time.Now()
		_, resp, err := wskClient.Actions.Invoke("hello-go", m, true, true)
		if err != nil {
			// msg := err.Error()
			// count, ok := errCount[msg]
			// if ok {
			// 	errCount[msg] = count + 1
			// } else {
			// 	errCount[msg] = 1
			// }
			fmt.Printf("Error: %s\n", err.Error())
			errCount++
			continue
		}
		// Close to enable reuse of TCP connection
		resp.Body.Close()

		e := time.Now()
		sinceStart := e.Sub(start)

		if sinceStart < time.Duration(warmUp)*time.Second || sinceStart > time.Duration(expLength-coolDown)*time.Second {
			continue
		}

		counter++
		latency := e.Sub(s).Milliseconds()
		latencies = append(latencies, latency)
	}

	var avgLatency float64 = 0
	for _, lat := range latencies {
		avgLatency += float64(lat)
	}
	avgLatency /= float64(len(latencies))

	args := os.Args[1:]
	name := ""
	if len(args) > 0 {
		name = fmt.Sprintf("-%s", args[0])
	}
	f, err := os.OpenFile(fmt.Sprintf("result%s.txt", name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	throughputLog := fmt.Sprintf("throughput: %f\n", float64(counter)/float64(expLength-warmUp-coolDown))
	latencyLog := fmt.Sprintf("latency: %f\n", avgLatency)
	errLog := fmt.Sprintf("errors: %d\n", errCount)

	f.WriteString(throughputLog)
	f.WriteString(latencyLog)
	f.WriteString(errLog)
}
