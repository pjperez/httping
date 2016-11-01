# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

### Latest release

- [Download from Github](https://github.com/pjperez/httping/releases)

### Requirements
Golang >1.3 ::  http.Client Timeout wasn't available in previous versions

### Installing
#### From source
```
go get github.com/pjperez/httping
go install github.com/pjperez/httping
```

You'll then find httping.exe in your $GOPATH/bin directory

#### Binary
Check the [latest release](https://github.com/pjperez/httping/releases) and download the standalone binary from there, it doesn't need installation and it's portable.

### Usage
httping -url requested_url [OPTIONS]

#### Options
```
-h
  Help

-url *https://github.com*
  Requested URL. If no protocol is specified with http:// or https:// the system will use http://

-count *10*
  Number of requests to send.
  Default: 10
  
-httpverb *GET*
  Verb to use for the HTTP request: GET or HEAD.
  Default: GET
```

#### Example

```
PS C:\temp> ./httping.exe -url https://wormhole.network --count 10

httping 0.6 - A tool to measure RTT on HTTP/S requests
Help: httping -h
HTTP GET to wormhole.network (https://wormhole.network):
connected to https://wormhole.network, seq=1, httpVerb=GET, httpStatus=200, bytes=10991, RTT=147.75 ms
connected to https://wormhole.network, seq=2, httpVerb=GET, httpStatus=200, bytes=10991, RTT=61.73 ms
connected to https://wormhole.network, seq=3, httpVerb=GET, httpStatus=200, bytes=10991, RTT=79.19 ms
connected to https://wormhole.network, seq=4, httpVerb=GET, httpStatus=200, bytes=10991, RTT=63.85 ms
connected to https://wormhole.network, seq=5, httpVerb=GET, httpStatus=200, bytes=10991, RTT=67.34 ms
connected to https://wormhole.network, seq=6, httpVerb=GET, httpStatus=200, bytes=10991, RTT=70.56 ms
connected to https://wormhole.network, seq=7, httpVerb=GET, httpStatus=200, bytes=10991, RTT=60.03 ms
connected to https://wormhole.network, seq=8, httpVerb=GET, httpStatus=200, bytes=10991, RTT=90.19 ms
connected to https://wormhole.network, seq=9, httpVerb=GET, httpStatus=200, bytes=10991, RTT=79.58 ms
connected to https://wormhole.network, seq=10, httpVerb=GET, httpStatus=200, bytes=10991, RTT=99.40 ms

Probes sent: 10
Successful responses: 10
% of requests failed: 0
Min response time: 60.0311ms
Average response time: 81.96285ms
Median response time: 74.87805ms
Max response time: 147.752ms

90% of requests were faster than: 123.5738ms
75% of requests were faster than: 90.194ms
50% of requests were faster than: 74.87805ms
25% of requests were faster than: 63.845ms
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
4. Just get in touch if you have other ideas :)
5. Do whatever you want actually! :)
