package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	//"net/url"
	//"errors"
	//"package/tools"
	//"sync"
	"io/ioutil"
	"os"
	//"time"
)

func main() {
	//http.HandleFunc("/regist", registHandler)
	http.HandleFunc("/v1/mail", requestHandler)
	http.HandleFunc("/v1/sms", requestHandler)
	http.HandleFunc("/v1/wechat", requestHandler)
	http.HandleFunc("/1", mimeHandler)
	//http.Handle("/", http.FileServer(http.Dir("/home/brant/temp")))
	http.HandleFunc("/download", downloadHandler)
	http.HandleFunc("/upload", uploadHandler)
	//http.HandleFunc("/", tokenHandler)
	http.HandleFunc("/health", healthChecker)
	http.ListenAndServe(":10086", nil)
}

func healthChecker(w http.ResponseWriter, r *http.Request) {
	<-r.Cancel
	time.Sleep(3 * 1e9)
	w.Write([]byte("hello world 2"))
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("========get it")
	r.ParseForm()
	code := r.Form.Get("app_auth_code")
	fmt.Println(code)

	return
}

func mimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("==================")
	f, err := os.Create("/home/brant/temp/xxx.jpg")
	if err != nil {
		fmt.Printf("failed create file:%v", err)
		return
	}
	defer f.Close()

	reader, err := r.MultipartReader()
	if err != nil {
		fmt.Printf("error:%s", err.Error())
		return
	}

	part, err := reader.NextPart()
	if err != nil {
		fmt.Printf("error:%s", err.Error())
		return
	}
	defer part.Close()

	body1, err := ioutil.ReadAll(part)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	fmt.Println("----", part.FileName(), string(body1))
	part2, err := reader.NextPart()
	if err != nil {
		fmt.Printf("error:%s", err.Error())
		return
	}

	body2, _ := ioutil.ReadAll(part2)
	fmt.Println("----", part2.FileName(), string(body2))
	part2.Close()

	w.Write([]byte("hello world"))
	f.Write(body1)
	//fmt.Fprintf(w, string(body1))
	return
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	f, err := os.Open("/home/brant/temp/" + r.Form.Get("name"))
	if err != nil {
		fmt.Printf("failed create file:%v", err)
		return
	}
	defer f.Close()
	body1, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	w.Header().Set("Content-Disposition", "attachment;filename=10001.lua")
	w.Write(body1)
	//fmt.Fprintf(w, string(body1))
	return
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/home/brant/Documents/kit接口.docx")
	if err != nil {
		fmt.Printf("failed create file:%v", err)
		return
	}
	defer f.Close()
	body1, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
	w.Write(body1)
	//fmt.Fprintf(w, string(body1))
	return
}

func registHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("---------\n")
	var xx string
	if r.Method == "GET" {
		//tm := fmt.Sprintf("%d", time.Now().UnixNano())
		//xx = tools.Md5Cal([]byte(tm))
		xx = "11111111111111111"

		t, err := template.ParseFiles("regist.gtpl")
		if err != nil {
			fmt.Printf("no find gtpl\n")
			return
		}
		t.Execute(w, xx)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token == "" {
			fmt.Printf("======xxxx\n")
			return
		}
		if token != xx {
			fmt.Printf("token=%s=%s=====----\n", token, xx)
			return
		}
		fmt.Printf("haha:%s==========\n", r.Form["email"])
		//fmt.Fprintf(w,"hello world")
		//http.ServeFile(w,r,"regres.html")
		template.HTMLEscape(w, []byte("hello the same"))
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fname := r.PostForm.Get("id")
	//fn := r.Form["id"]
	//fmt.Printf("get:%s===%s\n", fname, fn)
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("path:%s\n,data:%s", r.URL.Path, string(body))
	w.Write([]byte("hello world"))
}
