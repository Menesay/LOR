# Lightweight
# Onion
# Router
## http proxychain structure.

## Setup

1. Download go.

2. Clone this repo.

3. Open source code and config lor servers ip:port.

Default ip:ports:
* lor-0					ip:127.0.0.10 		host:port:1400		proxyport:1401
* lor-1					ip:127.0.0.20 		host:port:1400		proxyport:1401
* lor-2					ip:127.0.0.30 		host:port:1400		proxyport:1401

4. Compile 

```bash
go build lor-0.go
go build lor-1.go
go build lor-2.go
```


## Run

1. first run lor-2, second lor-1, third lor-0 because of the network structure

2. connect to lor-0

```bash
nc 127.0.0.10 1400
menesay.duckdns.org
80
```
Connect to the lor-0
1. send hostname
2. send port

Wait for a while and the http proxy will running at 127.0.0.10:1401
