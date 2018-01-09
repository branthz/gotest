package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"github.com/google/tcpproxy"
)

func server() net.Listener {
	ln, err := net.Listen("tcp", selfAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	return ln
}

const selfAddr = "127.0.0.1:8000"
const testFrontAddr = "1.2.3.9:567"

func mlstenFunc(ln net.Listener) func(network, laddr string) (net.Listener, error) {
	return func(network, laddr string) (net.Listener, error) {
		if network != "tcp" {
			fmt.Printf("got Listen call with network %q, not tcp", network)
			return nil, errors.New("invalid network")
		}
		if laddr != testFrontAddr {
			fmt.Printf("got Listen call with laddr %q, want %q", laddr, testFrontAddr)
			panic("bogus address")
		}
		return ln, nil
	}
}

var dst1 = "192.168.34.41:8888"
var dst2 = "192.168.34.41:8889"

func main() {
	ln := server()
	defer ln.Close()
	p := &tcpproxy.Proxy{ListenFunc: mlstenFunc(ln)}
	p.AddHTTPHostRoute(testFrontAddr, "foo.com", tcpproxy.To(dst1))
	p.AddHTTPHostRoute(testFrontAddr, "bar.com", tcpproxy.To(dst2))
	if err := p.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
