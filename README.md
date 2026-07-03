# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

### Latest release

- [Download from Github](https://github.com/pjperez/httping/releases)

### Requirements
Golang >=1.21 ::  needed for signal.NotifyContext and io.ReadAll (the previous '>1.3' note was stale)
External library requirement: [github.com/montanaflynn/stats](https://github.com/montanaflynn/stats)

### Installing
#### From source
```
go install github.com/pjperez/httping@latest
```

You'll then find httping.exe in your $GOPATH/bin directory

#### Binary
Check the [latest release](https://github.com/pjperez/httping/releases) and download the standalone binary from there, it doesn't need installation and it's portable.

### Usage
#### Server mode (listener)
httping -listen port [OPTIONS]

It will start a listener on *port* and reply to any GET request as follows:

    {
      "Hostname":"Server_Hostname",
      "ClientIP":"Client_IP",
      "Time":"Current_Timestamp"
    }

#### Client mode (original httping)
httping -url requested_url [OPTIONS]

##### Options
```
 -count int
    	Number of requests to send [0 means infinite] (default 10)
  -followredirects
    	If true, follows redirects, which may result in higher RTT
  -hostheader string
    	Optional: Host header
  -httpverb string
    	HTTP Verb: Only GET or HEAD supported at the moment (default "GET")
  -insecure
    	Skip TLS certificate verification
  -json
    	If true, produces output in json format
  -listen int
    	Enable listener mode on specified port, e.g. '-r 80'
  -noproxy
    	If true, ignores system proxy settings
  -timeout int
    	Timeout in milliseconds (default 2000)
  -url string
    	Requested URL
```

#### Example

```
$ ./httping -url https://www.bing.com -count 10
INFO: [2026-07-03 12:33:23 CEST] httping 0.2.0 starting
INFO: [2026-07-03 12:33:23 CEST] HTTP GET to https://www.bing.com
INFO: [2026-07-03 12:33:23 CEST] Use -h for help
INFO: [2026-07-03 12:33:23 CEST] Starting HTTP GET to www.bing.com (https://www.bing.com)
INFO: [2026-07-03 12:33:23 CEST] Connected to https://www.bing.com, proxy=None, seq=1, status=200, bytes=63478, rtt=128.73ms
INFO: [2026-07-03 12:33:24 CEST] Connected to https://www.bing.com, proxy=None, seq=2, status=200, bytes=63517, rtt=54.77ms
INFO: [2026-07-03 12:33:26 CEST] Connected to https://www.bing.com, proxy=None, seq=3, status=200, bytes=63988, rtt=58.94ms
INFO: [2026-07-03 12:33:27 CEST] Connected to https://www.bing.com, proxy=None, seq=4, status=200, bytes=63472, rtt=76.83ms
INFO: [2026-07-03 12:33:28 CEST] Connected to https://www.bing.com, proxy=None, seq=5, status=200, bytes=64430, rtt=51.98ms
INFO: [2026-07-03 12:33:29 CEST] Connected to https://www.bing.com, proxy=None, seq=6, status=200, bytes=63148, rtt=47.09ms
INFO: [2026-07-03 12:33:30 CEST] Connected to https://www.bing.com, proxy=None, seq=7, status=200, bytes=64280, rtt=53.34ms
INFO: [2026-07-03 12:33:31 CEST] Connected to https://www.bing.com, proxy=None, seq=8, status=200, bytes=63563, rtt=48.30ms
INFO: [2026-07-03 12:33:32 CEST] Connected to https://www.bing.com, proxy=None, seq=9, status=200, bytes=64304, rtt=55.96ms
INFO: [2026-07-03 12:33:33 CEST] Connected to https://www.bing.com, proxy=None, seq=10, status=200, bytes=63124, rtt=60.93ms
INFO: [2026-07-03 12:33:33 CEST] Results - Probes: 10, Success: 10, Failed: 0.0%
INFO: [2026-07-03 12:33:33 CEST] Timing - Min: 47.0925ms, Avg: 63.68689ms, Med: 55.3653ms, Max: 128.7281ms
INFO: [2026-07-03 12:33:33 CEST] Percentiles - P90: 76.8331ms, P75: 59.9334ms, P50: 55.3653ms, P25: 50.14035ms
INFO: [2026-07-03 12:33:33 CEST] httping completed
```

#### Example 2:

```
$ ./httping -url https://www.bing.com -count 5 -json
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":1,"httpStatus":200,"bytes":64375,"rtt":112.3862}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":2,"httpStatus":200,"bytes":63565,"rtt":49.0889}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":3,"httpStatus":200,"bytes":63350,"rtt":48.025}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":4,"httpStatus":200,"bytes":63350,"rtt":48.778}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":5,"httpStatus":200,"bytes":63597,"rtt":43.863}
```

#### Example 3:

Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1
INFO: [2026-07-03 12:33:38 CEST] httping 0.2.0 starting
INFO: [2026-07-03 12:33:38 CEST] HTTP GET to https://self-signed.badssl.com
INFO: [2026-07-03 12:33:38 CEST] Use -h for help
INFO: [2026-07-03 12:33:38 CEST] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
WARN: [2026-07-03 12:33:38 CEST] Request failed to https://self-signed.badssl.com | proxy=None | Error: Get "https://self-signed.badssl.com": tls: failed to verify certificate: x509: certificate signed by unknown authority
ERROR: [2026-07-03 12:33:38 CEST] All probes failed
```

Bypass Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1 -insecure
INFO: [2026-07-03 12:34:00 CEST] httping 0.2.0 starting
INFO: [2026-07-03 12:34:00 CEST] HTTP GET to https://self-signed.badssl.com
INFO: [2026-07-03 12:34:00 CEST] Use -h for help
INFO: [2026-07-03 12:34:00 CEST] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
INFO: [2026-07-03 12:34:01 CEST] Connected to https://self-signed.badssl.com, proxy=None, seq=1, status=200, bytes=502, rtt=637.19ms
INFO: [2026-07-03 12:34:01 CEST] Results - Probes: 1, Success: 1, Failed: 0.0%
INFO: [2026-07-03 12:34:01 CEST] Timing - Min: 637.1897ms, Avg: 637.1897ms, Med: 637.1897ms, Max: 637.1897ms
INFO: [2026-07-03 12:34:01 CEST] Percentiles - P90: 637.1897ms, P75: 637.1897ms, P50: 637.1897ms, P25: 637.1897ms
INFO: [2026-07-03 12:34:01 CEST] httping completed
```

### Help
httping help

### Warranty
This is just a learning exercise, hence it is distributed AS IS, with no warranty.

You should not use this software in production under any circumstances. It's not intended to and it has not been thoroughly tested and problably it's not very accurate.

# Contributing

In order to contribute you can do any of the below:

1. Open a new Issue
2. Fork the project, add a feature or fix and send us a PR
3. Go to the [Projects](https://github.com/pjperez/httping/projects) section and comment on a card/issue from the TODO list.
4. Open an issue
5. Just get in touch if you have other ideas :)
6. Do whatever you want actually! :)
