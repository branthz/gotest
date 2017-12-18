package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	js "github.com/go-simplejson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	//"package/tools"
	//"path/filepath"
	"time"
)

func Caltoken(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	s := base64.StdEncoding.EncodeToString(h.Sum(nil))

	h = md5.New()
	h.Write([]byte(s))
	checksum := hex.EncodeToString(h.Sum(nil))

	return checksum
}

func firmwareDelete() {
	client := http.Client{}
	urlstr := "http://192.168.1.205:17001/firmware/delete?dev_type=1&version_solid=1&version_soft=212"
	request, err := http.NewRequest("", urlstr, nil)
	request.Header.Add("timestamp", "123")
	token := CalTokenEx("123")
	request.Header.Add("sign", token)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
func backup_download() {
	client := http.Client{}
	urlstr := "http://115.29.28.145/rest/1.0/backup?method=download&pathname=/00/bb68a22f6bf87fba955ae64817b520/2015-01-03_17:35:51"
	urlstrxx := fmt.Sprintf("%s&timestamp=123&token=%s", urlstr, Caltoken("Broadlink:290123"))
	request, err := http.NewRequest("", urlstrxx, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func deployQuery() {
	client := http.Client{}
	urlstr := "http://192.168.3.111:9000/dnakit/server/query?vendor_id=10001"
	request, err := http.NewRequest("", urlstr, nil)
	tm := fmt.Sprintf("%d", time.Now().Unix())
	request.Header.Add("timestamp", tm)
	sign := CalTokenEx("/dnakit/server/query" + tm)
	request.Header.Add("sign", sign)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

//multipart/form-data
func uploadfile() {
	xxx := "./123.png"
	file, err := os.Open(xxx)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer file.Close()

	//var buf bytes.Buffer

	//io.ReadFull(file, buf.Bytes())
	body := bytes.NewBuffer(make([]byte, 0))
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "file")
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)
	ctype := writer.FormDataContentType()

	writer.Close()

	//data := "123"
	//buf := make([]byte, 1024*1024)
	//n, _ := io.ReadFull(file, buf)
	//fmt.Printf("##%v##\n", buf[:n])
	client := http.Client{}
	//fmd5 := tools.Md5Cal(buf[:n])
	//urlstr := fmt.Sprintf("http://192.168.3.111:17001/firmware/update?dev_type=11&version_solid=1&version_soft=23&filemd5=%s&infoen=haha&infocn=haha", fmd5)
	urlstr := "http://localhost:8080/dnakit/v2/user/upload?filename=123.png"
	request, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		fmt.Println("post %v", err)
		os.Exit(0)
	}
	request.Header.Set("Content-Type", ctype)
	request.Header.Add("timestamp", "123")
	token := CalTokenEx("123")
	request.Header.Add("sign", token)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("nono:%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func postJson() {
	data := "{\"Type\":1,\"MsgID\":20,\"ShopId\":20,\"Data\":300}"
	//data := "{\"main\":{\"user\":\"root\",\"password\":\"BroadlinkFL1201\",\"oip\":\"115.29.245.194\",\"iip\":\"10.165.32.59\"},\"backup\":{\"user\":\"roota\",\"password\":\"BroadlinkFL1201\",\"oip\":\"115.29.202.219\",\"iip\":\"10.161.214.215\"},\"vendor_id\":\"10001\"}"
	body := bytes.NewBuffer([]byte(data))
	client := http.Client{}
	urlstr := "http://127.0.0.1:8999/new_message"
	request, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(-1)
	}
	tm := fmt.Sprintf("%d", time.Now().Unix())
	request.Header.Add("timestamp", tm)
	sign := CalTokenEx("/dnakit/server/deploy" + tm)
	request.Header.Add("sign", sign)

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("nono:%s\n", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}

}

func postRegist() {
	data := "{\"phone\":\"15133992232\",\"passwd\":\"123456\",\"email\":\"15199993444@126.com\",\"verifycode\":\"2290\",\"countrycode\":\"+86\"}"
	//databs64 := "{\"type\":\"ui\",\"action\":\"download\",\"extrainfo\":{\"platform\":\"bl\",\"uiid\":\"343412345\"} }"
	//databs64 := base64.StdEncoding.EncodeToString([]byte(data))
	body := bytes.NewBuffer([]byte(data))
	client := http.Client{}

	urlstr := fmt.Sprintf("http://localhost:8080/dnakit/exempt/regist")
	request, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		fmt.Println("post %v", err)
		os.Exit(0)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("nono:%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}
}

func postLogin() {
	//data := "phone=15158134537&passwd=123456&email=163@163&verifycode=1932&countrycode=+86"
	data := "{\"phone\":\"15158134537\",\"passwd\":\"123456\",\"email\":\"163@163\",\"verifycode\":\"1932\",\"countrycode\":\"+86\"}"
	//databs64 := "{\"type\":\"ui\",\"action\":\"download\",\"extrainfo\":{\"platform\":\"bl\",\"uiid\":\"343412345\"} }"
	//databs64 := base64.StdEncoding.EncodeToString([]byte(data))
	body := bytes.NewBuffer([]byte(data))

	client := http.Client{}
	//fmd5 := tools.Md5Cal([]byte(data))

	urlstr := fmt.Sprintf("http://localhost:8080/dnakit/exempt/login")
	request, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		fmt.Println("post %v", err)
		os.Exit(0)
	}
	//request.Header.Add("uid", "1451319903")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("nono:%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		fmt.Printf("%s--%s--%s\n", string(contents), response.Header.Get("reqUserId"), response.Header.Get("reqUserSession"))
	}
}

