package usecases

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/ilhammhdd/go-toolkit/errorkit"
)

type Mail struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Body    []byte
}

func (m Mail) SendWithTLS() error {
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")

	header := make(map[string]string)
	header["From"] = m.From.String()
	header["To"] = m.To.String()
	header["Subject"] = m.Subject
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString(m.Body)

	servername := host + ":" + os.Getenv("MAIL_PORT")
	auth := smtp.PlainAuth("", m.From.Address, password, host)

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	client, err := smtp.Dial(servername)
	errorkit.ErrorHandled(err)

	err = client.StartTLS(tlsConf)
	errorkit.ErrorHandled(err)

	err = client.Auth(auth)
	errorkit.ErrorHandled(err)

	err = client.Mail(m.From.Address)
	errorkit.ErrorHandled(err)

	err = client.Rcpt(m.To.Address)
	errorkit.ErrorHandled(err)

	mailWriter, err := client.Data()
	errorkit.ErrorHandled(err)

	_, err = mailWriter.Write([]byte(message))
	errorkit.ErrorHandled(err)

	err = mailWriter.Close()
	errorkit.ErrorHandled(err)

	defer func() {
		client.Quit()
		client.Close()
	}()

	return err
}
