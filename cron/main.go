package main

import (
	"fmt"

	"github.com/robfig/cron"
)

//Readme:https://godoc.org/github.com/robfig/cron#Cron.AddJob
func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every seconds call") })
	c.AddFunc("0 49 * * * *", func() { fmt.Println("Every hour call on minute 49") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	defer c.Stop()
	ch := make(chan int)
	<-ch
}

//no contain returned values
func myjob() {
	fmt.Println("job done!")
	return
}