func getPhoneCode() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8088/user/select_coupon?phone=18767105477&controycode=+86"
	request, err := http.NewRequest("", urlstr, nil)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func postCompanyApply() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8080/dnakit/v2/user/company/apply"
	admin := js.New()
	admin.Set("name", "hahah")
	admin.Set("post", "something else")
	admin.Set("familyAddr", "杭州市长王大道32号")
	admin.Set("personfile", 10)
	admin.Set("personType", "licenseId")
	admin.Set("headPortrait", 11)

	firm := js.New()
	firm.Set("firmName", "google")
	firm.Set("firmAddr", "hangzhou jianghong road.")
	firm.Set("website", "www.google.com")
	firm.Set("license", "10000000000000")
	firm.Set("firmPhone", "1911919191")
	firm.Set("firmScale", 100)
	firm.Set("firmDesc", "the best company")
	firm.Set("corpIntent", "hello world")
	firm.Set("licenseFile", 100)
	firm.Set("certificate", 120)

	data := js.New()
	data.Set("admin", admin)
	data.Set("firm", firm)

	bodydata, err := data.Encode()
	if err != nil {
		fmt.Println(err)
		return
	}
	bodybuf := bytes.NewBuffer(make([]byte, 0))
	bodybuf.Write(bodydata)
	request, err := http.NewRequest("POST", urlstr, bodybuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("reqUserId", "1451329585")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func postDeveloperApply() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8081/dnakit/v2/user/developer/apply"
	admin := js.New()
	admin.Set("name", "brant")
	admin.Set("post", "something else")
	admin.Set("familyAddr", "杭州市长王大道32号")
	admin.Set("personfile", 10)
	admin.Set("headPortrait", 11)
	bodydata, _ := admin.Encode()
	bodybuf := bytes.NewBuffer(make([]byte, 0))
	bodybuf.Write(bodydata)
	request, err := http.NewRequest("POST", urlstr, bodybuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("reqUserId", "1451329585")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}

func queryCompany() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8081/dnakit/v2/user/company/query"
	request, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
func getCompanyApply() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8080/dnakit/v2/user/company/apply?companyid=1451529011"
	request, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
func firmList() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8081/dnakit/v2/user/firmlist"
	request, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("reqUserId", "1451329585")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
func userInfo() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8081/dnakit/v2/user/userinfo"
	request, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("reqUserId", "1451329585")
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))
}
func fileDown() {
	client := http.Client{}
	urlstr := "http://127.0.0.1:8080/dnakit/v2/user/download?fileid=2"
	request, err := http.NewRequest("GET", urlstr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%v\n%d\n", contents, len(contents))

}

func main() {
	//push_device()
	//getPhoneCode()
	//postRegist()
	//postLogin()
	//postCompanyApply()
	//getCompanyApply()
	//postDeveloperApply()
	//queryCompany()
	//firmList()
	//firmwareDelete()
	//uploadfile()
	//userInfo()
	//fileDown()
	postJson()
	//deployQuery()
}

func CalTokenEx(times string) string {
	h := sha1.New()
	//str := "dna.broadlink.com.cn" + times + "BroadLinkDNA@"
	str := times + "broadlinkDNA@"
	fmt.Printf("%s\n", str)
	h.Write([]byte(str))
	checksum := hex.EncodeToString(h.Sum(nil))
	return checksum
}
