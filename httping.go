// httping 0.9 - A tool to measure RTT on HTTP/S requests
// This software is distributed AS IS and has no warranty. This is merely a learning exercise and should not be used in production under any circumstances.
// This is my own work and not that of my employer, not is endorsed or supported by them in any conceivable way.
// Pedro Perez - pjperez@outlook.com
// Based on https://github.com/giigame/httping (Thanks!)

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"

	"strconv"

	"github.com/montanaflynn/stats"
)

const httpingVersion = "0.9"

// Reply is a data structure for the server replies
type Reply struct {
	Hostname string
	ClientIP string
	Time     time.Time
}

func main() {
	urlPtr := flag.String("url", "", "Requested URL")
	httpverbPtr := flag.String("httpverb", "GET", "HTTP Verb: GET or HEAD")
	countPtr := flag.Int("count", 10, "Number of requests to send")
	listenPtr := flag.Int("listen", 0, "Enable listener mode on specified port, e.g. '-r 80'")

	flag.Parse()

	urlStr := *urlPtr
	httpVerb := *httpverbPtr

	fmt.Println("\nhttping " + httpingVersion + " - A tool to measure RTT on HTTP/S requests")
	fmt.Println("Help: httping -h")

	// If listener mode is selected, ignore the rest of the args
	if *listenPtr > 0 {
		listenPort := strconv.Itoa(*listenPtr)
		fmt.Println("Listening on port " + listenPort)

		http.HandleFunc("/", serverRESPONSE)
		http.ListenAndServe(":"+listenPort, nil)

	}

	// Exit if URL is not specified, print usage
	if len(urlStr) < 1 {
		flag.Usage()
		fmt.Printf("\nYou haven't specified a URL to test!\n\n")

		os.Exit(1)
	}

	// Exit if the number of probes is zero, print usage
	if *countPtr < 1 {
		flag.Usage()
		fmt.Printf("\nNumber of probes has to be greater than 0!\n\n")

		os.Exit(1)
	}

	// Check what protocol has been specified in the URL by checking the first 7 or 8 chars.
	// If none specified, fall back to HTTP
	if len(urlStr) > 6 {
		if urlStr[:7] != "http://" {
			if urlStr[:8] != "https://" {
				if strings.Contains(urlStr, "://") {
					fmt.Println("\n\nWrong protocol specified, httping only supports HTTP and HTTPS")
					os.Exit(1)
				}
				fmt.Printf("\n\nNo protocol specified, falling back to HTTP\n\n")

				urlStr = "http://" + urlStr

			}
		}
	} else {
		fmt.Println()
		os.Exit(1)
	}

	// Parse URL and fail if the host can't be resolved.
	url, err := url.Parse(urlStr)

	if err != nil {
		fmt.Println("Cannot resolve: " + urlStr)
		os.Exit(1)
		return
	}

	fmt.Printf("HTTP %s to %s (%s):\n", httpVerb, url.Host, urlStr)
	ping(httpVerb, url, *countPtr)
}

func ping(httpVerb string, url *url.URL, count int) {
	// This function is responsible to send the requests, count the time and show statistics when finished

	// Initialise needed variables
	timeTotal := time.Duration(0)
	i := 1
	successfulProbes := 0
	var responseTimes []float64
	fBreak := 0

	// Change request timeout to 2 seconds
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	result, err := client.Get(url.String())

	// Send requests for url, "count" times
	for i = 1; count >= i && fBreak == 0; i++ {
		// Initialise variables
		timeStart := time.Now()
		responseTime := time.Since(timeStart)

		// Send GET or HEAD request, depending on httpVerb received
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
			fmt.Println("Timeout when connecting to", url)

		} else {
			// Add all the response times to calculate the average later
			timeTotal += responseTime

			// Calculate the downloaded bytes
			body, _ := ioutil.ReadAll(result.Body)
			bytes := len(body)

			// Print result on screen
			fmt.Printf("connected to %s, seq=%d, httpVerb=%s, httpStatus=%d, bytes=%d, RTT=%.2f ms\n", url, i, httpVerb, result.StatusCode, bytes, float32(responseTime)/1e6)

			// Count how many probes are successful, i.e. how many get a 200 HTTP StatusCode - If successful also add the result to a slice "responseTimes"
			if result.StatusCode == 200 {
				successfulProbes++
				responseTimes = append(responseTimes, float64(responseTime))
			}
		}

		time.Sleep(1e9)

		c := make(chan os.Signal, 1)

		signal.Notify(c, os.Interrupt)
		go func() {
			for sig := range c {
				_ = sig
				fBreak = 1
			}
		}()

	}

	// Let's calculate and spill some results
	// 1. Average response time
	timeAverage := time.Duration(int64(0))
	if successfulProbes > 0 {
		timeAverage = time.Duration(int64(timeTotal) / int64(successfulProbes))
	} else {
		fmt.Println("All probes failed")
		os.Exit(1)
	}

	// 2. Min and Max response times
	var biggest float64
	smallest := float64(1000000000)

	for _, v := range responseTimes {

		if v > biggest {
			biggest = v
		}

		if v < smallest {
			smallest = v
		}

	}

	// 3. Median response time
	median, _ := stats.Median(responseTimes)

	// 4. Percentile
	percentile90, _ := stats.Percentile(responseTimes, float64(90))
	percentile75, _ := stats.Percentile(responseTimes, float64(75))
	percentile50, _ := stats.Percentile(responseTimes, float64(50))
	percentile25, _ := stats.Percentile(responseTimes, float64(25))

	fmt.Println("\nProbes sent:", i-1, "\nSuccessful responses:", successfulProbes, "\n% of requests failed:", float64(100-(successfulProbes*100)/(i-1)), "\nMin response time:", time.Duration(smallest), "\nAverage response time:", timeAverage, "\nMedian response time:", time.Duration(median), "\nMax response time:", time.Duration(biggest))

	fmt.Println("\n90% of requests were faster than:", time.Duration(percentile90), "\n75% of requests were faster than:", time.Duration(percentile75), "\n50% of requests were faster than:", time.Duration(percentile50), "\n25% of requests were faster than:", time.Duration(percentile25))
}

func serverRESPONSE(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname() // Get the local hostname

	// Get the client's IP address.
	// RemoteAddr returns the client IP address with the port after a colon
	// We split the client IP + port based on colon(s) and only remove
	// after the last one, so we don't break IPv6
	clientsocket := r.RemoteAddr
	clientipMap := strings.Split(clientsocket, ":")
	clientipMap = clientipMap[:len(clientipMap)-1]
	clientip := strings.Join(clientipMap, ":")

	response := Reply{hostname, clientip, time.Now()} // Construct the response with the gathered data

	// Convert to json
	jsonRESPONSE, err := json.Marshal(response)
	if err != nil {
		log.Output(0, "json conversion failed")
	}

	io.WriteString(w, string(jsonRESPONSE)) // Send response back to client
}
