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
-h
  Help

-url *https://github.com*
  Requested URL. If no protocol is specified with http:// or https:// the system will use http://

-count *10*
  Number of requests to send [0 means infinite].
  Default: 10

-timeout 2000
  Timeout in milliseconds
  Default: 2000 (2 seconds)

-httpverb *GET*
  Verb to use for the HTTP request: GET or HEAD.
  Default: GET

-hostheader "github.com"
  Specify a custom Host: header

-json
  If specified, outputs the results in json format

-noProxy
  If specified, ignores system proxy settings
```

#### Example

```
PS C:\temp> ./httping.exe -url https://wormhole.network -count 10

httping 0.9.1 - A tool to measure RTT on HTTP/S requests
Help: httping -h
HTTP GET to wormhole.network (https://wormhole.network):
connected to https://wormhole.network, seq=1, httpVerb=GET, httpStatus=200, bytes=10991, RTT=381.99 ms
connected to https://wormhole.network, seq=2, httpVerb=GET, httpStatus=200, bytes=10991, RTT=169.03 ms
connected to https://wormhole.network, seq=3, httpVerb=GET, httpStatus=200, bytes=10991, RTT=94.19 ms
connected to https://wormhole.network, seq=4, httpVerb=GET, httpStatus=200, bytes=10991, RTT=106.94 ms
connected to https://wormhole.network, seq=5, httpVerb=GET, httpStatus=200, bytes=10991, RTT=78.16 ms
connected to https://wormhole.network, seq=6, httpVerb=GET, httpStatus=200, bytes=10991, RTT=121.95 ms
connected to https://wormhole.network, seq=7, httpVerb=GET, httpStatus=200, bytes=10991, RTT=103.13 ms
connected to https://wormhole.network, seq=8, httpVerb=GET, httpStatus=200, bytes=10991, RTT=81.28 ms
connected to https://wormhole.network, seq=9, httpVerb=GET, httpStatus=200, bytes=10991, RTT=125.78 ms
connected to https://wormhole.network, seq=10, httpVerb=GET, httpStatus=200, bytes=10991, RTT=81.77 ms

Probes sent: 10
Successful responses: 10
% of requests failed: 0
Min response time: 78.1639ms
Average response time: 134.42187ms
Median response time: 105.035ms
Max response time: 381.9932ms

90% of requests were faster than: 275.51005ms
75% of requests were faster than: 125.7822ms
50% of requests were faster than: 105.035ms
25% of requests were faster than: 81.7667ms
```

#### Example 2:

```
PS C:\temp> .\httping.exe -url https://wormhole.network -count 5 -json
{"host":"wormhole.network","httpVerb":"GET","hostHeader":"wormhole.network","seq":1,"httpStatus":200,"bytes":10991,"rtt":415.5466}
{"host":"wormhole.network","httpVerb":"GET","hostHeader":"wormhole.network","seq":2,"httpStatus":200,"bytes":10991,"rtt":120.0931}
{"host":"wormhole.network","httpVerb":"GET","hostHeader":"wormhole.network","seq":3,"httpStatus":200,"bytes":10991,"rtt":75.6925}
{"host":"wormhole.network","httpVerb":"GET","hostHeader":"wormhole.network","seq":4,"httpStatus":200,"bytes":10991,"rtt":121.3327}
{"host":"wormhole.network","httpVerb":"GET","hostHeader":"wormhole.network","seq":5,"httpStatus":200,"bytes":10991,"rtt":71.4523}
```

#### Example 3

Continuous monitoring of the connection quality,  
```
$ httping.exe -url http://detectportal.firefox.com/success.txt -count 0 -timeout 1000

httping 0.9.1 - A tool to measure RTT on HTTP/S requests
Help: httping -h
HTTP GET to detectportal.firefox.com (http://detectportal.firefox.com/success.txt):
Timeout when connecting to http://detectportal.firefox.com/success.txt
Timeout when connecting to http://detectportal.firefox.com/success.txt
connected to http://detectportal.firefox.com/success.txt, seq=3, httpVerb=GET, httpStatus=200, bytes=8, RTT=882.24 ms
Timeout when connecting to http://detectportal.firefox.com/success.txt
Timeout when connecting to http://detectportal.firefox.com/success.txt
Timeout when connecting to http://detectportal.firefox.com/success.txt
connected to http://detectportal.firefox.com/success.txt, seq=7, httpVerb=GET, httpStatus=200, bytes=8, RTT=928.17 ms
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
