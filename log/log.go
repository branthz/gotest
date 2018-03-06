package main

import (
	"time"

	"github.com/branthz/utarrow/lib/log"
)

type dd struct {
	name [30]byte
}

func main() {
	log.SetupRotate("./xx", 4)
	log.Error("nihao")
	var a [10]byte
	a[5] = 10
	var b string = "usdfjskjdfiojajdlk       fjsdkjfi      sljdkfjajdlsakfsdfklsdjfksdfjhgjlsajaiejfkjaklsdjkkdsjfjksjiwjfkjskdjfiwojflsj"
	for {
		log.Debug("---%v", a)
		log.Debug("===%s", b)
		time.Sleep(1e8)
	}
}
