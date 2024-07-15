package mail

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed templates/*.html
var templates embed.FS

var parsedTemplates = template.Must(template.ParseFS(templates, "templates/*.html"))

func RenderEmailVerificationTemplate(data EmailVerification) (string, error) {
	var buf bytes.Buffer
	err := parsedTemplates.ExecuteTemplate(&buf, "email_verification.html", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
