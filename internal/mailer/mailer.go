package mailer

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	mailer "github.com/xhit/go-simple-mail/v2"

	"github.com/shurco/litecart/internal/models"
)

var EncryptionTypes = map[string]mailer.Encryption{
	"None":     mailer.EncryptionNone,
	"SSL/TLS":  mailer.EncryptionSSL,
	"STARTTLS": mailer.EncryptionTLS,
}

// SendMail is ...
func SendMail(smtp *models.SMTP, mail *models.Mail) error {
	server := mailer.NewSMTPClient()
	server.Host = smtp.Host
	server.Port = smtp.Port
	server.Username = smtp.Username
	server.Password = smtp.Password
	server.Encryption = EncryptionTypes[smtp.Encryption]

	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mailer.NewMSG()
	email.SetFrom(mail.From).
		AddTo(mail.To).
		SetSubject(mail.Letter.Subject)

	bodyText, err := textTemplate(mail.Letter.Text, mail.Data)
	if err != nil {
		return err
	}
	email.SetBodyData(mailer.TextPlain, bodyText)
	// email.AddAlternativeData(mail.TextPlain, "Hello Gophers!")

	if len(mail.Files) > 0 {
		for _, file := range mail.Files {
			email.Attach(&mailer.File{
				FilePath: fmt.Sprintf("./lc_digitals/%s.%s", file.Name, file.Ext),
				Name:     file.OrigName,
			})
		}
	}

	if err := email.Send(smtpClient); err != nil {
		return err
	}

	return nil
}

func textTemplate(tmp string, data any) ([]byte, error) {
	tmpl, err := template.New("").Parse(tmp)
	if err != nil {
		return nil, err
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}
