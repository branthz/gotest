package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type config struct {
	Name    string
	Version string
	DB      database `toml:"database"`
	Ports   map[string]ports
}

type database struct {
	Server         string
	Connection_max int
	Ports          []int
	Enabled        bool
}

type ports struct {
	Desc string
}

var Lconfig config

func main() {
	_, err := toml.DecodeFile("./test.toml", &Lconfig)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Lconfig)
	for k, v := range Lconfig.Ports {
		fmt.Println(k, v)
	}
}
