package models

type DbENRecord struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	NotificationType int    `json:"type"`
	Template         string `json:"template"`
	Email            string `json:"email"`
	Pwd              string `json:"pwd"`
	SmtpServer       string `json:"smtp_server"`
	SmtpPort         int16  `json:"smtp_port"`
	Description      string `json:"description"`
}

func (record DbENRecord) SetConfigToNotification(notification *EmailNotification) {
	notification.Auth = EmailAuth{
		user:           record.Email,
		psw:            record.Pwd,
		smtpServerAddr: record.SmtpServer,
		smtpServerPort: record.SmtpPort,
	}
	notification.BaseNotification = Notification{
		Type:        0,
		Template:    record.Template,
		Description: record.Description,
	}
}
