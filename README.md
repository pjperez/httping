# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

### Latest release

- [Download from Github](https://github.com/pjperez/httping/releases)

### Requirements
Golang >1.3 ::  http.Client Timeout wasn't available in previous versions    
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
$ ./httping -url example.com -count 10
INFO: [2025-03-26 12:03:28 CET] httping 0.10.0 starting
INFO: [2025-03-26 12:03:28 CET] HTTP GET to example.com
INFO: [2025-03-26 12:03:28 CET] Use -h for help
WARN: [2025-03-26 12:03:28 CET] No protocol specified, defaulting to HTTP
INFO: [2025-03-26 12:03:28 CET] Starting HTTP GET to example.com (http://example.com)
INFO: [2025-03-26 12:03:28 CET] Connected to http://example.com, proxy=None, seq=1, status=200, bytes=1256, rtt=229.63ms
INFO: [2025-03-26 12:03:30 CET] Connected to http://example.com, proxy=None, seq=2, status=200, bytes=1256, rtt=366.04ms
INFO: [2025-03-26 12:03:31 CET] Connected to http://example.com, proxy=None, seq=3, status=200, bytes=1256, rtt=221.98ms
INFO: [2025-03-26 12:03:32 CET] Connected to http://example.com, proxy=None, seq=4, status=200, bytes=1256, rtt=360.16ms
INFO: [2025-03-26 12:03:34 CET] Connected to http://example.com, proxy=None, seq=5, status=200, bytes=1256, rtt=357.55ms
INFO: [2025-03-26 12:03:35 CET] Connected to http://example.com, proxy=None, seq=6, status=200, bytes=1256, rtt=227.71ms
INFO: [2025-03-26 12:03:36 CET] Connected to http://example.com, proxy=None, seq=7, status=200, bytes=1256, rtt=216.90ms
INFO: [2025-03-26 12:03:38 CET] Connected to http://example.com, proxy=None, seq=8, status=200, bytes=1256, rtt=349.76ms
INFO: [2025-03-26 12:03:39 CET] Connected to http://example.com, proxy=None, seq=9, status=200, bytes=1256, rtt=368.16ms
INFO: [2025-03-26 12:03:39 CET] Connected to http://example.com, proxy=None, seq=10, status=200, bytes=1256, rtt=221.70ms
INFO: [2025-03-26 12:03:39 CET] Results - Probes: 10, Success: 10, Failed: 0.0%
INFO: [2025-03-26 12:03:39 CET] Timing - Min: 216.897026ms, Avg: 0s, Med: 289.693866ms, Max: 368.162583ms
INFO: [2025-03-26 12:03:39 CET] Percentiles - P90: 366.039083ms, P75: 358.853235ms, P50: 229.631897ms, P25: 221.840585ms
INFO: [2025-03-26 12:03:39 CET] httping completed
```

#### Example 2:

```
$ ./httping -url example.com -count 5 -json
{"host":"example.com","httpVerb":"GET","hostHeader":"example.com","seq":1,"httpStatus":200,"bytes":1256,"rtt":362.5289}
{"host":"example.com","httpVerb":"GET","hostHeader":"example.com","seq":2,"httpStatus":200,"bytes":1256,"rtt":215.49931}
{"host":"example.com","httpVerb":"GET","hostHeader":"example.com","seq":3,"httpStatus":200,"bytes":1256,"rtt":389.4083}
{"host":"example.com","httpVerb":"GET","hostHeader":"example.com","seq":4,"httpStatus":200,"bytes":1256,"rtt":361.3451}
{"host":"example.com","httpVerb":"GET","hostHeader":"example.com","seq":5,"httpStatus":200,"bytes":1256,"rtt":217.03917}
```

#### Example 3:

Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1
INFO: [2025-03-26 12:06:21 CET] httping 0.10.0 starting
INFO: [2025-03-26 12:06:21 CET] HTTP GET to https://self-signed.badssl.com
INFO: [2025-03-26 12:06:21 CET] Use -h for help
INFO: [2025-03-26 12:06:21 CET] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
WARN: [2025-03-26 12:06:22 CET] Request failed to https://self-signed.badssl.com | proxy=None | Error: Get "https://self-signed.badssl.com": tls: failed to verify certificate: x509: certificate signed by unknown authority
ERROR: [2025-03-26 12:06:22 CET] All probes failed
```

Bypass Bad SSL:
```
$ ./httping -url https://self-signed.badssl.com -count 1 -insecure
INFO: [2025-03-26 12:06:25 CET] httping 0.10.0 starting
INFO: [2025-03-26 12:06:25 CET] HTTP GET to https://self-signed.badssl.com
INFO: [2025-03-26 12:06:25 CET] Use -h for help
INFO: [2025-03-26 12:06:25 CET] Starting HTTP GET to self-signed.badssl.com (https://self-signed.badssl.com)
INFO: [2025-03-26 12:06:25 CET] Connected to https://self-signed.badssl.com, proxy=None, seq=1, status=200, bytes=502, rtt=559.65ms
INFO: [2025-03-26 12:06:25 CET] Results - Probes: 1, Success: 1, Failed: 0.0%
INFO: [2025-03-26 12:06:25 CET] Timing - Min: 559.64504ms, Avg: 0s, Med: 559.64504ms, Max: 559.64504ms
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
