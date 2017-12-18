package main

import (
	"fmt"
	db "github.com/influxdb/influxdb/client"
	"net/url"
)

func main() {
	u,err:=url.Parse("192.168.1.205:")
	if err!=nil{
		return
	}
	config :=&db.Config{
		URL:u
	}
	db.NewClient(config)
}
