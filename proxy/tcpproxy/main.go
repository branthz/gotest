package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"github.com/branthz/margin-cache/proxy"
	"github.com/branthz/utarrow/lib/log"
)

var mlog *log.Logger 

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
		mlog.Info("local addr:%s",laddr)
		return ln, nil
	}
}

var dst1 = "192.168.34.41:8888"
var dst2 = "192.168.34.41:8889"

func main() {
	mlog,_=log.New("",log.Debug)
	//ln := server()
	//defer ln.Close()
	p := &proxy.Proxy{LocalHost:selfAddr}
	//TO returs DialProxy which impletes Target
	//Target has handleConn methon
	//p.AddHTTPHostRoute("foo.com", proxy.To(dst1))
	//p.AddHTTPHostRoute( "bar.com", proxy.To(dst2))
	p.AddRoute(dst1)
	p.AddRoute(dst2)
	if err := p.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
