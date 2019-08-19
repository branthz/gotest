package main

import(
	"regexp"
	"fmt"
	"encoding/binary"
	"net"
)

func main(){
	addrCheck()
}


func ip2int(ips string) int {
	ip:=net.ParseIP(ips)	
	if len(ip) == 16 {
		return int(binary.BigEndian.Uint32(ip[12:16]))
	}
	return int(binary.BigEndian.Uint32(ip))
}

func int2ip(nn int) string {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, uint32(nn))
	return ip.String() 
}

//regexp match
//regexp split
func addrCheck(){
	addr :="192.168.34.999/14"
	//reip:=regexp.MustCompile(`^\d{1,3}\.\d{1,3}.\d{1,3}.\d{1,3}$`)		
	reipprefix ,err:=regexp.Compile(`(^\d{1,3}\.\d{1,3}.\d{1,3}.\d{1,3})/(\d{1,2}$)`) 
	if err!=nil{
		fmt.Println(err)
		return
	}
	get:=reipprefix.MatchString(addr)
	ss:=regexp.MustCompile("/").Split(addr,2)
	fmt.Println(get,ss)
}
