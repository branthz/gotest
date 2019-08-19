package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main(){
	v:=viper.New()	

	v.SetConfigType("toml")
	v.SetConfigFile("./logbase.toml")
	err:=v.ReadInConfig()
	if err!=nil{
		fmt.Println(err)
		return
	}
	v.SetDefault("version","v1.2.3")	
	fmt.Println(v.GetString("version"))
	fmt.Println(v.GetStringMap("ports"))
}
