// httping 0.2 - A tool to measure RTT on HTTP/S requests
// This software is distributed AS IS and has no warranty. This is merely a learning exercise and should not be used in production under any circumstances.
// This is my own work and not that of my employer, not is endorsed or supported by them in any conceivable way.
// Pedro Perez - pjperez@outlook.com
// Based on https://github.com/giigame/httping (Thanks!)

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {

	httpverbPtr := flag.String("httpverb", "GET", "HTTP Verb: GET or HEAD")
	countPtr := flag.Int("count", 10, "Number of requests to send")

	flag.Parse()

	fmt.Println("\nhttping 0.2 - A tool to measure RTT on HTTP/S requests")
	fmt.Println("Help: httping -h")

	urlStr := flag.Args()[0]

	if urlStr[:7] != "http://" {
		if urlStr[:8] != "https://" {
			if strings.Contains(urlStr, "://") {
				fmt.Println("\n\nWrong protocol specified, httping supports HTTP and HTTPS")
				os.Exit(1)
			}
			fmt.Printf("\n\nNo protocol specified, falling back to HTTP\n\n")
			urlStr = "http://" + urlStr
		}
	}

	url, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Cannot resolve: " + urlStr)
		os.Exit(1)
		return
	}

	httpVerb := *httpverbPtr

	fmt.Printf("HTTP %s to %s (%s):\n", httpVerb, url.Host, urlStr)
	ping(httpVerb, url, *countPtr)
}

func ping(httpVerb string, url *url.URL, count int) {
	// This function loops indefinitely (TODO- Select number of iterations) and prints result on screen after each loop

	timeTotal := time.Duration(0)
	timeout := time.Duration(2 * time.Second)
	i := 1
	client := http.Client{
		Timeout: timeout,
	}
	successfulProbes := 0
	result, err := client.Get(url.String())

	for i = 1; count >= i; i++ {

		timeStart := time.Now()
		responseTime := time.Since(timeStart)
		// Send GET or HEAD request
		if httpVerb == "GET" {
			timeStart = time.Now()
			result, err = client.Get(url.String())
			responseTime = time.Since(timeStart)
		} else {
			timeStart = time.Now()
			result, err = client.Head(url.String())
			responseTime = time.Since(timeStart)
		}

		if err != nil {
			fmt.Println("Fatal error!", err)
		}

		timeTotal += responseTime

		// Calculate the downloaded bytes
		body, _ := ioutil.ReadAll(result.Body)
		bytes := len(body)

		// Print result on screen
		fmt.Printf("connected to %s, seq=%d, httpVerb=%s, httpStatus=%d, bytes=%d, RTT=%.2f ms\n", url, i, httpVerb, result.StatusCode, bytes, float32(responseTime)/1e6)

		// Count how many probes are successful, i.e. how many get a 200 HTTP StatusCode
		if result.StatusCode == 200 {
			successfulProbes++
		}

		time.Sleep(1e9)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				timeAverage := time.Duration(int64(timeTotal) / int64(i))
				_ = sig
				fmt.Println("\nProbes sent:", i, "\nSuccessful responses:", successfulProbes, "\nAverage response time:", timeAverage)
				os.Exit(1)
			}
		}()
	}

	timeAverage := time.Duration(int64(timeTotal) / int64(i))

	fmt.Println("\nProbes sent:", i-1, "\nSuccessful responses:", successfulProbes, "\nAverage response time:", timeAverage)
}
