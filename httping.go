// httping 0.10.0 - A tool to measure RTT on HTTP/S requests
// This software is distributed AS IS and has no warranty. This is merely a learning exercise and should not be used in production under any circumstances.
// This is my own work and not that of my employer, not is endorsed or supported by them in any conceivable way.
// Pedro Perez - pjperez@outlook.com
// Based on https://github.com/giigame/httping (Thanks!)

package main

import (
	"crypto/tls"
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
	"github.com/rapid7/go-get-proxied/proxy"
)

const (
	httpingVersion = "0.10.0"
	timeFormat     = "2006-01-02 15:04:05 MST"
)

// Custom logger type
type AppLogger struct {
	*log.Logger
}

var (
	infoLogger  *AppLogger
	warnLogger  *AppLogger
	errorLogger *AppLogger
)

// Initialize loggers
func initLoggers(jsonOutput bool) {
	flags := log.Lmsgprefix
	output := io.Discard // Default to no output for JSON mode
	if !jsonOutput {
		output = os.Stdout
	}

	infoLogger = &AppLogger{log.New(output, "INFO: ", flags)}
	warnLogger = &AppLogger{log.New(output, "WARN: ", flags)}
	errorLogger = &AppLogger{log.New(os.Stderr, "ERROR: ", flags)} // Keep errors in stderr
}

// Log functions with consistent formatting
func (l *AppLogger) Info(format string, v ...interface{}) {
	l.Printf("[%s] %s", time.Now().Format(timeFormat), fmt.Sprintf(format, v...))
}

func (l *AppLogger) Warn(format string, v ...interface{}) {
	l.Printf("[%s] %s", time.Now().Format(timeFormat), fmt.Sprintf(format, v...))
}

func (l *AppLogger) Error(format string, v ...interface{}) {
	l.Printf("[%s] %s", time.Now().Format(timeFormat), fmt.Sprintf(format, v...))
}

// Reply is a data structure for the server replies
type Reply struct {
	Hostname string
	ClientIP string
	Time     time.Time
}

// Result is the struct to generate the metadata json results
type Result struct {
	Host        string  `json:"host"`
	HTTPVerb    string  `json:"httpVerb"`
	HostHeaders string  `json:"hostHeader"`
	Seq         int     `json:"seq"`
	HTTPStatus  int     `json:"httpStatus"`
	Bytes       int     `json:"bytes"`
	RTT         float32 `json:"rtt"`
}

func main() {
	// Available flags
	urlPtr := flag.String("url", "", "Requested URL")
	httpverbPtr := flag.String("httpverb", "GET", "HTTP Verb: Only GET or HEAD supported at the moment")
	countPtr := flag.Int("count", 10, "Number of requests to send [0 means infinite]")
	listenPtr := flag.Int("listen", 0, "Enable listener mode on specified port, e.g. '-r 80'")
	timeoutPtr := flag.Int("timeout", 2000, "Timeout in milliseconds")
	hostHeaderPtr := flag.String("hostheader", "", "Optional: Host header")
	jsonResultsPtr := flag.Bool("json", false, "If true, produces output in json format")
	followRedirectsPtr := flag.Bool("followredirects", false, "If true, follows redirects, which may result in higher RTT")
	noProxyPtr := flag.Bool("noproxy", false, "If true, ignores system proxy settings")
	insecureTLS := flag.Bool("insecure", false, "Skip TLS certificate verification")

	flag.Parse()

	initLoggers(*jsonResultsPtr)

	// Log startup information
	infoLogger.Info("httping %s starting", httpingVersion)
	defer infoLogger.Info("httping completed")

	urlStr := *urlPtr
	httpVerb := *httpverbPtr
	jsonResults := *jsonResultsPtr
	noProxy := *noProxyPtr
	followRedirects := *followRedirectsPtr

	if !jsonResults {
		infoLogger.Info("HTTP %s to %s", httpVerb, urlStr)
		infoLogger.Info("Use -h for help")
	}

	// If listener mode is selected, ignore the rest of the args
	if *listenPtr > 0 {
		listenPort := strconv.Itoa(*listenPtr)
		infoLogger.Info("Starting listener on port %s", listenPort)

		http.HandleFunc("/", serverRESPONSE)
		if err := http.ListenAndServe(":"+listenPort, nil); err != nil {
			errorLogger.Error("Listener failed: %v", err)
			os.Exit(1)
		}
		return
	}

	// Validate URL
	if len(urlStr) < 1 {
		flag.Usage()
		errorLogger.Error("No URL specified")
		os.Exit(1)
	}

	// Validate timeout
	if *timeoutPtr < 0 {
		flag.Usage()
		errorLogger.Error("Timeout must be greater than 0")
		os.Exit(1)
	}
	timeout := time.Duration(*timeoutPtr) * time.Millisecond

	// Handle protocol
	if len(urlStr) > 6 {
		if urlStr[:7] != "http://" && urlStr[:8] != "https://" {
			if strings.Contains(urlStr, "://") {
				errorLogger.Error("Unsupported protocol (only HTTP/HTTPS allowed)")
				os.Exit(1)
			}
			warnLogger.Warn("No protocol specified, defaulting to HTTP")
			urlStr = "http://" + urlStr
		}
	} else {
		errorLogger.Error("Invalid URL format")
		os.Exit(1)
	}

	// Parse URL
	url, err := url.Parse(urlStr)
	if err != nil {
		errorLogger.Error("URL parse error: %v", err)
		os.Exit(1)
	}

	// Set host header
	hostHeader := *hostHeaderPtr
	if hostHeader == "" {
		hostHeader = url.Host
	}

	infoLogger.Info("Starting HTTP %s to %s (%s)", httpVerb, url.Host, urlStr)
	ping(httpVerb, url, *countPtr, timeout, hostHeader, jsonResults, followRedirects, noProxy, *insecureTLS)
}

