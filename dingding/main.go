package main

import (
	"fmt"
	"log"
	"time"

	"github.com/toolkits/net/httplib"
)

func main() {
	bodystr := fmt.Sprintf("{\"msgtype\":\"text\",\"text\":{\"content\":\"%s\"}}", "hello      world")
	url := "https://oapi.dingtalk.com/robot/send?                                              access_token=ff48c3836685eddd334e16ebe99ce8be425f7d5a4b6dce18c15e11528f97cf9e"
	fmt.Println(bodystr)
	r := httplib.Post(url).Body(bodystr).SetTimeout(5*time.Second, 10*time.Second)
	r = r.Header("Content-Type", "application/json")
	resp, err := r.String()
	if err != nil {
		log.Fatalf("send im fail,error:%v", err)
	}

	log.Printf("send im:, resp:%v, url:%s", resp, url)
}
