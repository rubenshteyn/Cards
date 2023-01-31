package email

import (
	"bytes"
	"fmt"
	"github.com/k3a/html2text"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	template1 "html/template"
	"log"
	"os"
	"path/filepath"
)

type Data struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(emailFrom string, email string, data *Data, templateName string) error {
	// Sender data.
	from := emailFrom
	smtpUser := viper.Get("SMTP_USER").(string)
	smtpPass := viper.Get("SMTP_PASS").(string)
	to := email
	smtpHost := viper.Get("SMTP_HOST").(string)
	smtpPort := 465 //viper.Get("SMTP_PORT").(int)

	var body bytes.Buffer

	template, err := ParseTemplateDir("/go/Grats.Cards/backend/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template = template.Lookup(templateName)
	err = template.Execute(&body, &data)
	if err != nil {
		return err
	}

	fmt.Println(template.Name())
	fmt.Println(smtpUser)
	fmt.Println(smtpPass)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // insecure
	fmt.Println("here")
	// Send Email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ParseTemplateDir(dir string) (*template1.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	fmt.Println("Am parsing templates...")

	if err != nil {
		return nil, err
	}

	return template1.ParseFiles(paths...)
}
