package main

import (
	"net"
	"os"
	"time"

	"github.com/branthz/utarrow/lib/log"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type Pinger struct {
	source  string
	dst     string
	count   int
	timeout time.Duration
	recv    int
}

func NewPinger(src, dst string) *Pinger {
	return &Pinger{
		source:  src,
		dst:     dst,
		count:   3,
		timeout: time.Second * 2,
	}
}

func (p *Pinger) run() error {
	c, err := icmp.ListenPacket("ip4:icmp", p.source)
	if err != nil {
		return err
	}
	defer c.Close()
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		return err
	}
	target, _ := net.ResolveIPAddr("ip4", p.dst)
	rb := make([]byte, 1480)
	for i := 0; i < p.count; i++ {
		c.SetDeadline(time.Now().Add(p.timeout))
		_, err = c.WriteTo(wb, target)
		if err != nil {
			log.Warn("ping from %s write to %s :%v", p.source, p.dst, err)
			continue
		}
		n, _, err := c.ReadFrom(rb)
		if err != nil {
			log.Warn("ping local %s read from %s :%v", p.source, p.dst, err)
			continue
		}
		rm, err := icmp.ParseMessage(1, rb[:n])
		if rm.Type != ipv4.ICMPTypeEchoReply {
			log.Warn("ping from %s to %s want echo reply,but get:%+v", p.source, p.dst, rm)
			continue
		}
		p.recv++
	}
	return nil
}
