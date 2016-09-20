// httping 0.1 - A tool to measure RTT on HTTP/S requests
// This software is distributed AS IS and has no warranty. This is merely a learning exercise and should not be used in production under any circumstances.
// This is my own work and not that of my employer, not is endorsed or supported by them in any conceivable way.
// Pedro Perez - pjperez@outlook.com
// Based on https://github.com/giigame/httping (Thanks!)

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/mgutz/ansi"
)

func main() {

	fmt.Println("\nhttping 0.1 - A tool to measure RTT on HTTP/S requests")
	fmt.Println("Help: httping help")

	if (len(os.Args) <= 1) || (os.Args[1] == "help") {
		fmt.Println("\n\nUsage: httping url[:port] [GET|HEAD]")
		os.Exit(1)
	}

	urlStr := os.Args[1]

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

	httpVerb := "GET"

	if len(os.Args) >= 3 {
		if os.Args[2] == "HEAD" {
			httpVerb = os.Args[2]
		} else {
			httpVerb = "GET"
		}
	} else {
		fmt.Printf("No HTTP Verb specified, falling back to HTTP\n\n")
	}

	fmt.Printf("HTTP %s to %s (%s):\n", httpVerb, url.Host, urlStr)
	ping(httpVerb, url)
}

func ping(httpVerb string, url *url.URL) {
	// This function loops indefinitely (TODO- Select number of iterations) and prints result on screen after each loop

	timeTotal := time.Duration(0)
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	successfulProbes := 0
	ansiRed := ansi.ColorCode("red+h:black")
	ansiGreen := ansi.ColorCode("green+h:black")
	reset := ansi.ColorCode("reset")
	result, err := client.Get(url.String())

	for i := 1; ; i++ {

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
			fmt.Println(ansiRed, "Fatal error!", err, reset)
		}

		timeTotal += responseTime

		// Calculate the downloaded bytes
		body, _ := ioutil.ReadAll(result.Body)
		bytes := len(body)

		// Print result on screen
		fmt.Printf("connected to %s, seq=%d, httpVerb=%s, httpStatus=%d, bytes=%d, RTT=%.2f ms\n", url.Host, i, httpVerb, result.StatusCode, bytes, float32(responseTime)/1e6)

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
				fmt.Println(ansiGreen)
				fmt.Println("Probes sent:", i, "\nSuccessful responses:", successfulProbes, "\nAverage response time:", timeAverage, reset)
				os.Exit(1)
			}
		}()
	}
}
