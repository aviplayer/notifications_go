package models

type EmailAuth struct {
	user           string
	psw            string
	smtpServerAddr string
	smtpServerPort int16
}

func (EmailAuth) NewAuth(
	user string,
	psw string,
	smtpServerAddr string,
	smtpServerPort int16,
) EmailAuth {
	return EmailAuth{
		user, psw, smtpServerAddr, smtpServerPort,
	}
}
