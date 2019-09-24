package main

import (
	"fmt"
	"time"

	//"strconv"
	"strings"
	//"unsafe"
	"net/smtp"
)

const (
	emailUsername = "service@globalsdn.com"
	emailPasswd   = "supercxp@123"
	emailServ     = "smtp.mxhichina.com:465"
	emailDest     = "zzzzzlll@yeah.net"
)

func main() {
	MailSend("123456789")
}

func MailSend(lid string) {
	subject := "license Domain Apply"
	tm := time.Now().Format("2006-01-02 03:04pm")
	body := fmt.Sprintf("<html>\n<body>\n<h3>\n请部署license域名，lid:%s,如有不明，请联系小王.\n</h3>\n<h4>\"%s\"</h4>\n</body>\n</html>", lid, tm)
	err := sendmail(emailUsername, emailPasswd, emailServ, emailDest, subject, body, "html")
	if err != nil {
		fmt.Printf("send mail error!:%s", err.Error())
		return
	}
	fmt.Println("send mail success!")
}

func sendmail(user, password, host, to, subject, body, mailtype string) error {
	fmt.Println("=============================")
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\nFrom: " + user + "<" + user + ">\nSubject: " + subject + "\n" + content_type + "\n\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
