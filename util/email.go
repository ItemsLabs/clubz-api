package util

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

var smtpHost = os.Getenv("SMTP_HOST")
var smtpPort = os.Getenv("SMTP_PORT")
var smtpUser = os.Getenv("SMTP_USER")
var smtpPassword = os.Getenv("SMTP_PASSWORD")

// EmailData holds the data to be injected into the template
type EmailData struct {
	Email string
	Url   string
}

func ConstructEmailMessage(email, token string) []byte {
	var url string
	envName := os.Getenv("ENV_NAME")
	if envName == "prd" {
		url = "https://livefantasy.gameon.app/api/auth/verify-email?token=" + token
	} else if envName == "dev" {
		url = "https://laliga.gamebuild.co/api/auth/verify-email?token=" + token
	} else {
		url = "https://localhost:8080/api/auth/verify-email?token=" + token
	}

	data := EmailData{
		Email: email,
		Url:   url,
	}

	tmpl, err := template.ParseFiles("templates/email.html")
	if err != nil {
		fmt.Println("error parsing template:", err)
		return nil
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Println("error executing template:", err)
		return nil
	}

	return []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: Welcome to GameOn Live Fantasy\r\n"+
		"MIME-version: 1.0;\nContent-Type: text/html;charset=\"UTF-8\"\r\n\r\n"+
		"%s", email, body.String()))
}

func SendEmail(ctx context.Context, to []string, body []byte) error {
	auth := smtp.PlainAuth("", smtpUser, smtpPassword, smtpHost)
	from := "no-reply@hello.gameon.app"

	select {
	case <-ctx.Done():
		return fmt.Errorf("email sending cancelled: %w", ctx.Err())
	default:
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body)
		if err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}
	}
	return nil
}
