# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

### Latest release

- Version [0.5](https://github.com/pjperez/httping/releases)

### Requirements
Golang >1.3 ::  Otherwise it will fail with *unknown http.Client field 'Timeout' in struct literal*

### Installing
```
go get github.com/pjperez/httping
```
### Usage
httping [OPTIONS] url

#### Options
```
-h
  Help
  
--count=10
  Number of requests to send.
  Default: 10
  
--httpverb=GET
  Verb to use for the HTTP request: GET or HEAD.
  Default: GET
```

#### Example

```
PS D:\httping> httping.exe --count=100 google.com

httping 0.5 - A tool to measure RTT on HTTP/S requests
Help: httping -h


No protocol specified, falling back to HTTP

HTTP GET to google.com (http://google.com):
connected to http://google.com, seq=1, httpVerb=GET, httpStatus=200, bytes=10950, RTT=85.26 ms
connected to http://google.com, seq=2, httpVerb=GET, httpStatus=200, bytes=11035, RTT=588.18 ms
connected to http://google.com, seq=3, httpVerb=GET, httpStatus=200, bytes=10959, RTT=299.30 ms
connected to http://google.com, seq=4, httpVerb=GET, httpStatus=200, bytes=10935, RTT=100.60 ms
connected to http://google.com, seq=5, httpVerb=GET, httpStatus=200, bytes=10919, RTT=593.78 ms
connected to http://google.com, seq=6, httpVerb=GET, httpStatus=200, bytes=10966, RTT=285.03 ms
connected to http://google.com, seq=7, httpVerb=GET, httpStatus=200, bytes=11023, RTT=57.83 ms
connected to http://google.com, seq=8, httpVerb=GET, httpStatus=200, bytes=10933, RTT=338.61 ms
connected to http://google.com, seq=9, httpVerb=GET, httpStatus=200, bytes=11017, RTT=535.14 ms
connected to http://google.com, seq=10, httpVerb=GET, httpStatus=200, bytes=10993, RTT=108.15 ms
connected to http://google.com, seq=11, httpVerb=GET, httpStatus=200, bytes=10959, RTT=346.78 ms
connected to http://google.com, seq=12, httpVerb=GET, httpStatus=200, bytes=10944, RTT=534.66 ms
connected to http://google.com, seq=13, httpVerb=GET, httpStatus=200, bytes=10969, RTT=114.99 ms
connected to http://google.com, seq=14, httpVerb=GET, httpStatus=200, bytes=10965, RTT=339.64 ms
connected to http://google.com, seq=15, httpVerb=GET, httpStatus=200, bytes=10974, RTT=534.37 ms
connected to http://google.com, seq=16, httpVerb=GET, httpStatus=200, bytes=10951, RTT=119.27 ms
(...)
connected to http://google.com, seq=97, httpVerb=GET, httpStatus=200, bytes=11024, RTT=534.96 ms
connected to http://google.com, seq=98, httpVerb=GET, httpStatus=200, bytes=10983, RTT=115.59 ms
connected to http://google.com, seq=99, httpVerb=GET, httpStatus=200, bytes=11004, RTT=340.18 ms
connected to http://google.com, seq=100, httpVerb=GET, httpStatus=200, bytes=11008, RTT=535.81 ms

Probes sent: 100
Successful responses: 100
% of requests failed: 0
Min response time: 53.9967ms
Average response time: 337.575484ms
Median response time: 347.70375ms
Max response time: 621.5548ms

90% of requests were faster than: 581.8406ms
75% of requests were faster than: 535.29445ms
50% of requests were faster than: 347.70375ms
25% of requests were faster than: 117.43025ms
```

### Help
httping help

### Warranty
This is just a learning exercise, hence it is distributed AS IS, with no warranty.

You should not use this software in production under any circumstances. It's not intended to and it has not been thoroughly tested and problably it's not very accurate.

# Contributing

In order to contribute you can:

1. Open a new Issue
2. Fork the project, add a feature or fix and send us a PR
3. Go to the [Projects](https://github.com/pjperez/httping/projects) section and comment on a card/issue from the TODO list.
4. Just get in touch if you have other ideas :)
5. Do whatever you want actually! :)
