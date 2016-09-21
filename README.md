# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

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
PS D:\httping> .\httping.exe github.com HEAD

httping 0.1 - A tool to measure RTT on HTTP/S requests
Help: httping -h


No protocol specified, falling back to HTTP

HTTP HEAD to github.com (http://github.com):
connected to github.com, seq=1, httpVerb=HEAD, httpStatus=200, bytes=0, RTT=581.92 ms
connected to github.com, seq=2, httpVerb=HEAD, httpStatus=200, bytes=0, RTT=317.96 ms
connected to github.com, seq=3, httpVerb=HEAD, httpStatus=200, bytes=0, RTT=300.06 ms
connected to github.com, seq=4, httpVerb=HEAD, httpStatus=200, bytes=0, RTT=327.54 ms
connected to github.com, seq=5, httpVerb=HEAD, httpStatus=200, bytes=0, RTT=288.91 ms

Probes sent: 5
Successful responses: 5
Average response time: 363.27848ms
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
