package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"os"

	js "github.com/bitly/go-simplejson"
)

type info struct {
	Enabled     bool
	DeviceToken string
	Taglist     []string
	MobileType  int
	Encoding    string
}

func checkOverInfoUnmarshal(s string) {
	fo := new(info)
	err := json.Unmarshal([]byte(s), fo)
	if err != nil {
		fmt.Printf("======wrong\n")
	} else {
		fmt.Printf("success:%v\n", fo)
	}
}

func main() {
	var str = "{\"enabled\":true,\"deviceToken\":\"d256898932798de31e0b276174189dc9f0c6078c983ca89937f893a40e47e605\",\"taglist\":[],\"mac\":[\"b4430d30cf18\"],\"mobileType\":2,\"language\":\"zh_CN\",\"encoding\":\"utf8\"}"
	checkOverInfoUnmarshal(str)
}

func checkJosnMap() {
	var str = "{\"enabled\":true,\"deviceToken\":\"d256898932798de31e0b276174189dc9f0c6078c983ca89937f893a40e47e605\",\"taglist\":[],\"mac\":[\"b4430d30cf18\"],\"mobileType\":2,\"language\":\"zh_CN\",\"encoding\":\"utf8\"}"
	buf := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &buf)
	if err != nil {
		fmt.Printf("======wrong\n")
	} else {
		fmt.Printf("haha\n")
	}

	tt := make(map[string]interface{})
	tt["123"] = "xxxxx"
	tt["234"] = "xxxxx"
	tx, _ := json.Marshal(tt)
	fmt.Printf("%s\n", string(tx))
}

func jsonPack() {

	var str = "{\"code\":200,\"mssg\":\"ok\",\"enabled\":true,\"result\":[\"beijing\",\"hangzhou\"]}"
	obj, err := js.NewJson([]byte(str))
	checkErr(err)
	thing := obj.Get("msg").MustString()

	fmt.Printf("-------%d", obj.Get("code").MustInt())

	arr, err := obj.Get("result").StringArray()
	fmt.Println(arr)
	thing = obj.GetPath("mssg").MustString()
	fmt.Printf("--------%s\n", thing)
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(0)
	}
}
