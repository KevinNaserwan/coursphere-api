package mail

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func RenderEmailVerificationTemplate(data EmailVerification, frontEndUrl string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "email_verification.html", map[string]string{
		"Link": fmt.Sprintf("%s?token=%s", frontEndUrl, data.Token),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderPasswordResetTemplate(data PasswordReset, frontEndUrl string) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "reset_password.html", map[string]string{
		"Link": fmt.Sprintf("%s/reset-password?token=%s", frontEndUrl, data.Token),
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
