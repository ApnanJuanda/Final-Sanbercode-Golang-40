package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"project/model"
	
	"github.com/joho/godotenv"
	"github.com/k3a/html2text"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ParseTemplateDir(dir string) (*template.Template, error) {
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

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmail(customer *model.Customer, data *model.EmailData) (string, error) {
	// ENV Configuration
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environtment")
	}

	// template email
	var body bytes.Buffer
	template, err := ParseTemplateDir("template")
	if err != nil {
		return "Could not parse template", err
	}
	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	// Initialise the required mail message variables
	from := mail.NewEmail(os.Getenv("SEND_FROM_NAME"), os.Getenv("SEND_FROM_ADDRESS"))
	subject := "Welcome to Electronics Component Store"
	to := mail.NewEmail(data.FirstName, customer.Email)
	plainTextContent := html2text.HTML2Text(body.String())
	htmlContent := body.String()
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Attempt to send the email
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		fmt.Println("Unable to send your email")
		log.Fatal(err)
	}

	// Check if it was sent
	statusCode := response.StatusCode
	if statusCode == 200 || statusCode == 201 || statusCode == 202 {
		fmt.Println(statusCode)
		fmt.Println("Email sent!")
	}

	result := "We sent an email with a verification code to " + customer.Email
	return result, nil
} 