func ping(httpVerb string, url *url.URL, count int, timeout time.Duration, hostHeader string, jsonResults bool, followRedirects bool, noProxy bool, insecureTLS bool) {
	timeTotal := time.Duration(0)
	i := 1
	successfulProbes := 0
	var responseTimes []float64
	fBreak := 0

	// Setup redirect policy
	var checkRedirectFunc func(req *http.Request, via []*http.Request) error
	if !followRedirects {
		checkRedirectFunc = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	// Setup signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		warnLogger.Warn("Interrupt signal received, stopping...")
		fBreak = 1
	}()

	for i = 1; (count >= i || count < 1) && fBreak == 0; i++ {
		transport := &http.Transport{
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureTLS,
				MinVersion:         tls.VersionTLS12,
			},
		}

		proxyInformation := "proxy=None"
		if !noProxy {
			p := proxy.NewProvider("").GetProxy(httpVerb, url.String())
			if p != nil {
				proxyInformation = fmt.Sprintf("proxy=%s", p)
				transport.Proxy = http.ProxyURL(p.URL())
			}
		}

		client := http.Client{
			Timeout:       timeout,
			Transport:     transport,
			CheckRedirect: checkRedirectFunc,
		}

		request, err := http.NewRequest(httpVerb, url.String(), nil)
		if err != nil {
			errorLogger.Error("Request creation failed: %v", err)
			continue
		}
		request.Host = hostHeader
		request.Header.Set("User-Agent", "httping "+httpingVersion)

		timeStart := time.Now()
		result, errRequest := client.Do(request)
		responseTime := time.Since(timeStart)

		if err != nil || errRequest != nil {
			if tlsErr, ok := errRequest.(*tls.CertificateVerificationError); ok {
				errorLogger.Error("TLS verification failed: %v", tlsErr)
				for i, cert := range tlsErr.UnverifiedCertificates {
					errorLogger.Error("Cert %d: Subject: %s, Issuer: %s, Expires: %s",
						i+1,
						cert.Subject.CommonName,
						cert.Issuer.CommonName,
						cert.NotAfter.Format("2006-01-02"))
				}
				if !insecureTLS {
					errorLogger.Error("Use -insecure to bypass certificate validation")
					return
				}
				warnLogger.Warn("Proceeding with insecure connection")
			} else {
				warnLogger.Warn("Request failed to %s | %s | Error: %v", url, proxyInformation, errRequest)
			}
			continue
		}

		body, err := ioutil.ReadAll(result.Body)
		result.Body.Close()
		if err != nil {
			warnLogger.Warn("Failed to read response body: %v", err)
		}
		bytes := len(body)

		if jsonResults {
			results := &Result{
				Host:        url.Host,
				HTTPVerb:    httpVerb,
				HostHeaders: hostHeader,
				Seq:         i,
				HTTPStatus:  result.StatusCode,
				Bytes:       bytes,
				RTT:         float32(responseTime) / 1e6,
			}
			jsonData, _ := json.Marshal(results)
			fmt.Println(string(jsonData))
		} else {
			infoLogger.Info("Connected to %s, %s, seq=%d, status=%d, bytes=%d, rtt=%.2fms",
				url, proxyInformation, i, result.StatusCode, bytes, float32(responseTime)/1e6)
		}

		if result.StatusCode >= 100 && result.StatusCode < 400 {
			successfulProbes++
			responseTimes = append(responseTimes, float64(responseTime))
		}

		if ((count - i) > 1) || (count <= 0) {
			time.Sleep(1 * time.Second)
		}
	}

	if successfulProbes == 0 {
		errorLogger.Error("All probes failed")
		os.Exit(1)
	}

	// Calculate statistics
	timeAverage := time.Duration(int64(timeTotal) / int64(successfulProbes))
	min, max := calculateMinMax(responseTimes)
	median, _ := stats.Median(responseTimes)
	p90, _ := stats.Percentile(responseTimes, 90)
	p75, _ := stats.Percentile(responseTimes, 75)
	p50, _ := stats.Percentile(responseTimes, 50)
	p25, _ := stats.Percentile(responseTimes, 25)

	if !jsonResults {
		failureRate := float64(100 - (successfulProbes*100)/(i-1))
		infoLogger.Info("Results - Probes: %d, Success: %d, Failed: %.1f%%",
			i-1, successfulProbes, failureRate)
		infoLogger.Info("Timing - Min: %v, Avg: %v, Med: %v, Max: %v",
			time.Duration(min), timeAverage, time.Duration(median), time.Duration(max))
		infoLogger.Info("Percentiles - P90: %v, P75: %v, P50: %v, P25: %v",
			time.Duration(p90), time.Duration(p75), time.Duration(p50), time.Duration(p25))
	}
}

func calculateMinMax(times []float64) (min, max float64) {
	min = times[0]
	max = times[0]
	for _, t := range times {
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	return min, max
}

func serverRESPONSE(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	clientSocket := r.RemoteAddr
	clientParts := strings.Split(clientSocket, ":")
	clientIP := strings.Join(clientParts[:len(clientParts)-1], ":")

	response := Reply{
		Hostname: hostname,
		ClientIP: clientIP,
		Time:     time.Now(),
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		errorLogger.Error("JSON marshal failed: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		errorLogger.Error("Response write failed: %v", err)
	}
}
