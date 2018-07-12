package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	js "github.com/bitly/go-simplejson"
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
func gzipWrite(w io.Writer, data []byte) error {
	// Write gzipped data to the client
	gw, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
	defer gw.Close()
	gw.Write(data)
	return err
}

func postJson() {
	var data = `{ "version": "v1.3.1", "timestamp_seconds": "09:38:54", "tunnel_connected": [ ], "tunnel_speed": [ { "module": "tunnel", "event": "rate", "tunnel_id": 14, "step_sn": "STEP01R1201807030000000000000001", "peer_step_sn": "STEP01L1201807030000000000000009", "last_send_bytes_raw": 31, "last_recv_bytes_raw": 15, "last_send_bytes": 31, "last_recv_bytes": 15, "ts_seconds": 1531359526, "last_tx_ok_counts": 1, "last_tx_busy_counts": 0, "last_max_cache_size": 0, "cache_counts": 0, "tx_rate_kbps": 0, "peer_tx_rate_kbps": 0, "buy_bw_mbps": 100, "cache_bytes": 0, "total_send_kbytes_raw": 1, "total_recv_kbytes_raw": 0, "total_send_kbytes": 1, "total_recv_kbytes": 0, "session_count": 0, "total_mem_used": 29535935 }, { "module": "tunnel", "event": "rate", "tunnel_id": 32, "step_sn": "STEP01R1201807030000000000000001", "peer_step_sn": "STEP01L1201807030000000000000005", "last_send_bytes_raw": 327, "last_recv_bytes_raw": 2381, "last_send_bytes": 327, "last_recv_bytes": 2381, "ts_seconds": 1531359526, "last_tx_ok_counts": 29, "last_tx_busy_counts": 0, "last_max_cache_size": 0, "cache_counts": 0, "tx_rate_kbps": 0, "peer_tx_rate_kbps": 0, "buy_bw_mbps": 100, "cache_bytes": 0, "total_send_kbytes_raw": 9, "total_recv_kbytes_raw": 71, "total_send_kbytes": 9, "total_recv_kbytes": 71, "session_count": 0, "total_mem_used": 29535935 } ], "tunnel_disconnect": [ ], "tunnel_reconnected": [ ], "session_info": [ ], "mpath_timer": [ { "module": "mpath", "event": "info", "tunnel_id": 16, "step_sn": "STEP01R1201807030000000000000001", "peer_step_sn": "STEP01L1201807030000000000000007", "start_seconds": 1531359524, "duration_seconds": 30, "work_path_id": 2, "path_conf": "->10.0.33.13:8819:3", "path": [ { "type": 1, "link_num": 7, "path_rtt": 0, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 }, { "type": 1, "link_num": 7, "path_rtt": 8, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 }, { "type": 0, "link_num": 7, "path_rtt": 5, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 } ] }, { "module": "mpath", "event": "info", "tunnel_id": 14, "step_sn": "STEP01R1201807030000000000000001", "peer_step_sn": "STEP01L1201807030000000000000009", "start_seconds": 1531359524, "duration_seconds": 30, "work_path_id": 1, "path_conf": "->60.191.35.164:8819:1", "path": [ { "type": 1, "link_num": 7, "path_rtt": 0, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 }, { "type": 1, "link_num": 7, "path_rtt": 7, "busy_num": 0, "rx_bytes": 15, "tx_bytes": 31 }, { "type": 0, "link_num": 7, "path_rtt": 0, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 } ] }, { "module": "mpath", "event": "info", "tunnel_id": 32, "step_sn": "STEP01R1201807030000000000000001", "peer_step_sn": "STEP01L1201807030000000000000005", "start_seconds": 1531359525, "duration_seconds": 30, "work_path_id": 2, "path_conf": "->10.0.31.13:8819:3", "path": [ { "type": 1, "link_num": 7, "path_rtt": 0, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 }, { "type": 1, "link_num": 7, "path_rtt": 9, "busy_num": 0, "rx_bytes": 0, "tx_bytes": 0 }, { "type": 0, "link_num": 7, "path_rtt": 3, "busy_num": 0, "rx_bytes": 2381, "tx_bytes": 327 } ] } ] }`
	var buf bytes.Buffer
	err := gzipWrite(&buf, []byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(buf.Bytes())
	client := http.Client{}
	urlstr := "http://127.0.0.1:9898/v1/step/data"
	request, err := http.NewRequest("POST", urlstr, body)
	if err != nil {
		fmt.Printf("2222%v\n", err)
		os.Exit(-1)
	}
	tm := fmt.Sprintf("%d", time.Now().Unix())
	request.Header.Add("timestamp", tm)
	sign := CalTokenEx("/dnakit/server/deploy" + tm)
	request.Header.Add("sign", sign)
	request.Header.Add("Content-Encoding", "application/gzip")

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
