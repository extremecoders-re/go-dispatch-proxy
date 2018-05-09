# Go dispatch proxy

A SOCKS5 load balancing proxy to combine multiple internet connections into one. Works on Windows and Linux. Untested on macOS. Written in pure Go with no additional dependencies.

## Rationale

The idea for this project came from [dispatch-proxy](https://github.com/Morhaus/dispatch-proxy) which is written in NodeJS.
[NodeJS is not entirely harddisk friendly considering the multitude of files it creates even for very simple programs](https://medium.com/@jdan/i-peeked-into-my-node-modules-directory-and-you-wont-believe-what-happened-next-b89f63d21558). I needed something which was light without polluting the entire drive.

![](https://pbs.twimg.com/media/DEIV_1XWsAAlY29.jpg)

## Installation

No installation required. Grab the latest binary for your platform from the CI server and start speeding up your internet connection!

https://ci.appveyor.com/project/extremecoders-re/go-dispatch-proxy/build/artifacts

[![Build status](https://ci.appveyor.com/api/projects/status/nll4hvpdjlfsp7mu?svg=true)](https://ci.appveyor.com/project/extremecoders-re/go-dispatch-proxy/build/artifacts)

## Usage

The example below are shown on Windows. The steps are similar for other platforms.

The primary purpose of the tool is to combine multiple internet connections into one. For this we need to know the IP addresses of the interface we wish to combine. You can obtain the IP addresses using the `ipconfig` (`ifconfig` on linux) command. Alternatively run `go-dispatch-proxy -list`.

```
D:\>go-dispatch-proxy.exe -list
--- Listing the available adresses for dispatching
[+] Mobile Broadband Connection , IPv4:10.81.201.18
[+] Local Area Connection, IPv4:192.168.1.2
```

Start `go-dispatch-proxy` specifying the IP addresses of the load balancers obtained in the previous step. Along with the IP address you also need to provide the contention ratio as shown below.

```
D:\>go-dispatch-proxy.exe 10.81.201.18@3 192.168.1.2@2
2018/05/09 15:57:50 [+] Load balancer 1: 10.81.201.18, contention ratio: 3
2018/05/09 15:57:50 [+] Load balancer 2: 192.168.1.2, contention ratio: 2
2018/05/09 15:57:50 [+] SOCKS server started at 127.0.0.1:8080
```

Out of 5 consecutive connections, the first 3 are routed to `10.81.201.18` and the remaining 2 to `192.168.1.2`. The SOCKS server is started by default on `127.0.0.1:8080`. It can be changed using the `-lhost` and `lport` directive.

Now change the proxy settings of your browser, download manager etc to point to the above address (eg `127.0.0.1:8080`). Be sure to add this as a SOCKSv5 proxy and NOT as a HTTP/S proxy.

## Compiling (For Development)

Ensure that Go is installed and available on the system path.

```sh
$ git clone https://bitbucket.org/extremecoders-re/go-dispatch-proxy
$ cd go-dispatch-proxy

# Compile for Windows x86
$ GOOS=windows GOARCH=386 go build

# Compile for Windows x64
$ GOOS=windows GOARCH=amd64 go build

# Compile for Linux x86
$ GOOS=linux GOARCH=386 go build

# Compile for Linux x64
$ GOOS=linux GOARCH=amd64 go build

# Compile for macos x86
$ GOOS=darwin GOARCH=386 go build

# Compile for macos x64
$ GOOS=darwin GOARCH=amd64 go build
```

## Credits

- [dispatch-proxy](https://github.com/Morhaus/dispatch-proxy): A SOCKS5/HTTP load balancing proxy written in NodeJS.

## License

Licensed under MIT