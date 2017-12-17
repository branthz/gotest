package main

import (
	"fmt"
	"os"
)

func main() {
	envget()
}

//测试文件写同步
func fdsync() {
	fd, err := os.Create("./hah")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fd.Close()
	_, err = fd.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		return
	}
	fd.Sync()
}

//测试从环境变量获取变量值
func envget() {
	gp := os.Getenv("GOPATH")
	fmt.Printf("get env from host:%s\n", gp)
}
