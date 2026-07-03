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
INFO: [2026-07-03 12:11:13 CEST] httping 0.2.0 starting
INFO: [2026-07-03 12:11:13 CEST] HTTP GET to https://www.bing.com
INFO: [2026-07-03 12:11:13 CEST] Use -h for help
INFO: [2026-07-03 12:11:13 CEST] Starting HTTP GET to www.bing.com (https://www.bing.com)
INFO: [2026-07-03 12:11:13 CEST] Connected to https://www.bing.com, proxy=None, seq=1, status=200, bytes=64275, rtt=133.11ms
INFO: [2026-07-03 12:11:14 CEST] Connected to https://www.bing.com, proxy=None, seq=2, status=200, bytes=63476, rtt=56.63ms
INFO: [2026-07-03 12:11:15 CEST] Connected to https://www.bing.com, proxy=None, seq=3, status=200, bytes=63599, rtt=58.43ms
INFO: [2026-07-03 12:11:16 CEST] Connected to https://www.bing.com, proxy=None, seq=4, status=200, bytes=63478, rtt=56.08ms
INFO: [2026-07-03 12:11:17 CEST] Connected to https://www.bing.com, proxy=None, seq=5, status=200, bytes=63457, rtt=49.36ms
INFO: [2026-07-03 12:11:18 CEST] Connected to https://www.bing.com, proxy=None, seq=6, status=200, bytes=63472, rtt=46.21ms
INFO: [2026-07-03 12:11:20 CEST] Connected to https://www.bing.com, proxy=None, seq=7, status=200, bytes=64379, rtt=48.43ms
INFO: [2026-07-03 12:11:21 CEST] Connected to https://www.bing.com, proxy=None, seq=8, status=200, bytes=63176, rtt=51.43ms
INFO: [2026-07-03 12:11:22 CEST] Connected to https://www.bing.com, proxy=None, seq=9, status=200, bytes=63560, rtt=51.42ms
INFO: [2026-07-03 12:11:23 CEST] Connected to https://www.bing.com, proxy=None, seq=10, status=200, bytes=64295, rtt=44.43ms
INFO: [2026-07-03 12:11:23 CEST] Results - Probes: 10, Success: 10, Failed: 0.0%
INFO: [2026-07-03 12:11:23 CEST] Timing - Min: 44.4264ms, Avg: 59.55321ms, Med: 51.42815ms, Max: 133.1134ms
INFO: [2026-07-03 12:11:23 CEST] Percentiles - P90: 58.4312ms, P75: 56.35255ms, P50: 51.42815ms, P25: 47.31795ms
INFO: [2026-07-03 12:11:23 CEST] httping completed
```

#### Example 2:

```
$ ./httping -url https://www.bing.com -count 5 -json
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":1,"httpStatus":200,"bytes":63201,"rtt":132.3643}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":2,"httpStatus":200,"bytes":62969,"rtt":57.5249}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":3,"httpStatus":200,"bytes":63574,"rtt":49.6071}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":4,"httpStatus":200,"bytes":63349,"rtt":48.6354}
{"host":"www.bing.com","httpVerb":"GET","hostHeader":"www.bing.com","seq":5,"httpStatus":200,"bytes":63517,"rtt":52.891}
```

#### Example 3:

Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1
INFO: [2025-03-26 12:06:21 CET] httping 0.2.0 starting
INFO: [2025-03-26 12:06:21 CET] HTTP GET to https://self-signed.badssl.com
INFO: [2025-03-26 12:06:21 CET] Use -h for help
INFO: [2025-03-26 12:06:21 CET] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
WARN: [2025-03-26 12:06:22 CET] Request failed to https://self-signed.badssl.com | proxy=None | Error: Get "https://self-signed.badssl.com": tls: failed to verify certificate: x509: certificate signed by unknown authority
ERROR: [2025-03-26 12:06:22 CET] All probes failed
```

Bypass Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1 -insecure
INFO: [2025-03-26 12:06:25 CET] httping 0.2.0 starting
INFO: [2025-03-26 12:06:25 CET] HTTP GET to https://self-signed.badssl.com
INFO: [2025-03-26 12:06:25 CET] Use -h for help
INFO: [2025-03-26 12:06:25 CET] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
INFO: [2025-03-26 12:06:25 CET] Connected to https://self-signed.badssl.com, proxy=None, seq=1, status=200, bytes=502, rtt=559.65ms
INFO: [2025-03-26 12:06:25 CET] Results - Probes: 1, Success: 1, Failed: 0.0%
INFO: [2025-03-26 12:06:25 CET] Timing - Min: 559.64504ms, Avg: 559.64504ms, Med: 559.64504ms, Max: 559.64504ms
INFO: [2025-03-26 12:06:25 CET] Percentiles - P90: 559.64504ms, P75: 559.64504ms, P50: 559.64504ms, P25: 559.64504ms
INFO: [2025-03-26 12:06:25 CET] httping completed
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
