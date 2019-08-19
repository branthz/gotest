package main

import (
	"fmt"
	"net"
	"time"

	"github.com/miekg/dns"
)

const timeout = 1 * time.Second

func NewMsg(site string) (m *dns.Msg) {
	m = new(dns.Msg)
	m.SetQuestion(dns.Fqdn(site), dns.TypeA)
	m.RecursionDesired = true
	return
}

func main() {
	ips, err := digQuery("192.168.30.32", "8.8.8.8", "www.baidu.com", 3)
	if err != nil {
		fmt.Printf("dig err:%v", err)
	}
	fmt.Println(ips)
}

func digQuery(src, server, domain string, retries int) (ips []string, err error) {
	dst, err := net.ResolveUDPAddr("udp4", server+":53")
	if err != nil {
		return
	}
	srcaddr, err := net.ResolveUDPAddr("udp4", src+":0")
	if err != nil {
		return
	}
	conn, err := net.DialUDP("udp4", srcaddr, dst)
	if err != nil {
		return
	}
	defer conn.Close()
	m := NewMsg(domain)
	data, err := m.Pack()
	if err != nil {
		return
	}
	var rdata = make([]byte, 1460)
	var rn int
	var i =1
	for {
		conn.SetDeadline(time.Now().Add(timeout))
		_, err = conn.Write(data)
		if err != nil {
			if i<retries{ 
				i++
				continue	
			}else{
				return	
			}
		}
		rn, err = conn.Read(rdata)
		if err != nil {
			if i<retries {
				i++
				continue
			}else{
				return
			}
		}
		break
	}
	mm := new(dns.Msg)
	err = mm.Unpack(rdata[:rn])
	if err != nil {
		fmt.Printf("recv unpack datalen:%d\n",rn)
		return
	}
	if mm.Rcode != dns.RcodeSuccess {
		err = fmt.Errorf("dig no success")
		return
	}
	for _, record := range mm.Answer {
		if v, ok := record.(*dns.A); ok {
			ips = append(ips, v.A.String())
		}
	}

	return
}
