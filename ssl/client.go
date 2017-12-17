package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"crypto/x509"
	//"net"
)

func main(){
	roots := x509.NewCertPool()
	severCert, err := ioutil.ReadFile("server.pem")
        if err != nil {
                fmt.Printf("Could not load server certificate!")
        }
	
	ok := roots.AppendCertsFromPEM(severCert)
	if !ok {
    		panic("failed to parse root certificate")
	}
	scon,err :=tls.Dial("tcp", "127.0.0.1:6666", &tls.Config{RootCAs: roots,})
	if err != nil {
                fmt.Println(err)
                return
        }
        defer scon.Close()
	
	scon.Write([]byte("hello"))
        var buf [128]byte
        rn, err := scon.Read(buf[:])
        if err != nil {
                fmt.Println(err)
                return
        }
        fmt.Println(string(buf[:rn]))
}

func skipVerify() {
	scon, err := tls.Dial("tcp", "127.0.0.1:6666", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer scon.Close()

	scon.Write([]byte("hello"))
	var buf [128]byte
	rn, err := scon.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:rn]))

}
