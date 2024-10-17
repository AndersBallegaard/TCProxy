# TCProxy
A lightweight dual stack TCP proxy written in go

## Capabilities
* Proxy all kinds of TCP based protocols
    * tested on the following
        * HTTP
        * HTTPS
        * SSH

* IPv6 enable ipv4 only services
* Supports multiple concurrent connections


## Configuration
The application looks for a file called settings.yaml in the current working directory.
### Config samples
#### HTTP forwarding to ip
In this example we are listing on port 8080, and forwarding any connection the server is getting to 10.255.255.1:80
```yaml
server:
  port: 8080
target:
  address: 10.255.255.1
  port: 80
```

#### HTTPS forwarding with DNS
In this example we are acting as a proxy towards a proxmox web management interface. The server will lookup DNS names in the target address field. Note that the proxy is a direct passthough, so if the target server speaks HTTPS then it's passed though directly to the client. The server can't look inside the TLS connection for encrypted data
```yaml
server:
  port: 8080
target:
  address: proxmox.internal
  port: 8006
```

#### SSH proxy
Like the HTTPS example, this does not break the security of the protocol itself. The configuration is also almost identical to the other examples because the proxy isn't application aware
```yaml
server:
  port: 2222
target:
  address: 10.10.10.10
  port: 22
```

## How to build
Make sure your version of go is new enough
```bash
go version
go version go1.22.2 linux/amd64 # You should have 1.22 or newer
```
Download the source code using git, and enter the directory
```
git clone https://github.com/AndersBallegaard/TCProxy.git
cd TCProxy
```
Run the following build command
```bash
go build .
```
Run the proxy
```
./TCProxy
```

