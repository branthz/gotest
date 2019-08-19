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
	emailServ     = "smtp.mxhichina.com:25"
	emailDest     = "zhanglei@globalsdn.com"
)

type mailer struct {
	dst    string
	passwd string
	user   string
	server string
	ssl    bool
}

func newMailer() *mailer {
	m := new(mailer)
	m.dst = emailDest
	m.server = emailServ
	m.user = emailUsername
	m.passwd = emailPasswd
	return m
}
func (m *mailer) setSSL() {
	m.ssl = true
}

func (m *mailer) send(subject, body string) error {
	hp := strings.Split(m.server, ":")
	auth := smtp.PlainAuth("", m.user, m.passwd, hp[0])
	content_type := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + m.dst + "\nFrom: " + m.user + "<" + m.user + ">\nSubject: " + subject + "\n" + content_type + "\n\n" + body)
	send_to := strings.Split(m.dst, ";")
	err := smtp.SendMail(m.server, auth, m.user, send_to, msg)
	return err
}

func main() {
	m := newMailer()
	tm := time.Now().Format("2006-01-02 03:04pm")
	body := fmt.Sprintf("<html>\n<body>\n<h3>\n请部署license域名，lid:,如有不明，请联系小王.\n</h3>\n<h4>\"%s\"</h4>\n</body>\n</html>", tm)
	err := m.send("module down alarm", body)
	if err != nil {
		fmt.Println(err)
	}
}

func MailSend(lid string) {
	subject := "license Domain Apply"
	tm := time.Now().Format("2006-01-02 03:04pm")
	body := fmt.Sprintf("<html>\n<body>\n<h3>\n请部署license域名，lid:%s,如有不明，请联系窦德厚.\n</h3>\n<h4>\"%s\"</h4>\n</body>\n</html>", lid, tm)
	err := sendmail(emailUsername, emailPasswd, emailServ, emailDest, subject, body, "html", false)
	if err != nil {
		fmt.Printf("send mail error!:%s", err.Error())
		return
	}
	fmt.Println("send mail success!")
}

func sendmail(user, password, host, to, subject, body, mailtype string, tls bool) error {
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
