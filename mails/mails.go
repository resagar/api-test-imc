package mails

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"
	"os"
)

type Dest struct {
	Name string
}

type Acount struct {
	serverName string
	host       string
	username   string
	password   string
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
func SendEmailContact(toName string, toAddress string, subject string) {
	from := mail.Address{Name: os.Getenv("NAME_COMPANY"), Address: os.Getenv("USER_NAME")}
	to := mail.Address{Name: toName, Address: toAddress}

	newEmail(from, to, subject, "templateContact")
}

func SendEmailTestResult(toName string, toAddress string, subject string) {
	from := mail.Address{Name: os.Getenv("NAME_COMPANY"), Address: os.Getenv("USER_NAME")}
	to := mail.Address{Name: toName, Address: toAddress}

	newEmail(from, to, subject, "templateTestResult")
}

func newEmail(from mail.Address, to mail.Address, subject string, fileTemplate string) {
	acount := Acount{
		serverName: os.Getenv("SERVER_NAME"),
		host:       os.Getenv("HOST"),
		username:   os.Getenv("USER_NAME"),
		password:   os.Getenv("PASSWORD")}

	message := newHeaders(from, to, subject, fileTemplate)

	auth := smtp.PlainAuth("", acount.username, acount.password, acount.host)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         acount.host,
	}
	conn, err := tls.Dial("tcp", acount.serverName, tlsConfig)
	check(err)

	client, err := smtp.NewClient(conn, acount.host)
	check(err)

	err = client.Auth(auth)
	check(err)

	err = client.Mail(from.Address)
	check(err)

	err = client.Rcpt(to.Address)
	check(err)

	w, err := client.Data()
	check(err)

	_, err = w.Write([]byte(message))
	check(err)

	err = w.Close()
	check(err)

	client.Quit()
}

func newHeaders(from mail.Address, to mail.Address, subject string, fileTemplate string) string {
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += parseTemplate(fileTemplate, to)
	return message
}

func parseTemplate(fileTemplate string, to mail.Address) string {
	t, err := template.ParseFiles(fmt.Sprintf("./mails/templates/%s.html", fileTemplate))
	check(err)
	buf := new(bytes.Buffer)
	err = t.Execute(buf, Dest{Name: to.Address})
	check(err)
	return buf.String()
}
