package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

//参数不能合在一个变量里，需要分开
func showRoute() {
	cmd := exec.Command("/usr/sbin/ip", "route", "show", "table", "254")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("in all caps:\n%s\n", out.String())
	strs := strings.Split(string(out.Bytes()), "\n")
	var line string
	var lines []string
	for i := 0; i < len(strs); i++ {
		line = strings.Trim(strs[i], "\r\n ")
		if line == "" {
			continue
		}
		lines = append(lines, line)

	}
	fmt.Println(lines)
}

func main() {
	showRoute()
}

func cp() {
	cmd := exec.Command("/bin/cp", "-r", "/home/brant/temp/openab", "/home/brant/temp/debian")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
