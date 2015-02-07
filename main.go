package main

import (
	"fmt"
	"os"
	"net/smtp"
	"flag"
	"bufio"
	"log"
	"strings"
)

var subject, cc, bcc, server, from *string
var recipients []string

func init() {
	server = flag.String("server", "localhost:25", "smtp server")
	from = flag.String("from", "do-no-reply@example.com", "your from address")
	subject = flag.String("subject", "<no subject>", "mail subject")
	cc = flag.String("cc", "", "CC list")
	bcc = flag.String("bcc", "", "BCC list")
}


func main() {
	flag.Parse()
	recipients = flag.Args()
	var cc_list, bcc_list []string
	if len(*cc) > 0 {
		cc_list = strings.Split(*cc, ",")
	}
	if len(*bcc) > 0 {
		bcc_list = strings.Split(*bcc, ",")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Email text: ")
	message, _ := reader.ReadString('\x04')
	message = fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nCC: %s\r\nBCC: %s\r\nSubject: %s\r\n\r\n%s",
		*from,
		strings.Join(recipients, ","),
		*cc,
		*bcc,
		*subject,
		message,
	)
	err := smtp.SendMail(
		*server, nil, *from,
		append(append(recipients, cc_list...), bcc_list...),
		[]byte(message),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Email successfully sent!")
}