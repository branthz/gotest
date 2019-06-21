package main

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	routeShow()
	routeAdd("10.12.0.0/16", "192.168.34.254", "192.168.41.0/24")
}

func routeShow() {
	links, err := netlink.LinkList()
	if err != nil {
		fmt.Println(err)
	}
	for _, link := range links {
		routes, _ := netlink.RouteList(link, netlink.FAMILY_V4)
		for _, route := range routes {
			fmt.Println(route)
		}
	}
}
func routeAdd(dst string, gateway string, src string) {
	//link, err := netlink.LinkByName("ens160")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//route := netlink.Route{LinkIndex: link.Attrs().Index}
	var err error
	route := netlink.Route{}
	if dst != "default" {
		dstAddr, _ := netlink.ParseAddr(dst)
		route.Dst = dstAddr.IPNet
	}
	if src != "" {
		srcIP := net.ParseIP(src)
		route.Src = srcIP
	}
	if gateway != "" {
		gatewayIP := net.ParseIP(gateway)
		route.Gw = gatewayIP
	}
	err = netlink.RouteAdd(&route)
	if err != nil {
		fmt.Printf("----%v\n", err)
	}
}

func addrAdd() {
	lo, err := netlink.LinkByName("ens192")
	checkErr(err)
	addr, err := netlink.ParseAddr("192.168.200.254/32")
	checkErr(err)
	err = netlink.AddrAdd(lo, addr)
	checkErr(err)
}
