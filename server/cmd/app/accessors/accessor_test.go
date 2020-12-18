package accessors

import (
	. "andrew.com/notifications/cmd/app/models"
	"log"
	"testing"
)

var (
	tpl = `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
	        "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html>
	</head>
	<body>
	<p>
	    Build: {{.Name}} / {{.Version}} - {{.Status}}
	</p>
	<p>
		Last commit: {{.LastCommit}}	
	</p
	</body>
	</html>
	`
	n = Notification{}.NewNotification(
		"Test notification",
		Email,
		"test",
		"Notification for testing",
	)

	buildInfo = Build{
		LastCommit: "@andrew",
		Version:    "1.0.0.0",
		Status:     Success,
		Name:       "Data Audit",
	}

	auth = EmailAuth{}.NewAuth("vasya.pupkin@test.com", "909090", "smtp.gmail.com", 465)

	emailN = EmailNotification{
		BaseNotification: *n,
		BuildInfo:        buildInfo,
		Subject:          "Build executed",
		To:               []string{"vasya.pupkin@test.com"},
		Auth:             auth,
	}

	ID = 0
)

func TestInsertion(t *testing.T) {
	conn := CreateConnection()
	emailAccessor := EmailAccessor{&conn}
	noti, err := emailAccessor.InsertNotification(emailN)
	if err != nil {
		log.Fatalf("--TEST FAIL-- \n %s", err)
	}
	ID = int(noti.BaseNotification.Id)
}

func TestUpdate(t *testing.T) {
	conn := CreateConnection()
	emailN.BaseNotification.Template = tpl
	emailN.BaseNotification.Id = int8(ID)
	emailAccessor := EmailAccessor{&conn}
	if err := emailAccessor.upsertNotification(emailN); err != nil {
		log.Fatalf("--TEST FAIL-- \n %s", err)
	}
}

func TestGet(t *testing.T) {
	conn := CreateConnection()
	emailAccessor := EmailAccessor{&conn}
	notification, err := emailAccessor.GetNotification(ID)
	if err != nil {
		log.Fatalf("--TEST FAIL-- \n %s", err)
	}
	log.Println("Notification is {}", notification)
	if emailN.Auth != notification.Auth {
		t.Fatalf("Auth structs are not equals \n %v \n %v",
			emailN.Auth, notification.Auth)
	}
	if emailN.BaseNotification.Template != notification.BaseNotification.Template {
		t.Fatalf("Auth structs are not equals \n %s \n %s",
			emailN.BaseNotification.Template, notification.BaseNotification.Template)
	}
}

//func TestDelete(t *testing.T) {
//	conn := CreateConnection()
//	emailAccessor := EmailAccessor{&conn}
//	if err := emailAccessor.deleteNotification(); err != nil {
//		log.Fatalf("--TEST FAIL-- \n %s", err)
//	}
//}
