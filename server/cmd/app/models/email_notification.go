package models

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var mime = "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"

type EmailNotification struct {
	Subject          string
	BaseNotification Notification
	Auth             EmailAuth
	To               []string
	BuildInfo        Build
}

func NewEmailNotification(subj string, no Notification, auth EmailAuth, to []string, build Build) *EmailNotification {
	return &EmailNotification{
		subj,
		no,
		auth,
		to,
		build,
	}
}

func (n EmailNotification) GetAuth() (string, string, string, int16) {
	return n.Auth.user, n.Auth.psw, n.Auth.smtpServerAddr, n.Auth.smtpServerPort
}

func (n EmailNotification) sendMessage() (bool, error) {
	cred := n.Auth
	body, err := n.createMessage()
	if err != nil {
		return false, err
	}
	msg := []byte(n.Subject + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%d", cred.smtpServerAddr, cred.smtpServerPort)
	auth := smtp.PlainAuth("", cred.user, cred.psw, cred.smtpServerAddr)
	if err := smtp.SendMail(addr, auth, cred.user, n.To, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (n EmailNotification) createMessage() (string, error) {
	t, err := template.New("email").Parse(n.BaseNotification.Template)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, n.BuildInfo); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

func (n *EmailNotification) GetDbNotification() DbENRecord {
	email, pwd, server, port := n.GetAuth()
	return DbENRecord{
		Id:               0,
		Title:            n.BaseNotification.Title,
		NotificationType: Types[Email],
		Template:         n.BaseNotification.Template,
		Email:            email,
		Pwd:              pwd,
		SmtpServer:       server,
		SmtpPort:         port,
		Description:      n.BaseNotification.Description,
	}
}
