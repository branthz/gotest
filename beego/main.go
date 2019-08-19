package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/open-falcon/falcon-plus/common/model"
)

func init() {
	beego.Handler("/v1/push", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.ContentLength == 0 {
			http.Error(w, "body is blank", http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(req.Body)
		var metrics []*model.MetricValue
		err := decoder.Decode(&metrics)
		if err != nil {
			http.Error(w, "connot decode body", http.StatusBadRequest)
			return
		}
		fmt.Println("-----", metrics)
		//g.SendToTransfer(metrics)
		w.Write([]byte("success"))
	}))
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
func main() {
	beego.Run()
	beego.Router("/", &MainController{})
	beego.Run()
}
