# httping
httping - A tool to measure RTT on HTTP/S requests 

This tool should be able to run on Windows, Linux and Mac OS/X, but it has only been tested in Windows 10.

### Requirements
Golang >1.3 ::  Otherwise it will fail with *unknown http.Client field 'Timeout' in struct literal*

### Installing
go get github.com/pjperez/httping

### Usage
httping url[:port] [GET|HEAD]

### Help
httping help

### Warranty
This is just a learning exercise, hence it is distributed AS IS, with no warranty.

You should not use this software in production under any circumstances. It's not intended to and it has not been thoroughly tested.
